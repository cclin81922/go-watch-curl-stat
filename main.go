package main

import (
    "time"
)

func main() {
    Parse()
    model := NewModel()
    for now := range time.Tick(*interval) {
        Get(model)
        Update(now, model)
    }
}
