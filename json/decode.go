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

/*

	“convert a variable into a coded form” = “marshaling”
	“convert a coded form into a variable” = “unmarshaling”
	“convert a data stream into a coded form” = “encoding”
	“convert a coded form into a data stream” = “decoding”

	To convert a variable into JSON/XML, we can use json.Marshal / xml.Marshal
	To convert JSON/XML into a variable, we can use json.Unmarshal / xml.Unmarshal

*/
