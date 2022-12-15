package main

import "fmt"

type language interface {
	say_hello()
}

type french struct {
	name string
}

type german struct {
	name string
}

func (f french) say_hello() {
	fmt.Println(f.name + " speaking french!!")
}

func (g german) say_hello() {
	fmt.Println(g.name + " speaking german!!")
}

func main() {

	f := french{
		name: "ramesh",
	}

	g := german{
		name: "german",
	}

	f.say_hello()
	g.say_hello()
}

// output:
// ramesh speaking french!!
// german speaking german!!
