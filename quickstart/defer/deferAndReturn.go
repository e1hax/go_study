package main

import "fmt"

//defer跟return，谁先谁后
func deferFunc() int {
	fmt.Println("deferFunc....")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc....")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

//结论：return的优先级大于defer
func main() {
	returnAndDefer()
}
