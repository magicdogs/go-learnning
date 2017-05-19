package main

import (
	"fmt"
	"pkg/cal"
	"pkg/wtc"
)

var (
	 dogs wtc.Watchdog
)

func main() {
	dogs = wtc.NewInit("d:")
	dogs.ListError();
	r := cal.Add(1,2)
	fmt.Println("r= ",r)
}


