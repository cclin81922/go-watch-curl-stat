package main

import (
    "os"
    "time"
)

func loop(model Model, done <-chan bool) {
    for {
        select {
        case <-done:
            return
        case now := <-time.Tick(*interval):
            HttpGet(model)
            UpdateScreen(now, model)
        }
    }
}

func main() {
    parseArgs()
    done := registerSignal()
    model := newModel()
    loop(model, done)
    statusCode := writeLog(model)

    os.Exit(statusCode)
}
