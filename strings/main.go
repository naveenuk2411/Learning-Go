package main

import (
	"fmt"
	"strings"
)

// There are significant differences between strings in golang and strings in different languages.
// In Go, strings are made of "characters", where the concept of charactes is a rune.
// Strings are essentially read-only slice of bytes.

// What is a rune ?
// It is an integer that represents a Unicode code point.


func main() {
	fmt.Println("Learning strings in golang")
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
