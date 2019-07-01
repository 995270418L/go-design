package hello

import "fmt"


func Hello(name string) string {
	return fmt.Sprintf("hello %s", name)
}

func hello(name string){
	fmt.Printf("hello %s", name)
}

