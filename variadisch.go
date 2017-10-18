package main

import "fmt"

func print( x interface{} ) {
	//fmt.Println(x)

	//x soll abhängig vom type behandelt werden
	switch v := x.(type) {
	case int:
		fmt.Println("integer")
	case float64:
		fmt.Printf("float from type %T\n", v)
	case string:
		fmt.Println("string")
	}

	//oder wir versuchen x in int zu trnsformieren
	/*
	v , ok := x.(int)
	if !ok {
		return
	}
	fmt.Println(v)
	*/
}

func main() {
	print(1)
	print(3.1415)
	print("hello")
}
