package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/docopt/docopt-go"
	"bufio"
	"io"
	"os"
	"fmt"
	"net"
	"log"
)


func main() {
	usage := `Naval Fate.

	Usage:
	  run path <name>
	  run -c <conf>
	  run file shoot <x> <y>
	  run -h | --help
	  run --version

	Options:
	  -h --help     Show this screen.
	  --version     Show version.
	  --speed=<kn>  Speed in knots [default: 10].
	  --moored      Moored (anchored) mine.
	  --drifting    Drifting mine.`

	arguments, _ := docopt.Parse(usage, nil, true, "Naval Fate 2.0", false)


	conn,error := net.Dial("tcp","localhost:8000")
	if(error != nil){
		log.Fatal(error)
	}
	fmt.Println("hello word!")
	log := logs.NewLogger(10000)
	log.SetLogger("console", "")
	log.Info("info")
	log.Warn("warning")
	fmt.Println(arguments)
	var path string = arguments["<name>"].(string)
	log.Info(path)
	f,err := os.Open(path);
	if(err != nil){
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	var count int = 0
	for {
		buf,err := bfRd.ReadBytes('\n')
		if(err != nil){
			if(err == io.EOF){
				break
			}
			break
		}
		os.Stdout.Write(buf)
		conn.Write(buf)
		count ++
	}
	log.Info("count: {}",count)
}
