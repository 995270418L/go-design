package example

import "fmt"

func Map_Demo(){
	a_map := map[string]int{
		"steve": 22,  // , 不可少
	}
	a_map["demo"] = 23
	for k, v := range a_map {
		fmt.Printf("key : %s, \t value: %d \n", k, v)
	}
}

/**
  Go 语言字典的键类型不可以是函数类型、字典类型和切片类型。
 */
func Map_Demo_2(){
	//var a  map[string]int // 申明
	b_map := map[interface{}]int{
		"1": 1,
		[]int{2}: 2,
		3: 3,
	}
	for k, v := range b_map {
		fmt.Printf("key : %s, \t value: %d \n", k, v)
	}
}

func Map_Demo_3(){
	var c_map map[interface{}]int // make(map[string]int)
	print(c_map["steve"])
	print(c_map[nil])
	//c_map["steve"] = 4
	//print(c_map["steve"])
}