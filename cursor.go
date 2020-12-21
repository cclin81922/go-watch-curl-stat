package main

import (
    "fmt"
    "github.com/cclin81922/go-3rd-party/go-cursor"
)

func Update(message... interface{}) {
    fmt.Print(cursor.ClearEntireScreen())
    fmt.Print(cursor.MoveTo(0, 0))
    fmt.Println(message...)
}
