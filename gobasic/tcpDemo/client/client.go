package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client dail err=", err)
		return
	}

	//发送数据
	reader := bufio.NewReader(os.Stdin) //标准输入

	//从终端输入,并发送给服务器
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		//输入exit退出
		readString = strings.Trim(readString, "\r\n")
		if readString == "exit" {
			fmt.Println("客户端退出!")
			break
		}

		//将readString发送给服务器
		_, err = conn.Write([]byte(readString))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}
	//fmt.Printf("客户端发送了 %d 字节的数据,并退出", write)

}
