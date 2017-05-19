package wtc

import (
	"time"
	"os"
	"fmt"
	"io/ioutil"
	"regexp"
)

type  Watchdog struct{
	rootPath string
	errorList map[int]error
	fileModTime map[string]time.Time
	fileObjects map[string]os.FileInfo
	syncMain chan string
	timerJob <- chan time.Time
}

func NewInit(path string) (result Watchdog){
	var watcher = Watchdog{}
	watcher.rootPath = path
	watcher.errorList = make(map[int]error)
	watcher.fileModTime = make(map[string]time.Time)
	watcher.fileObjects = make(map[string]os.FileInfo)
	watcher.syncMain = make(chan string)
	watcher.timerJob = time.Tick(10 * time.Second)
	//watcher.timerJob <- time.Now()
	fmt.Println("timeTick: ",(<- watcher.timerJob).Format("2006-01-02 15:04:05.999"))
	watcher.List(path)
	return watcher
}

func (p *Watchdog) GetSyncMain()(c chan string){
	return p.syncMain
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
				fs := path + "/" + fso.Name()
				reg := regexp.MustCompile(`.+\.log$`)
				if(reg.MatchString(fs)){
					fmt.Println(fs)
					p.fileObjects[fs] = fso
					p.fileModTime[fs] = fso.ModTime()
				}
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
