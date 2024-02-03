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
	timeElapsecedTime()
	iteration()

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
	fmt.Printf("process took %s", elapsed)
}

func comparision() {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	firstFebruary1990 := time.Date(1990, 2, 1, 0, 0, 0, 0, location)

	timeToParse := "2019-02-15T07:33-02:00"
	t, err := time.Parse("2006-01-02T15:04-07:00", timeToParse)
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

func iteration() {
	start, err := time.Parse("2006-01-02", "2022-02-19")
	if err != nil {
		panic(err)
	}
	end, err := time.Parse("2006-01-02", "2024-07-17")
	if err != nil {
		panic(err)
	}
	for i := start; i.Unix() < end.Unix(); i = i.AddDate(0, 0, 1) {
		fmt.Println(i.Format(time.RFC3339))
	}
}
