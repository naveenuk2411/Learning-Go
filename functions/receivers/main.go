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
	person.name = "javascript"
	fmt.Println(person.name, person.age)
}

func main() {
	fmt.Println("Learning receivers")
	person := Person{
		name: "golang",
		age:  50,
	}
	person.printPersonDetails()
	fmt.Println("Inside main", person.name, person.age) // -> Pass by value
	person.printPersonDetailsP()
	fmt.Println("Inside main", person.name, person.age) // -> Pass by reference
}
