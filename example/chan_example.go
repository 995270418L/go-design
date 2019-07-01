package example

func Chan_Demo(){
	c := make(chan int, 1) // 有缓冲的通道，无缓冲！
	c <- 2  // <-
	example := <- c
	print(example)

	//time.Sleep(1 * time.Second)
}
