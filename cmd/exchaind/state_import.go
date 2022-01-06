package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	sm "github.com/okex/exchain/libs/tendermint/state"

	"github.com/spf13/viper"

	"github.com/okex/exchain/libs/cosmos-sdk/server"
	"github.com/okex/exchain/libs/tendermint/node"
	"github.com/spf13/cobra"
)

const (
	FlagExportFile = "export-state-file"
	FlagImportFile = "import-state-file"
)

func exportStateCmd(ctx *server.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-state",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("--------- export state start ---------")
			filePath := viper.GetString(FlagExportFile)
			if filePath == "" {
				panic("must set export-state-file")
			}
			dataDir := filepath.Join(ctx.Config.RootDir, "data")
			// load state
			stateDB, err := openDB(stateDB, dataDir)
			if err != nil {
				panic(err)
			}
			genesisDocProvider := node.DefaultGenesisDocProviderFunc(ctx.Config)
			state, _, err := node.LoadStateFromDBOrGenesisDocProvider(stateDB, genesisDocProvider)
			if err != nil {
				panic(err)
			}
			stateBytes, err := sm.ModuleCodec.MarshalJSON(state)
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(filePath, stateBytes, fs.ModePerm)
			if err != nil {
				panic(err)
			}

			log.Println("--------- export state success ---------")
		},
	}
	cmd.Flags().StringP(FlagExportFile, "f", "", "")
	return cmd
}

func importStateCmd(ctx *server.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import-state",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("--------- import state start ---------")
			filePath := viper.GetString(FlagImportFile)
			if filePath == "" {
				panic("must set import-state-file")
			}

			stateBytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}
			var state sm.State
			err = sm.ModuleCodec.UnmarshalJSON(stateBytes, &state)
			if err != nil {
				panic(err)
			}

			dataDir := filepath.Join(ctx.Config.RootDir, "data")
			// load state
			stateDB, err := openDB(stateDB, dataDir)
			if err != nil {
				panic(err)
			}
			sm.SaveState(stateDB, state)

			log.Println("--------- import state success ---------")
		},
	}
	cmd.Flags().StringP(FlagImportFile, "i", "", "")
	return cmd
}
