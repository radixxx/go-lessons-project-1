package main

import "fmt"

type Vechicle struct {
	Seats int
	Color string
}

// ! important
type Car struct {
	Vechicle
}

type MotorCycle struct {
	Base Vechicle
}

func main() {
	carImpl()
	motoImpl()
}

func carImpl() {
	car := &Car{
		Vechicle{
			Seats: 4,
			Color: "black",
		},
	}

	fmt.Println(car.Seats, car.Color)
}

func motoImpl() {
	motoCycle := &MotorCycle{
		Vechicle{
			Seats: 2,
			Color: "red",
		},
	}
	fmt.Println(motoCycle.Base.Seats, motoCycle.Base.Color)
}
