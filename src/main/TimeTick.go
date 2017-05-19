package main

import (
	"time"
	"fmt"
	"net"
	"log"
	"io"
	"os"
)

func main() {
	tick := time.Tick(1 * time.Microsecond)

	for countdown := 10; countdown >= 0; countdown-- {
		fmt.Println("countdown: ",countdown,",tickVal: ",<- tick)
	}

	listenner,error := net.Listen("tcp","localhost:8000")
	if(error != nil){
		log.Fatal(error)
	}
	for {
		conn,error := listenner.Accept()
		if(error != nil){
			log.Fatal(error)
			continue
		}

		go handleConn(conn)

	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	io.WriteString(conn,"Server: bfe/1.0.8.18\r\n")
	io.WriteString(conn,"\r\n")
	io.WriteString(conn,"\r\n")
	go io.Copy(os.Stdout,conn)
	time.Sleep(1 * time.Second)
}
