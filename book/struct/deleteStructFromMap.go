package main

import "fmt"

type girl struct {
	id   int
	name string
	age  int
}

func main() {
	ce := make(map[int]girl)
	ce[1] = girl{1, "xiaolizi", 22}
	ce[2] = girl{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}
