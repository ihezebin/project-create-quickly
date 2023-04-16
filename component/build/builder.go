package build

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Builder interface {
	Build(string) error
	Rename(string) error
}

var workDir string

func init() {
	workDir, _ = os.Getwd()
}

func cloneGitProject(origin string, projectName string) error {
	// 判断目录是否存在, 已存在为了防止覆盖原目录，直接报错
	dirPath := filepath.Join(workDir, projectName)
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		return fmt.Errorf("[%s] already exists under the current directory", dirPath)
	}
	_, err := exec.Command("git", "clone", origin, dirPath).CombinedOutput()
	return err
}
