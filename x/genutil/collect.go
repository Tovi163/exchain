package genutil

// DONTCOVER

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	stakingtypes "github.com/okex/exchain/x/staking/types"

	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	authexported "github.com/okex/exchain/libs/cosmos-sdk/x/auth/exported"
	authtypes "github.com/okex/exchain/libs/cosmos-sdk/x/auth/types"
	"github.com/okex/exchain/libs/cosmos-sdk/x/genutil/types"
	cfg "github.com/okex/exchain/libs/tendermint/config"
	tmtypes "github.com/okex/exchain/libs/tendermint/types"
)

// GenAppStateFromConfig gets the genesis app state from the config
func GenAppStateFromConfig(cdc *codec.Codec, config *cfg.Config,
	initCfg InitConfig, genDoc tmtypes.GenesisDoc,
	genAccIterator types.GenesisAccountsIterator,
) (appState json.RawMessage, err error) {

	// process genesis transactions, else create default genesis.json
	appGenTxs, persistentPeers, err := CollectStdTxs(
		cdc, config.Moniker, initCfg.GenTxsDir, genDoc, genAccIterator)
	if err != nil {
		return appState, err
	}

	config.P2P.PersistentPeers = persistentPeers

	var nodeKeyWhiteList []string
	for _, nodeAddr := range strings.Split(persistentPeers, ",") {
		nodeKey := strings.Split(nodeAddr, "@")[0]
		nodeKeyWhiteList = append(nodeKeyWhiteList, nodeKey)
	}
	config.Mempool.NodeKeyWhitelist = nodeKeyWhiteList
	cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

	// if there are no gen txs to be processed, return the default empty state
	if len(appGenTxs) == 0 {
		return appState, errors.New("there must be at least one genesis tx")
	}

	// create the app state
	appGenesisState, err := GenesisStateFromGenDoc(cdc, genDoc)
	if err != nil {
		return appState, err
	}

	appGenesisState, err = SetGenTxsInAppGenesisState(cdc, appGenesisState, appGenTxs)
	if err != nil {
		return appState, err
	}
	appState, err = codec.MarshalJSONIndent(cdc, appGenesisState)
	if err != nil {
		return appState, err
	}

	genDoc.AppState = appState
	err = ExportGenesisFile(&genDoc, config.GenesisFile())
	return appState, err
}

// CollectStdTxs processes and validates application's genesis StdTxs and returns
// the list of appGenTxs, and persistent peers required to generate genesis.json.
func CollectStdTxs(cdc *codec.Codec, moniker, genTxsDir string,
	genDoc tmtypes.GenesisDoc, genAccIterator types.GenesisAccountsIterator,
) (appGenTxs []authtypes.StdTx, persistentPeers string, err error) {

	var fos []os.FileInfo
	fos, err = ioutil.ReadDir(genTxsDir)
	if err != nil {
		return appGenTxs, persistentPeers, err
	}

	// prepare a map of all accounts in genesis state to then validate
	// against the validators addresses
	var appState map[string]json.RawMessage
	if err := cdc.UnmarshalJSON(genDoc.AppState, &appState); err != nil {
		return appGenTxs, persistentPeers, err
	}

	addrMap := make(map[string]authexported.Account)
	genAccIterator.IterateGenesisAccounts(cdc, appState,
		func(acc authexported.Account) (stop bool) {
			addrMap[acc.GetAddress().String()] = acc
			return false
		},
	)

	// addresses and IPs (and port) validator server info
	var addressesIPs []string

	for _, fo := range fos {
		if fo == nil {
			continue
		}
		filename := filepath.Join(genTxsDir, fo.Name())
		if !fo.IsDir() && (filepath.Ext(filename) != ".json") {
			continue
		}

		// get the genStdTx
		var jsonRawTx []byte
		if jsonRawTx, err = ioutil.ReadFile(filename); err != nil {
			return appGenTxs, persistentPeers, err
		}
		var genStdTx authtypes.StdTx
		if err = cdc.UnmarshalJSON(jsonRawTx, &genStdTx); err != nil {
			return appGenTxs, persistentPeers, err
		}
		appGenTxs = append(appGenTxs, genStdTx)

		// the memo flag is used to store
		// the ip and node-id, for example this may be:
		// "528fd3df22b31f4969b05652bfe8f0fe921321d5@192.168.2.37:26656"
		nodeAddrIP := genStdTx.GetMemo()
		if len(nodeAddrIP) == 0 {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"couldn't find node's address and IP in %s", fo.Name())
		}

		// genesis transactions must be single-message
		msgs := genStdTx.GetMsgs()
		if len(msgs) != 1 {
			return appGenTxs, persistentPeers, errors.New(
				"each genesis transaction must provide a single genesis message")
		}

		// TODO abstract out staking message validation back to staking
		msg := msgs[0].(stakingtypes.MsgCreateValidator)
		// validate delegator and validator addresses and funds against the accounts in the state
		delAddr := msg.DelegatorAddress.String()
		valAddr := sdk.AccAddress(msg.ValidatorAddress).String()

		delAcc, delOk := addrMap[delAddr]
		if !delOk {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"account %v not in genesis.json: %+v", delAddr, addrMap)
		}

		_, valOk := addrMap[valAddr]
		if !valOk {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"account %v not in genesis.json: %+v", valAddr, addrMap)
		}

		if delAcc.GetCoins().AmountOf(msg.MinSelfDelegation.Denom).LT(msg.MinSelfDelegation.Amount) {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"insufficient fund for delegation %v: %v < %v",
				delAcc.GetAddress(), delAcc.GetCoins().AmountOf(msg.MinSelfDelegation.Denom), msg.MinSelfDelegation.Amount,
			)
		}

		// exclude itself from persistent peers
		if msg.Description.Moniker != moniker {
			addressesIPs = append(addressesIPs, nodeAddrIP)
		}
	}

	sort.Strings(addressesIPs)
	persistentPeers = strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}
