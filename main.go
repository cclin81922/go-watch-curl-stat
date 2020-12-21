package main

import (
    "time"
)

func main() {
    Parse()
    model := NewModel()
    for now := range time.Tick(time.Duration(*interval) * time.Second) {
        Get(model)
        Update(now, model)
    }
}
