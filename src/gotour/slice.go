package gotour

import "fmt"

type base int

func (b base) print() {
	fmt.Println("base: ", b)
}

func main() {
	var b base
	b.print()
}


