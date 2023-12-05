package main

import "fmt"

const pi float32 = 3.14
const carMaxSpeed = 310

func main() {
	fmt.Println(pi, carMaxSpeed)
	fmt.Println(threeArguments())
	fmt.Println(nonDeclaredReturn())
	fmt.Println(loop())
	fmt.Println(anotherLoop())
	statement()
	swithStatement(2)
}

func threeArguments() (string, string, string) {
	a := "Hello"
	b := "programmer"
	c := "mundo"

	return a, b, c
}

func nonDeclaredReturn() (f string) {
	f = "Return Type !"
	return
}

func loop() (sum int) {
	sum = 0
	count := 10
	for i := 0; i < count; i++ {
		sum += i
	}
	return sum
}

func anotherLoop() (sum int) {
	// like while loop
	sum = 0
	for sum < 100 {
		sum += 5
	}
	return
}

func statement() {
	a := 1
	for a < 10 {
		if a == 5 {
			fmt.Println("*** a == 5!")
		} else {
			fmt.Printf("\na == %d", a)
		}
		a++
	}
}

func isTrue(a int) int {
	if a == 3 {
		return 1
	} else if a == 5 {
		return 2
	}
	return 0
}

func swithStatement(x int) {
	switch x {
	case 1:
		fmt.Println("\nShow: 1")
	case 2:
		fmt.Println("\nShow: 2")
	default:
		fmt.Print("\nby Default !")
	}
}

// TODO: implement func with if statements
func resolve() {}
