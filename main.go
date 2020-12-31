package main

import (
    "os"
    "time"
)

func loop(model Model) {
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
    registerSignal()
    model := newModel()
    loop(model)
    statusCode := writeLog(model)

    os.Exit(statusCode)
}
