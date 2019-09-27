package main

import (
	"fmt"
	"learngo/factory/model"
)

func main() {
	var stu = model.NewStudent("tom", 18, 100)
	fmt.Println(*stu)

	fmt.Println(stu.GetScore())
}
