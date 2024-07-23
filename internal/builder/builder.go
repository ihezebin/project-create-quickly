package builder

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type Builder interface {
	Build() error
}

type BaseBuilder struct {
	WorkDir             string
	ProjectName         string
	ReplaceContentTable map[string]string
}

func NewBaseBuilder(workDir, projectName string, renameTable map[string]string) *BaseBuilder {
	return &BaseBuilder{
		WorkDir:             workDir,
		ProjectName:         projectName,
		ReplaceContentTable: renameTable,
	}
}

func (b *BaseBuilder) Build() error {
	// 删除.git 等文件，保持文件目录结构整洁
	if err := exec.Command("rm", "-rf",
		filepath.Join(b.WorkDir, b.ProjectName, ".git"),
	).Run(); err != nil {
		return errors.Wrap(err, "remove redundant files err")
	}

	fmt.Println("Renaming project files...     ")

	path := filepath.Join(b.WorkDir, b.ProjectName)
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, "read file %s err", path)
		}
		for oldContent, newContent := range b.ReplaceContentTable {
			data = []byte(strings.ReplaceAll(string(data), oldContent, newContent))
		}
		err = os.WriteFile(path, data, os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "write file %s err", path)
		}
		rel, _ := filepath.Rel(b.WorkDir, path)
		fmt.Println("[Renamed!] ", rel)

		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "walk %s err", path)
	}

	return nil
}
