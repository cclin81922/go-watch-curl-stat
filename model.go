package main

type Model map[int]int

func (m Model) Inc(key int) {
    m[key]++
}

func NewModel() Model {
    return Model{}
}
