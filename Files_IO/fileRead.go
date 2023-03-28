package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Reading an entire file into memory
/*
To read an entire file into a variable, use the ReadFile
function from the ioutil library.
*/

func OpenAndReadFile() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	// fmt.Println("Reading an entire file Contents:\n")
	fmt.Println(string(data))
}

// Reading file in bytes
/*
Sometimes, it’s not feasible to read the whole file in one go,
especially when the file size is too large. So, instead,​ we can
read the file in chunks of set bytes. The following code reads
the file in 4-byte chunks.
*/
func ReadFileInBytes() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	p := make([]byte, 4)
	for {
		n, err := file.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}

// Reading a file line by line
func ReadFileLineByLine() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	fmt.Println("********OpenAndReadFile********")
	OpenAndReadFile()
	fmt.Println("********ReadFileInBytes********")
	ReadFileInBytes()
	fmt.Println("********ReadFileLineByLine********")
	ReadFileLineByLine()

}

/*
Reference:

1. https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185
2. https://gobyexample.com/reading-files
3. https://gobyexample.com/reading-files
*/
