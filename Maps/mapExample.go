package main

import (
	"fmt"
)

func main() {
	fmt.Println("helllo from map....")
	// 1. Declaration typ1
	var map1 map[int]int
	map2 := map[int]string{}
	fmt.Println("Map Type1:", map1, map2)

	// 2. Declaration type2
	map3 := map[int]string{1: "hello", 2: "how", 3: "are", 4: "you"}
	fmt.Println("Map Type2:", map3)

	// 3. Declaration Type3
	var map4 = make(map[int]string)
	fmt.Println("Map Type3:", map4)

	// 4. Insertion in map
	map2[1] = "Pravin"
	map2[2] = "Ranjan"
	fmt.Println("Map after intert:", map2)

	// 5. To check key available or not
	key1, ok := map3[2]
	fmt.Println("the key available:", key1, ok)

	// 6. Delete a key in hashing
	delete(map3, 2)
	fmt.Println("Map Type2:", map3)

	// 7. Iterate the map
	fmt.Println("itereting now...")
	for key, value := range map3 {
		fmt.Printf("key:%+v and value:%+v\n", key, value)
	}

}
