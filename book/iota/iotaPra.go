package main

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
const (
	n1 = iota //0
	n2        //1
	_
	n4 //3
)
const (
	i1 = iota //0
	i2 = 100  //100
	i3 = iota //2
	i4        //3
)
const n5 = iota //0

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {

}
