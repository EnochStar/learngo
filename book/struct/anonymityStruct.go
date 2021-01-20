package main

import "fmt"

func main() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "Enoch"
	user.Age = 11
	fmt.Printf("%#v", user)
}
