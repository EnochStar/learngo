package face

// 语法糖
// 如果Speak设置的为值接收者
// 那么无论传入类型使Student{} 或 &Student{} 均可以
// 反之 必须为指针类型

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
