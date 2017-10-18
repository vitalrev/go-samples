package main

import "fmt"

type yourType complex128

func (y yourType) bla() int {
	return 42
}

func (y yourType) blup() int {
	return 32
}

func (y yourType) foo() int {
	return 67
}


func main() {
	//abstract function - method expresion
	var fn func(yourType) int
	//map function to bla implementation
	fn = yourType.bla

	var y yourType

	fmt.Println(fn(y))

	//und jetzt mal andere implementierung verwenden
	fn = yourType.foo
	fmt.Println(fn(y))

}