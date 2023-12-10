package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const hotelName = "GoLang Hotel"
	const totalRooms = 134
	const firstRoomNumber = 110

	rand.Int63n(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := occupacyRate(roomsOccupied, totalRooms)
	occupancyLevel := occupancyLevel(occupancyRate)

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("                  Occupancy Level:", occupancyLevel)
	fmt.Printf("                   Occupancy Rate: %0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}
}

func showRoomDetails(roomNumber, size, nights int) {
	fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}

func occupancyLevel(occupancyRate float32) string {
	// 0% - 30% occup: Low
	// 30% - 60% occup: Medium
	// 60% - 100% occup:  High

	switch {
	case occupancyRate > 70:
		return "Hight"
	case occupancyRate > 20:
		return "Medium"
	default:
		return "Low"
	}
}

func occupacyRate(roomsOccupied, totalRooms int) float32 {
	return (float32(roomsOccupied) / float32(totalRooms)) * 100
}
