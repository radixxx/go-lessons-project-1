package main

import "fmt"

type Human struct {
	name     string
	lastName string
	age      int
	country  string
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GiveAffection(to Human)
}

func main() {
	var ion Human
	ion.name = "Ion"

	var cat Cat
	cat.name = "Mein-Katze"

	var dog Dog
	dog.name = "Mein-Hund"

	Pet(cat, ion)
	Pet(dog, ion)
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GiveAffection(human)
	animal.ReceiveAffection(human)
}

func (c Cat) ReceiveAffection(from Human) {
	fmt.Printf("The cat named %s has received affection from Human named %s\n", c.name, from.name)
}

func (c Cat) GiveAffection(to Human) {
	fmt.Printf("The cat named %s has given affection to Human named %s\n", c.name, to.name)
}

func (d Dog) ReceiveAffection(from Human) {
	fmt.Printf("The cat named %s has received affection to Human named %s\n", d.name, from.name)
}

func (d Dog) GiveAffection(to Human) {
	fmt.Printf("The dog named %s has given affection to Human named %s\n", d.name, to.name)
}
