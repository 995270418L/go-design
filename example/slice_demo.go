package example

import "fmt"

func Test_one(){
	//s := []int{0,0,0,0,0}
	//s_array := [5]int{} // 数组
	s := make([]int, 5, 8)
	fmt.Printf("s 长度: %d \t s 容量: %d, \t s value : %v", len(s), cap(s), s)
}

func Test_two(){
	s_n := make([]int, 3, 8)
	s := s_n[3:8] // [3, 8) 切片索引
	s = s_n[0 : cap(s_n)]
	fmt.Printf("s 长度: %d \t s 容量: %d, \t s value : %v", len(s), cap(s), s)
}