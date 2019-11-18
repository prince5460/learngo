package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

var onlineMap map[string]Client

var message = make(chan string)

//新开一个协程,转发消息
func Manager() {
	onlineMap = make(map[string]Client)

	for {
		msg := <-message //没有消息前会阻塞

		//遍历map

	}
}

//处理用户连接
func HandleConn(conn net.Conn) {

}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	defer listener.Close()

	go Manager()

	//主协程,循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err=", err)
			continue
		}

		go HandleConn(conn)
	}
}
