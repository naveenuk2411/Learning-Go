package main

import "fmt"

// A slice, is a dynamically-sized, flexible view of the elements of an array.
// A slice is written like []T with T being the type of the elements in the slice
func main() {
	fmt.Println("Learning arrays")

	// We need to specify the type of the array elements.
	// We also need to specify the size of the array, however this is optional if the size is not known before hand.
	var fruitArr [2]string
	fmt.Println(fruitArr)           // Prints [  ] -> 2 empty strings
	fmt.Printf("%T\n", fruitArr[0]) // Prints string
	fruitArr[0] = "Apple"           // Prints [Apple ] -> 1 empty string
	fmt.Println(fruitArr)

	var scoresArr [2]int
	fmt.Println(scoresArr) // Prints [0 0]

	// Declare and assign
	newFruitsArr := [2]string{"Orange", "Apple"} // Prints [Orange Apple]
	fmt.Println(newFruitsArr)

	// Size not known before hand, also known as slices
	newScoresArr := []int{1, 2, 3}
	fmt.Println(newScoresArr)
	fmt.Println(len(newScoresArr), newScoresArr[1:2]) // Prints 3 [2]

	// Creating a empty slice
	var emptySlice []int
	// Self append to the same slice
	emptySlice = append(emptySlice, 2, 3, 4)
	fmt.Println(emptySlice)

	// Creating a new slice from exisiting slice
	var newSlice = emptySlice[1:3]
	fmt.Println(newSlice) // [3 4]
	newSlice = emptySlice[:3]
	fmt.Println(newSlice) // [2 3 4]
	newSlice = emptySlice[1:]
	fmt.Println(newSlice) // [3 4]
	newSlice = emptySlice[:]
	fmt.Println(newSlice)

	// Merging two slices
	mergedSlice := append(newSlice, emptySlice...)
	fmt.Println(mergedSlice)
}
