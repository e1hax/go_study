package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		//用go创建承载一个形参为空，返回值为空的一个函数
		go func() {
			defer fmt.Println("A.defer")

			//return //只会执行	A.defer
			func() {
				defer fmt.Println("B.defer")
				//return //不会执行 B，其它正常执行
				runtime.Goexit() //退出当前的 goroutine，不会退出外面的goroutine
				fmt.Println("B")
			}() //加()表示执行匿名函数

			fmt.Println("A")
		}()
	*/

	//go 创建一个有形参，有返回值的函数的协程
	//如果要获取到该函数的返回值，需要通过channel来实现
	go func(a int, b int) bool {
		fmt.Println("a = ", a, " b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
