package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 为函数实现接口
func fibonacci() intGen {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func main() {
	f := fibonacci()
	printFileContents(f)
}
