package main

import "fmt"

type Hero struct {
	Name  string
	Ad    string
	Level int
}

//函数名大写，表示该函数包外可以引用
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)

}

// 使用*Hero才表示引用传递，否则的话相当于值拷贝一个副本
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(Name string) {
	this.Name = Name
}

func main() {
	myHero := Hero{Name: "zhang3", Ad: "剑", Level: 100}

	myHero.Show()

	myHero.SetName("li4")

	myHero.Show()
}
