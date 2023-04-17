package build

import (
	"github.com/ihezebin/project-create-quickly/component/rename"
	"github.com/pkg/errors"
	"os/exec"
	"path/filepath"
	"strings"
)

type DDDBuilder struct {
	Origin string
}

const defaultDDDOrigin = "https://gitee.com/ihezebin/go-template-ddd.git"

func NewDDDBuilder(origin string) *DDDBuilder {
	if origin == "" {
		origin = defaultDDDOrigin
	}
	return &DDDBuilder{
		Origin: origin,
	}
}

func (b *DDDBuilder) Build(projectName string) error {
	// 如果项目名是mod名，则取最后一个路径名为项目名
	if strings.Contains(projectName, "/") {
		mnSplit := strings.Split(projectName, "/")
		projectName = mnSplit[len(mnSplit)-1]
	}

	if err := cloneGitProject(b.Origin, projectName); err != nil {
		return err
	}
	// 删除.git 等文件，保持文件目录结构整洁
	if err := exec.Command("rm", "-rf",
		filepath.Join(workDir, projectName, "ddd"),
		filepath.Join(workDir, projectName, ".git"),
		filepath.Join(workDir, projectName, "log"),
		filepath.Join(workDir, projectName, "go.sum"),
		filepath.Join(workDir, projectName, "cmd", "ddd"),
		filepath.Join(workDir, projectName, "script", "shell"),
	).Run(); err != nil {
		return errors.Wrap(err, "remove redundant files err")
	}
	return nil
}

func (b *DDDBuilder) Rename(newProjectName string) error {
	// 如果项目名是mod名，则取最后一个路径名为项目名
	const defaultOldName = "github.com/ihezebin/go-template-ddd"
	// 如果项目名是mod名，则取最后一个路径名为项目名
	modName := newProjectName
	if strings.Contains(newProjectName, "/") {
		mnSplit := strings.Split(newProjectName, "/")
		newProjectName = mnSplit[len(mnSplit)-1]
	}
	return rename.Rename(newProjectName, defaultOldName, modName)
}
