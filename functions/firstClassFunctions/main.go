package main

import (
	"fmt"
)

// Functions are first-class citizens meaning they can be treated as values.
// Functions can be assigned to a variable
// Functions can be passed to some function as argument
// Functions can accept functions as parameter
// Return functions from functions similar to values

// Simple function
func greeting(name string) string {
	return "Good morning/evening " + name
}

// A function returning another function
// Also this is a classic example of closure where the computed value of sum is in the scope of returnFunc
// But since the func() is also defined within the same scope or nested scope, so it is able to get the value of the sum variable.
func returnFunc(value int) func() {
	sum := (value * (value + 1)) / 2
	return func() {
		fmt.Printf("Sum of first %d natural numbers is %d\n", value, sum)
	}
}

// A function taking a function as a paramter
func takeAnotherFunc(callbackFn func(string) string, message string) string {
	return callbackFn(message)
}

// Closures in functions
func intSeq() (func() int, func() int) {
	i := 0
	return func() int {
			i++
			return i
		}, func() int {
			i--
			return i
		}
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	fmt.Println("Learning First class functions")

	// A anonymous function being assigend as a value to a variable
	assignedFunc := func(name string) {
		fmt.Println("I am being assigned to a variable function yay!! " + name)
	}

	// Or
	var myAssignedFn func()
	myAssignedFn = func() {
		fmt.Println("I am being assigned to a variable function yay!! ")
	}

	myAssignedFn()

	// Self calling function, IIFE
	func() {
		fmt.Println("Self calling function")
	}()

	sumOfFirstN := returnFunc(20)
	sumOfFirstN()
	assignedFunc("Naveen")
	// A function being passed as argument
	fmt.Println(takeAnotherFunc(greeting, "Naveen"))

	myAddSeq, mySubSeq := intSeq()
	// myIntSeq will be able to refer the variable i defined in the scope of it's outer function intSeq.
	// Any updates done to that variable will be read by myIntSeq fn, as it refers the value of the variable by reference.
	fmt.Println(myAddSeq(), myAddSeq(), mySubSeq(), myAddSeq())

	myNewAddSeq, _ := intSeq()
	fmt.Println(myNewAddSeq(), myNewAddSeq())

	// Q. What happens in case when outer function exits i.e it returns and the inner function within the outer function references
	// some variables. Example: Inner func is a goroutine which references the variables defined in the outer function and the outer
	// function returns.

	// Ans. Go compiler analyzes the closures, and allocates a heap memory for the variables that are accessed as part of closure, i.e
	// Go creates a memory for such variables in heap instead of stack, so even if the outer function call exits, inner functions could
	// still refer the same variable in the heap memory.
	// Once the variable is no longer accessed by any of inner functions accessing it as part of closure, tha gc takes care of clearing
	// up the heap memory allocated to that variable.

	// Recursive function call
	var fib func(n int) int

	fib = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	nthFibonacciNumber := fib(5)
	fmt.Println("Nth fibonacci number", nthFibonacciNumber)
	fmt.Println("Factorial number", factorial(10))
}
