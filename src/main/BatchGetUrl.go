package main

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:]{
		go fetch(url,ch)
	}
	for  range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed \n",time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp,err := http.Get(url)
	if(err != nil){
		ch <-  fmt.Sprintf("fetch url error: %v",err)
		return
	}
	b,err := io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()
	if(err != nil){
		ch <- fmt.Sprintf("Copy error: %v",b)
		return
	}
	ch <-  fmt.Sprintf("url: %s , data: %7d , time: %.2fs",url,b,time.Since(start).Seconds())
}
