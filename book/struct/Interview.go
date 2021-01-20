package main

import "fmt"

type student struct {
	name string
	age  int
}

// 实际上一直指向同一个指针，最终该指针的值为遍历的最后一个struct的值拷贝
func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		fmt.Printf(" stu.name : %v, &stu : %p, *stu: %v\n", stu.name, &stu, stu)
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
