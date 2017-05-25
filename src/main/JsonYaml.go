package main

import (
	"encoding/json"
	"fmt"
)


type Student struct {
	Name    string
	Age     int
	Guake   bool `true`
	Classes []string
	Price   float32 `16.5`
}


func main() {

	st := &Student {
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}

	b,_ := json.Marshal(st)

	js := string(b)
	fmt.Println(js)

	ft := &Student{}
	json.Unmarshal([]byte(js),&ft)
	fmt.Println(ft.Classes)


}
