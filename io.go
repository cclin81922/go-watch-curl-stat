package main

import (
    "os"
    "fmt"
)

func WriteLog(model Model) int {
    f, e := os.OpenFile("/tmp/go-watch-curl-stat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if e != nil {
	    fmt.Println(e)
        return 1
    }

    defer f.Close()

    if _, e := f.WriteString(model.String() + "\n"); e != nil {
	    fmt.Println(e)
        return 1
    }

    fmt.Println("Generated log file at /tmp/go-watch-curl-stat.log")
    return 0
}
