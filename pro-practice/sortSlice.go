package main

import (
	"fmt"
	"sort"
)

type Person struct {
	person string
	height int
	weight int
}

func main() {
	personSlice := make([]Person, 0)
	personSlice = append(personSlice, Person{"Tom", 180, 71})
	personSlice = append(personSlice, Person{"Kate", 160, 55})
	personSlice = append(personSlice, Person{"Gabriel", 175, 64})
	personSlice = append(personSlice, Person{"John", 185, 80})

	fmt.Println(personSlice)

	sort.Slice(personSlice, func(i, j int) bool {
		return personSlice[i].height < personSlice[j].height
	})
	fmt.Println("<:", personSlice)
	sort.Slice(personSlice, func(i, j int) bool {
		return personSlice[i].height > personSlice[j].height
	})
	fmt.Println(">:", personSlice)
}
