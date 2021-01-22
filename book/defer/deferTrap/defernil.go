package deferTrap

import "fmt"

//解释：
//名为 DeferNil 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。
//然而值得注意的是，run() 的声明是没有问题，因为在 DeferNil 函数运行完成后它才会被调用。
func DeferNil() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

func RunDeferNil() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	DeferNil()
}
