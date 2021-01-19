package main

import "fmt"

// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
// Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode的文本处理更为方便，
// 也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾

// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
// rune类型，代表一个 UTF-8字符。
// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 修改字符串
//要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
//无论哪种转换，都会重新分配内存，并复制字节数组。
func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

// 类型转换
// go语言只有强转
func main() {
	traversalString()
	changeString()
}
