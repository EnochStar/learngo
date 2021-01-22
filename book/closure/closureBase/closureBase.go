package closureBase

import "fmt"

//闭包复制的是原对象指针
func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func AClosure() {
	var c = a()
	c()
	c()
	c()
	var c2 = a()
	c2()
	c2()
	c2()

}

func Test() func() {
	x := 100
	fmt.Printf("fis - > x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("second -> x (%p) = %d\n", &x, x)
	}
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}
func AddClosure() {
	temp1 := add(10)
	fmt.Println(temp1(5))
	fmt.Println(temp1(6))

	temp2 := add(10)
	fmt.Println(temp2(6))

}

// base 相当于该环境下的全局变量
func test01(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}
func Test01() {
	f1, f2 := test01(5)
	fmt.Println("f1 ---> ", f1(2))
	fmt.Println("f2 ---> ", f2(3))

}
