package main

import (
	"fmt"

	"cmp"
	"slices"
	"sort"
)

// Go's slices package implements sorting for builtins and user defined types
// Sorting functions work for any orderered type.
// Ordered types:
// type Ordered interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64 |
// 		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
// 		~float32 | ~float64 |
// 		~string
// }

type Person struct {
	name string
	age  int
}

// Custom comparator function to sort strings by their length instead of default lexicographical order
func stringByLenSort(str1, str2 string) int {
	return cmp.Compare(len(str1), len(str2))
}

type myCustomType []Person

// Implementing the interface of sort package on my own custom type for the Sort method to work
// We need to implement the following methods for implementing the interface in order for the Sort method to work
// 1. Len() int -> Returns the length
// 2. Less(i, j int) bool -> Custom comparator function
// 3. Swap(i, j) -> Swap i, j

func (customTypeVar myCustomType) Len() int {
	return len(customTypeVar)
}

func (customTypeVar myCustomType) Swap(i, j int) {
	customTypeVar[i], customTypeVar[j] = customTypeVar[j], customTypeVar[i]
}

func (customTypeVar myCustomType) Less(i, j int) bool {
	return customTypeVar[i].name < customTypeVar[j].name
}

func main() {
	stringSlice := []string{"cat", "bulldog", "ant"}
	numbersSlice := []int{3, 1, 2}
	personsSlice := []Person{{
		name: "Sudhanshu",
		age:  34,
	}, {
		name: "Akshay",
		age:  43,
	}, {
		name: "Gagan",
		age:  12,
	}}

	fmt.Println(stringSlice)
	slices.Sort(stringSlice)
	slices.Reverse(stringSlice)
	fmt.Println(stringSlice)

	fmt.Println(numbersSlice)
	slices.Sort(numbersSlice)
	fmt.Println(numbersSlice)

	fmt.Println(stringSlice)
	slices.SortFunc(stringSlice, stringByLenSort)
	fmt.Println(stringSlice)

	newStringSlice := []string{"cat", "ant", "bulldog", "hey"}
	fmt.Println(newStringSlice)

	// Using comparator function along with stable sort
	slices.SortStableFunc(newStringSlice, stringByLenSort)
	fmt.Println(newStringSlice)

	// fmt.Println(personsSlice)
	// This will throw an error as a user defined type or a struct type is not a part of an ordered type interface which is
	// mentioend at the top of this file.
	// slices.Sort(personsSlice)
	// fmt.Println(personsSlice)

	fmt.Println(personsSlice)
	slices.SortFunc(personsSlice, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(personsSlice)

	if slices.IsSorted(numbersSlice) {
		fmt.Println("Numbers slice is already sorted")
	} else {
		fmt.Println("Numbers slice is not sorted, lol!")
	}

	indexPos, found := slices.BinarySearch(numbersSlice, 2)
	fmt.Println(indexPos, found)

	// The following code will demonstrate the sorting of slices using the "sort" package.
	myNumbersSlice := []int{24, 12, 5, 24}
	fmt.Println(myNumbersSlice)
	// Sorting ints
	sort.Ints(myNumbersSlice)
	fmt.Println(myNumbersSlice)

	myStringSlice := []string{"Naveen", "Gojo", "Yuta"}
	fmt.Println(myStringSlice)
	// Sorting strings
	sort.Strings(myStringSlice)
	fmt.Println(myStringSlice)

	myPersonSlice := []Person{{
		name: "Naveen",
		age:  24,
	}, {
		name: "Gojo",
		age:  31,
	}, {
		name: "Yuta",
		age:  21,
	}}

	fmt.Println(myPersonSlice)
	// Sorting user defined types with custom comparator function
	sort.Slice(myPersonSlice, func(i, j int) bool {
		return myPersonSlice[i].age < myPersonSlice[j].age
	})
	fmt.Println(myPersonSlice)

	myCustomPersonSlice := myCustomType(myPersonSlice)
	// Sorting type defs
	sort.Sort(myCustomPersonSlice)
	fmt.Println(myCustomPersonSlice)
}
