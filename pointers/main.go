package main

import "fmt"

// Pointers store the address of a variable, struct etc.
// Pointers can drastically decrease the memory footprint of the application, when values need to be passed around functions.
// The zero value of a pointer is nil(similar to nullptr in C++).
// To find the address of a variable, struct etc, use & operator
// Note: Before dereferencing a pointer, always check if its value is nil or not

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Learning pointers")

	// Declaring a pointer
	var pointer *int
	value := 10
	pointer = &value
	fmt.Println(pointer)

	// Dereferencing a pointer
	pointerValue := *pointer
	fmt.Println(pointerValue)

	*pointer = 100
	fmt.Println(value)        // value variable was pointing to pointer address, so value becomes 100
	fmt.Println(pointerValue) // -> This is a new variable when it was created in memory, and was just assigned the value, hence 10

	person1 := Person{
		name: "golang",
		age:  50,
	}

	personPointer1 := &person1
	var personPointer2 *Person = &person1

	fmt.Println(*personPointer1, *personPointer2)

	var personPointer3 *Person = &Person{
		"javascript",
		500,
	}
	fmt.Println(*personPointer3)

	// Dereferencing a struct pointer
	// Go automatically dereferences the struct pointer to directly access properties
	fmt.Println(personPointer3.name, personPointer3.age)

	// Maps and slices are already acting as pointers, so we really do not need separate pointers for them.
	// Hence when passed to functions, they really are passed by reference.
	// However append in case of slice works differently, append just creates a new slice from scratch giving it new memory,
	// even if you append elements to the same slice.
	// Refer: functions in Go here for more info
}
