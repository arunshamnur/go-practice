//Launch 5 goroutines that each print their ID (1 to 5) concurrently.

package main

import (
	"fmt"
	"sync"
)

func gosum(arr []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	var sum int
	for _, val := range arr {
		sum += val
	}
	ch <- sum
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	maxgo := 5
	var ch = make(chan int, maxgo)
	chunks := (len(arr) + maxgo - 1) / maxgo
	for i := 0; i < len(arr); i += chunks {
		var end int
		if i+chunks < len(arr) {
			end = i + chunks
		} else {
			end = len(arr)
		}
		wg.Add(1)
		go gosum(arr[i:end], &wg, ch)
	}
	wg.Wait()
	close(ch)
	var finalsum int
	for i := range ch {
		finalsum += i
	}
	fmt.Println(finalsum)
}
