package main

import (
	"flag"
	"awesomeProject/example"
)
var name  = flag.String("name", "steve", "programming owner") // Ctrl + 鼠标左键 查看方法的具体定义
var ip string
var port int
var port_val = flag.Int64("port_value", 8080, "port value demo")

func init(){
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip string address") // 传入的是地址
	flag.IntVar(&port, "port", 1, "port value")
	//flag.IntVar(&port_val, "port_val", 8080, "hh")
}

func main(){
	//flag.Parse()
	//fmt.Printf("Hello: %s!\n", *name)
	//fmt.Printf("ip address: %s\n", ip)
	//fmt.Printf("port: %d\n", port)
	//fmt.Printf("port_value: %d\n", *port_val)
	example.Chan_Demo()
}