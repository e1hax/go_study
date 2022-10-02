package lib2

import "fmt"

//L大写，才表示该方法允许其它文件调用
func Lib2Test() {
	fmt.Println("Lib2Test()....")
}

func init() {
	fmt.Println("lib2,init()....")
}
