package main

import (
	"io/ioutil"
	"fmt"
)

var (
	errorList map[int]error = make(map[int]error)
)

func main() {
	listpath("d:",errorList)
	for i:=0; i< len(errorList); i++ {
		fmt.Println("errorIndex: ",i," errorInfo: ",errorList[i])
	}
}

func listpath(path string,errorList map[int]error){
	root,error := ioutil.ReadDir(path)
	if(error != nil){
		//log.Fatal(error)
		errorList[len(errorList)] = error
		//fmt.Println("error : ",error)
	}else{
		for _,fso := range root{
			if(fso.IsDir()){
				listpath(path + "/" + fso.Name(),errorList)
			}else{
				fmt.Println(path + fso.Name())
			}
		}
	}

}
