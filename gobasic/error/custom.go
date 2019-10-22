package main

import (
	"errors"
	"fmt"
)

//自定义错误

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test2() {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发生错误,就输出这个错误并终止程序
		panic(err)
	}
	fmt.Println("test2...")

}

func main() {
	test2()

	fmt.Println("main...")
}
