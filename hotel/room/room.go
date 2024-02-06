package room

import "fmt"

const roomText = "%d : %d people / %d nights"

func PrintDetails(roomNumber, size, nights int) {
	fmt.Println(roomNumber, ":", size, "people /", nights, "nights")
}
