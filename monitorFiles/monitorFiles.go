package monitorFiles

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func MonitorFiles(directory string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// log.Println("event:", event)
				// if event.Op&fsnotify.Write == fsnotify.Write {
				// 	log.Println("modified file:", event.Name)
				// }
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("file added:", event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("file removed:", event.Name)
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("file removed:", event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(directory)
	if err != nil {
		log.Fatal(err)
	}
	<-done

}
