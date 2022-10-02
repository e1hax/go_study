package main

import (
	"fmt"
)

// 一个匿名返回参数的函数
func foo1(a string, b int) int {
	fmt.Println("-----foo1-----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

//两个匿名返回参数的函数
func foo2(a string, b int) (int, int) {
	fmt.Println("-----foo2-----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

//两个返回参数的函数
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("-----foo3-----")

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//r1,r2属于 foo3的形参，初始化默认值为0
	//r1,r2作用域是 foo3 函数体
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	//给有名称的返回参数赋值
	r1, r2 = 100, 200

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("abc", 888)
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)

	ret3, ret4 := foo3("abc", 222)
	fmt.Println("ret3 = ", ret3)
	fmt.Println("ret4 = ", ret4)

}
