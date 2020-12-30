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
        model.Inc(599)
        log.Println(err)
    } else {
        model.Inc(res.StatusCode)
    }
}
