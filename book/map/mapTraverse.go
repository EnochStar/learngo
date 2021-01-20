package main

import "fmt"

func main() {
	ScoreMap := make(map[string]int, 8)
	ScoreMap["小明"] = 10
	ScoreMap["张三"] = 19

	for k, v := range ScoreMap {
		fmt.Println(k, v)
	}
}
