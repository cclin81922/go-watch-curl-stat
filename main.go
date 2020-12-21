package main

import (
    "time"
)

func main() {
    for now := range time.Tick(5 * time.Second) {
        Update(now)
    }
}
