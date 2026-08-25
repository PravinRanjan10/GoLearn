package main

import "fmt"

type print interface {
	Print()
}

type scan interface {
	Scan()
}
type fax interface {
	Fax()
}

type BasicPrinter struct{}

func (b BasicPrinter) Print() {
	fmt.Println("Printing from base printer..")
}

type AdvancedPrinter struct{}

func (a AdvancedPrinter) Print() {
	fmt.Println("Printing from advanced printer..")
}

func (a AdvancedPrinter) Scan() {
	fmt.Println("Scaning from advanced printer..")
}

func (a AdvancedPrinter) Fax() {
	fmt.Println("Fax from advanced printer..")
}

func DocPrinter(p print) {
	p.Print()
}

func DocScan(s scan) {
	s.Scan()
}

func DocFax(f fax) {
	f.Fax()
}

func main() {

	b := BasicPrinter{}
	DocPrinter(b)

	a := AdvancedPrinter{}
	DocPrinter(a)
	DocScan(a)
	DocFax(a)

}
