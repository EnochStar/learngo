package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type CommonUse struct {
	Contents string
}

//实现stringer接口 相当于java的toString
func (c *CommonUse) String() string {
	return fmt.Sprintf("Retriever:{Contents=%s}", c.Contents)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	c := CommonUse{"hello"}
	fmt.Println(&c)
	printFile("abs.txt")
	s := `abc"d"
	kkkkk
	123`
	printFileContents(strings.NewReader(s))
}
