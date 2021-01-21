package main

import "fmt"

//Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s 会动！\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 嵌套匿名结构
}

func (d *Dog) wang() {
	fmt.Printf("%s 汪汪汪\n", d.name)

}

func main() {
	d1 := Dog{
		3, &Animal{
			"taotao",
		},
	}
	d1.wang()
	d1.move()
}
