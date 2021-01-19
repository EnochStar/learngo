package main

import (
	"fmt"
	"learngo/video/testing"
)

func getRetriever() retriever {
	return testing.Retriever{}
}

// ?: Something that can Get
// 接口即降低耦合
type retriever interface {
	Get(string) string
}

func main() {
	var retriever retriever = getRetriever()
	bytes := retriever.Get("https://www.imooc.com")
	fmt.Printf(bytes)
}
