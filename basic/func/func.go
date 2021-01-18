package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (int, int) {
	return a / b, a * b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("callhing function %s with agrs %d, %d", opName, a, b)
	return op(a, b)
}

func sum(num ...int) int {
	s := 0
	for i := range num {
		s += num[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	x, _ := div(15, 5)
	fmt.Println(x)
	fmt.Println(apply(pow, 16, 4))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		},
		3, 4))
	fmt.Println(sum(12, 32, 432, 12, 4234, 343, 4))
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
