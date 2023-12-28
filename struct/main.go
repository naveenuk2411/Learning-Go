package main

import "fmt"

// Structs are just user defined types similar to classes
// Declared using type and struct keywords
// type Shape struct {fieldName fieldType fieldName fieldType}

// NOTE: Field names in structs follow the Go convention - fields whose name starts with a lower case letter are only visible
// to code in the same package, whereas those whose name starts with an upper case letter are visible in other packages.
type Shape struct {
	name string
	size int
}

// type Shape struct {
// 	name, typeOfShape string
// 	size, radius int
// }

// Just an ordinary function similar to a constructor call, for creating an instance of a Shape
// Usefule for:
// 1. Validating the passed values.
// 2. Handling of default values
// 3. Instead of exporting the struct, we can export a function, and thus can initialize the required fields.
func NewShape(name string, size int) Shape {
	return Shape{
		name: name,
		size: size,
	}
}

func main() {
	fmt.Println("Learning structs")
	// Omitted files are 0 values
	shape1 := Shape{}
	// You can specify the name of the properties of struct when initializing
	shape2 := Shape{
		name: "Rect",
		size: 30,
	}
	// You can choose not to specify the names of the properties, but order should be maintained.
	shape3 := Shape{
		"Square",
		50,
	}
	fmt.Println(shape1)
	// Structs are mutable
	shape1.name = "Circle"
	shape1.size = 30
	fmt.Println(shape1)
	fmt.Println(shape2)
	fmt.Println(shape3)

	shape4 := NewShape("Oval", 50)
	fmt.Println(shape4)

	// Anonymous struct type
	// If a struct type will be used just one time, i.e it will just have a single value, then not required,
	// to give that struct a name, the value can have anonymous struct type.
	dog := struct {
		name string
		age  int
	}{
		name: "Tom",
		age:  12,
	}

	var shapePointer *Shape = &shape4
	// Go autmatically dereferences the pointer to struct type for us, so we can directly use the dot notaion to access the
	// properties of the struct instead of using * for dereferencing.
	fmt.Println(shapePointer, shapePointer.name, shapePointer.size)

	fmt.Println("Dog struct", dog)
}
