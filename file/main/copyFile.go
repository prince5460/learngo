package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将文件A内容导入到文件B
	filePath1 := "hello.txt"
	filePath2 := "golang.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}

	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Println("write file err:", err)
	}
}
