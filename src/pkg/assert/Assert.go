package assert

import "os"

func NotError(err error){
	if(err != nil){
		panic(err.Error())
		os.Exit(-1)
	}
}
