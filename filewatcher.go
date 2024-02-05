package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func HandelPaths(settings *Settings, watcher *fsnotify.Watcher) error {
	return filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if shouldIgnorePath(path, settings) {
			return nil
		}

		log.Printf("add path %s", path)
		return watcher.Add(path)
	})
}

func shouldIgnorePath(path string, settings *Settings) bool {
	for i := 0; i < len(settings.Ignore); i++ {
		if strings.HasPrefix(path, settings.Ignore[i]) {
			return true
		}
	}

	for i := 0; i < len(settings.IgnoreRegex); i++ {
		rp := regexp.MustCompile(settings.IgnoreRegex[i])
		if rp.MatchString(path) {
			return true
		}
	}

	return false
}
