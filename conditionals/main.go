package main

import "fmt"

// If statements can also include a short initialization statement that can be used to initialize one or
// more variables for the if statement. Note: any variables created in the initialization statement go out
// of scope after the end of the if statement.

// NOTE: if condition {

// } else {   -> else / else if should always come in the same line as the closing bracket

// }

func main() {
	fmt.Println("Learning Conditionals")
	// >, >=, ==, <, <=, !=
	// In case of strings, lexicographical comparison
	// "apple" < "banana" -> true

	var value string
	if value == "val" {
		fmt.Println("Equals val")
	} else {
		fmt.Println("Not equals val")
	}

	var number int
	var result string = "The resultant number is "
	if number > 0 {
		result += "positive"
	} else if number < 0 {
		result += "negative"
	} else {
		result += "zero"
	}
	fmt.Println(result)

	num := 7
	if val := num * 2; val > 10 {
		fmt.Println("val is greater than 10")
	} else {
		fmt.Println("val is less than 10")
	}

	operatingSystem := "Windows"
	res := 10

	// Instead of large if else if, use switch against a value or expression
	switch operatingSystem {
	case "Windows":
		{
			res *= 10
			res *= 10
			if res >= 100 {
				res /= 10
			}
			fmt.Println(res)
		}
	case "Linux":
		fmt.Println("Linux")
		fmt.Println("Operating system is Linux")
	default:
		fmt.Println("Default")
	}

	age := 100
	// When not using a value or expression for a switch, you can evaluate booleans for cases
	switch {
	case age >= 10 && age < 100:
		fmt.Println("age is less than 100")
	case age >= 100:
		fmt.Println("age is greater than or equal to 100")
	default:
		fmt.Println("default")
	}
}
