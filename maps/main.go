package main

import (
	"fmt"
)

// Key value pair, with keys being unique
// Creation of Map Syntax:
// 1. myMap := map[keyType]valueType{} -> Example: myMap := map[string]int{}
// 2. myMap := make(map[keyType]valueType) -> Example: myMap := make(map[string]int)
func main() {
	fmt.Println("Learning maps")

	myMap := make(map[string]int)
	myMap["a"] = 0
	myMap["b"] = 1
	myMap["c"] = 2
	myMap["d"] = 3

	myNewMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(myNewMap)

	// Retreive a key
	fmt.Println(myMap["a"], len(myMap))

	// Update a key
	myMap["a"] = 10
	fmt.Println(myMap["a"])

	// Store the value of map
	value := myMap["a"]
	fmt.Println(value)

	// Delete a key
	delete(myMap, "a")
	fmt.Println(myMap["a"]) // -> Even if the key does not exist, it will return the "zero value" for the value

	// Check if the key exists
	value, exists := myMap["a"]
	fmt.Println(value, exists)

	_, existsNew := myMap["b"]
	fmt.Println(existsNew)

	var newValue, newExist = myMap["a"]
	fmt.Println(newValue, newExist, len(myMap))

	// Clear all key value pairs of map -> Only available in go 1.21, will not run for now in this PC
	fmt.Println(myNewMap, len(myNewMap))
	// clear(myNewMap)
	// fmt.Println(myNewMap)

	// Alternate way to delete all keys in map
	for key := range myNewMap {
		delete(myNewMap, key)
	}
	fmt.Println(myNewMap, len(myNewMap))
}
