package main

import (
	"encoding/json"
	"fmt"
)

// type character struct {
// 	First string
// 	Last  string
// 	Age   int
// }

// If you want the parameters to be in lower case, you need to specify it manually
type character struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func mainJson() {
	marshalJson()
	unmarshalJson()
}

func marshalJson() {
	c1 := character{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	c2 := character{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	characters := []character{c1, c2}
	fmt.Println(characters)

	bs, err := json.Marshal(characters)

	if err != nil {
		fmt.Println("Err - ", err)
	}

	fmt.Println("Marshal - ", string(bs))
}

func unmarshalJson() {
	var characters []character
	// or characters := []person{}

	content := []byte(`[{"first":"James","last":"Bond","age":32},{"first":"Miss","last":"Moneypenny","age":27}]`)

	err := json.Unmarshal(content, &characters)

	if err != nil {
		fmt.Println("Err - ", err)
	}

	fmt.Println(characters)
}
