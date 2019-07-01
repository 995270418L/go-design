package test

import (
	"testing"
	"awesomeProject/hello"
	"fmt"
)

// 功能测试函数
func TestDemo(t *testing.T){
	t.Log(hello.Hello("steve"))
}

// 性能测试函数
func BenchmarkHello(b *testing.B){
	for i:=0; i< b.N; i++{
		b.Log(hello.Hello("steve"))
	}
}

// 示例测试函数
func ExampleHello() {
	fmt.Println("hello, world")
	// Output: hello, world
}
