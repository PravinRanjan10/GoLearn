// This is used for controlling concurrency — instead of spawning thousands of goroutines
// (which can overwhelm CPU/memory), you use a limited pool of workers.

// The fixed number of workers (goroutines) are allowed to run at anytime (known as controller concurrency).

// This is helpful in batch data processing

// It helps in resource management by limiting the number of goroutines and re-use them for multiple tasks

package main

import (
	"fmt"
)

func worker(id int, jobs chan int) {
	for job := range jobs {
		fmt.Printf("worker:%d, processing job:%d\n", id, job)
	}
}

func main() {

	jobs := make(chan int)

	// start 3 workers
	for w := range 3 {
		go worker(w, jobs)
	}

	// there are 10 jobs, which can be processed among three goroutines
	for j := range 10 {
		jobs <- j
	}

	close(jobs)
}

// One more example: https://www.youtube.com/watch?v=ZWMiKQXmh9s&t=318s
