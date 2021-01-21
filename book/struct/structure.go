package main

import "fmt"

type people struct {
	name string
	age  int
	city string
}

// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，
// 所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int) *people {
	return &people{
		name: name,
		city: city,
		age:  age,
	}
}
func main() {
	p := newPerson("prof.cn", "测试", 90)
	fmt.Printf("%#v\n", p)
}
