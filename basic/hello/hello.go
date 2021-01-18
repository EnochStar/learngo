package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func variableTypeDeduction() {
	var a, b, flag = 3, 4, true
	var s = "abc"
	fmt.Println(a, b, s, flag)
}
func variableShorter() {
	a, b, flag := 3, 4, true
	s := "abc"
	fmt.Println(a, b, s, flag)
}
func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}
func consts() {
	const filename = "hello"
	const a, b = 3, 4
	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)
}
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	enums()
}
