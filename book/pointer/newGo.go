package main

import "fmt"

//    func new(Type) *Type
//    1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
//    2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
//    相当于给一个引用变量初始化
func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}
