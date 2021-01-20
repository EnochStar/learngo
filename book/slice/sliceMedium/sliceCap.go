package main

import (
	"fmt"
)

// cap重新分配规律
func capRule() {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

//超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。

func capOut() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	fmt.Printf("s : %v\n", s)
	fmt.Printf("s  len : %v\n", len(s))
	fmt.Printf("s cap : %v\n", cap(s))
	// 一次 append 两个值，超出 s.cap 限制。
	s = append(s, 100, 200)
	fmt.Printf("s cap:%v\n", cap(s))
	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
}
func main() {

	capRule()

}
