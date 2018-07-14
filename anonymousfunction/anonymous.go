package main

import "fmt"

func main()  {
	//直接赋值
	f  := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1,2))
	//直接调用
	go func(ch chan int){
		ch <- 48
	}(47)
}
