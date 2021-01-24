package useInterface

import "fmt"

// WashingMachine 洗衣机

// 并且一个接口的方法，不一定需要由一个类型完全实现，
// 接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
type WashingMachine interface {
	Wash()
	Dry()
}

// 甩干器
type Dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d Dryer) Dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type Haier struct {
	Dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h Haier) Wash() {
	fmt.Println("洗刷刷")
}
