package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello Mudila Bleat!")

	array := make([]byte, 2) //chunk size
	//array := make([]byte, 11) //chunk size

	for {
		n, err := r.Read(array)
		fmt.Printf("n = %d error = %v array = %v\n", n, err, array)
		fmt.Printf("array n bytes: %q\n", array[:])
		if err == io.EOF {
			break
		}
	}
}
