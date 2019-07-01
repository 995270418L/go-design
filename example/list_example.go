package example

import (
	"container/list"
	"fmt"
	"container/ring"
)

func List_Demo(){
	l := list.List{}
	//l_n := list.New()
	l.PushFront("steve")
	fmt.Printf("%+v", l)
}

func Ring_example(){
	var r ring.Ring
	fmt.Printf("%+v", r)
}
