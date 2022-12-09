package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(i int) {
	fmt.Printf("worker:%d starting\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("worker:%d finishing\n", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker1(i)
		}()
	}

	wg.Wait()
}
