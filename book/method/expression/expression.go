package expression

import "fmt"

type User struct {
	id   int
	name string
}

//instance.method(args...) ---> <type>.func(instance, args...)
// 前者为 method value   会复制receiver
// 后者为 method expression
func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}
func (self User) Test1() {
	fmt.Println(self)
}
func Expression() {
	//u := User{1, "Tom"}
	//u.Test()
	//
	//mValue := u.Test
	//mValue() // 隐式传递
	//mExoression := (*User).Test
	//mExoression(&u)

	u := User{1, "tom"}
	myValue := u.Test1 // 立即复制 receiver，因为不是指针类型，不受后续修改影响

	u.id = 2
	u.name = "enoch"
	u.Test1()
	myValue()
}
