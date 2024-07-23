package builder

import (
	"os/exec"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

type GoDDDBuilder struct {
	*BaseBuilder
	ModName string
}

func NewGoDDDBuilder(workDir, projectName, modName string) *GoDDDBuilder {
	const templateModName = "github.com/ihezebin/go-template-ddd"

	return &GoDDDBuilder{
		ModName: modName,
		BaseBuilder: NewBaseBuilder(workDir, projectName,
			RenameKv{
				Old: templateModName, New: modName,
			}, RenameKv{
				Old: path.Base(templateModName), New: projectName,
			}),
	}
}

func (b *GoDDDBuilder) Build() error {
	// 删除.git 等文件，保持文件目录结构整洁
	if err := exec.Command("rm", "-rf",
		filepath.Join(b.WorkDir, b.ProjectName, "ddd"),
		filepath.Join(b.WorkDir, b.ProjectName, "log"),
	).Run(); err != nil {
		return errors.Wrap(err, "remove redundant files err")
	}

	return b.BaseBuilder.Build()
}
