package main

import (
	"fmt"
	"sync"
)

func G1(g1, g2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-g1
		if val > 10 {
			close(g2)
			return
		}
		if !ok {
			return
		}
		fmt.Println("g1-->", val)
		g2 <- val + 1
	}
}

func G2(g2, g3 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-g2
		if val > 10 {
			close(g2)
			close(g3)
			return
		}
		if !ok {
			return
		}
		fmt.Println("g2-->", val)
		g3 <- val + 1
	}
}

func G3(g1, g3 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-g3
		if val > 10 {
			close(g3)
			return
		}
		if !ok {
			close(g1)
			return
		}
		fmt.Println("g3-->", val)
		g1 <- val + 1
	}
}

func main() {

	wg := new(sync.WaitGroup)

	g1 := make(chan int)
	g2 := make(chan int)
	g3 := make(chan int)

	wg.Add(3)
	go G1(g1, g2, wg)
	go G2(g2, g3, wg)
	go G3(g1, g3, wg)

	g1 <- 1

	wg.Wait()

	fmt.Println("I am done")

}

/*
Output:

g1--> 1
g2--> 2
g3--> 3
g1--> 4
g2--> 5
g3--> 6
g1--> 7
g2--> 8
g3--> 9
g1--> 10
I am done
*/
