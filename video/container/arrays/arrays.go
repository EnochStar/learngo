package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4, 5}

	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	//for i := 0; i < len(arr3); i++ {
	//	fmt.Print(arr3[i])
	//}
	printArray(&arr1)
	fmt.Println(arr1[0])
}

// 数组为值类型 拷贝数组 与 java不同
func printArray(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
	fmt.Println(arr[0])
}
