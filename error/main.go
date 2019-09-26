package main

import (
	"fmt"
)

// defer + recover来捕获和处理异常
func main() {
	test()

}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	n1 := 10
	n2 := 0
	res := n1 / n2
	fmt.Println("res=", res)
}
