package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送消息
		//fmt.Printf("等待%s发送消息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("%s退出,err = %v\n", conn.RemoteAddr().String(), err)
			return
		}
		//显示客户端发送的内容
		fmt.Printf("%s:%s\n", conn.RemoteAddr().String(), string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listener err=,", err)
		return
	}
	defer listener.Close()

	//循环等待客户端来连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() err=,", err)
		} else {
			fmt.Printf("Accept() success=%v,客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}

		go process(conn)
	}

}
