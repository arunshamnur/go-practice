//Create a program where a producer sends numbers to a
//channel and a consumer prints them.

package main

import (
	"fmt"
	"sync"
)

func produce(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	var ch = make(chan int)
	wg.Add(2)
	go produce(ch, &wg)
	go consume(ch, &wg)
	wg.Wait()
}
