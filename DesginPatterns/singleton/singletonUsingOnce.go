/*
There are other methods of creating a singleton instance in Go:

init function
We can create a single instance inside the init function. This is only applicable if the early initialization
of the instance is ok. The init function is only called once per file in a package, so we can be sure that
only a single instance will be created.

func init(){
	instance = &Connection{}
}

sync.Once
The sync.Once will only perform the operation once. See the code below:

*/

package main

import (
	"fmt"
	"sync"
)

type Connection struct {
	conn string
}

var instance *Connection
var once sync.Once

func getInstance(wg *sync.WaitGroup) *Connection {
	defer wg.Done()

	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating new instance...")
				instance = &Connection{
					conn: "newConnection",
				}
			})
	} else {
		fmt.Println("Instace is already created!!")
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
