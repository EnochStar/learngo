package main

import (
	"fmt"
)

// data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8
// a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x
func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[6:8]
	fmt.Println(d1, len(d1), cap(d1))
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))
}
