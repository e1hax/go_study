package main

import "fmt"

func main() {
	//声明一个数组，并进行初始化
	s := []int{1, 2, 3}
	fmt.Println(s)

	//对s进行截取
	s1 := s[0:2]
	fmt.Println(s1)

	//修改截取后的值，原数组s也会改变
	s1[0]=100
	fmt.Println(s)
	fmt.Println(s1)

	//copy可以对数组进行拷贝
	fmt.Println("-----copy-----")
	s2 := make([]int,3)
	copy(s2,s)
	fmt.Println(s)
	fmt.Println(s2)

	//改变s2的值，s的值不会改变
	s2[1]=200
	fmt.Println(s)
	fmt.Println(s2)
}
