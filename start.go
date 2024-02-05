package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

var ignoreEvent = false

func waitFor(fn func (), wg *sync.WaitGroup) {
	fn()
	wg.Done()
}

func startCommands(settings *Settings) {
	ignoreEvent = true
	for i := 0; i < len(settings.SetupCommands); i++ {
		InitProcess(settings.SetupCommands[i])
	}

	wg := sync.WaitGroup{}
	for i := 0; i < len(settings.ParallelCommands); i++ {
		wg.Add(1)
		cmd := settings.ParallelCommands[i]
		toRun := func() {
			InitProcess(cmd)
		}
		go waitFor(toRun, &wg)
	}
	wg.Wait()

	for i := 0; i < len(settings.BackgroundCommands); i++ {
		go InitProcess(settings.BackgroundCommands[i])
	}
	ignoreEvent = false
}

func Start(settings *Settings, watcher *fsnotify.Watcher) {
	startCommands(settings)

	fmt.Println("Watching for file changes...")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if ignoreEvent {
				continue
			}

			ignoreEvent = true

			go func() {
				time.Sleep(time.Duration(settings.Delay) * time.Millisecond)
				ignoreEvent = false
			}()

			fmt.Println("Event:", event)

			if event.Has(fsnotify.Write) {
				fmt.Println("Modified file:", event.Name)
				go handelEvent(settings)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error:", err)

		}
	}
}


func handelEvent(settings *Settings) {
	KillAllProcesses()
	fmt.Println("Commands killed. Restarting...")
	startCommands(settings)
}
