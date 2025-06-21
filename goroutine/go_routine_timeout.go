package main

import (
	"fmt"
	"time"
)

func task1(done chan bool) {
	time.Sleep(3 * time.Second)
	done <- true
}

func main() {
	var done = make(chan bool)
	go task1(done)
	select {
	case <-done:
		fmt.Println("go routine done it task")
	case <-time.After(2 * time.Second):
		fmt.Println("timout while waiting for goroutine")

	}
}
