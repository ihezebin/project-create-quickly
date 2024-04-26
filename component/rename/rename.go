package rename

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var workDir string

func init() {
	workDir, _ = os.Getwd()
}

func Rename(dir string, oldName, newName string) error {
	oldNameSuffix := filepath.Base(oldName)
	newNameSuffix := filepath.Base(newName)

	aimDir := filepath.Join(workDir, dir)
	files, err := os.ReadDir(aimDir)
	if err != nil {
		return errors.Wrapf(err, "read dir err, dir: %s", aimDir)
	}
	for _, file := range files {
		path := filepath.Join(aimDir, file.Name())
		if file.IsDir() {
			if file.Name() == ".git" {
				continue
			}
			if err = Rename(filepath.Join(dir, file.Name()), oldName, newName); err != nil {
				return err
			}
		} else {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			str := strings.ReplaceAll(string(data), oldName, newName)
			str = strings.ReplaceAll(str, oldNameSuffix, newNameSuffix)
			err = os.WriteFile(path, []byte(str), os.ModePerm)
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(workDir, path)
			fmt.Println("[Success] ", rel)
		}
	}
	return nil
}
