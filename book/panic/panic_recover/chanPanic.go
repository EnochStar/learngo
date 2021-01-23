package panic_recover

import "fmt"

//向已关闭的通道发送数据会引发panic
func ChanPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var ch = make(chan int, 10)
	close(ch)
	ch <- 1
}
