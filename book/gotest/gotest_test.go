package gotest

import "testing"

func TestDivision(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除数函数测试没通过")
	} else {
		t.Log("第一个测试通过了")
	}
}
func TestDivision2(t *testing.T) {
	if _, e := Division(6, 0); e == nil { //try a unit test on function
		t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
	} else {
		t.Log("one test passed.", e) //记录一些你期望记录的信息
	}
}
