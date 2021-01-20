package main

import (
	"fmt"
)

func main() {

	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	fmt.Printf("s : %v\n", s)
	fmt.Printf("s  len : %v\n", len(s))
	fmt.Printf("s cap : %v\n", cap(s))
	// 一次 append 两个值，超出 s.cap 限制。
	s = append(s, 100, 200)

	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。

}
