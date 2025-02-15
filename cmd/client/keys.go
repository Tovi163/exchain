package client

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/okex/exchain/libs/tendermint/p2p"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/okex/exchain/libs/cosmos-sdk/client/flags"
	clientkeys "github.com/okex/exchain/libs/cosmos-sdk/client/keys"
	"github.com/okex/exchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"

	"github.com/okex/exchain/app/crypto/hd"
)

const (
	flagDryRun = "dry-run"
)

// KeyCommands registers a sub-tree of commands to interact with
// local private key storage.
func KeyCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Add or view local private keys",
		Long: `Keys allows you to manage your local keystore for tendermint.

    These keys may be in any format supported by go-crypto and can be
    used by light-clients, full nodes, or any other application that
    needs to sign with a private key.`,
	}

	// support adding Ethereum supported keys
	addCmd := clientkeys.AddKeyCommand()

	// update the default signing algorithm value to "eth_secp256k1"
	algoFlag := addCmd.Flag("algo")
	algoFlag.DefValue = string(hd.EthSecp256k1)
	err := algoFlag.Value.Set(string(hd.EthSecp256k1))
	if err != nil {
		panic(err)
	}
	addCmd.RunE = runAddCmd

	cmd.AddCommand(
		clientkeys.MnemonicKeyCommand(),
		addCmd,
		clientkeys.ExportKeyCommand(),
		clientkeys.ImportKeyCommand(),
		clientkeys.ListKeysCmd(),
		clientkeys.ShowKeysCmd(),
		flags.LineBreak,
		clientkeys.DeleteKeyCommand(),
		clientkeys.ParseKeyStringCommand(),
		clientkeys.MigrateCommand(),
		flags.LineBreak,
		UnsafeExportEthKeyCommand(),
		ExportEthCompCommand(),
		extractNodeKey(),
	)
	return cmd
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())
	kb, err := getKeybase(viper.GetBool(flagDryRun), inBuf)
	if err != nil {
		return err
	}

	return clientkeys.RunAddCmd(cmd, args, kb, inBuf)
}

func getKeybase(transient bool, buf io.Reader) (keys.Keybase, error) {
	if transient {
		return keys.NewInMemory(
			hd.EthSecp256k1Options()...,
		), nil
	}

	return keys.NewKeyring(
		sdk.KeyringServiceName(),
		viper.GetString(flags.FlagKeyringBackend),
		viper.GetString(flags.FlagHome),
		buf,
		hd.EthSecp256k1Options()...,
	)
}

func extractNodeKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extract-node-key [filename] ",
		Short: "extract current node key or from specificed file",
		RunE: func(cmd *cobra.Command, args []string) error {
			var filename string
			if len(args) >= 1 {
				filename = args[0]
			}
			nodekey, err := p2p.LoadNodeKey(filename)
			if err != nil {
				return err
			}

			//fmt.Printf("base64: %s\n", base64.StdEncoding.EncodeToString(nodekey.PubKey().Bytes()))
			fmt.Printf("hex: %s\n", hex.EncodeToString(nodekey.PubKey().Bytes()))

			return nil
		},
	}
	return cmd
}
