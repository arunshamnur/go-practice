package main

import (
	"fmt"
	"sync"
)

func odd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println("odd", i)
		i++
		if i == 10 {
			close(ch)
			break
		}
		ch <- i
	}
}
func even(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println("even", i)
		i++
		if i == 10 {
			close(ch)
			break
		}
		ch <- i
	}
}
func main() {
	var ch1 = make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go odd(ch1, &wg)
	ch1 <- 0
	go even(ch1, &wg)

	wg.Wait()
}
