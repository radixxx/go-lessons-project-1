package main

import "fmt"

func main() {
	aMap := make(map[string]int, 0)
	aMap["r1"] = 12
	aMap["r2"] = 13
	aMap["r2"] = 14
	aMap["r2"] = 15
	aMap["r2"] = 16
	aMap["r2"] = 17
	fmt.Println("map : ", aMap)

	anotherMap := map[string]int{
		"p1": 3,
		"p2": 4,
		"p3": 5,
		"p4": 6,
	}

	fmt.Println("anotherMap: ", anotherMap)
	delete(anotherMap, "p1")
	delete(anotherMap, "p1")
	delete(anotherMap, "p1")
	fmt.Println("anotherMap: ", anotherMap)

	_, ok := aMap["doesIsExist"]
	if ok {
		fmt.Println("Exists !")
	} else {
		fmt.Println("Does not exists")
	}

	for key, value := range aMap {
		fmt.Println(key, value)
	}
}
