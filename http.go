package main

import (
    "log"
    "net/http"
    "time"
)

var (
    client = http.Client{
        Timeout: time.Duration(*timeout) * time.Second,
    }
)

func Get(model Model) {
    res, err := client.Get(*url)

    if err != nil {
        log.Fatal(err)
    }

    model.Inc(res.StatusCode)
}
