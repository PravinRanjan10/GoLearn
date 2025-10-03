/*
🔹 When to Use Fan-out

Use Fan-out when you want to parallelize tasks across multiple goroutines.

You have a single source of jobs (channel, list, or generator).

Multiple goroutines can work on jobs independently.

Each worker takes one job at a time, speeding up processing.

✅ Examples:

Web scraping multiple URLs concurrently.

Processing files (e.g., resizing images, parsing logs).

Handling many API requests.

CPU-intensive tasks split across cores.

🔹 When to Use Fan-in

Use Fan-in when you want to merge results from multiple sources (goroutines) into a single channel for easier consumption.

Multiple goroutines produce results independently.

A single goroutine needs to consume results in one place.

✅ Examples:

Merging results of multiple API calls.

Combining logs from different services into one channel.

Collecting sensor data from multiple devices.

Gathering processed chunks of data into a single stream.
*/

// Thumb Rule=====:
// Use Fan-out if:
//      “I have one source of jobs, and I want them processed faster by splitting work.”

// Use Fan-in if:
//      “I have multiple sources of results, and I want to consume them in one place.”

// Use Fan-out + Fan-in if:
//      “I have jobs to distribute and then need to merge results back.”

package main

import (
	"fmt"
	"sync"
	"time"
)

func simulateUrlFetch(url string) string {
	time.Sleep(time.Second)
	return fmt.Sprintf("URL:%s is fetched", url)
}

func Worker(w int, urls, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range urls {
		output := simulateUrlFetch(url)
		results <- fmt.Sprintf("Worker: %d:  %s", w, output)
	}
}

func main() {

	urlCh := make(chan string, 10)
	results := make(chan string, 10)

	var wg sync.WaitGroup

	// FAN-OUT: Three worker
	for w := range 3 {
		wg.Add(1)
		go Worker(w, urlCh, results, &wg)
	}

	// feed jobs (URLs)
	urlList := []string{
		"https://site1.com",
		"https://site2.com",
		"https://site3.com",
		"https://site4.com",
		"https://site5.com",
	}

	for _, url := range urlList {
		urlCh <- url
	}
	close(urlCh)

	// FAN-IN: Merge the result and print
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
