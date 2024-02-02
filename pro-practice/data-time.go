package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("%s\n", now)

	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		panic(err)
	}

	nowChi := now.In(loc)
	fmt.Println("%s\n", nowChi)

	comparision()

	fmt.Println(time.Now().Format(time.RFC1123))
	fmt.Println(time.Now().Format(time.RFC3339))
}

func timeElapsecedTime() {
	start := time.Now()
	err := os.WriteFile("/tmp/thisIsATest", []byte("TEST"), 0777)
	if err != nil {
		panic(err)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("process took %s", elapsed)
}

func comparision() {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	firstFebruary1990 := time.Date(1990, 2, 1, 0, 0, 0, 0, location)

	timeToParse := "2024-02-02 15:10:35.264477 -0600 CST"
	t, err := time.Parse("2007-01-02T15:04-07:00", timeToParse)
	//TODO Fix this error !
	if err != nil {
		panic(err)
	}

	now := time.Now()
	if t.After(firstFebruary1990) && t.Before(now) {
		fmt.Println(t, "is Between", firstFebruary1990, "and", now)
	} else {
		fmt.Println("now in between")
	}

	now = now.Add(time.Hour * 5)
	now = now.Add(time.Minute * 5)
	now = now.Add(time.Microsecond * 5)
	now = now.Add(time.Millisecond * 5)
}
