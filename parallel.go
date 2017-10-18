package main

import (
	"fmt"
	"time"
)

func main() {

	//channel anlegen
	ch := make(chan int, 2) //2 ist kapazität

	go func(z int) {
		x := <- ch
		y := <- ch
		fmt.Println("hello", x, y, z)
	}(50)

	ch <- 32
	ch <- 40

	//sonst würde man nicht sehen, denn
	//wenn das programm beendet wird, werden sämtliche neben laufende prozesse beendet
	time.Sleep(time.Second / 2)

	ch2 := make(chan int, 2) //2 ist kapazität
	done := make(chan struct{})

	go func() {
		defer close(done)
		for x := range ch2 {
			fmt.Println("hillo", x)
		}
		fmt.Println("after")

	}()

	ch2 <- 35
	ch2 <- 45
	//sonst wartet meine for schleife auf weitere inputs im channel
	//wenn ich das auskommentiere, dann wird "after" nicht angezegt
	close(ch2)

	//time.Sleep(time.Second / 2)
	//statt sleep
	<- done
	//oder wenn result kommt:
	/*
	for range done {

	}
	*/

}
