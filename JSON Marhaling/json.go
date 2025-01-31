package main

import (
	"encoding/json"
	"fmt"
)

// json marshal
type T struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
	Value4 string `json:"-"`
}

func main() {
	fmt.Println("json marshalling")
	str := T{"abc", "def", "", "klm"}

	jsonString, _ := json.Marshal(str) //second argument for ignoring the error

	// fmt.Println(jsonString)
	fmt.Printf("%s\n", jsonString)
}
