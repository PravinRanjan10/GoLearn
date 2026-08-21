/*
Create 100 goroutines, and each goroutine should increment the counter 1,000 times. So output should be: 100000
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter = 0
var aCounter int64

//var lock sync.Mutex

func Increment(lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	i := 0
	for i < 1000 {
		counter += 1
		i += 1
	}
	lock.Unlock()
}

func IncrementUsingAtomic(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for i < 1000 {
		atomic.AddInt64(&aCounter, 1)
		i += 1
	}
}

func main() {

	i := 0

	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i < 100 {
		wg.Add(1)
		go Increment(&lock, &wg)
		i += 1
	}

	for i < 100 {
		wg.Add(1)
		go IncrementUsingAtomic(&wg) // here no need to pass locks
		i += 1
	}

	wg.Wait()
	fmt.Println("Counter==:", counter)
	fmt.Println("Using atomic package Counter==:", aCounter)
}
