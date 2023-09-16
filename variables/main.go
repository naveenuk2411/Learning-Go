package main

// fmt stands for format package
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Learning variables")
	// Statically typed language, hence all variables must have a defined type at compile time.
	// Logical operators
	// &&, ||, !

	// Variables using var
	var name string = "Go lang"
	fmt.Println(name)

	// This will throw an error if you do not use the variable
	// var age int = 30

	// Go automatically infers the type as string
	var inferredName = "Infers the type as string"
	fmt.Println(inferredName)
	fmt.Printf("%T\n", inferredName)

	var age = 30
	fmt.Printf("%T\n", age) // int
	var newAge int32 = 30
	fmt.Printf("%T\n", newAge) // int32

	var isCool = true
	fmt.Println(isCool) // true
	isCool = false      // This will throw an error if you use const isCool = true
	fmt.Println(isCool) // false

	// Variable shorthand declarations
	// These declarations do not work outside the functions, only declarations using var keywork works outside the functions, which obviously
	// act as global variables
	coolName := "Go lang learning" // Again automatically infers the type
	fmt.Println(coolName)          // Go lang learning
	fmt.Printf("%T\n", coolName)   // string

	userName, userEmail := "goLang", "goLang@gmail.com"
	fmt.Println(userName, userEmail) // golang goLang@gmail.com

	// Constants are defined using the const keyword and can be numbers, characters, strings or booleans
	const oldAge int = 45
	fmt.Println(oldAge)

	// Int variables, and type conversions
	// int32: -2e9 to 2e9
	// int64: -9e18 to 9e18
	// float64, uint, int
	// +, -, /, *, %
	// a += 5, a++, a--

	// Type conversions
	var a float64 = 3.5
	var b int = 30
	// var c int = b * a  -> This will throw an error, automatic type conversion is not followed
	var c int = b * int(a)
	fmt.Println(c)

	// Strings
	const stringName string = "Go lang"
	fmt.Println(stringName)
	const concatString string = "Learning" + " " + "golang."
	fmt.Println(concatString)
	var newStringName string = strings.ReplaceAll(stringName, "Go", "Gone")
	fmt.Println(newStringName)
	newStringName = strings.TrimSpace(newStringName)
	fmt.Println(newStringName)
}
