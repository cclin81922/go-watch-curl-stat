package main

import (
    "time"
)

func main() {
    model := NewModel()
    model.Inc(200)
    model.Inc(200)
    model.Inc(500)
    for now := range time.Tick(1 * time.Second) {
        Update(now, model)
    }
}
