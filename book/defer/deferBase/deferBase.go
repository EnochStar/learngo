package deferBase

import "fmt"

//    defer 特性
//    1. 关键字 defer 用于注册延迟调用。
//    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
//    3. 多个defer语句，按先进后出的方式执行。
//    4. defer语句中的变量，在defer声明时就决定了。

//    defer 用途
//    1. 关闭文件句柄
//    2. 锁资源释放
//    3. 数据库连接释放

// defer对应是数据结构 堆栈
func DeferStack() {
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println(i)
	}
}

// 由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.
func DeferClosure() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

// 延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
func DeferDelay() {
	x, y := 10, 20

	defer func(i int) {
		println("defer: ", i, y) // 闭包引用
	}(x) // x被复制

	x += 10
	y += 100
	println("x = ", x, "y = ", y)
}
