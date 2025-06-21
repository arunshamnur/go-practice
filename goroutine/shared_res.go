//Implement a function that launches N goroutines but limits only M
//to run at a time.

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var mu sync.RWMutex
	var counter int
	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}(&wg, &mu)
	}
	wg.Wait()
	fmt.Println(counter)
}
