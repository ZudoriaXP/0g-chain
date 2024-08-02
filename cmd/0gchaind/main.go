package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/0glabs/0g-chain/chaincfg"
)

func main() {
	fmt.Println("main")
	chaincfg.SetSDKConfig().Seal()
	fmt.Println("main2")
	rootCmd := NewRootCmd()
	if err := svrcmd.Execute(rootCmd, chaincfg.EnvPrefix, chaincfg.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			fmt.Println("error")
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
	fmt.Println("main3")
}
