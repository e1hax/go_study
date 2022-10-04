package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	// t := reflect.TypeOf(str).Elem()
	t := reflect.TypeOf(str)
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, ", doc: ", tagdoc)
	}

}

func main() {
	var re resume
	//传地址  上面方法中需要使用到.Elem()
	// findTag(&re)
	findTag(re)

}
