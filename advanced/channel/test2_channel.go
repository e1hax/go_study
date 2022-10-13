package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)

	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))

	go func() {

		defer fmt.Println("子go程结束")

		//存数据
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程正在运行，发送元素：", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	//存完数据,延时2秒取数据
	time.Sleep(2 * time.Second)

	//取数据
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num=", num)
	}

	fmt.Println("main goroutine结束。。。")
}
