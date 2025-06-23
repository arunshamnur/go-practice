//Build a worker pool using goroutines and
//channels to process tasks concurrently.

package main

import (
	"fmt"
	"time"
)

func tasktorun(id int, job chan int, res chan int) {
	for i := range job {
		fmt.Printf("job %d run started job %d\n", id, i)
		time.Sleep(2 * time.Second)
		fmt.Printf("job %d  run completed job %d\n", id, i)
		res <- i * 2
	}
}

func main() {
	var jobs = 5
	var job = make(chan int, jobs)
	var result = make(chan int, jobs)
	for i := 0; i < 3; i++ {
		go tasktorun(i, job, result)
	}
	for i := 0; i < jobs; i++ {
		job <- i
	}
	close(job)
	for i := 0; i < jobs; i++ {
		<-result
	}
}
