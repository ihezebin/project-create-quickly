package build

import (
	"github.com/ihezebin/project-create-quickly/component/rename"
	"github.com/pkg/errors"
	"os/exec"
	"path/filepath"
)

type ReactTsBuilder struct {
	Origin string
}

const defaultReactOrigin = "https://gitee.com/ihezebin/react-template-ts.git"

func NewReactTsBuilder(origin string) *ReactTsBuilder {
	if origin == "" {
		origin = defaultReactOrigin
	}

	return &ReactTsBuilder{
		Origin: origin,
	}
}

func (b *ReactTsBuilder) Build(projectName string) error {
	if err := cloneGitProject(b.Origin, projectName); err != nil {
		return err
	}
	// 删除.git 等文件，保持文件目录结构整洁
	if err := exec.Command("rm", "-rf",
		filepath.Join(workDir, projectName, ".git"),
	).Run(); err != nil {
		return errors.Wrap(err, "remove redundant files err")
	}
	return nil
}

func (b *ReactTsBuilder) Rename(newProjectName string) error {
	const defaultOldName = "react-template-ts"
	return rename.Rename(newProjectName, defaultOldName, newProjectName)
}
