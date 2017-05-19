package wtc

import (
	"time"
	"os"
	"fmt"
	"io/ioutil"
)

type  Watchdog struct{
	rootPath string
	errorList map[int]error
	fileModTime map[string]time.Time
	fileObjects map[string]os.File
}

func NewInit(path string) (result Watchdog){
	var watcher = Watchdog{}
	watcher.rootPath = path
	watcher.errorList = make(map[int]error)
	watcher.fileModTime = make(map[string]time.Time)
	watcher.fileObjects = make(map[string]os.File)
	watcher.List(path)
	return watcher
}

func (p *Watchdog) Release(){
	p.errorList = nil
	p.rootPath = ""
	p.fileObjects = nil
	p.fileModTime = nil
}

func (p *Watchdog)List(path string){
	root,error := ioutil.ReadDir(path)
	if(error != nil){
		//log.Fatal(error)
		p.errorList[len(p.errorList)] = error
		//fmt.Println("error : ",error)
	}else{
		for _,fso := range root{
			if(fso.IsDir()){
				p.List(path + "/" + fso.Name())
			}else{
				fmt.Println(path + fso.Name())
			}
		}
	}

}


func Listpath(path string,errorList map[int]error){
	root,error := ioutil.ReadDir(path)
	if(error != nil){
		//log.Fatal(error)
		errorList[len(errorList)] = error
		//fmt.Println("error : ",error)
	}else{
		for _,fso := range root{
			if(fso.IsDir()){
				Listpath(path + "/" + fso.Name(),errorList)
			}else{
				fmt.Println(path + fso.Name())
			}
		}
	}

}

func (p *Watchdog)ListError(){
	for i:=0; i< len(p.errorList); i++ {
		fmt.Println("errorIndex: ",i," errorInfo: ",p.errorList[i])
	}
}
