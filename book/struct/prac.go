package main

import "fmt"

type man struct {
	id   int
	name string
	age  int
}

func demo(ce []man) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// append 返回的是一个新的slice
	ce = append(ce, man{3, "xiaowang", 56})
	//return ce
}
func main() {
	var ce []man //定义一个切片类型的结构体
	ce = []man{
		man{1, "xiaoming", 22},
		man{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}
