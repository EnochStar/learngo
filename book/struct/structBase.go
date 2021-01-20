package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

func main() {
	var p1 person
	p1.name = "enoch"
	p1.age = 10
	p1.city = "WenZhou"

	fmt.Printf("p1 = %v\n", p1)
	fmt.Printf("p1 = %#v\n", p1)

}
