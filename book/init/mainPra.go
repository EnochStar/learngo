package main

import (
	"fmt"
	_ "learngo/book/init/_hello"
)

// 	  1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
//
//    2 每个包可以拥有多个init函数
//
//    3 包的每个源文件也可以拥有多个init函数
//
//    4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
//
//    5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
//
//    6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
func main() {
	fmt.Println("main")
}
