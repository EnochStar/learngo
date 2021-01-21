package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]string, 5)
	map1[1] = "enoch"
	map1[2] = "coll"
	map1[3] = "linda"
	map1[4] = "zero"
	map1[5] = "star"
	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])

	}
}
