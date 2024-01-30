package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type MyJson struct {
	Cat `json:"cat"`
}

func main() {
	myJSON := []byte(`{"cat":{ "name":"Joey", "age":8}}`)
	c := MyJson{}
	err := json.Unmarshal(myJSON, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Age)
}
