package main

import (
    "flag"
    "time"
)

var (
    url = flag.String("url", "http://httpstat.us/200", "Target url")
    interval = flag.Duration("interval", 10 * time.Second, "Request interval (seconds)")
    timeout = flag.Int("timeout", 5, "Request timeout (seconds)")
    note = flag.String("note", "", "A short message appending the log")
)

func parseArgs() {
    flag.Parse()
}
