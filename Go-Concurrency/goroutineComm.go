package main

import (
	"fmt"
	"sync"
)

func writer(wg *sync.WaitGroup, ch1 chan string) {
	defer wg.Done()
	ch1 <- "messages from writer"
}

func reader(wg *sync.WaitGroup, ch1 chan string) {
	defer wg.Done()
	fmt.Println("recieved messages in reader:", <-ch1)
}

func main() {

	ch1 := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go reader(wg, ch1)
	go writer(wg, ch1)
	defer close(ch1)
	wg.Wait()

}

// output: recieved messages in reader: messages from writer
