package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	structures()

	p_objc := Point{
		X: 32,
		Y: 32,
	}

	p_objc.pStructureMethod()
	p_obj_reference := &p_objc

	p_obj_reference.pStructureMethod()

}

func (pMethod Point) pStructureMethod() {
	//TODO add logic
	fmt.Println(pMethod.X)
	fmt.Println(pMethod.Y)
}

func structures() {
	p := Point{
		X: 1,
		Y: 3,
	}
	fmt.Println(p)
	fmt.Println(p.X)
	fmt.Println(p.Y)

	p.X = 3
	fmt.Println(p)

	pp := Point{
		Y: 13,
	}
	fmt.Println(pp)

	g := &p
	fmt.Println("1. ", g)
	fmt.Println("2. ", &g)
	fmt.Println("3. ", *g)
}
