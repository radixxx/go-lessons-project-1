package main

import "fmt"

func main() {
	var name = "TOM"
	fmt.Print(name)

	var firstName, age = "Dann", 77
	var x = fmt.Sprintln("\n My name is %s. and age %d: ", firstName, age)
	fmt.Println(x)
}
