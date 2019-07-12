package main

import "fmt"

/**
指针：golang只有值传递
 */

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a

}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	c, d := 3, 4
	c, d = swap2(c, d)
	fmt.Println(c, d)
}
