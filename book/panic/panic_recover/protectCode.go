package panic_recover

import "fmt"

func ProductCode(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()
	fmt.Printf("x/y = %d\n", z)
}
