package watchdog

import (
	"time"
	"os"
)

type  watchdog struct{
	rootPath string
	errorList map[int]string
	fileModTime map[string]time.Time
	fileObjects map[string]os.File
}

func newInit(path string) watchdog{
	var watcher = watchdog{}
	watcher.rootPath = path
	watcher.errorList = make(map[int]string)
	watcher.fileModTime = make(map[string]time.Time)
	watcher.fileObjects = make(map[string]os.File)
	return watcher
}

func (* watchdog) release(){
	watchdog.errorList = nil
	watchdog.rootPath = nil
	watchdog.fileObjects = nil
	watchdog.fileModTime = nil
}
