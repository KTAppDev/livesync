package getDirs

import (
	"fmt"
	"os"
	"syscall"
)

func GetDirectories(cwd string) {

	entries, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("Directory name:", entry.Name()) // Print the directory name

			info, err := entry.Info() // Get the FileInfo object
			if err != nil {
				fmt.Printf("Error reading directory: %v\n", err)
				return
			}

			fmt.Println("START DIR STATS")
			fmt.Println("Size:", info.Size())                     // Size in bytes
			fmt.Println("Permissions:", info.Mode())              // File permissions
			fmt.Println("Last Modified:", info.ModTime())         // Last modification time
			fmt.Println("Is Directory:", info.IsDir())            // Is it a directory?
			fmt.Printf("System specific info: %+v\n", info.Sys()) // OS-specific info; might be nil

			if stat, ok := info.Sys().(*syscall.Stat_t); ok {
				fmt.Println("Device:", stat.Dev)
				fmt.Println("Inode:", stat.Ino)
				fmt.Println("UID:", stat.Uid)
				fmt.Println("GID:", stat.Gid)
				// fmt.Println("Creation Date: ", unix_time_converter.ConvertUnixTimeToReadable(stat.Birthtimespec.Sec))
				// fmt.Println("Modified Date: ", unix_time_converter.ConvertUnixTimeToReadable(stat.Mtimespec.Sec))
				fmt.Println("END DIR STATS")
				fmt.Println("")
			} else {
				fmt.Println("System specific info could not be asserted to *syscall.Stat_t")
			}
		}
	}

}
