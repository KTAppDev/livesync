package getFiles

import (
	"github.com/ktappdev/filesync/models" // Assuming the FileInfo struct and related logic are in the model package
	"io/fs"
	"path/filepath"
)

func GetFiles(cwd string) ([]models.FileInfo, error) {
	var filesInfo []models.FileInfo

	err := filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			fileInfo := models.FileInfo{
				Path:        path,
				Size:        info.Size(),
				Permissions: info.Mode(),
				Modified:    info.ModTime(),
			}
			filesInfo = append(filesInfo, fileInfo)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filesInfo, nil
}
