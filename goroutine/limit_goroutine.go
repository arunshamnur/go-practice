//Implement a function that launches N goroutines but limits only M
//to run at a time.

package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- struct{}{} //aquire channel
	//fmt.Printf("goroutine %d running\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine  %d done\n", id)
	<-ch
}

func main() {
	var M = 10
	var N = 3
	var wg sync.WaitGroup
	ch := make(chan struct{}, N)
	for i := 0; i < M; i++ {
		wg.Add(1)
		go task(i, ch, &wg)
	}
	wg.Wait()
}
