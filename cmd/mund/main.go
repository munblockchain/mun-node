package main

import (
	"os"

	"github.com/munblockchain/mun-node/app"

	"github.com/munblockchain/mun-node/cmd/mund/cmd"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	// cmdOptions := GetWasmCmdOptions()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
