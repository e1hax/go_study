package main

import "fmt"

func main() {

	//声明slice1是一个切片，并初始化，默认值是1,2,3，len是3
	// slice1 := []int{1, 2, 3}

	//声明一个slice，但是没有分配空间
	// var slice1 []int

	//声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值为0
	// var slice1 []int = make([]int, 3)

	//声明slice1是一个切片，同时给sllice1分配空间，3个空间，初始化值为0，通过:=推导出slice是一个切片
	slice1 := make([]int, 3, 5)

	fmt.Printf("len = %d ,cap=%d, slice=%v\n", len(slice1), cap(slice1), slice1)

	//当append的长度大于cap时，cap切片空间会直接在原来的基础上x2
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 3)
	fmt.Printf("len = %d ,cap=%d, slice=%v\n", len(slice1), cap(slice1), slice1)
}
