package main

import (
	"net/http"
	"log"
	"fmt"
	"sync"
)

var count int = 0
var mu sync.Mutex
func main() {
	http.HandleFunc("/",handler1)
	http.HandleFunc("/hello",handler2)
	http.HandleFunc("/counter",counter)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}

func handler1(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "handler1 URL.Path = %q\n", r.URL.Path)
	print(w,r)
	count++
	mu.Unlock()
}

func handler2(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "handler2 URL.Path = %q\n", r.URL.Path)
	print(w,r)
	count++
	mu.Unlock()
}

func print(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}


func counter(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Count = %d\n", count)
}