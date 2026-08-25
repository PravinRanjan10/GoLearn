package main

import "fmt"

type Machine interface {
	Print()
	Scan()
	Fax()
}

type BasicPrinter struct{}

func (b BasicPrinter) Print() {
	fmt.Println("Printing basic printer...")
}

func (b BasicPrinter) Scan() {
	fmt.Println("Not supported!")
}

func (b BasicPrinter) Fax() {
	fmt.Println("Not supported!")
}

func main() {
	b := &BasicPrinter{}
	b.Print()
	b.Fax()
	b.Scan()
}

/*
ISP says:

A type should not be forced to implement methods that it does not need.

A common way to achieve this in Go is to create small, focused interfaces rather than one large interface.

❌ Bad example — Large interface
type Machine interface {
	Print()
	Scan()
	Fax()
}

Now we have two printers:

type BasicPrinter struct{}

func (b BasicPrinter) Print() {
	fmt.Println("Printing...")
}

func (b BasicPrinter) Scan() {
	panic("not supported")
}

func (b BasicPrinter) Fax() {
	panic("not supported")
}

BasicPrinter only supports printing, but because Machine requires all three methods, it is forced to implement Scan() and Fax().

This violates ISP.
*/
