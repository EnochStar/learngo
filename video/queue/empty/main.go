package main

import (
	"fmt"
	"learngo/video/queue"
)

func main() {
	q := queue.Queue{}

	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Print(q.IsEmpty())
}
