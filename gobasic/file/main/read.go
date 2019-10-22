package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/home/zhou/hello.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("file:", file)

	//err = file.Close()
	//if err != nil{
	//	fmt.Println("close file err:",err)
	//}

	defer file.Close() //关闭

	reader := bufio.NewReader(file) //带缓冲方式
	for {
		str, err := reader.ReadString('\n') //读到换行就结束
		if err == io.EOF {                  //io.EOF表示文件末尾
			break
		}
		fmt.Print(str)
	}

	fmt.Println("文件读取结束")
}
