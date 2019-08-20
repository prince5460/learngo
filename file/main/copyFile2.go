package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	//通过srcfile，获取Reader
	reader := bufio.NewReader(srcFile)
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//通过dstFile，获取Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {
	srcFile := "2.jpg"
	dstFile := "animal.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Printf("拷贝错误 err = %v\n", err)
	}
	fmt.Println("拷贝成功！")
}
