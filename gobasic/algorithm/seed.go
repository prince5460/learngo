package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		//产生随机数
		//fmt.Println("rand =",rand.Int())//随机很大的数
		fmt.Println("rand =", rand.Intn(100)) //100以内
	}
}
