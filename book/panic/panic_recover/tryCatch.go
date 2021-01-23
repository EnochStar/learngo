package panic_recover

import "fmt"

// 什么时候用panic，error
// 惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error。
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(nil)
		}
	}()
	fun()
}
func Catch() {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}
