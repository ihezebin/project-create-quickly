package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Just print the version.",
	Action: func(v *cli.Context) error {
		version := v.App.Version
		fmt.Println("pcq version", version)
		return nil
	},
}
