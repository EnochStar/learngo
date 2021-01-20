package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	if v, ok := scoreMap["张三"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
