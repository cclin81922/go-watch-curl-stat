package main

import (
    "os"
    "time"
)

func Loop(model Model) {
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
    ParseArgs()
    RegisterSignal()
    model := NewModel()
    Loop(model)
    statusCode := WriteLog(model)

    os.Exit(statusCode)
}
