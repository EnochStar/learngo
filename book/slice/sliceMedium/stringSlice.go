package main

//string底层就是一个byte的数组，因此，也可以进行切片操作。
import (
	"fmt"
)

func main() {
	str := "Hello world"
	s := []byte(str) //中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
}
