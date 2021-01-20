package main

//使用 make 动态创建slice，避免了数组必须用常量做长度的麻烦。还可用指针直接访问底层数组，退化成普通数组操作。

import "fmt"

func main() {
	s := []int{0, 1, 2, 3}
	s[0] = 20
	p := &s[2] // *int, 获取底层数组元素指针。
	*p += 100

	fmt.Println(s)
}
