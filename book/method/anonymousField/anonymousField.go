package anonymousField

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	tittle string
}

func (self *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v", self, self)
}
func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

// 通过匿名字段，可获得和继承类似的复用能力。依据编译器查找次序，只需在外层定义同名方法，就可以实现 "override"。
func AnonymousField() {
	m := Manager{User{1, "Tom"}, "Administrator"}

	fmt.Println(m.ToString())

	fmt.Println(m.User.ToString())
}
