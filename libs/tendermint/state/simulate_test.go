package state

import (
	"fmt"
	"github.com/okex/exchain/libs/cosmos-sdk/baseapp"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/okex/exchain/libs/cosmos-sdk/types/errors"
	"github.com/okex/exchain/libs/iavl"
	abci "github.com/okex/exchain/libs/tendermint/abci/types"
	"github.com/okex/exchain/libs/tendermint/libs/log"
	"github.com/okex/exchain/libs/tendermint/mempool"
	"github.com/okex/exchain/libs/tendermint/proxy"
	"github.com/okex/exchain/libs/tendermint/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
	"math/big"
	"testing"
)

const (
	routeMsgCounter  = "msgCounter"
	routeMsgCounter2 = "msgCounter2"
)

var (
	capKey1 = sdk.NewKVStoreKey("key1")
	capKey2 = sdk.NewKVStoreKey("key2")
)

// Simple tx with a list of Msgs.
type txTest struct {
	Msgs       []sdk.Msg
	Counter    int64
	FailOnAnte bool
}

func (tx *txTest) setFailOnAnte(fail bool) {
	tx.FailOnAnte = fail
}

func (tx *txTest) setFailOnHandler(fail bool) {
	for i, msg := range tx.Msgs {
		tx.Msgs[i] = msgCounter{msg.(msgCounter).Counter, fail}
	}
}

// Implements Tx
func (tx txTest) GetMsgs() []sdk.Msg   { return tx.Msgs }
func (tx txTest) ValidateBasic() error { return nil }

func (tx txTest) GetTxInfo(ctx sdk.Context) mempool.ExTxInfo {
	return mempool.ExTxInfo{
		Sender:   "",
		GasPrice: big.NewInt(0),
		Nonce:    0,
	}
}

func (tx txTest) GetGasPrice() *big.Int {
	return big.NewInt(0)
}

func (tx txTest) GetTxFnSignatureInfo() ([]byte, int) {
	return nil, 0
}


// ValidateBasic() fails on negative counters.
// Otherwise it's up to the handlers
type msgCounter struct {
	Counter       int64
	FailOnHandler bool
}

// Implements Msg
func (msg msgCounter) Route() string                { return routeMsgCounter }
func (msg msgCounter) Type() string                 { return "counter1" }
func (msg msgCounter) GetSignBytes() []byte         { return nil }
func (msg msgCounter) GetSigners() []sdk.AccAddress { return nil }
func (msg msgCounter) ValidateBasic() error {
	if msg.Counter >= 0 {
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrInvalidSequence, "counter should be a non-negative integer")
}


// Another counter msg. Duplicate of msgCounter
type msgCounter2 struct {
	Counter int64
}

// Implements Msg
func (msg msgCounter2) Route() string                { return routeMsgCounter2 }
func (msg msgCounter2) Type() string                 { return "counter2" }
func (msg msgCounter2) GetSignBytes() []byte         { return nil }
func (msg msgCounter2) GetSigners() []sdk.AccAddress { return nil }
func (msg msgCounter2) ValidateBasic() error {
	if msg.Counter >= 0 {
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrInvalidSequence, "counter should be a non-negative integer")
}

// a msg we dont know how to route
type msgNoRoute struct {
	msgCounter
}

func (tx msgNoRoute) Route() string { return "noroute" }

func newTxCounter(counter int64, msgCounters ...int64) *txTest {
	msgs := make([]sdk.Msg, 0, len(msgCounters))
	for _, c := range msgCounters {
		msgs = append(msgs, msgCounter{c, false})
	}

	return &txTest{msgs, counter, false}
}

func registerTestCodec(cdc *codec.Codec) {
	// register Tx, Msg
	sdk.RegisterCodec(cdc)

	// register test types
	cdc.RegisterConcrete(&txTest{}, "cosmos-sdk/baseapp/txTest", nil)
	cdc.RegisterConcrete(&msgCounter{}, "cosmos-sdk/baseapp/msgCounter", nil)
	cdc.RegisterConcrete(&msgCounter2{}, "cosmos-sdk/baseapp/msgCounter2", nil)
	cdc.RegisterConcrete(&msgNoRoute{}, "cosmos-sdk/baseapp/msgNoRoute", nil)
}

// amino decode
func testTxDecoder(cdc *codec.Codec) sdk.TxDecoder {
	return func(txBytes []byte, _ ...int64) (sdk.Tx, error) {
		var tx txTest
		if len(txBytes) == 0 {
			return nil, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "tx bytes are empty")
		}

		err := cdc.UnmarshalBinaryLengthPrefixed(txBytes, &tx)
		if err != nil {
			return nil, sdkerrors.ErrTxDecode
		}

		return tx, nil
	}
}

func newBaseApp(name string, options ...func(*baseapp.BaseApp)) *baseapp.BaseApp {
	logger := log.TestingLogger()
	db := dbm.NewMemDB()
	codec := codec.New()
	registerTestCodec(codec)
	return baseapp.NewBaseApp(name, logger, db, testTxDecoder(codec), options...)
}

// simple one store baseapp
func setupBaseApp(t *testing.T, options ...func(*baseapp.BaseApp)) *baseapp.BaseApp {
	app := newBaseApp(t.Name(), options...)
	require.Equal(t, t.Name(), app.Name())

	// no stores are mounted
	require.Panics(t, func() {
		app.LoadLatestVersion(capKey1)
	})

	app.MountStores(capKey1, capKey2)

	// stores are mounted
	err := app.LoadLatestVersion(capKey1)
	require.Nil(t, err)
	return app
}

func setupSimulateBaseApp(t *testing.T) *baseapp.BaseApp {
	gasConsumed := uint64(5)

	anteOpt := func(bapp *baseapp.BaseApp) {
		bapp.SetAnteHandler(func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
			newCtx = ctx.WithGasMeter(sdk.NewGasMeter(gasConsumed))
			return
		})
	}

	routerOpt := func(bapp *baseapp.BaseApp) {
		bapp.Router().AddRoute(routeMsgCounter, func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
			ctx.GasMeter().ConsumeGas(gasConsumed, "test")
			return &sdk.Result{}, nil
		})
	}

	app := setupBaseApp(t, anteOpt, routerOpt)
	app.InitChain(abci.RequestInitChain{})

	return app
}

func TestDelta(t *testing.T) {
	iavl.SetProduceDelta(true)
	types.UploadDelta = true
	viper.Set("fast-query", true)

	app := setupSimulateBaseApp(t)

	// Create same codec used in txDecoder
	cdc := codec.New()
	registerTestCodec(cdc)
	gasConsumed := uint64(5)
	count := int64(1)

	header := abci.Header{Height: count}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	tx := newTxCounter(count, count)
	txBytes, err := cdc.MarshalBinaryLengthPrefixed(tx)
	require.Nil(t, err)

	// simulate a message, check gas reported
	gInfo, result, err := app.Simulate(txBytes, tx, 0)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, gasConsumed, gInfo.GasUsed)

	app.EndBlock(abci.RequestEndBlock{})

	res := app.Commit(abci.RequestCommit{})
	wd, _ := getWatchDataFunc()
	fmt.Println(wd)
	fmt.Println(res.Deltas)
}

func TestExec(t *testing.T) {
	app := setupSimulateBaseApp(t)
	cc := proxy.NewLocalClientCreator(app)
	proxyApp := proxy.NewAppConns(cc)
	err := proxyApp.Start()
	require.Nil(t, err)
	defer proxyApp.Stop()

//	height1, idx1, val1 := int64(8), 0, state.Validators.Validators[0].Address
//	height2, idx2, val2 := int64(3), 1, state.Validators.Validators[1].Address
//	ev1 := types.NewMockEvidence(height1, time.Now(), idx1, val1)
//	ev2 := types.NewMockEvidence(height2, time.Now(), idx2, val2)

	blocks, stateDB := produceBlock()

	for _, block := range blocks {
		deltas, _, err := execCommitBlockDelta(proxyApp.Consensus(), block, log.TestingLogger(), stateDB)
		require.Nil(t, err)
		fmt.Println("delta:", deltas)
	}
}