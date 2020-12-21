package main

import (
    "log"
    "net/http"
    "time"
)

var (
    client = http.Client{
        Timeout: 5 * time.Second,
    }
)

func Get(model Model) {
    res, err := client.Get("http://httpstat.us/200")

    if err != nil {
        log.Fatal(err)
    }

    model.Inc(res.StatusCode)
}
