package main

import "fmt"

func main() {
	var pointerVar *int
	fmt.Println(*pointerVar) // to dereference pointer
}

func addCity(cities map[string]string) {
	cities["France"] = "Paris"
	fmt.Println(cities)
}
