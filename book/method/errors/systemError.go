package errors

import "fmt"

// 系统抛错
func test01() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	index := 10
	a[index] = 10
	fmt.Println(a)
}
func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		panic("半径不能为负")
	}
	return 3.14 * radius * radius
}

func test03() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束   （2）发生异常时
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}
func test04() {
	test03()
	fmt.Println("test04")
}
func SystemError() {
	test04()
}
