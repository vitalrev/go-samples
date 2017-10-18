package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	n := runtime.NumCPU()
	fmt.Println(n)

	//kann man ändern mit environment var, aber funzt nicht immer
	//GOMAXPROCS=1

	ch := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		//i := i  //sonst übergibt man nur die erste zahl - 0 an die go routine
		//oder lieber als parameted (id)
		go func(id int) {
			defer wg.Done()
			for x := range ch {
				fmt.Println(id, x)
			}
		}(i)
	}

	for j := 0; j < 10000; j++ {
		ch <- j
	}
	close(ch)  //wichtig!

	wg.Wait()

	//time.Sleep(10 * time.Second)
}
