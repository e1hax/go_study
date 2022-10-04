package main

import "fmt"

//定义一个父类
type Human struct {
	Name string
	Sex  string
}

//父类eat方法
func (this *Human) Eat() {
	fmt.Println("Human eat()...")
}

//父类walk方法
func (this *Human) Walk() {
	fmt.Println("Human Walk()...")
}

//定义一个子类，继承父类
type SuperMan struct {
	Human //继承父类的属性
	Level int
}

//重写父类的eat方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan eat()...")
}

func (this *SuperMan) print() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Sex = ", this.Sex)
	fmt.Println("Level = ", this.Level)

}

func main() {
	//定义一个父类对象
	human := Human{Name: "zhang3", Sex: "male"}
	human.Eat()
	human.Walk()

	//定义一个子类对象
	superman := SuperMan{Human{Name: "zhang3", Sex: "male"}, 19}
	superman.print()

}
