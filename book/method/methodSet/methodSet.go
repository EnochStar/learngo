package methodSet

import "fmt"

//    • 类型 T 方法集包含全部 receiver T 方法。
//    • 类型 *T 方法集包含全部 receiver T + *T 方法。
//    • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
//    • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
//    • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。
type T struct {
	int
}

func (t T) testT() {
	fmt.Println("类型 *T 方法集包含全部 receiver T 方法。")
}

func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

func MethodSet() {
	t1 := T{1}
	t2 := &t1
	fmt.Printf("t2 is : %v\n", t2)
	t2.testT()
	t2.testP()
}
