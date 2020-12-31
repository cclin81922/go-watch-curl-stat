package main

import (
    "os"
    "fmt"
)

func writeLog(model Model) int {
    log := fmt.Sprintf("%s%s\n", model.String(), *note)

    f, e := os.OpenFile("/tmp/go-watch-curl-stat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if e != nil {
	    fmt.Println(e)
        return 1
    }

    defer f.Close()

    if _, e := f.WriteString(log); e != nil {
	    fmt.Println(e)
        return 1
    }

    fmt.Println("Generated log file at /tmp/go-watch-curl-stat.log")
    return 0
}
