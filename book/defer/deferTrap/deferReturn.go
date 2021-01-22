package deferTrap

import "fmt"

// 输出2
func DeferReturn() (i int) {
	i = 0
	defer func() {
		fmt.Println(i)
	}()
	return 2
}

// 输出0
func DeferReturnAnonymous() int {
	i := 0
	defer func() {
		fmt.Println(i)
	}()
	return 2
}
