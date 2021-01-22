package deferClose

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func Close(t Test) {
	t.Close()
}

// defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。
// 但是并没有说struct这里的this指针如何处理，
// 通过这个例子可以看出go语言并没有把这个明确写出来的this指针当作参数来看待。

func DeferClose() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		// 此时调用的结果永远为 c closed
		// defer t.Close()

		// 方案一： 使用Close(t Test) 函数 将变量作为参数 放入Close函数中
		//defer Close(t)

		// 方案二： 用变量先存储t的当前值 然后defer
		t2 := t
		defer t2.Close()
	}
}

//多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。
//哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。
func Panic(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")
}
