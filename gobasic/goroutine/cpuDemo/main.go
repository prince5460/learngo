package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU() //主机逻辑cpu个数
	fmt.Println(num)

	runtime.GOMAXPROCS(num - 1) //设置num-1个cpu运行go程序
	fmt.Println("ok")
}
