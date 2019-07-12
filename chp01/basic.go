package main

/**
	1.定义变量
	2.内建变量类型
	3.常量与枚举
 */
import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	//包内变量
	aa = 3
	bb = "abc"
	cc = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d,%q\n", a, s)
}

func variableInitialValue() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	//根据初始值自动推断变量类型
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	//:=只能在函数内使用
	a, b, c, d := 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func euler() {
	//欧拉公式
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	//fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	//强制类型转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	//常量
	//const数值可作为各种类型使用
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

func enums()  {
	//枚举
	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		//iota:自增值种子
		cpp = iota
		java
		python
		golang
	)
	const (
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)


}

func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)

	euler()

	triangle()

	consts()

	enums()
}
