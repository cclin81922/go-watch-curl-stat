package main

import (
    "fmt"
)

type Model map[int]int

func (m Model) Inc(key int) {
    m[key]++
}

func (m Model) String() string {
    var s = ""

    for k,v := range m {
        s += fmt.Sprintf("%v:%v ", k,v)
    }

    return s
}

func newModel() Model {
    return Model{}
}
