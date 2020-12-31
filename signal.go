package main

import (
    "os"
    "os/signal"
    "syscall"
)

var (
    done = make(chan bool, 1)
)

func registerSignal() {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-sigs
        done <- true
    }()
}
