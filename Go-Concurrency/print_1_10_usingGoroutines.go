package main

import (
	"fmt"
	"sync"
)

func main() {
	oddChan := make(chan bool)
	evenChan := make(chan bool)
	var wg sync.WaitGroup

	oddFn := func() {
		defer wg.Done()
		for i := 1; i <= 9; i += 2 {
			<-oddChan // Wait for signal to print odd
			fmt.Printf("odd: %d\n", i)
			evenChan <- true // Signal evenFn to print
		}
		close(evenChan)
	}

	evenFn := func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-evenChan // Wait for signal to print even
			fmt.Printf("even: %d\n", i)
			if i < 10 {
				oddChan <- true // Signal oddFn to print
			}
		}
		close(oddChan)
	}

	wg.Add(2)
	go oddFn()
	go evenFn()

	oddChan <- true // Start with oddFn
	wg.Wait()
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func oddPrint(wg *sync.WaitGroup, oddChan, evenChan chan bool) {
// 	defer wg.Done()
// 	for i := 1; i <= 9; i += 2 {
// 		<-oddChan // Wait for signal to print odd
// 		fmt.Println(i)
// 		evenChan <- true // Signal evenFn to print
// 	}
// 	close(evenChan)
// }

// func evenPrint(wg *sync.WaitGroup, oddChan, evenChan chan bool) {
// 	defer wg.Done()
// 	for i := 2; i <= 10; i += 2 {
// 		<-evenChan // Wait for signal to print odd
// 		fmt.Println(i)
// 		oddChan <- true // Signal evenFn to print
// 	}
// 	close(oddChan)
// }

// func main() {
// 	oddChan := make(chan bool)
// 	evenChan := make(chan bool)
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go oddPrint(&wg, oddChan, evenChan)
// 	go evenPrint(&wg, oddChan, evenChan)

// 	oddChan <- true // Start with oddFn
// 	wg.Wait()
// }

/*

package main

import (
	"fmt"
	"sync"
)

func Odd(odd, even chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val := <-odd
		if val > 10 {
			close(even)
			return
		}
		fmt.Println(val)
		even <- val + 1
	}
}

func Even(odd, even chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-even
		if !ok {
			return
		}
		fmt.Println(val)
		odd <- val + 1
	}
}

func main() {
	fmt.Println("Hello...")
	wg := new(sync.WaitGroup)

	odd := make(chan int)
	even := make(chan int)

	wg.Add(2)
	go Odd(odd, even, wg)
	go Even(odd, even, wg)

	odd <- 1

	wg.Wait()

	fmt.Println("I am done")

}

*/
