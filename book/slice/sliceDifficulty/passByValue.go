package main

import "fmt"

// 由内存地址可知，Go语言中无论是数组赋值还是函数传参
// 均为值传递
// 这样会导致每次数组传参 都有极大的内存消耗
func main() {
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}
