package cmd

import (
	"fmt"
	"github.com/ihezebin/project-create-quickly/component/rename"
	"github.com/ihezebin/sdk/cli/command"
	"github.com/pkg/errors"
)

var renameCmd = command.NewCommand(
	command.WithName("rename"),
	command.WithUsageText("pcq rename [old name] [new name]"),
).WithAction(func(v command.Value) error {
	var oldProjectName string
	var newProjectName string
	switch v.NArg() {
	case 0, 1:
		return errors.New("[old name] and [new name] are must, use pcq rename --help to get help")
	case 2:
		oldProjectName = v.Args().Get(0)
		newProjectName = v.Args().Get(1)
	default:
		return errors.New("Args num can not be greater than 2, use pcq rename --help to get help")
	}

	if oldProjectName == "" || newProjectName == "" {
		return errors.New("old or new name can not be empty")
	}

	fmt.Println("\nStart rename...\n")

	if err := rename.Rename("./", oldProjectName, newProjectName); err != nil {
		return errors.Wrap(err, "rename failed")
	}

	fmt.Println("\nRename succeed!\n")

	return nil
})
