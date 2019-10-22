package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{45, 7, 9, 1, 2, 54}
	sort.Ints(a)

	for _, v := range a {
		fmt.Println(v)
	}
}
