package commitix

import (
	"fmt"
	"os"
	"path"
	"commitix/utils"
)

// TODO: Change Permission Mode for the directories and files
func InitHandler(args []string) error {

	if err := os.MkdirAll(utils.commitixPath, os.ModePerm); err != nil {
		return fmt.Errorf("error in creating .commitix: %v", err)
	}
	subdirs := []string{"objects", "refs"}
	for i := range subdirs {
		if err := os.MkdirAll(path.Join(utils.commitixPath, subdirs[i]), os.ModePerm); err != nil {
			return fmt.Errorf("error in creating %s; %v", subdirs, err)
		}
	}

	headPath := path.Join(utils.commitixPath, "HEAD")
	if err := os.WriteFile(headPath, []byte{}, 0664); err != nil {
		return fmt.Errorf("error in creating HEAD: %v", err)
	}

	fmt.Printf("Initialized empty commitix repository in %s\n", utils.commitixPath)
	return nil
}
