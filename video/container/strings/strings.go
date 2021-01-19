package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Enoch我不喜欢月亮!"
	// UTF8 中3字节 英1字节
	for _, c := range []byte(s) {
		fmt.Printf("( %X)", c)
	}
	fmt.Println()
	for i, ch := range s { //ch is rune utf8转unicode
		fmt.Printf("(%d %x)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	fmt.Println()
	bytes := []byte(s)
	for len(bytes) > 0 {
		decodeRune, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", decodeRune)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
