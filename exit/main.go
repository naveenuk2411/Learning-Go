package main

import (
	"fmt"
	"os"
)

func willDeferWork() {
	// This defer will be called because it will be executed by main first, only the defers present within main
	// will not be called in prescence of exit
	defer fmt.Println("Defer inside func call")

	// If you enable to below exit, then again the whole program will exit with the given status without the above
	// defer to work
	// os.Exit(1)
}

// Note that unlike e.g. C, Go does not use an integer return value from main to
// indicate exit status. If you’d like to exit with a non-zero status you should use os.Exit.
func main() {
	willDeferWork()
	fmt.Println("Will I be able to print") // -> This will not be printed in presence of any exits called by methods or anything above it.
	defer fmt.Println("Will this defer work in presence of exit, No!!!")

	// Exit code 0 means success, any other exit code represents an error

	os.Exit(3)
}
