package assumeMethod

import "fmt"

type Data struct {
	x int
}

func (self Data) ValueTest() {
	fmt.Printf("value: %p\n", &self)
}
func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}

func UseData() {
	d := Data{}
	p := &d
	d.ValueTest()
	d.PointerTest()
	p.PointerTest()
	p.ValueTest()
}
