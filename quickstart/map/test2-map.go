package main

import (
	"fmt"
)

//通过引用传递
func printMap(citymap map[string]string){
	for key, value := range citymap {
			fmt.Println("key = ",key)
			fmt.Println("value = ",value)
	}
}

//通过引用传递，能够修改原本map的值	
func changeMap(cityMap map[string]string){
	cityMap["England"] = "London"
}

func main(){

	//声明一个map
	cityMap := make(map[string]string)

	//添加
	cityMap["Japan"] = "Tokyo"
	cityMap["China"] = "Beijing"
	cityMap["USA"] = "Washington, dc,"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap,"USA")
	fmt.Println("---------------")
	printMap(cityMap)

	//修改
	cityMap["Japan"]="TTT"
	changeMap(cityMap)
	fmt.Println("--------------")
	printMap(cityMap)
}