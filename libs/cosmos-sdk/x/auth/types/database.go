package types

import (
	"github.com/ethereum/go-ethereum/core/rawdb"
	ethstate "github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/okex/exchain/libs/cosmos-sdk/client/flags"
	"github.com/spf13/viper"
	"path/filepath"
	"sync"
)

var (
	gEvmMptDatabase ethstate.Database = nil

	initOnce sync.Once

	TrieDirtyDisabled = false
	TrieCacheSize uint = 1554 // MB
)

const (
	EvmDataDir = "data"
	EvmSpace   = "evm"

	FlagTrieDirtyDisabled = "trie-dirty-disabled"
	FlagTrieCacheSize = "trie-cache-size"
)

func InstanceOfEvmStore() ethstate.Database {
	initOnce.Do(func() {
		homeDir := viper.GetString(flags.FlagHome)
		file := filepath.Join(homeDir, EvmDataDir, EvmSpace+".db")
		//freezerPath := filepath.Join(homeDir, EvmDataDir, FreezerSpace)

		kvdb, err := leveldb.New(file, 128, 1024, EvmSpace, false)
		if err != nil {
			panic("fail to open level database: " + err.Error())
		}

		db := rawdb.NewDatabase(kvdb)
		//frdb, err := rawdb.NewDatabaseWithFreezer(kvdb, freezerPath, EvmSpace, false)
		//if err != nil {
		//	kvdb.Close()
		//	panic(fmt.Sprintf("fail to init evm mpt database: %v", err))
		//}

		gEvmMptDatabase = ethstate.NewDatabaseWithConfig(db, &trie.Config{
			Cache:     int(TrieCacheSize),
			Journal:   "",
			Preimages: true,
		})
	})

	return gEvmMptDatabase
}
