package main

import "fmt"

//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，
//而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

//make和new的区别：
//    1.二者都是用来做内存分配的。
//    2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//    3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
func main() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}
