package main

import (
	"fmt"
	"encoding/json"
)


type Student struct {
	Name    string
	Age     int
	Guake   bool `true`
	Classes []string
	Price   float32 `16.5`
}


func main() {


	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a)
	fmt.Println("&a = ",&a)
	fmt.Println("*&a = ",*&a)
	fmt.Println("b = ",b)
	fmt.Println("&b = ",&b)
	fmt.Println("*&b = ",*&b)
	fmt.Println("*b = ",*b)
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
	fmt.Println("x = ",x)


	stx := &Student {
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}


	fmt.Println("=====================================")
	fmt.Println(stx.Age)
	fmt.Println(&stx.Age)
	//fmt.Println(*stx)
	fmt.Println("=====================================")

	bx,_ := json.Marshal(stx)

	js := string(bx)
	//fmt.Println(js)

	ft := &Student{}
	json.Unmarshal([]byte(js),&ft)
	//fmt.Println(ft.Classes)


}
