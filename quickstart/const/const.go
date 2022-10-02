package main

import "fmt"

func main() {
	//定义常量
	const a int = 100
	

	//多常量定义
	const (
		a1 = 100
		a2 = "abc"
		a3 = true
	)

	//多常量定义，使用iota枚举,iota只能配合const使用
	const (
		BeiJing = iota
		ShangHai
		TianJin
	)

	fmt.Println("a = ", a)
	fmt.Printf("a type is %T\n", a)

	fmt.Println("a1 = ", a1)
	fmt.Printf("a1 type is %T\n", a1)
	fmt.Println("a2 = ", a2)
	fmt.Printf("a2 type is %T\n", a2)
	fmt.Println("a3 = ", a3)
	fmt.Printf("a3 type is %T\n", a3)

	fmt.Println("BeiJing is ", BeiJing)
	fmt.Println("ShangHai is ", ShangHai)
	fmt.Println("TianJin is ", TianJin)

}
