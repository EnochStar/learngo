package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Employee struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	e := Employee{
		1,
		"boy",
		"xiaoming",
	}
	marshal, err := json.Marshal(e)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str : %s\n", marshal)
}
