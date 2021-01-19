package main

import (
	"errors"
	"fmt"
)

// panic停止当前函数执行
// 一直向上返回，执行每一层的defer
// 如果没有遇到recover，程序退出

// recover 仅在defer调用中使用
// 获取panic的值
// 无法处理则重新panic
func tryRecover() {
	defer func() {
		r := recover()
		if errors, ok := r.(error); ok {
			fmt.Println("Error occurred:", errors)
		} else {
			panic(r)
		}
	}()
	panic(errors.New("this is an error"))
}
func main() {
	tryRecover()
}
