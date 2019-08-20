package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(PathExist("abc.txt"))
}

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
