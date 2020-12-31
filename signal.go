package main

import (
    "os"
    "os/signal"
    "syscall"
)

func registerSignal() <-chan bool{
    done := make(chan bool, 1)

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-sigs
        done <- true
    }()

    return done
}
