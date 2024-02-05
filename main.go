package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating watcher:", err)
	}

	settings, err := LoadSetting()
	if err != nil {
		log.Fatal(err)
	}

	err = HandelPaths(settings, watcher)
	if err != nil {
		fmt.Println("Error adding directory to watcher:", err)
		return
	}

	Start(settings, watcher)
}
