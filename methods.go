package main

import "fmt"

type myType int

func (m myType) bla() {
	fmt.Println("bla")
}

func main() {
	var m myType
	m.bla()

	//kurzform
	fn := m.bla
	fn()

	//das gleiche wie kurzform
	//var fn func() = func() { m.bla() }
}