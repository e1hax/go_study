package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name   string   `json:"name"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 1994, 20, []string{"zxc", "zbz"}}

	//编码过程 结构体 ---> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json Marshal is error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码过程 json ----> 结构体
	//声明一个对象
	myMovie := Movie{}
	err2 := json.Unmarshal(jsonStr, &myMovie)
	if err2 != nil {
		fmt.Println("json Unmarshal is error", err2)
		return
	}

	fmt.Printf("%v", myMovie)
}
