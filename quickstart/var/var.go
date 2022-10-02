package main

import "fmt"

//声明全局变量不能使用方法四
//使用方法一，方法二声明全局变量
var gA int = 100
var gB = "bdcd"

//多变量声明

var gg, jj = 100, "hj"

var (
	gC = 300
	gD = 3.14
)

func main() {

	//声明变量
	//方法一： var 来声明变量
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("a type is %T\n", a)

	//方法二：var 声明变量并初始化
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("b type is %T\n", b)

	//方法三：初始化的时候，省去变量类型
	var c = 200
	fmt.Println("c = ", c)
	fmt.Printf("c type is %T\n", c)

	var s = "abcd"
	fmt.Println("s = ", s)
	fmt.Printf("s type is %T\n", s)

	//方法四：省去var关键字，直接自动匹配，只能在函数体中使用
	aa := 300
	fmt.Println("aa = ", aa)
	fmt.Printf("aa type is %T\n", aa)

	bb := "abcd"
	fmt.Println("bb = ", bb)
	fmt.Printf("bb type is %T\n", bb)

	//全局变量
	fmt.Println("gA = ", gA)
	fmt.Printf("gA type is %T\n", gA)
	fmt.Println("gB = ", gB)
	fmt.Printf("gB type is %T\n", gB)
	fmt.Println("gC = ", gC)
	fmt.Printf("gC type is %T\n", gC)
	fmt.Println("gD = ", gD)
	fmt.Printf("gD type is %T\n", gD)

	fmt.Println("gg = ", gg)
	fmt.Printf("gg type is %T\n", gg)

	fmt.Println("jj = ", jj)
	fmt.Printf("jj type is %T\n", jj)
}
