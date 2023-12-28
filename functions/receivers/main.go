package main

import "fmt"

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between func keyword and the name of the method.
// func (receiver type) MethodName(parameters) (returnTypes) {}
// You can only define a method with a receiver whose type is defined in the same package as the method.
type Person struct {
	name string
	age  int
}

// Pass by value receiver
func (person Person) printPersonDetails() {
	person.name = "javascript"
	fmt.Println(person.name, person.age)
}

// Pass by reference receiver, changes the value of the struct obj as well
func (person *Person) printPersonDetailsP() {
	person.name = "golang"
	person.age += 2
	fmt.Println(person.name, person.age)
}

func main() {
	fmt.Println("Learning receivers")
	person := Person{
		name: "golang",
		age:  50,
	}
	var personPointer *Person = &person
	person.printPersonDetails()
	fmt.Println("Inside main", person.name, person.age) // -> Pass by value
	person.printPersonDetailsP()
	fmt.Println("Inside main", person.name, person.age) // -> Pass by reference

	// Go is intelligent enough to do conversion between values and pointers for method calls, that's why even if you make the call
	// with a pointer type, based on the receiver type it will either pass by value or pass by reference.

	personPointer.printPersonDetails()
	fmt.Println("Inside main", person.name, person.age) // -> Pass by value
	personPointer.printPersonDetailsP()
	fmt.Println("Inside main", person.name, person.age) // -> Pass by reference
}
