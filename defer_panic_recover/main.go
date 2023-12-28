package main

import "fmt"

// Panic: Something when unexpectedly wrong, and we are not prepared to such errors. Abort a function to execute further
// and returns an error value that we do not know how or want to handle -> Similar to exiting with a non-zero exit code.
// Panic is not exactly exit, as it will make the call to defer function call in the enclosing function's scope, which is
// not done by exit call.

// Defer: Ensures that an operation or a function call is performed in the end of a function's execution scope.
// Useful for performing any cleanups, basically a finally function call in context of try catch.
// Defining multiple defer statements within a funcion will execute those statements in the reverse order in which they are defined.
// Hence if you have defined defer Statement1, defer Statement2, defer Statement3. The orde of execution becomes: Statement3, Statement2, Statement1

// Recover: How to recover during a panic? Basically encountering a panic will lead your program to crash from that point.
// A recover can stop a panic from aborting the program and let it continue with it's execution, hence allowing us to handle
// the error for which we did not expect to occur or we did not handle.
// Example: a server wouldn’t want to crash if one of the client connections exhibits a critical error. Instead, the server
// would want to close that connection and continue serving other clients. In fact, this is what Go’s net/http does by default
// for HTTP servers.

// recover() must be called within a defer function, when the "enclosing" function panics, then the call goes to defer and recover
// call is made to handle the panic, it catches the error returned from panic.
// Return value: The error raised in the call to panic

func iWillPanic() {
	panic("Sorry!, something wrong occurred which must not have happened")
}

func main() {
	fmt.Println("Learning defer, panic and recover in golang")

	defer func() {
		fmt.Println("Inside defer call")
		if rec := recover(); rec != nil {
			fmt.Println("Recovering from a panic, the panic message is", rec)
		}
	}()

	iWillPanic()

	fmt.Println("Did the function call went well??, If I am printed, everything is good. Yay!!")
}
