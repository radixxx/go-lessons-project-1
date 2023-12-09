package main

import (
	"fmt"
	"unsafe"
)

func main() {
	pointer()
	unsafePointer()
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

func unsafePointer() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", p1)
	fmt.Println("*p2: ", p2)
	*p1 = 5434123412312431212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}
