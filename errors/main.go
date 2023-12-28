package main

import (
	"errors"
	"fmt"
)

// Go defines an in-built error type basically an interface, and uses these "error" values to indicate any type of error.
// By convention errors are the last return values, and have the type error, a built-in interface
// errors.New("Error message here") constructs a basic error value with given error message

// How does the error interface look like ?
// type error interface {
// 	Error() string
// }

// errors package which just allows you to construct an error with given message internally implements the error interface
// type errorString struct {
//   s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

// You can construct one of these values with the errors.New function. It takes a string that it converts to an errors.errorString
// and returns as an error value.
// New returns an error that formats as the given text.
// func New(text string) error {
//   return &errorString{text}
// }

// Hence when you do error.New("Error message here"), it essentialy returns you a the errorString type which has Error() method
// implemented.

// When you do fmt.Println(err), it formats an error value by calling its Error() string method.

func checkError1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Arg not supported")
	}
	return arg, nil
}

// Implement your own custom error, you just need to implement the Error() method for implementing the error interface for the
// defined type.
type customError struct {
	value   int
	message string
}

func (cError *customError) Error() string {
	return cError.message
}

func checkError2(arg int) (res int, err error) {
	if arg == 42 {
		return -1, &customError{value: arg, message: "We do not support this argument"}
	}
	return arg, nil
}

func main() {
	fmt.Println("Learning errors in go lang")

	for _, value := range []int{20, 42} {
		if res, err := checkError1(value); err != nil {
			fmt.Println("Error encountered in checkError1", err)
		} else {
			fmt.Printf("Success! %d\n", res)
		}
	}

	for _, value := range []int{20, 42} {
		if res, err := checkError2(value); err != nil {
			fmt.Println("Error encountered in checkError2", err)
		} else {
			fmt.Printf("Success! %d\n", res)
		}
	}

	// If you want to access different properties of custom error object, you can do so by the following
	_, err := checkError2(42)
	if ae, ok := err.(*customError); ok {
		fmt.Println(ae.message)
		fmt.Println(ae.value)
	}
}
