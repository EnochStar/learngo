package main

import "fmt"

func AssumeMap() {
	ScoreMap := make(map[string]int, 8)
	ScoreMap["小明"] = 10
	ScoreMap["张三"] = 19
	fmt.Println(ScoreMap)
	fmt.Println(ScoreMap["小明"])
	fmt.Printf("type of a:%T\n", ScoreMap)
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
