package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
)

// An interface type is a set of method signatures. Or a named colletion of method signatures.
// Interface allows us to encapsulate the logic of the methods defined on the user defined types.
// Interface provides abstraction for higher level functions with guarantee on behaviour of the underlying concrete type, as if
// any concrete type is passed as the interface type, then it must have implemented the methods defined in the interface.

// Interfaces are implicitly implemented:
// In order to implement an interface a type/struct needs to implement all it's methods, there is nothing like
// implements or any other keyword, if the struct/type implements all the methods that are part of the interface
// it implicitly implements the interface.

// An interface is useful for providing useful abstraction and dependency inversion, as any type can take place which
// implements the interface.

// Special interface: Empty interface -> It is defined using "any keyword" or just interface{}, since it has no methods,
// so every type in go implicitly implements it, hence any function which can accept any type of value or return any type
// of value can define this as the input/output type.
// Some things to note about empty interfaces:
// 1. Empty interface gives us no knowledge about the type of data coming in.
// 2. Benefits of static type lang is nullified.
// 3. Need to use reflect library to turn arbitary structs into specific type.

// For implementing an interface it does not matter if you add receiver methods using pass by value or pass by reference,
// in the sense if you pass by value or pass by reference it is considered that the type is implementing the interface.
// But depending on if the receiver method is using pass by value, or pass by reference the underlying instance passed to receiver
// method gets affected accordingly.

// Some practices to follow when looking to define interfaces and use abstraction
// 1. We are not ideally looking into specifying deep levels of abstractions at the start of any source code.
// 2. We can define interfaces and bring abstraction into pitcure as and when required.
// 3. Being able to implicitly implement an interface allows us to define the required interfaces later, that get satisfied
// by exisiting concrete implementations defining different methods, without changing these concrete exisiting types. Here types
// refers to the user defined type which can have properties and receiver methods defined on them.
// 4. Keep interfaces simple and short.
// 5. Define an interface when there are 2 more concrety types (again referring to user defined types) that share the same behaviour
// functionalities in some way
// 6. Create smaller interfaces with fewer, simple methods.
type CarDriver interface {
	drive(int)                 // -> Methods can just define the types which they expect to be implemented.
	setModel(modelName string) // -> For more readability, the methods can also add a paramater name along with it's type
	getDistance() int
	reset()
}

type Car struct {
	speed    int
	distance int
	model    string
}

func (car *Car) drive(distance int) {
	car.distance += distance
}

// Again if you pass just by value, the model name for the car will not take any effect.
func (car *Car) setModel(modelName string) {
	car.model = modelName
}

func (car *Car) getDistance() int {
	return car.distance
}

func (car *Car) reset() {
	car.distance = 0
	car.speed = 0
	car.model = ""
}

// Metal - mass and volume information
type Metal struct {
	mass   float64
	volume float64
}

// Density - return density of metal
func (m Metal) Density() float64 {
	return m.mass / m.volume
}

// Gas - measurements for Gas
type Gas struct {
	pressure      float64
	temperature   float64
	molecularMass float64
}

// Density - return density of liquid
func (g Gas) Density() float64 {
	var density float64
	density = (g.molecularMass * g.pressure) / (0.0821 * (g.temperature + 273))
	return density
}

// Instead of defining two repetitive methods isDenser(a, b Metal) and isDenser(a, b Gas), we can use abstraction by defining an
// interface Dense which allows the types Metal, Gas to implement this interface having method density, and hence now we can just
// define a single method IsDenser which relies on the abstracted interface type i.e Dense, on which we can call the underlying concrete
// implementation of density.

// Dense - interface
type Dense interface {
	Density() float64
}

// IsDenser - compare density of two objects
func IsDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

// A function taking the empty interface, hence can receive a value of any type.
func describeValue(value interface{}) {
	// We are using type switch(a special type of switch) which helps us compare the types and process them.
	// We can also get the underlying concrete value i.e the dynamic value of the interface which is wrapped by the empty interface.
	fmt.Println("Conrete type of the empty interface", reflect.TypeOf(value))
	fmt.Println("Concrete value of the empty interface", reflect.ValueOf(value))
	fmt.Println("Zero value of the empty interface", reflect.Zero(reflect.TypeOf(value)))
	switch receivedConcreteValue := value.(type) {
	case int:
		fmt.Println("We received an int type of value", value)
	case string:
		fmt.Printf("We received a string type of value %s with length %d\n", value, len(receivedConcreteValue))
	default:
		fmt.Println("We do not know how to evaulate this type")
	}
}

// How go lang standard packages take advantages of abstraction using interfaces.
// Fprintf function takes in io.Writer interface that allows it to deal with any data type as input that has implemented the interface
// io.Writer, and hence will be able to call the method Write on the io.Writer interface without worrying on the specific concrete type
// that gets passed to it.
// io.Writer interface provides an abstraction to how different types can defined their own Write methods to control how bytes are
// written.

// type Writer interface {
// 	Write(value []byte) (n int, err error)
// }

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type Closer interface {
// 	Close() error
// }
// In order to implement this interface you will have to implement both the Reader and Writer interface.
// type ReadWriter interface {
// 	Reader,
// 	Writer
// }

// Another standard interface is Stringer interface that is used by Println method to format the type and print in the desired string format.
// type Stringer interface {
// 	String() string
// }

type ByteCounter int

// Defining a write method on our custom type ByteCounter, having the definition that is required io.Writer interface.
// Hence now our custom type ByteCounter will start implementing the io.Writer interface and can be safely passed to fmt.Fprintf(),
// which expects the arg to be of type io.Writer interface.
func (byteCounter *ByteCounter) Write(value []byte) (n int, err error) {
	*byteCounter += ByteCounter(len(value))
	return int(*byteCounter), nil
}

type user struct {
	name  string
	email string
}

// TODO: Implement custom formating for user struct values.
// We are defining the method String() string on our custom type user, so that it implements the Stringer interface which is used by
// format package's Println method to print any type that implements the stringer interface.

func (usr user) String() string {
	return fmt.Sprintf("name: %s, email: %s", usr.name, usr.email)
}

// Interface satisfaction
// A user defined type satisfies an interface if it implements all the methods defined in the interface. The type can implement multiple
// interfaces, or an interface that embeds multiple interfaces.

// Can we assign some value to an interface type or a var declared of interface type??
// Yes. We can!!!.

// Assignability rule:
// An expression can be assigned to an interface only if it's type satisfies the interface.

// Wrapping of the concrete type and concealing it's methods:
// When you assign an expression or a instance type to interface value, then it wraps this concrete type, and only the methods that are
// defined by this interface are revealed and can be called, irrespective of if this concrety type has multiple methods or has implemented
// multiple interfaces.

// Example:
// printer(os.Stdout, "hello")

// Here calling Close() will throw an error because of the way interfaces wrap around the concrete types.
// Here we pass os.Stdout which is an instance of os.File type, and has access to methods of Write(), Read(), and Close() as os.File
// implements io.Writer, io.Read, io.Close interfaces.

// But when we pass the os.Stdout type to a io.Writer interface type, it wraps the underlying concrete implementation of os.Stdout, and
// only exposes the methods that the interface exposes irrespective of the many extra methods that are defined on the os.Stdout/underlying
// concrete implementation.
// func printer(w io.Writer, str string) {
// w.Write([]byte(str))
// w.Close() // -> This will throw error as the io.Writer interface type has no method Close defined on it.
// w.Close undefined (type io.Writer has no field or method Close)
// }

// Type assertion in interfaces:
// Since interfaces can be assigned values of any type that implements that interface, and since it wraps the concrete type/value, and
// only allows the methods defined in the interface to be called, what if we want the actual underlying value which was assigned to this
// interface type i.e the concrete value which can be any one of the types which implements this interface ??

// Ans: Using type assertion

// Using type assertion, we can extract out the underlying concret value that was passed to the interface value.
// Syntax: concreteValue, ok := myInterface.(ConcreteType)
// This checks if the underlying concrete dynamic value which is passed to the interface matches the type of the concrete type/user-defined type
// then it gives back the actual instance of that type that was passed to this interface.
// Else it usually panics, for that, we can have an ok bool variable which tells us if the match was successful or not, instead of panicking.
// Based on the value of ok bool variable we can choose to call the methods that were defined on the ConcreteType, even of the other interfaces, or
// the other methods which where not a part of this interface as well.

type circle struct {
	radius float64
}
type triangle struct {
	a, b, c float64 // lengths of the sides of a triangle.
}
type rectangle struct {
	h, w float64
}

type shape interface {
	area() float64
}

func (c circle) String() string {
	return fmt.Sprint("Circle (Radius: ", c.radius, ")")
}
func (t triangle) String() string {
	return fmt.Sprint("Triangle (Sides: ", t.a, ", ", t.b, ", ", t.c, ")")
}
func (r rectangle) String() string {
	return fmt.Sprint("Rectangle (Sides: ", r.h, ", ", r.w, ")")
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t triangle) area() float64 {
	// Heron's formula
	p := (t.a + t.b + t.c) / 2.0
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (r rectangle) area() float64 {
	return r.h * r.w
}

func (t triangle) angles() []float64 {
	return []float64{angle(t.b, t.c, t.a), angle(t.a, t.c, t.b), angle(t.a, t.b, t.c)}
}

func angle(a, b, c float64) float64 {
	return math.Acos((a*a+b*b-c*c)/(2*a*b)) * 180.0 / math.Pi
}

func main() {
	fmt.Println("Learning interfaces")

	mercedesCar := Car{
		speed:    400,
		distance: 500,
		model:    "S-Class",
	}
	fmt.Println(mercedesCar)
	mercedesCar.drive(1000)
	fmt.Println(mercedesCar)
	mercedesCar.setModel("A-Class")
	fmt.Println(mercedesCar)
	fmt.Println(mercedesCar.getDistance())
	mercedesCar.reset()
	fmt.Println(mercedesCar)

	describeValue(45)
	describeValue("Hello")
	describeValue(true)

	gold := Metal{478, 24}
	silver := Metal{100, 10}

	// result := isDenser(gold, siver) will also work
	// Reason: IsDenser method in turn calls the method density on type Metal has a value receiver type instead of a value receiver type.
	// Here go autmatically dereferences the pointer when calling the method defined on Metal type.
	// But do keep in mind the values will still be passed by value type and not reference type no matter what type(value or pointer), you
	// pass to the method when calling it on the custom type.
	// Go will automatically dereference the pointer type and pass it as value to the receiver on the method.

	// However if it's a pointer receiver type, then you need to explicity pass the addresses of the Metal instance, when calling the
	// density method.
	result := IsDenser(&gold, &silver)
	if result {
		fmt.Println("gold has higher density than silver")
	} else {
		fmt.Println("silver has higher density than gold")
	}

	oxygen := Gas{pressure: 5,
		temperature:   27,
		molecularMass: 32}

	hydrogen := Gas{pressure: 1,
		temperature:   0,
		molecularMass: 2}

	result = IsDenser(oxygen, hydrogen)

	if result {
		fmt.Println("oxygen has higher density than hydrogen")
	} else {
		fmt.Println("hydrogen has higher density than oxygen")
	}

	var buf bytes.Buffer
	var countBytes ByteCounter
	fmt.Fprintf(&countBytes, "hello world")
	fmt.Println(countBytes)
	fmt.Fprintf(os.Stdout, "hello \n")
	fmt.Fprintf(&buf, "world")

	user := user{
		name:  "Gojo Satoru",
		email: "nil",
	}
	fmt.Println(user)

	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	// w = time.Second -> This will throw an error, as the type time.Second has not implemented the io.Writer interface i.e it does
	// not have the implementation of Write() method on it's type.
	fmt.Println(w)

	shapes := []shape{
		circle{
			radius: 32,
		},
		triangle{
			a: 20,
			b: 40,
			c: 50,
		},
		rectangle{
			h: 10,
			w: 20,
		},
	}
	for _, shape := range shapes {
		fmt.Println(shape)
		fmt.Println(shape.area())
		// Since angles() method is specific to underlying concrete implementation of triangle type instead of the shape interface
		// which wraps the concrete implementation and only exposes methods which are available to the interface i.e area().

		// Hence using type assertion, we get the dynamic value of the interface and check if the type of that dynamic value in the
		// interface matches with the triangle type, and if yes we get that underlying concrete implmentation i.e the concrete type
		// of triangle and call the angles() method defined on it.
		if concreteType, ok := shape.(triangle); ok {
			fmt.Println(concreteType.angles())
		}
		fmt.Println(shape)
		fmt.Println("Conrete type", reflect.TypeOf(shape))
		fmt.Println("Concrete value", reflect.ValueOf(shape))
		fmt.Println("Zero value", reflect.Zero(reflect.TypeOf(shape)))
	}
}
