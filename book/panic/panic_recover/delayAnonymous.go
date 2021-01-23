package panic_recover

import "fmt"

func except() {
	fmt.Println(recover())
}

func DelayAnonymous() {
	defer except()
	panic("test panic")
}
