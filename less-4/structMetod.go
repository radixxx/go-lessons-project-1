package main

import "fmt"

type Point struct {
	X, Y int
}

func (point Point) movePointBleat(x, y int) Point {
	point.X += x
	point.Y += y
	return point
}

func (p *Point) movePointPtr(x, y int) {
	p.X = x
	p.Y = y
}

func main() {
	p := Point{2, 3}
	fmt.Println(p)
	fmt.Println(p.movePointBleat(4, 7))
	fmt.Println(p)

	p.movePointPtr(2, 6)
	fmt.Println(p)

	v := &p
	fmt.Println(v.movePointBleat(3, 4))
	fmt.Println(p)
}
