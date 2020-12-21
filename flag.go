package main

import (
    "flag"
)

var (
    url = flag.String("url", "http://httpstat.us/200", "Target url")
    interval = flag.Int("interval", 10, "Request interval (seconds)")
    timeout = flag.Int("timeout", 5, "Request timeout (seconds)")
)

func Parse() {
    flag.Parse()
}
