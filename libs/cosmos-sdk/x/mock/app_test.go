package mock

import (
	"testing"

	"github.com/okex/exchain/libs/cosmos-sdk/x/supply"
	abci "github.com/okex/exchain/libs/tendermint/abci/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/okex/exchain/libs/cosmos-sdk/types/errors"
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth"
	"github.com/okex/exchain/libs/cosmos-sdk/x/supply/exported"
)

const msgRoute = "testMsg"

var (
	numAccts                 = 2
	genCoins                 = sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 77)}
	accs, addrs, _, privKeys = CreateGenAccounts(numAccts, genCoins)
)

// testMsg is a mock transaction that has a validation which can fail.
type testMsg struct {
	signers     []sdk.AccAddress
	positiveNum int64
}

func (tx testMsg) Route() string                      { return msgRoute }
func (tx testMsg) Type() string                       { return "test" }
func (tx testMsg) GetMsg() sdk.Msg                    { return tx }
func (tx testMsg) GetMemo() string                    { return "" }
func (tx testMsg) GetSignBytes() []byte               { return nil }
func (tx testMsg) GetSigners() []sdk.AccAddress       { return tx.signers }
func (tx testMsg) GetSignatures() []auth.StdSignature { return nil }
func (tx testMsg) ValidateBasic() error {
	if tx.positiveNum >= 0 {
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "positiveNum should be a non-negative integer")
}

// getMockApp returns an initialized mock application.
func getMockApp(t *testing.T) *App {
	mApp := NewApp()

	mApp.Router().AddRoute(msgRoute, func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		return &sdk.Result{}, nil
	})
	require.NoError(t, mApp.CompleteSetup())

	return mApp
}

func TestCheckAndDeliverGenTx(t *testing.T) {
	mApp := getMockApp(t)
	mApp.Cdc.RegisterConcrete(testMsg{}, "mock/testMsg", nil)
	mApp.Cdc.RegisterInterface((*exported.ModuleAccountI)(nil), nil)
	mApp.Cdc.RegisterConcrete(supply.ModuleAccount{}, "cosmos-sdk/ModuleAccount", nil)

	SetGenesis(mApp, accs)
	ctxCheck := mApp.BaseApp.NewContext(true, abci.Header{})

	msg := testMsg{signers: []sdk.AccAddress{addrs[0]}, positiveNum: 1}

	acct := mApp.AccountKeeper.GetAccount(ctxCheck, addrs[0])
	require.Equal(t, accs[0], acct.(*auth.BaseAccount))

	header := abci.Header{Height: mApp.LastBlockHeight() + 1}
	SignCheckDeliver(
		t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{msg},
		[]uint64{accs[0].GetAccountNumber()}, []uint64{accs[0].GetSequence()},
		true, true, privKeys[0],
	)

	// Signing a tx with the wrong privKey should result in an auth error
	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
	_, _, err := SignCheckDeliver(
		t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{msg},
		[]uint64{accs[1].GetAccountNumber()}, []uint64{accs[1].GetSequence() + 1},
		true, false, privKeys[1],
	)

	// Will fail on SetPubKey decorator
	space, code, log := sdkerrors.ABCIInfo(err, false)
	require.Equal(t, sdkerrors.ErrInvalidPubKey.ABCICode(), code, log)
	require.Equal(t, sdkerrors.ErrInvalidPubKey.Codespace(), space)

	// Resigning the tx with the correct privKey should result in an OK result
	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
	SignCheckDeliver(
		t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{msg},
		[]uint64{accs[0].GetAccountNumber()}, []uint64{accs[0].GetSequence() + 1},
		true, true, privKeys[0],
	)
}

func TestCheckGenTx(t *testing.T) {
	mApp := getMockApp(t)
	mApp.Cdc.RegisterConcrete(testMsg{}, "mock/testMsg", nil)
	mApp.Cdc.RegisterInterface((*exported.ModuleAccountI)(nil), nil)

	SetGenesis(mApp, accs)

	msg1 := testMsg{signers: []sdk.AccAddress{addrs[0]}, positiveNum: 1}
	CheckGenTx(
		t, mApp.BaseApp, []sdk.Msg{msg1},
		[]uint64{accs[0].GetAccountNumber()}, []uint64{accs[0].GetSequence()},
		true, privKeys[0],
	)

	msg2 := testMsg{signers: []sdk.AccAddress{addrs[0]}, positiveNum: -1}
	CheckGenTx(
		t, mApp.BaseApp, []sdk.Msg{msg2},
		[]uint64{accs[0].GetAccountNumber()}, []uint64{accs[0].GetSequence()},
		false, privKeys[0],
	)
}
