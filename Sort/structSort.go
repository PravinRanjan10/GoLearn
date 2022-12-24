package main

import (
	"fmt"
	"sort"
)

// NameSorter sorts planets by name.
type NameSorter []Planet

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

// AgeSorter sorts planets by name.
type AgeSorter []Planet

func (a AgeSorter) Len() int           { return len(a) }
func (a AgeSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AgeSorter) Less(i, j int) bool { return a[i].Age < a[j].Age }

type Planet struct {
	Name string
	Age  int
}

func main() {
	var mars Planet
	mars.Name = "Mars"
	mars.Age = 20

	var earth Planet
	earth.Name = "Earth"
	earth.Age = 10

	var venus Planet
	venus.Name = "Venus"
	venus.Age = 15

	planets := []Planet{mars, venus, earth}
	fmt.Println("unsorted planets:", planets)

	sort.Sort(NameSorter(planets))
	fmt.Println("sorted by name:", planets)

	sort.Sort(AgeSorter(planets))
	fmt.Println("sorted by age:", planets)
}

/*
Output:
unsorted planets: [{Mars 20} {Venus 15} {Earth 10}]
sorted by name: [{Earth 10} {Mars 20} {Venus 15}]
sorted by age: [{Earth 10} {Venus 15} {Mars 20}]
*/
