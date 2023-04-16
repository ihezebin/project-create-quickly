package rename

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var workDir string

func init() {
	workDir, _ = os.Getwd()
}

func Rename(dir string, oldName, newName string) error {
	aimDir := filepath.Join(workDir, dir)
	files, err := ioutil.ReadDir(aimDir)
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
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			str := strings.ReplaceAll(string(data), oldName, newName)
			err = ioutil.WriteFile(path, []byte(str), os.ModePerm)
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(workDir, path)
			fmt.Println("[Success] ", rel)
		}
	}
	return nil
}
