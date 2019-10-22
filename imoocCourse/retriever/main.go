package main

import (
	"fmt"
	"learngo/course/retriever/mock"
	"learngo/course/retriever/real"
	"time"
)

//接口的定义和实现

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "https://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "zhou",
			"passwd": "123",
		})
}

//接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "what are you doing?",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T,%v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	fmt.Println()
}

func main() {
	var r Retriever

	retriever := mock.Retriever{"hello world!"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Second,
	}
	inspect(r)

	//Type assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a real real retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

	//fmt.Println(download(r))
}
