package main

import "fmt"

//交换两个变量的值
func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 100
	var b int = 200

	var p *int = &a

	//将a,b的地址传入swap函数
	swap(&a, &b)
	fmt.Println(p)
	fmt.Println("a = ", a, "\nb = ", b)

}
