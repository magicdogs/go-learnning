package main

import (
	"pkg/wtc"
)

var (
	 dogs wtc.Watchdog
)

func main() {
	//r := cal.Add(1,2)
	//fmt.Println("r= ",r)
	dogs = wtc.NewInit("d:/tmp/logs")
	dogs.ListError();

}


