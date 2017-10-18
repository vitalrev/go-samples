package main

import (
	"fmt"
	"sync"
)

func main() {
	//Lösung
	var mu sync.Mutex

	var problem int

	var wg sync.WaitGroup

	for i:=0; i<5; i++ {
		wg.Add(1)  //warte auf 1 process
		go func() {
			defer wg.Done()  //process beendet, nicht mehr warten
			for j:=0; j< 10000; j++ {
				mu.Lock()  //synchronise
				problem++  //Problemstelle
				mu.Unlock()
			}
		}()
	}

	//time.Sleep(time.Second)
	wg.Wait()

	//bei jedem run anderes ergebnis
	// go build -race
	fmt.Println(problem)
}
