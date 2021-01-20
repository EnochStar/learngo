package main

import "fmt"

//在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，
//否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
func main() {

	var a int
	a = 100
	fmt.Println(a)
	//var a *int
	//*a = 100
	//fmt.Println(*a)
	//
	//var b map[string]int
	//b["测试"] = 100
	//fmt.Println(b)
}
