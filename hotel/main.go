package main

import (
	"fmt"
	"thisIsATest/hotel/occupancy"
	"thisIsATest/hotel/room"
)

func init() {
	fmt.Println("launch initialization")
}

func main() {
	fmt.Println("launch the program !")
	room.PrintDetails(123, 3, 2)
	occupancy.Level(21.3)
}
