package lib1

import "fmt"

//L大写，才表示该方法允许其它文件调用
func Lib1Test() {
	fmt.Println("Lib1Test()....")
}

//init函数
func init() {
	fmt.Println("lib1,init()....")
}
