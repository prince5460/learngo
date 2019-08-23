package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i
		//time.Sleep(time.Second)
		fmt.Println("writeData 写数据,", i)
	}

	close(intChan)
}

func readChan(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Println("readData 读数据,", v)
	}
	//readData读完数据后任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readChan(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
