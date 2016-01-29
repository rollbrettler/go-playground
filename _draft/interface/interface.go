package main

import (
	"fmt"
)

type message interface {
	Show() string
}

type greeting struct {
	foreword, name string
}

type letter struct {
	content string
}

func (g *greeting) Show() string {
	return fmt.Sprintf("%v, %v", g.foreword, g.name)
}

func (l *letter) Show() string {
	return fmt.Sprintf("%v", l.content)
}

func printMessage(m message) {
	fmt.Printf("%v\n", m.Show())
}

func main() {
	g := &greeting{"Hello", "Tim"}
	l := &letter{"have a nice day"}

	printMessage(g)
	printMessage(l)
}
