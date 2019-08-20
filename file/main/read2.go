package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "hello.txt"
	content, err := ioutil.ReadFile(file) //文件不能太大
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Printf("%v", string(content))
}
