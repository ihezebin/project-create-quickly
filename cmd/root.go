package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/ihezebin/project-create-quickly/internal/builder"
	"github.com/ihezebin/project-create-quickly/internal/constant"
)

var (
	template   string
	repository string
	workDir    string
)

func init() {
	workDir, _ = os.Getwd()
}

func Run() error {

	app := &cli.App{
		Name:        "pcq",
		Version:     "v1.0.6",
		Usage:       "A script to create and init template project quickly",
		UsageText:   "pcq [-t | --template=<value>] [-r | --repository=<value>] <project name>",
		Description: "This application relies on Git",
		Authors: []*cli.Author{
			{Name: "hezebin", Email: "ihezebin@qq.com"},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "template", Aliases: []string{"t"},
				Value:       "",
				Usage:       fmt.Sprintf("point the template of project which you want to create, support: %s. java-ddd can be 'com.domain.project' or 'project', go-ddd can be 'github.com/user/project' or 'project'", strings.Join(constant.SupportTemplates, "、")),
				Required:    true,
				Destination: &template,
			},
			&cli.StringFlag{
				Name: "repository", Aliases: []string{"r"},
				Value:       "",
				Usage:       "customize a git repository url",
				Destination: &repository,
			},
		},
		Commands: cli.Commands{
			versionCmd,
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return errors.New("Args num must be 1, use pcq -h to get help")
			}

			projectName := c.Args().Get(0)
			if projectName == "" {
				return errors.New("project name can not be empty")
			}

			if template == "" {
				return errors.Errorf("must point out the project template, support: %s", strings.Join(constant.SupportTemplates, "、"))
			}

			var err error
			if repository == "" {
				repository, err = constant.GetDefaultRepository(template)
				if err != nil {
					return errors.Wrap(err, "get default repository error")
				}
			}

			// create builder to handle especial template
			var b builder.Builder
			switch template {
			case constant.TemplateGoDDD:
				modName := projectName
				projectName = path.Base(projectName)
				// project 只能包含字母、数字、下划线、中划线，且不能以数字开头
				if matched, err := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_-]+$", projectName); !matched || err != nil {
					return errors.New("project name can only contain letters, numbers, underscores, and hyphens, and must not start with a number")
				}
				fmt.Printf("\nGolang DDD project name: %s, Mod name: %s\n\n", projectName, modName)
				b = builder.NewGoDDDBuilder(workDir, projectName, modName)
			case constant.TemplateCraTs:
				// project 只能包含字母、数字、下划线、中划线，且不能以数字开头
				if matched, err := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_-]+$", projectName); !matched || err != nil {
					return errors.New("project name can only contain letters, numbers, underscores, and hyphens, and must not start with a number")
				}
				fmt.Printf("\nReact Cra TS project name: %s\n\n", projectName)
				b = builder.NewBaseBuilder(workDir, projectName, builder.RenameKv{
					Old: "react-template-ts", New: projectName,
				})
			case constant.TemplateVite:
				// project 只能包含字母、数字、下划线、中划线，且不能以数字开头
				if matched, err := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_-]+$", projectName); !matched || err != nil {
					return errors.New("project name can only contain letters, numbers, underscores, and hyphens, and must not start with a number")
				}
				fmt.Printf("\nReact Vite TS project name: %s\n\n", projectName)
				b = builder.NewBaseBuilder(workDir, projectName, builder.RenameKv{
					Old: "react-template-vite", New: projectName,
				})
			case constant.TemplateJavaDDD:
				groupName := projectName
				groupNames := strings.Split(groupName, ".")
				if len(groupNames) != 3 && len(groupNames) > 1 {
					return errors.New("group name must be like 'com.example.project'")
				}

				if len(groupNames) == 3 {
					projectName = groupNames[2]
				}
				// project 只能包含字母、数字、下划线，不包含中划线，且不能以数字开头，数字和划线不是必须的
				if matched, err := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]+$", projectName); !matched || err != nil {
					return errors.New("project name can only contain letters, numbers, underscores, and must not start with a number")
				}

				fmt.Printf("\nJava DDD project name: %s\n\n", projectName)

				// 新项目名不能是 test
				if strings.Contains(projectName, "test") {
					return errors.New("project name can not be 'test'")
				}

				kvs := []builder.RenameKv{
					{
						Old: "java-template-ddd", New: projectName,
					},
					{
						Old: "template", New: projectName,
					},
				}

				if len(groupNames) == 3 {
					comNew := groupNames[0]
					hezebinNew := groupNames[1]
					comHezebinNew := groupNames[0] + "." + groupNames[1]
					kvs = append(kvs, builder.RenameKv{
						Old: "com.hezebin", New: comHezebinNew,
					})
					kvs = append(kvs, builder.RenameKv{
						Old: "com", New: comNew, JustDir: true,
					})
					kvs = append(kvs, builder.RenameKv{
						Old: "hezebin", New: hezebinNew, JustDir: true,
					})
				}

				b = builder.NewBaseBuilder(workDir, projectName, kvs...)
			}

			// 判断目录是否存在, 已存在为了防止覆盖原目录，直接报错
			projectDir := filepath.Join(workDir, projectName)
			if _, err = os.Stat(projectDir); !os.IsNotExist(err) {
				return fmt.Errorf("[%s] already exists under the current directory", projectDir)
			}

			buildChan := make(chan struct{})
			var buildErr error
			go func() {
				defer func() {
					buildChan <- struct{}{}
				}()
				// clone repository
				_, err = git.PlainClone(projectDir, false, &git.CloneOptions{
					URL:      repository,
					Progress: os.Stdout,
				})
				if err != nil {
					buildErr = errors.Wrapf(err, "git clone from repository %s error", repository)
					return
				}

				// handle build
				if b != nil {
					if err = b.Build(); err != nil {
						buildErr = errors.Wrapf(err, "build project error")
					}
				}
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

			fmt.Println("\nInit project success!")
			fmt.Printf("\nNow: cd %s\n\n", projectName)
			return nil
		},
	}

	return app.Run(os.Args)
}
