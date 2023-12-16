package main

import (
	"fmt"
)

func main() {
	i := -10
	j := 25

	pointerI := &i
	pointerJ := &j

	fmt.Println("pI memory:", pointerI)
	fmt.Println("pI memory:", pointerJ)
	fmt.Println("pI value:", *pointerI)
	fmt.Println("pI value:", *pointerJ)
	fmt.Println("pI value:", &pointerI)
	fmt.Println("pI value:", &pointerJ)

	*pointerI = 99885577
	*pointerI--
	fmt.Println("i:", i)

	getPointer((pointerJ))
	fmt.Println("j:", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}
