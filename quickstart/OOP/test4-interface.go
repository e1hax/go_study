package main

import "fmt"

//定义一个Animal接口，并声明三个方法
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

//定义一个Cat类
type Cat struct {
	color string
}

//实现Animal接口的全部方法，即为实现接口
func (this *Cat) Sleep() {
	fmt.Println("cat sleep...")
}

func (this *Cat) GetType() string {
	return "Cat.."
}

func (this *Cat) GetColor() string {
	return this.color
}

//定义一个Dog类
type Dog struct {
	color string
}

//实现Animal接口的全部方法，即为实现接口
func (this *Dog) Sleep() {
	fmt.Println("dog sleep...")
}

func (this *Dog) GetType() string {
	return "Dog.."
}

func (this *Dog) GetColor() string {
	return this.color
}

//show函数，接受实现Animal接口的类
func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("type = ", animal.GetType())
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("-----------")

}

func main() {

	//第一种
	var animal AnimalIF
	animal = &Cat{"green"}
	animal.Sleep()

	animal = &Dog{"brown"}
	animal.Sleep()

	fmt.Println("--------")

	//第二种
	cat := Cat{"green"}
	dog := Dog{"yellow"}

	showAnimal(&cat)
	showAnimal(&dog)

}
