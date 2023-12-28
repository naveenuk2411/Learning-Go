package main

import "fmt"

// If statements can also include a short initialization statement that can be used to initialize one or
// more variables for the if statement. Note: any variables created in the initialization statement go out
// of scope after the end of the if statement.

// NOTE: if condition {

// } else {   -> else / else if should always come in the same line as the closing bracket

// }

// NOTE: There is no ternary in go, you will have to use basic if else.

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

	if 0*0 > 0 {
		fmt.Println("0 is greater than 0")
	} else {
		fmt.Println("0 is not greater than 0")
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
	} else if val < 0 {
		fmt.Println("val is less than 0")
	} else {
		fmt.Println("val is between 0 and 10")
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

	day := "Sunday"
	switch day {
	case "Sunday", "Saturday": // Or of matches
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	// Type switch
	// A type switch compares the types instead of values of the provided variable
	inlineFn := func(passedValue interface{}) {
		switch typeOfPassedValue := passedValue.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", typeOfPassedValue)
		}
	}
	inlineFn(true)
	inlineFn(1)
	inlineFn("hey")

}
