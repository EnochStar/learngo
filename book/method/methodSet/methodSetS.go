package methodSet

import "fmt"

type S struct {
	T
}

func (t T) testS() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。")
}

func MethodSetS() {
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testS()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testS()
	s1.testP()
	s1.testT()
}
