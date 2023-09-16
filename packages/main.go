package main

// By convention packages are named to be the same as the folder they are located in.

// Go determines if an item can be called by code in other packages through how it is declared.
// To make a function, type, variable, constant or struct field externally visible (known as exported)
// the name must start with a capital letter.

// If you import the package use it, else it will be removed
// In order for go to track the dependencies and import packages from the directoy use go.mod file
// -> go mod init github.com/naveen19991124/packages
// -> go mod tidy

// When a package is imported, only entities (functions, types, variables, constants) whose names
// start with a capital letter can be used / accessed.

// Identifiers will be named using camelCase, except for those meant to be accessible across packages which
// should be PascalCase.
import (
	"fmt"
	"math"

	"github.com/naveen19991124/packages/strutil"
)

func main() {
	fmt.Println("Learning packages")
	fmt.Println(math.Floor(2.7), math.Ceil(2.7), math.Sqrt(17))
	fmt.Println(strutil.Reverse("hello"))
}
