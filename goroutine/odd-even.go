//Use two goroutines to print even
//and odd numbers from 1 to 10 in order.

package main

import (
	"fmt"
	"sync"
)

func even(n int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
		if i == n {
			close(ch)
			break
		}
		i++
		ch <- i
	}
}

func odd(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
		i++
		ch <- i
	}
}

func main() {
	var ch = make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go odd(&wg, ch)
	go even(10, &wg, ch)
	ch <- 0
	wg.Wait()
}
