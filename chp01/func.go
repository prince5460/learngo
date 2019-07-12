package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
函数
 */

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operator:%s", op)
	}
}

//多个返回值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数作为参数
func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s

}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(eval(1, 2, "+"))

	q, r := div(12, 3)
	fmt.Println(q, r)

	//fmt.Println(apply(pow, 3, 4))

	//匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
