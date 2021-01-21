package funcReturn

import "fmt"

// 返回值的名称应当具有一定的意义，可以作为文档使用。
// Golang 返回值不能用容器对象接收。只能用多个变量
// 命名返回参数可被同名局部变量遮蔽，此时需要显式返回。
func add(x, y int) (z int) {
	{ // 不能在一个级别，引发 "z redeclared in this block" 错误。
		var z = x + y
		// return   // Error: z is shadowed during return
		return z // 必须显式返回。
	}
}
func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}
func FunctionReturn() {
	var a, b int = 1, 2
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
}

func test() (int, int) {
	return 1, 2
}
func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}

	return x
}

//多返回值可直接作为其他函数调用实参。
func FunctionInUse() {
	println(add(test()))
	println(sum(test()))
}

func deferAdd(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

func FunctionDefer() {
	fmt.Println(deferAdd(1, 29))
}
