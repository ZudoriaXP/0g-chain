package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"

	tmcfg "github.com/cometbft/cometbft/config"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	ethermintclient "github.com/evmos/ethermint/client"
	"github.com/evmos/ethermint/crypto/hd"
	ethermintserver "github.com/evmos/ethermint/server"
	servercfg "github.com/evmos/ethermint/server/config"
	"github.com/spf13/cobra"

	"github.com/0glabs/0g-chain/app"
	"github.com/0glabs/0g-chain/app/params"
	"github.com/0glabs/0g-chain/chaincfg"
	"github.com/0glabs/0g-chain/cmd/opendb"
	"github.com/0glabs/0g-chain/crypto/vrf"
)

func customKeyringOptions() keyring.Option {
	return func(options *keyring.Options) {
		options.SupportedAlgos = append(hd.SupportedAlgorithms, vrf.VrfAlgo)
		options.SupportedAlgosLedger = append(hd.SupportedAlgorithmsLedger, vrf.VrfAlgo)
	}
}

// NewRootCmd creates a new root command for the 0g-chain blockchain.
func NewRootCmd() *cobra.Command {
	encodingConfig := app.MakeEncodingConfig()
	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
<<<<<<< HEAD
		WithBroadcastMode(flags.FlagBroadcastMode).
		WithHomeDir(app.DefaultNodeHome).
		WithKeyringOptions(hd.EthSecp256k1Option()).
		WithViper(EnvPrefix)
=======
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(chaincfg.DefaultNodeHome).
		WithKeyringOptions(customKeyringOptions()).
		WithViper(chaincfg.EnvPrefix)
>>>>>>> be1cd76f (add vrf)

	rootCmd := &cobra.Command{
		Use:   chaincfg.AppName,
		Short: "Daemon and CLI for the 0g-chain blockchain.",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err = client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			customAppTemplate, customAppConfig := servercfg.AppConfig(chaincfg.BaseDenom)

			return server.InterceptConfigsPreRunHandler(
				cmd,
				customAppTemplate,
				customAppConfig,
				tmcfg.DefaultConfig(),
			)
		},
	}

	addSubCmds(rootCmd, encodingConfig, chaincfg.DefaultNodeHome)

	return rootCmd
}

// addSubCmds registers all the sub commands used by 0g-chain.
func addSubCmds(rootCmd *cobra.Command, encodingConfig params.EncodingConfig, defaultNodeHome string) {
	gentxModule, ok := app.ModuleBasics[genutiltypes.ModuleName].(genutil.AppModuleBasic)
	if !ok {
		panic(fmt.Errorf("expected %s module to be an instance of type %T", genutiltypes.ModuleName, genutil.AppModuleBasic{}))
	}

	rootCmd.AddCommand(
		StatusCommand(),
		ethermintclient.ValidateChainID(
			genutilcli.InitCmd(app.ModuleBasics, defaultNodeHome),
		),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, defaultNodeHome, gentxModule.GenTxValidator),
		AssertInvariantsCmd(encodingConfig),
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, defaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(defaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true), // TODO add other shells, drop tmcli dependency, unhide?
		// testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}), // TODO add
		debug.Cmd(),
		config.Cmd(),
	)

	ac := appCreator{
		encodingConfig: encodingConfig,
	}

	opts := ethermintserver.StartOptions{
		AppCreator:      ac.newApp,
		DefaultNodeHome: chaincfg.DefaultNodeHome,
		DBOpener:        opendb.OpenDB,
	}
	// ethermintserver adds additional flags to start the JSON-RPC server for evm support
	ethermintserver.AddCommands(
		rootCmd,
		opts,
		ac.appExport,
		ac.addStartCmdFlags,
	)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		newQueryCmd(),
		newTxCmd(),
		keyCommands(app.DefaultNodeHome),
		rocksdb.RocksDBCmd,
		newShardCmd(opts),
	)
}
