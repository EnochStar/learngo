package main

import (
	"fmt"
)

//append ：向 slice 尾部添加数据，返回新的 slice 对象。
func newSlice() {
	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1)

	s2 := append(s1, 1)
	fmt.Printf("%p\n", &s2)

	fmt.Println(s1, s2)
}
func main() {
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)
	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)
	s := append(a, b...)
	fmt.Printf("slice s : %v\n", s)
	fmt.Printf("slice c : %v\n", s)
	d := append(s, 7)
	fmt.Printf("slice d : %v\n", d)
	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)
	newSlice()
}
