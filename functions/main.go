package main

import "fmt"

// Paramters: names defined in the function's signature
// Arguments: concrete values passed to the function parameters when we invoke the function
// There are no default values for parameters so all function parameters are required.
// Parameters must be explicitly typed, there is no type inference.
// Return type: The function can return in one of the following ways:
// 1. nothing -> do not mention the return type
// 2. a single value -> func greet(message string) string {}
// 3. multiple values -> func greet(messsage string) (string, string) {}

func greeting(name string) string {
	return "Hello " + name
}

func helloAndGoodbye(name string) (string, string) {
	return "Hello " + name, "Goodbye " + name
}

// Instead of add(num1 int, num2 int) we can do the following as both are same types
func add(num1, num2 int) int {
	return num1 + num2
}

// A variadic function is a function that accepts a variable number of arguments.

// If the type of the last parameter in a function definition is prefixed by ellipsis ...,
// then the function can accept any number of arguments for that parameter.

// 1. The variadic parameter must be the last parameter of the function.
// 2. The variadic parameter in itself is compiled to a slice, so we can pass a slice to it as well, and can interpret and
// use it as slice within the function
func printValues(num1, num2 int, nums ...int) {
	slice := []int{num1, num2}
	slice = append(slice, nums...)
	fmt.Println(slice)
}

// Named return values, in case of named return values we can return without specifying the values
func sumAndMultiply(a, b, c int) (sum, mult int) {
	sum, mult = a+b, a*b // Not sum, mult := a+b, a*b
	sum -= c
	mult -= c
	return
}

func calculateSum(nums ...int) (sum int) {
	sum = 0
	for _, value := range nums {
		sum += value
	}
	return
}

// Pass by value vs Pass by reference
// Strictly speaking, all arguments passed to the functions are pass by values, until we use pointers to pass
// them by reference
// NOTE: Exceptions to this rule are maps and slices, when we pass a slice or map as argument, then they are treated
// as pointer types and are passed by reference.

// Using pointers to pass by reference
// Pointers are denoted by *. Example *int

func multiplyBy2(val int) int {
	val *= 2
	return val
}

func multiplyBy2P(val *int) int {
	*val *= 2
	return *val
}

func modifySlice(slice []int) ([]int, []int) {
	slice[0] = 24                      // This will modify the value in the original slice passed to it as argument
	newSlice := append(slice, 2, 3, 4) // This will create a new slice and return this new slice
	return slice, newSlice
}

func main() {
	fmt.Println("Learning functions")
	fmt.Println(greeting("go lang"))
	fmt.Println(helloAndGoodbye("go lang"))
	fmt.Println(add(3, 4))
	printValues(2, 3, 4, 5, 6)
	printValues(2, 3)

	greetingString := greeting("javascript")
	fmt.Println(greetingString)

	newGreetingString1, newGreetingString2 := helloAndGoodbye("javascript")
	fmt.Println(newGreetingString1, newGreetingString2)

	fmt.Println(sumAndMultiply(1, 2, 3))

	slice := []int{1, 2, 3}
	printValues(2, 3, slice...)

	a, b := 2, 3
	fmt.Println(a, b)

	cur := 2
	fmt.Println(multiplyBy2(cur), cur)
	fmt.Println(multiplyBy2P(&cur), cur)

	// You can always use the blank operator to omit the values you do not want to have stored in a variable.
	slice, newSlice := modifySlice(slice)
	fmt.Println(slice, newSlice)

	overallSum := calculateSum(1, 2, 3, 4, 5)
	fmt.Println(overallSum)
}
