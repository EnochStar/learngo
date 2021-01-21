package anonymousFunction

import (
	"fmt"
	"math"
)

//匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
func Anonymous() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
}
func MoreAnonymous() {
	// --- function variable ---
	fn := func() { println("Hello, World!") }
	fn()
	fns := []func(x int) int{
		func(x int) int {
			return x + 1
		},
		func(x int) int {
			return x + 2
		},
	}
	println(fns[0](100))
	// --- function as field ---
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	println(d.fn())

	// --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())
}
