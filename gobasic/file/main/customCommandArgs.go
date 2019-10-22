package main

import (
	"flag"
	"fmt"
)

func main() {
	//定义几个变量用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//1:参数值,2:-u指定参数,3:默认值,4:说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机地址")
	flag.IntVar(&port, "p", 8080, "端口号")

	flag.Parse()

	fmt.Printf("user:%v,pwd:%v,host:%v,port:%v\n", user, pwd, host, port)

}
