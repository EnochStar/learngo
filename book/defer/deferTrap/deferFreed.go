package deferTrap

import (
	"fmt"
	"io"
	"os"
)

//  1、f.Close 可能会报错
//  2、使用相同的变量释放不同的资源 闭包的延迟读取
func Free() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(f)
	}

	return nil
}
