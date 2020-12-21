package main

import (
    "time"
)

func main() {
    model := NewModel()
    for now := range time.Tick(1 * time.Second) {
        Get(model)
        Update(now, model)
    }
}
