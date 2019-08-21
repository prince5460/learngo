package main

import (
	"fmt"
	"strconv"
	"time"
)

func hello() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello World," + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

func main() {
	go hello() //开启一个协程

	for i := 1; i <= 10; i++ {
		fmt.Println("Hello Golang," + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
