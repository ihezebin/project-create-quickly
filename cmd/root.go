package cmd

import (
	"fmt"
	"github.com/ihezebin/project-create-quickly/component/build"
	"github.com/ihezebin/sdk/cli"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func Run() error {
	return app.Run()
}

var app = cli.NewApp(
	cli.WithVersion("v1.0"),
	cli.WithName("pcq"),
	cli.WithAuthor("hezebin"),
	cli.WithUsage("A script to create and init template project quickly"),
	cli.WithDescription("This application relies on Git"),
	cli.WithUsageText("pcq <project name> [-t | --template=<value>] [--git] [-o | --origin=<value>]"),
).
	WithFlagString("template, t", "", "point the template of project which you want to create", false).
	WithFlagString("origin, o", "", "Customize a git repository url", false).
	WithAction(func(v cli.Value) error {
		if v.NArg() != 1 {
			return errors.New("Args num must be 1, use pcq -h to get help")
		}

		projectName := v.Args().Get(0)
		if projectName == "" {
			return errors.New("project name can not be empty")
		}

		template := v.String("template")
		if template == "" {
			return errors.New("must point out the project template")
		}

		origin := v.String("origin")
		var outProjectName string
		var builder build.Builder
		switch template {
		case "go":
			builder = build.NewDDDBuilder(origin)
			mnSplit := strings.Split(projectName, "/")
			outProjectName = mnSplit[len(mnSplit)-1]
			fmt.Printf("\nProject name: %s, Mod name: %s\n\n", outProjectName, projectName)
		case "react":
			builder = build.NewReactTsBuilder(origin)
			fmt.Printf("\nProject name: %s\n\n", projectName)
		default:
			return errors.Errorf("template type %s is not supported", template)
		}

		buildChan := make(chan struct{})
		var buildErr error
		go func() {
			if err := builder.Build(projectName); err != nil {
				buildErr = err
			}
			buildChan <- struct{}{}
		}()

		spin := `-\|/`
		for i, loading := 0, true; loading; i++ {
			select {
			case <-buildChan:
				loading = false
				if buildErr != nil {
					fmt.Println("Generating project Failed!         ")
					return errors.Wrap(buildErr, "Build project failed")
				}
				fmt.Println("Generating project Success!         ")
			default:
				fmt.Printf("Generating project file...  %s\r", string(spin[i%len(spin)]))
				time.Sleep(100 * time.Millisecond)
			}
		}

		fmt.Println("\nOrganizing project files...")
		if err := builder.Rename(projectName); err != nil {
			return errors.Wrap(err, "Rename project failed")
		}

		fmt.Println("\nInit project success!")
		fmt.Printf("\nNow: cd %s\n\n", outProjectName)
		return nil
	}).
	AddCommand(versionCmd).
	AddCommand(renameCmd)
