package main

import "fmt"

//interface{} 万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// 区分interface{}的数据类型
	// interface{} 的 “类型断言” 机制
	value, ok := arg.(string)
	//判断arg 数据类型
	if !ok {
		fmt.Println("arg is not string type...")
	} else {
		fmt.Println("arg is string type , value = ", value)
		fmt.Printf("value type = %T\n", value)
	}

	fmt.Println("------------")
}

type Book struct {
	name string
}

func main() {
	book := Book{name: "Golang"}

	//传入一个Book类
	myFunc(book)
	//传入一个整数
	myFunc(100)
	//传入一个字符串
	myFunc("hello")
	//传入一个浮点型
	myFunc(3.1415)
}
