package main

import (
	"context"
	"fmt"
	"time"
)

func query(ctx context.Context, quiz string, ch chan string) {
	var ans string
	fmt.Println(quiz)
	_, _ = fmt.Scanf("%s", &ans)
	fmt.Println("The answer:", ans)
	ch <- ans
}

func main() {
	fmt.Println("Hello in quiz!!")
	quizes := map[string]string{
		"Capital of India":  "delhi",
		"Capital of kntaka": "blr",
		"Capital of US":     "newyork",
	}

	var ch chan string
	var wrong, correct int
	ch = make(chan string)

	for quiz := range quizes {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		go query(ctx, quiz, ch)

		select {
		case ans := <-ch:
			if ans != quizes[quiz] {
				wrong += 1
			} else {
				correct += 1
			}
		case <-ctx.Done():
			fmt.Println("Timeout!!")
		}
	}
	fmt.Println("The total correct ans:", correct)

}
