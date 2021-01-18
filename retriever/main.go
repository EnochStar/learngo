package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

const url = "http://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "hello world"}
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute}
	inspect(r)
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	//fmt.Println(download(r))
	fmt.Println("try a session")
	fmt.Println(session(&retriever))
}
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
