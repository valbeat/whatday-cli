package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/valbeat/whatday-cli/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "random",
		Usage:  "",
		Action: command.CmdRandom,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "help",
		Usage:  "",
		Action: command.CmdHelp,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "version",
		Usage:  "",
		Action: command.CmdVersion,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "config",
		Usage:  "",
		Action: command.CmdConfig,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
