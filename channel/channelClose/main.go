package main

import "fmt"

func main() {
	intChan := make(chan int, 3)

	intChan <- 10
	intChan <- 20
	close(intChan)
	//intChan <- 30
	//channel关闭后不能写入,可以读取
	n1 := <-intChan
	fmt.Println(n1)

	intChan2 := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		intChan2 <- i
	}
	//遍历前需关闭
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
