package main

import (
	"fmt"
)

func prime(n int) []int {
	a := make([]int, 0)
	for i := 2; i <= n; i++ {
		for j := 2; j <= i; j++ {
			if j == i {
				a = append(a, i)
			}
			if i%j == 0 && j < i {
				break
			}
		}
	}
	return a
}

func prime2(n int) {
	a := make([]int, 0)
	for x := 2; x <= n; x++ {
		for y := 2; y <= x; y++ {
			if x%y == 0 && y < x {
				break
			}
			if y == x {
				a = append(a, x)
			}
		}
	}
	fmt.Println(a)
}

func main() {
	//var a int = 100
	//var b int = 200
	//b, a = a, b
	//fmt.Println(a, b)

	//fmt.Println(prime(100))
	prime2(100000)
}
