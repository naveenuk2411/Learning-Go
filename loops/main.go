package main

import "fmt"

// for init; condition; post { // body }
// Note: Unlike other languages, there are no parentheses () surrounding the three components of the header.
// In fact, inserting such parenthesis is a compilation error. However, the braces { } surrounding the loop
// body are always required.
func main() {
	init := 1
	fmt.Println("Basic for loop")
	for init <= 10 {
		fmt.Println(init)
		init++
	}

	fmt.Println("Complete for loop")
	for init := 1; init < 3; init++ {
		fmt.Println(init)
	}

	fmt.Println("For loop with break")
	for {
		fmt.Println("Break statement")
		break
	}

	fmt.Println("For loop with continue")
	for value := 1; value <= 10; value += 1 {
		if value%2 == 1 {
			continue
		}
		fmt.Println(value)
	}
}
