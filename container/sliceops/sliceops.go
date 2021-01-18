package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s = %d,len = %d,cap = %d \n", s, len(s), cap(s))
}

func main() {
	var s []int //Zero value for slice is nil
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
		printSlice(s)
	}
	fmt.Println(s)
	s1 := []int{1, 2, 3, 4}
	// 创建长度自定义的slice
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)

	fmt.Println("copying slice")
	copy(s2, s1)
	printSlice(s2)
	copy(s3, s2)
	printSlice(s3)
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)

	fmt.Println("poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)
}
