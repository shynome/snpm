package run

import (
	"github.com/shynome/snpm"
	"github.com/shynome/snpm/cmd/common"
	"github.com/urfave/cli"
)

// Command exports the run command
var Command = cli.Command{
	Name:  "run",
	Usage: "run package script",
}

func init() {
	subCommands := []cli.Command{}
	for n, c := range common.Pkg.Scripts {
		cmd := cli.Command{
			Name:   n,
			Usage:  c,
			Action: Run(n),
		}
		subCommands = append(subCommands, cmd)
	}
	Command.Subcommands = subCommands
}

// Run npm script
func Run(stage string) func(*cli.Context) {
	return func(c *cli.Context) {
		snpm.Exec(stage, c.Args(), common.Pkg)
	}
}
