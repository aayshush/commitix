package utils

import (
	"os"
	"path"
)

// Directories to ignore
var IgnoreDirs = []string{".commitix"}

var (
	commitixPath   string
	WorkindDir string
)

func Init() error {
	if wd, err := os.Getwd(); err != nil {
		return err
	} else {
		WorkindDir = wd
	}
	commitixPath = path.Join(WorkindDir, ".commitix")
	return nil
}
