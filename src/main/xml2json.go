package main

import (
	"fmt"
	"strings"
	xj "github.com/basgys/goxml2json"
	"io/ioutil"
	"os"
)

func main() {
	// xml is an io.Reader
	if(len(os.Args) != 2){
		fmt.Println("Args must be xml file path .")
		os.Exit(-1)
	}
	buf,err := ioutil.ReadFile(os.Args[1])
	xml := strings.NewReader(string(buf))
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(json.String())
}
