package main

import "fmt"

// map中的key 必须可以比较相等 除了slice，map，function的内建类型都可以作为key
// struct类型不包含上述字段也可以作为key
func main() {
	m := map[string]string{
		"name":    "tom",
		"course":  "math",
		"site":    "zust",
		"quality": "great",
	}
	m2 := make(map[string]int) // m2 == empty

	var m3 map[string]int // m3 == nill

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	// 乱序 和Java一样
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	// 尽管该值在map中不存在 也能收到空串，可用额外的值进行接收来判断是否存在
	if causeName, ok := m["nill"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println("before delete")
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("after delete")
	fmt.Println(name, ok)

}
