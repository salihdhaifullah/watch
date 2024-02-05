package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
)

var Cancels []context.CancelFunc
var mutex = sync.Mutex{}

func InitProcess(command string) {
	log.Printf("\n\n running command %s \n\n", command)
	ctx, cancel := context.WithCancel(context.Background())
	cancelOnTermination(cancel)
	cmd := exec.CommandContext(ctx, "bash", "-c", command)

	mutex.Lock()
	Cancels = append(Cancels, cancel)
	mutex.Unlock()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Printf("err was not nil: %v", err)
	}

	log.Println("Process terminated normally")
}


func KillAllProcesses() {
	for _, fn := range Cancels {
		fn()
	}
}

func cancelOnTermination(cancel context.CancelFunc) {
	log.Println("setting up a signal handler")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM)

	go func() {
		log.Printf("received SIGTERM %v\n", <-s)
		cancel()
	}()
}



