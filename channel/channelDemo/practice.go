package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Student struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var strChan chan interface{}
	strChan = make(chan interface{}, 10)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		student := Student{
			Name:    "tom" + strconv.Itoa(rand.Intn(100)),
			Age:     12,
			Address: "aaaaa",
		}
		strChan <- student

		fmt.Println(student)
	}

}
