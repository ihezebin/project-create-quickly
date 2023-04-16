package cmd

import (
	"fmt"
	"github.com/ihezebin/sdk/cli/command"
)

var versionCmd = command.NewCommand(
	command.WithName("version"),
	command.WithUsage("Just print the version."),
).WithAction(func(v command.Value) error {
	version := v.Kernel().App.Version
	fmt.Println("qcp version", version)
	return nil
})
