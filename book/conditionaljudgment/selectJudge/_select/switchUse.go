package _select

import (
	"fmt"
	"time"
)

//比如在下面的场景中，使用全局resChan来接受response，如果时间超过3S,resChan中还没有数据返回，则第二条case将执行

func SelectOutTime() {
	var resChan = make(chan int)
	select {
	case data := <-resChan:
		fmt.Println("接收到数据", data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}

//var shouldQuit=make(chan struct{})
//func ShouldQuit(){
//{
////loop
//}
////...out of the loop
//select {
//case <-c.shouldQuit:
//cleanUp()
//return
//default:
//}
////...
//}
//
////再另外一个协程中，如果运行遇到非法操作或不可处理的错误，就向shouldQuit发送数据通知程序停止运行
//close(shouldQuit)
