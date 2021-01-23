package assumeMethod

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() {
	fmt.Printf("%v : %v\n", u.Name, u.Email)
}

//当接受者不是一个指针时，该方法操作对应接受者的值的副本(意思就是即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。

//当接受者是指针时，即使用值类型调用那么函数内部也是对指针的操作。
func UseNotify() {
	u1 := User{
		"enoch",
		"213855@qq.com",
	}
	u1.Notify()

	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()
}
