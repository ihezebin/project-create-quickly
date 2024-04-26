package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ihezebin/project-create-quickly/component/build"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func Run() error {

	app := &cli.App{
		Name:        "pcq",
		Version:     "v1.0.2",
		Usage:       "A script to create and init template project quickly",
		UsageText:   "pcq <project name> [-t | --template=<value>] [--git] [-o | --origin=<value>]",
		Description: "This application relies on Git",
		Authors: []*cli.Author{
			{Name: "hezebin", Email: "ihezebin@qq.com"},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "template", Aliases: []string{"t"},
				Value: "",
				Usage: "point the template of project which you want to create, support: react、go",
			},
			&cli.StringFlag{
				Name: "origin", Aliases: []string{"o"},
				Value: "",
				Usage: "customize a git repository url",
			},
		},
		Commands: cli.Commands{
			versionCmd,
			renameCmd,
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return errors.New("Args num must be 1, use pcq -h to get help")
			}

			projectName := c.Args().Get(0)
			if projectName == "" {
				return errors.New("project name can not be empty")
			}

			template := c.String("template")
			if template == "" {
				return errors.New("must point out the project template")
			}

			origin := c.String("origin")
			var outProjectName string
			var builder build.Builder
			switch template {
			case "go":
				builder = build.NewDDDBuilder(origin)
				mnSplit := strings.Split(projectName, "/")
				outProjectName = mnSplit[len(mnSplit)-1]
				fmt.Printf("\nGolang project name: %s, Mod name: %s\n\n", outProjectName, projectName)
			case "react":
				builder = build.NewReactTsBuilder(origin)
				outProjectName = projectName
				fmt.Printf("\nReact project name: %s\n\n", projectName)
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
		},
	}

	return app.Run(os.Args)
}
