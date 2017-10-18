package main

import (
	"fmt"
)


func myRoutine(ch <- chan int, done chan <- struct{}) {
	defer close(done)
	for x := range ch {
		fmt.Println("hillo", x)
	}
	fmt.Println("after")
}

func main() {

	ch := make(chan int, 2) //2 ist kapazität
	done := make(chan struct{})

	go myRoutine(ch, done)

	ch <- 35
	ch <- 45
	//sonst wartet meine for schleife auf weitere inputs im channel
	//wenn ich das auskommentiere, dann wird "after" nicht angezegt
	close(ch)

	<- done
}
