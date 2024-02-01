package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	s := [][]string{
		{"age", "genre", "name"},
		{"23", "M", "Hendrik"},
		{"65", "F", "Stephany"},
	}
	//Open file
	f, err := os.Create("data.csv")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	var buffer bytes.Buffer
	for _, data := range s {
		buffer.WriteString(fmt.Sprintf("%s, %s, %s\n", data[0], data[1], data[2]))
	}

	n, err := f.Write(buffer.Bytes())
	fmt.Println("%d bytes written\n", n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
