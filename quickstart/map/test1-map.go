package main

import "fmt"

func main(){
	//==>第一种方式声明map
	//声明map的key是string类型，value是string类型
	var myMap1 map[string]string
	if myMap1 == nil{
		fmt.Println("myMap1 是一个空map")
	}

	//使用map前 需要先给map分配空间
	myMap1 = make(map[string]string,10)
	myMap1["one"]="java"
	myMap1["two"]="c++"
	myMap1["three"]="python"

	fmt.Println(myMap1)


	//==>第二种声明方式
	myMap2 := make(map[int]string)

	myMap2[1]="java"
	myMap2[2]="python"
	myMap2[3]="go"

	fmt.Println(myMap2)

	//==>第三种声明方式
	myMap3 := map[string]string{
		"one" : "c++",
		"two" : "java",
		"three" : "go",
	}

	fmt.Println(myMap3)


}