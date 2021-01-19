package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6]
	// slice相当于 arr的视图
	fmt.Println("arr[2:6]", arr[2:6])
	s1 := arr[2:]
	fmt.Println("s1", s1)
	fmt.Println("arr[:6]", arr[:6])
	s2 := arr[:]
	fmt.Println("s2", s2)
	//updateSlice(s2)
	//fmt.Println(s2)
	fmt.Println("after update")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println(s2)

	fmt.Println("after reslices")
	s2 = s2[:4]
	fmt.Println(s2)
	s2 = s1[:3]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("arr = ", arr)
	fmt.Printf("s1 = %v,len(s1) = %d,cap(s1) = %d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v,len(s2) = %d,cap(s2) = %d\n",
		s2, len(s2), cap(s2))
	fmt.Println("after append")
	// 将切片中对应剩余的数组内容改变成append的值，当append的值超过数组的长度时
	// 就不再是原数组的视图了 而是分配一个更大的底层数组，拷贝进去
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5", s3, s4, s5)
	fmt.Println(arr)
}
