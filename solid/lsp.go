package main

import (
	"fmt"
)

// This principle states that objects of a superclass should be replaceable with objects of a subclass
// without affecting the correctness of the program. This helps to ensure that the relationships
// between classes are well-defined and maintainable.

type Animal struct {
	Name string
}

type Bird struct {
	Animal
}

func (a Animal) MakeSound() {
	fmt.Println("BLEAT animal sound !")
}

func (b Bird) MakeSound() {
	fmt.Println("Cirik Pizdik !")
}

type AnimalBehavior interface {
	MakeSound()
}

func MakeSound(ab AnimalBehavior) {
	ab.MakeSound()
}

func main() {
	a := Animal{}
	b := Bird{}
	MakeSound(a)
	MakeSound(b)
}
