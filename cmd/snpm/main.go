package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/shynome/snpm/cmd/run"
)

var pkgScriptAlias = map[string]string{
	"start":   "normally start application",
	"stop":    "normally stop application",
	"restart": "normally restart application",
}

func addPkgScriptAliasCommands(commands []cli.Command) []cli.Command {
	for n, u := range pkgScriptAlias {
		cmd := cli.Command{
			Name:   n,
			Usage:  u,
			Action: run.Run(n),
		}
		commands = append(commands, cmd)
	}
	return commands
}

func main() {

	app := cli.NewApp()

	app.Name = "snpm"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	commands := []cli.Command{
		run.Command,
	}
	commands = addPkgScriptAliasCommands(commands)

	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
