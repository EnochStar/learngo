package main

import "fmt"

//不过传指针会有一个弊端，从打印结果可以看到，第一行和第三行指针地址都是同一个，万一原数组的指针指向更改了，那么函数里面的指针指向都会跟着更改。
//
//切片的优势也就表现出来了。用切片传数组参数，既可以达到节约内存的目的，也可以达到合理处理好共享内存的问题。
func main() {
	arrayA := [2]int{100, 200}
	testArrayPoint(&arrayA) // 1.传数组指针
	arrayB := arrayA[:]
	testSlicePoint(&arrayB) // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}

func testArrayPoint(x *[2]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}
func testSlicePoint(x *[]int) {
	fmt.Printf("func Slice : %p , %v\n", x, *x)
	(*x)[1] += 100
}
