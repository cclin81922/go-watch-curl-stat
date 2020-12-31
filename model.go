package main

import (
    "fmt"
    "strings"
)

type Model map[int]int

func (m Model) Inc(key int) {
    m[key]++
}

func (m Model) String() string {
    var b strings.Builder

    for k,v := range m {
        fmt.Fprintf(&b, "%v:%v ", k,v)
    }

    return b.String()
}

func newModel() Model {
    return Model{}
}
