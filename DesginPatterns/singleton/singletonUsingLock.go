package main

import (
	"fmt"
	"sync"
)

type Connection struct {
	conn string
}

var lock sync.Mutex
var instance *Connection

func getInstance(wg *sync.WaitGroup) *Connection {
	defer wg.Done()
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &Connection{
				conn: "newConnection",
			}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Instance is already created!!!")
	}
	return instance
}

func main() {
	//var wg sync.WaitGroup
	wg := new(sync.WaitGroup)
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go getInstance(wg)
	}
	wg.Wait()
}
