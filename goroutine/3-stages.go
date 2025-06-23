//Create a pipeline of 3 stages using goroutines
//and channels (e.g., generate → square → print).

package main

import (
	"fmt"
	"sync"
)

func generate(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Printf("generated number %d\n", i)
		ch <- i
	}
	close(ch)
}

func square(ch chan int, ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Printf("******computed square %d\n", i)
		ch1 <- i * i
	}
	close(ch1)

}

func print(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Printf("______recieved square %d\n", i)
	}
}

func main() {
	var ch = make(chan int)
	var ch2 = make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go generate(ch, &wg)
	go square(ch, ch2, &wg)
	go print(ch2, &wg)
	wg.Wait()
}
