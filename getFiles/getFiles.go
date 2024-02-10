package getFiles

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"
)

func GetFiles(cwd string) {
	err := filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			fmt.Printf("File: %s, Size: %d, Permissions: %s, Modified: %s\n",
				path, info.Size(), info.Mode(), info.ModTime().Format(time.RFC3339))
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
	}
}
