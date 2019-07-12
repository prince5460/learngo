package main

import (
	"fmt"
	"io/ioutil"
)

func testIf() {
	const filename = "abc.txt"

	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	//if的条件里可以赋值
	//if的条件里赋值的变量作用域就在这个if语句里
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func grade(score int) string {
	//switch会自动break，除非使用fallthrough
	//switch后面可以没有表达式
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score :%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	return g

}

func main() {
	testIf()

	fmt.Println(eval(9, 2, "+"))

	fmt.Println(
		grade(20),
		grade(70),
		grade(80),
		grade(99),
		grade(100),
		grade(-3),
	)

}
