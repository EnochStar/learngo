package main

import "fmt"

type persons struct {
	name string
	age  int
	city string
}

//通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
func main() {
	//var p = new(persons)
	//p.name = "Enoch"
	//p.age = 18
	//p.city = "北京"
	//fmt.Printf("p2 = %#v\n", p)

	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	p3 := &persons{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}
}
