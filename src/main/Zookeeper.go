package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"pkg/assert"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	ListPath("/",c)
}

func ListPath(path string,c *zk.Conn){
	children,_,error := c.Children(path)
	assert.NotError(error)
	if(len(children) == 0){
		bt,_,error := c.Get(path)
		assert.NotError(error)
		fmt.Println(path)
		fmt.Println(string(bt))
		return
	}
	for _,n := range children{
		var fullPath string = ""
		if(path == "/"){
			fullPath =  path + n
		}else{
			fullPath = path + "/" + n
		}
		//println(fullPath)
		ListPath(fullPath,c)
	}

}

func Demo(){
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
