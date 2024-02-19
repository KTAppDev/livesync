package getFiles

import (
	// "fmt"
	"io/fs"
	"log"
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
			log.Println(err)
			return err
		}
		if !d.IsDir() {
			relPath, err := filepath.Rel(cwd, path)
			if err != nil {
				log.Println(err)
				return err
			}

			info, err := d.Info()
			if err != nil {
				log.Println(err)
				return err
			}

			fileInfo := models.FileInfo{}
			fileInfo.NewFileInfoWithDefaults()
			fileInfo.Name = strings.SplitAfter(info.Name(), ".")[0][:len(strings.SplitAfter(info.Name(), ".")[0])-1]
			fileInfo.Path = cwd + relPath
			fileInfo.Size = uint32(info.Size())
			fileInfo.Permissions = info.Mode()
			fileInfo.ModifiedAt = info.ModTime()

			if stat, ok := info.Sys().(*syscall.Stat_t); ok {
				creationTime := time.Unix(stat.Birthtimespec.Sec, 0)
				fileInfo.CreatedAt = creationTime
			} else {
				// fmt.Println("System specific info could not be asserted to *syscall.Stat_t")
			}

			filesInfo = append(filesInfo, fileInfo)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return filesInfo, nil
}
