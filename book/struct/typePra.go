package main

import "fmt"

type NewInt int
type MyInt = int

// type 和 typeAlias的区别
func main() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a : %T\n", a)
	fmt.Printf("type of b : %T\n", b)
}
