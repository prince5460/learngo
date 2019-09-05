package main

import (
	"fmt"
	"regexp"
)

const text = "My email is abc@123.com,qwe@123.com what are you doing ?aaa@123.com   sss@123.com.cn"

func main() {
	re := regexp.MustCompile(`([a-zA-z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)

	for _, v := range match {
		fmt.Println(v)
	}
}
