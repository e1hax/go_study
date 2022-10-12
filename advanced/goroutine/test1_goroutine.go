package main

import (
	"fmt"
	"time"
)

//子 goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

//主 goroutine
func main() {
	//创建一个go程，去执行newTask()流程
	go newTask()

	//-------第二次例子-------
	// main goroutine退出，子goroutine也直接退出
	fmt.Println("main goroutine exit。。。")

	//-------第一次例子-------
	//主goroutine与 子goroutine交替执行
	/*
		i := 0
		for {
			i++
			fmt.Printf("main goroutine i=%d\n", i)
			time.Sleep(1 * time.Second)
		}
	*/
}
