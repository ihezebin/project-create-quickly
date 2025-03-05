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
	WorkDir     string
	ProjectName string
	Renames     []RenameKv
}

type RenameKv struct {
	Old     string
	New     string
	JustDir bool
}

func NewBaseBuilder(workDir, projectName string, renames ...RenameKv) *BaseBuilder {
	return &BaseBuilder{
		WorkDir:     workDir,
		ProjectName: projectName,
		Renames:     renames,
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

	dirPaths := make([]string, 0)

	path := filepath.Join(b.WorkDir, b.ProjectName)
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			dirPaths = append(dirPaths, path)
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, "read file %s err", path)
		}
		for _, rename := range b.Renames {
			if rename.JustDir {
				continue
			}
			data = []byte(strings.ReplaceAll(string(data), rename.Old, rename.New))
		}
		err = os.WriteFile(path, data, os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "write file %s err", path)
		}
		rel, _ := filepath.Rel(b.WorkDir, path)
		fmt.Println("[Renamed File Content!] ", rel)

		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "walk %s err", path)
	}

	// rename dirs, 逆序遍历
	for i := len(dirPaths) - 1; i >= 0; i-- {
		dirPath := dirPaths[i]
		for _, rename := range b.Renames {
			// 从路径中提取文件名和 old 比较
			if filepath.Base(dirPath) == rename.Old {
				// 如果目录名是Old，则重命名，只重命名最后一部分
				// newDirPath := strings.ReplaceAll(dirPath, rename.Old, rename.New)
				newDirPath := filepath.Join(filepath.Dir(dirPath), rename.New)
				err = os.Rename(dirPath, newDirPath)
				if err != nil {
					return errors.Wrapf(err, "rename %s to %s err", dirPath, newDirPath)
				}
				oldRel, _ := filepath.Rel(b.WorkDir, dirPath)
				newRel, _ := filepath.Rel(b.WorkDir, newDirPath)
				fmt.Printf("[Renamed Dir Name!] %s -> %s\n", oldRel, newRel)
			}
		}
	}

	return nil
}
