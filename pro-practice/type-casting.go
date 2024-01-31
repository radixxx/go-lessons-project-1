package main

import (
	"fmt"
	"strconv"
)

func main() {
	numericCastring()
	stringCasting()
	bytesStringCasting()
}

func numericCastring() {
	var a int = 42
	f := float64(a)
	fmt.Println(f)

	pi := 3.144469
	i := int(f)
	fmt.Println(i, pi)

	ii := 71
	ff := float64(ii)
	fmt.Println(ff)

	c := 6 / 3
	d := 6.3 / 3

	fmt.Println(c, d)

}

func stringCasting() {
	var s string = "43"

	// convert string to int
	v, _ := strconv.Atoi(s)
	fmt.Println(v)

	// convert int to string
	var i int = 42
	str := strconv.Itoa(i)
	fmt.Println(str)
}

func bytesStringCasting() {
	var s string = "Hello M"
	var b []byte = []byte(s)

	fmt.Println(b)
	ss := string(b)
	fmt.Println(ss)
}
