package main

import "fmt"

//声明一种类型，是int类型的别名
type myint int

//定义一个结构体
type book struct {
	title string
	auth  string
}

//传递一个副本
func changeBook(book book) {
	book.auth = "go"
	fmt.Printf("%v\n", book)
}

//引用传递
func changeBook2(book *book) {
	book.auth = "go1"
}

func main() {
	/*
		var a myint = 10
		fmt.Println("a = ", a)
		fmt.Printf("a type = %T\n", a)
	*/

	var book1 book
	book1.title = "GoLang"
	book1.auth = "Google"

	fmt.Printf("%v\n", book1)

	//调用changeBook函数
	changeBook(book1)
	fmt.Printf("%v\n", book1)
	fmt.Println("-------------")

	//调用changeBook2函数
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
	fmt.Println("hello")
}
