package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/PlatONEnetwork/PlatONE-Go/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	app = utils.NewApp("", "PlatONE interaction command line interface")
)

func init() {

	// Initialize the CLI app
	app.Commands = []cli.Command{
		AccountCmd,  // see cmd_account.go
		ContractCmd, // see cmd_contract.go
		// AdminCmd,    // see cmd_admin.go, deprecated
		CnsCmd,       // see cmd_cns.go
		FwCmd,        // see cmd_firewall.go
		RoleCmd,      // see cmd_role.go
		NodeCmd,      // see cmd_node.go
		SysConfigCmd, // see cmd_sysconfig.go
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.After = func(ctx *cli.Context) error {
		return nil
	}

}

//go:generate go-bindata -pkg precompile -o precompiled/bindata.go ../../release/linux/conf/contracts/...
func main() {
	// Initialize the related file
	configInit()
	/// abiInit()
	/// utl.LogInit()

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}