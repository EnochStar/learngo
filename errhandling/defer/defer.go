package main

import (
	"bufio"
	"fmt"
	"os"
)

// defer 确保在函数结束时发生
// 采取栈的结构
// 参数在defer语句时计算
// PrintHeader PrintFooter  Lock  unLock Open close
func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed to many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, i)
	}
}

func main() {
	tryDefer()
}
