package main

import "fmt"

// An interface type is a set of method signatures.

// In order to implement an interface a type/struct needs to implement all it's methods, there is nothing like
// implements or any other keyword, if the struct/type implements all the methods that are part of the interface
// it implicitly implements the interface.

// An interface is useful for providing useful abstraction and dependency inversion, as any type can take place which
// implements the interface.

// Special interface: Empty interface -> It is defined using "any keyword" or just interface{}, since it has no methods,
// so every type in go implicityl implements it, hence any function which can accept any type of value or return any type
// of value can define this as the input/output type.

// For implementing an interface it does not matter if you add receiver methods using pass by value or pass by reference,
// in the sense if you pass by value or pass by reference it is considered that the type is implementing the interface.
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
}
