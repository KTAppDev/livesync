package getFiles

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/ktappdev/filesync/models"
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
			fileInfo := models.FileInfo{}
			fileInfo.NewFileInfoWithDefaults()
			fileInfo.Name = strings.SplitAfter(info.Name(), ".")[0][:len(strings.SplitAfter(info.Name(), ".")[0])-1] //This just removes the .extension part
			fileInfo.Path = path
			fileInfo.Size = info.Size()
			fileInfo.Permissions = info.Mode()
			fileInfo.Modified = info.ModTime()
			if stat, ok := info.Sys().(*syscall.Stat_t); ok {
				t := time.Unix(stat.Birthtimespec.Sec, 0)
				fileInfo.CreatedAt = t
			} else {
				fmt.Println("System specific info could not be asserted to *syscall.Stat_t")
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
