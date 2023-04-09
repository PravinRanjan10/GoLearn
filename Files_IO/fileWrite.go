package main

import (
	"fmt"
	"log"
	"os"
)

func WriteString() {
	data := []string{
		"line-1",
		"line-2",
		"line-3",
		"line-4",
	}

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, d := range data {
		n, err2 := file.WriteString(d + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println("no of bytes:", n)
	}
}

func WriteByte() {

	var data byte
	data = 97

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err2 := file.Write([]byte{data})
	if err2 != nil {
		log.Fatal(err2)
	}

	val2 := " and red fox\n"
	data2 := []byte(val2)

	var idx int64 = int64(len(data2))

	_, err3 := file.WriteAt(data2, idx)

	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("done")

}

func main() {
	WriteString()
	WriteByte()
}

//Reference:
//	1. https://zetcode.com/golang/writefile/
