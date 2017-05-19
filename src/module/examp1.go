package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"time"
)

var ums string = "=========================="

func main() {

	path,_ := filepath.Abs(os.Args[1:][0])
	fileInfos,_ := ioutil.ReadDir(path)
	fileNums := len(fileInfos)
	result := fmt.Sprintf("%d",fileNums)
	fmt.Println("ResultSize: ",result)
	var counts map[string]string = make(map[string]string)
	for _,fileInfo := range fileInfos{
		var tm time.Time =fileInfo.ModTime()
		year,mon,day := tm.Date()
		hour,min,sec := tm.Clock()
		var mx string = fmt.Sprintf("%d%d%d%d%d%d",year,mon,day,hour,min,sec)
		counts[mx] = fileInfo.Name()
	}

	fmt.Println(ums)
	for _,item := range counts{
		fmt.Println(item)
	}


	/*var count int = len(os.Args)
	fmt.Println("hello world ")
	fmt.Println(count)
	fmt.Println("sss" + os.Args[1:][0])
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])*/
}