package main

import (
	"fmt"
	"time"
)

func main() {
	//构建一个通道
	ch := make(chan int)

	//开启一个并发匿名函数
	go func() {
		for i := 3; i >= 0; i-- {
			//发送3到0之间的数值
			ch <- i

			time.Sleep(time.Second)
		}
	}()

	//遍历接收通道数据
	for data := range ch {
		fmt.Println(data)

		if data == 0 {
			break
		}
	}

}
