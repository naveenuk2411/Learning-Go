package main

import "fmt"

// range can be used to iterate over the following:
// 1. slices
// 2. maps
// 3. strings
// Every iteration on map/slice returns 2 values i.e index/key, and the copy of the element/value
// Range on string iterates over unicode code points, where it returns 2 values i.e starting byte index of the rune, and the rune itself.

// You can declare type definitions of your own as non-struct types which can be used as an alias for built-in
// type declarations.
type Name string
type Names []string

func sayHello(name Name) {
	fmt.Println("Hello " + name)
}

func sayHelloToAll(names Names) {
	for curName := 0; curName < len(names); curName++ {
		fmt.Println(names[curName])
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
func main() {
	fmt.Println("Learning range")
	myMap := map[int]int{1: 10, 2: 20, 3: 30, 4: 40}
	// No guarantee of accessing the keys sequentially in the same order as of insertions.
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// This will access the index sequentially.
	slice := []int{10, 20, 30, 40}
	for index, element := range slice {
		fmt.Println(index, element)
	}

	for key := range myMap {
		fmt.Println(key)
	}

	for _, value := range myMap {
		fmt.Println(value)
	}

	myName := Name("go lang")
	sayHello(myName)

	names := Names([]string{"name1", "name2", "name3"})
	sayHelloToAll(names)

	for startByteIndexOfRune, currentRune := range myName {
		fmt.Println(startByteIndexOfRune, currentRune)
	}
}
