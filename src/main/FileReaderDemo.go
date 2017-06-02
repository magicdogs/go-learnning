package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
	"time"
)

func main() {
	fmt.Println("start read file : d:/tmp/1.txt")
	fso,error := os.Open("d:/tmp/1.txt")
	if(error != nil){
		log.Fatal(error)
	}
	reader := bufio.NewReader(fso)
	for {
		buf,err := reader.ReadBytes('\n')
		if(err != nil){
			if( err == io.EOF){
				time.Sleep(500 * time.Microsecond)
				continue
			}else{
				log.Fatal(err)
				break
			}
		}
		fmt.Println(string(buf))
	}

}
