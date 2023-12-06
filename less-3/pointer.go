package main

import "fmt"

func main() {
	pointer()
}

func pointer() {
	a := "Hola Mudila"
	b := 4
	fmt.Println(a)
	fmt.Println(b)

	pointer := &a

	fmt.Print("Reference: ", pointer)
	fmt.Print("\nValue by reference: ", *pointer)

	*pointer = "\nGet up Mudila !" //updated value
	fmt.Print("\n=================================", a)

	c := 52
	pointer_c := &c //get refence

	fmt.Printf("\nNumber: %d", c)
	fmt.Println("\nPointer& : ", &pointer_c)
	fmt.Println("\nPointer* : ", *pointer_c)

	*pointer_c = 52 / 2 // working with value
	fmt.Println("\nDivision by 2: ", *pointer_c)
}
