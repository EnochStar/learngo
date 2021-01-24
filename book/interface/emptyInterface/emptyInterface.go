package emptyInterface

import "fmt"

//空接口类型的变量可以存储任意类型的变量。
func EmptyInterface() {
	var x interface{}
	s := "pprof.cn"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

// 使用空接口实现可以接收任意类型的函数参数。
func Show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 使用空接口实现可以保存任意值的字典。
func MapUse() {
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

// 获取空接口的值
// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。
// 这两部分分别称为接口的动态类型和动态值。
func Assert() {
	var x interface{}
	x = "pprof.cn"
	//    x：表示类型为interface{}的变量
	//    T：表示断言x可能是的类型。
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	}
}
func JustifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
