package main

import (
	"time"
	"fmt"
)

var watcher *time.Timer
var count int = 0

func main() {
	watcher = time.AfterFunc(1 * time.Second,callback)
	time.Sleep(20 * time.Second)
}

func callback(){
	count ++
	fmt.Println("callback run .... ",count)
	watcher.Stop()
	watcher = time.AfterFunc(3 * time.Second,callback)
}
