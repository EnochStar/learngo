package main

import "fmt"

//	  1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
//    2.指针变量的值是指针地址。
//    3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
func main() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
