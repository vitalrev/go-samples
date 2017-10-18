package main

import "fmt"

type base int

type sub struct {
	base base
}

func (b base) print() {
	fmt.Println("base: ", b)
}

func (s sub) print() {
	s.base.print()
	fmt.Println("sub: ", s)
}

func main() {
	var s sub
	//b.print()
	s.print()

	var x struct{ base }
	x.print()  //delegate to base.print()
}

