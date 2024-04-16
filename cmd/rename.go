package cmd

import "C"
import (
	"fmt"

	"github.com/ihezebin/project-create-quickly/component/rename"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var renameCmd = &cli.Command{
	Name:      "rename",
	Aliases:   []string{"r"},
	Usage:     "rename project name",
	UsageText: "pcq rename [old name] [new name]",
	Action: func(c *cli.Context) error {
		var oldProjectName string
		var newProjectName string
		switch c.NArg() {
		case 0, 1:
			return errors.New("[old name] and [new name] are must, use pcq rename --help to get help")
		case 2:
			oldProjectName = c.Args().Get(0)
			newProjectName = c.Args().Get(1)
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
	},
}
