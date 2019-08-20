package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//writeFile()

	//coverFile()

	//addFile()

	readAndWrite()

}

func readAndWrite() {
	//读写文件
	fileName := "abc.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//读文件方式一
	//content, err := ioutil.ReadFile(fileName)
	//fmt.Printf("%v",string(content))

	//读文件方式二
	reader := bufio.NewReader(file) //带缓冲方式
	for {
		str, err := reader.ReadString('\n') //读到换行就结束
		if err == io.EOF {                  //io.EOF表示文件末尾
			break
		}
		fmt.Print(str)
	}

	defer file.Close()

	str := "你好， beijing!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func addFile() {
	//	追加内容
	fileName := "abc.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()
	str := "你好， Golang!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}

func coverFile() {
	//覆盖已有文件
	fileName := "abc.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()
	str := "你好， World!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func writeFile() {
	//创建新文件
	fileName := "abc.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()
	str := "Hello World!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
