package watcher_test

import (
	"github.com/ethereum/go-ethereum/common"
	ethcmn "github.com/ethereum/go-ethereum/common"
	jsoniter "github.com/json-iterator/go"
	"github.com/okex/exchain/app"
	"github.com/okex/exchain/app/crypto/ethsecp256k1"
	ethermint "github.com/okex/exchain/app/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	abci "github.com/okex/exchain/libs/tendermint/abci/types"
	"github.com/okex/exchain/libs/tendermint/crypto/secp256k1"
	"github.com/okex/exchain/libs/tendermint/crypto/tmhash"
	"github.com/okex/exchain/x/evm"
	"github.com/okex/exchain/x/evm/types"
	"github.com/okex/exchain/x/evm/watcher"
	"github.com/spf13/viper"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type KV struct {
	k []byte
	v []byte
}

func calcHash(kvs []KV) []byte {
	ha := tmhash.New()
	// calc a hash
	for _, kv := range kvs {
		ha.Write(kv.k)
		ha.Write(kv.v)
	}
	return ha.Sum(nil)
}

type WatcherTestSt struct {
	ctx     sdk.Context
	app     *app.OKExChainApp
	handler sdk.Handler
}

func setupTest() *WatcherTestSt {
	w := &WatcherTestSt{}
	checkTx := false
	viper.Set(watcher.FlagFastQuery, true)
	viper.Set(watcher.FlagDBBackend, "memdb")
	w.app = app.Setup(checkTx)
	w.ctx = w.app.BaseApp.NewContext(checkTx, abci.Header{Height: 1, ChainID: "ethermint-3", Time: time.Now().UTC()})
	w.handler = evm.NewHandler(w.app.EvmKeeper)

	params := types.DefaultParams()
	params.EnableCreate = true
	params.EnableCall = true
	w.app.EvmKeeper.SetParams(w.ctx, params)

	return w
}

func getDBKV(db *watcher.WatchStore) []KV {
	var kvs []KV
	it := db.Iterator(nil, nil)
	for it.Valid() {
		kvs = append(kvs, KV{it.Key(), it.Value()})
		it.Next()
	}
	return kvs
}
func flushDB(db *watcher.WatchStore) {
	it := db.Iterator(nil, nil)
	for it.Valid() {
		db.Delete(it.Key())
		it.Next()
	}
}

func delDirtyAccount(wdBytes []byte, w *WatcherTestSt) error {
	wd := watcher.WatchData{}
	if err := json.Unmarshal(wdBytes, &wd); err != nil {
		return err
	}
	for _, account := range wd.DirtyAccount {
		w.app.EvmKeeper.Watcher.DeleteAccount(*account)
	}
	return nil
}

func testWatchData(t *testing.T, w *WatcherTestSt) {
	// produce WatchData
	w.app.EvmKeeper.Watcher.Commit()
	time.Sleep(time.Second * 1)

	// get WatchData
	wd, err := w.app.EvmKeeper.Watcher.GetWatchData()
	require.Nil(t, err)
	require.NotEmpty(t, wd)
	err = delDirtyAccount(wd, w)
	require.Nil(t, err)

	store := watcher.InstanceOfWatchStore()
	pWd := getDBKV(store)
	flushDB(store)

	// use WatchData
	w.app.EvmKeeper.Watcher.UseWatchData(wd)
	time.Sleep(time.Second * 1)

	cWd := getDBKV(store)

	// compare db_kv of producer and consumer
	require.Equal(t, pWd, cWd, "compare len:", "pwd:", len(pWd), "cwd", len(cWd))
	pHash := calcHash(pWd)
	cHash := calcHash(cWd)
	require.NotEmpty(t, pHash)
	require.NotEmpty(t, cHash)
	require.Equal(t, pHash, cHash)
}

func TestHandleMsgEthereumTx(t *testing.T) {
	w := setupTest()
	privkey, err := ethsecp256k1.GenerateKey()
	require.NoError(t, err)
	sender := ethcmn.HexToAddress(privkey.PubKey().Address().String())

	var tx types.MsgEthereumTx

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"passed",
			func() {
				w.app.EvmKeeper.SetBalance(w.ctx, sender, big.NewInt(100))
				tx = types.NewMsgEthereumTx(0, &sender, big.NewInt(100), 3000000, big.NewInt(1), nil)

				// parse context chain ID to big.Int
				chainID, err := ethermint.ParseChainID(w.ctx.ChainID())
				require.NoError(t, err)

				// sign transaction
				err = tx.Sign(chainID, privkey.ToECDSA())
				require.NoError(t, err)
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			w = setupTest() // reset
			//nolint
			tc.malleate()
			w.ctx = w.ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
			res, err := w.handler(w.ctx, tx)

			//nolint
			if tc.expPass {
				require.NoError(t, err)
				require.NotNil(t, res)
				var expectedConsumedGas uint64 = 21000
				require.EqualValues(t, expectedConsumedGas, w.ctx.GasMeter().GasConsumed())
			} else {
				require.Error(t, err)
				require.Nil(t, res)
			}

			testWatchData(t, w)
		})
	}
}

func TestMsgEthermint(t *testing.T) {
	var (
		tx   types.MsgEthermint
		from = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
		to   = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	)
	w := setupTest()
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"passed",
			func() {
				tx = types.NewMsgEthermint(0, &to, sdk.NewInt(1), 100000, sdk.NewInt(2), []byte("test"), from)
				w.app.EvmKeeper.SetBalance(w.ctx, ethcmn.BytesToAddress(from.Bytes()), big.NewInt(100))
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			w = setupTest() // reset
			//nolint
			tc.malleate()
			w.ctx = w.ctx.WithIsCheckTx(true)
			w.ctx = w.ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
			res, err := w.handler(w.ctx, tx)

			//nolint
			if tc.expPass {
				require.NoError(t, err)
				require.NotNil(t, res)
				var expectedConsumedGas uint64 = 21064
				require.EqualValues(t, expectedConsumedGas, w.ctx.GasMeter().GasConsumed())
			} else {
				require.Error(t, err)
				require.Nil(t, res)
			}

			testWatchData(t, w)
		})
	}
}

func TestDeployAndCallContract(t *testing.T) {
	w := setupTest()

	// Deploy contract - Owner.sol
	gasLimit := uint64(100000000)
	gasPrice := big.NewInt(10000)

	priv, err := ethsecp256k1.GenerateKey()
	require.NoError(t, err, "failed to create key")

	sender := ethcmn.HexToAddress(priv.PubKey().Address().String())
	w.app.EvmKeeper.SetBalance(w.ctx, sender, big.NewInt(100))

	bytecode := common.FromHex("0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73560405160405180910390a36102c4806100dc6000396000f3fe608060405234801561001057600080fd5b5060043610610053576000357c010000000000000000000000000000000000000000000000000000000090048063893d20e814610058578063a6f9dae1146100a2575b600080fd5b6100606100e6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100e4600480360360208110156100b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061010f565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73560405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fea265627a7a72315820f397f2733a89198bc7fed0764083694c5b828791f39ebcbc9e414bccef14b48064736f6c63430005100032")
	tx := types.NewMsgEthereumTx(1, &sender, big.NewInt(0), gasLimit, gasPrice, bytecode)
	tx.Sign(big.NewInt(3), priv.ToECDSA())
	require.NoError(t, err)

	result, err := w.handler(w.ctx, tx)
	require.NoError(t, err, "failed to handle eth tx msg")

	resultData, err := types.DecodeResultData(result.Data)
	require.NoError(t, err, "failed to decode result data")

	testWatchData(t, w)

	// store - changeOwner
	gasLimit = uint64(100000000000)
	gasPrice = big.NewInt(100)
	receiver := common.HexToAddress(resultData.ContractAddress.String())

	storeAddr := "0xa6f9dae10000000000000000000000006a82e4a67715c8412a9114fbd2cbaefbc8181424"
	bytecode = common.FromHex(storeAddr)
	tx = types.NewMsgEthereumTx(2, &receiver, big.NewInt(0), gasLimit, gasPrice, bytecode)
	tx.Sign(big.NewInt(3), priv.ToECDSA())
	require.NoError(t, err)

	result, err = w.handler(w.ctx, tx)
	require.NoError(t, err, "failed to handle eth tx msg")

	resultData, err = types.DecodeResultData(result.Data)
	require.NoError(t, err, "failed to decode result data")

	testWatchData(t, w)

	// query - getOwner
	bytecode = common.FromHex("0x893d20e8")
	tx = types.NewMsgEthereumTx(2, &receiver, big.NewInt(0), gasLimit, gasPrice, bytecode)
	tx.Sign(big.NewInt(3), priv.ToECDSA())
	require.NoError(t, err)

	result, err = w.handler(w.ctx, tx)
	require.NoError(t, err, "failed to handle eth tx msg")

	resultData, err = types.DecodeResultData(result.Data)
	require.NoError(t, err, "failed to decode result data")

	getAddr := strings.ToLower(hexutils.BytesToHex(resultData.Ret))
	require.Equal(t, true, strings.HasSuffix(storeAddr, getAddr), "Fail to query the address")

	testWatchData(t, w)
}
