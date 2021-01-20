package main

import "fmt"

func AssumeMap() {
	scoreMap := make(map[string]int, 8)
	scoreMap["小明"] = 10
	scoreMap["张三"] = 19
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
}
func assumeMapAdvance() {
	userInfo := map[string]string{
		"username": "enoch",
		"password": "zero",
	}
	fmt.Println(userInfo)
}
func main() {
	assumeMapAdvance()
}
