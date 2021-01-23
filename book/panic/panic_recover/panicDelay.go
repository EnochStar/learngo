package panic_recover

import "fmt"

//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
func PanicDelay() {
	// 此处无法被捕获
	defer func() {
		panic("fis defer panic")
	}()
	defer func() {
		fmt.Println(recover())
	}()
	// 最后一个延迟调用引发的错误 覆盖前一个
	defer func() {
		panic("second defer panic")
	}()
	// 被覆盖
	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}
