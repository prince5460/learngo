package main

import "fmt"

//快速排序

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[0]
	var less []int
	var greater []int

	for _, v := range list[1:] {
		if pivot > v {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	less = append(quickSort(less), pivot)
	greater = quickSort(greater)
	return append(less, greater...)
}

func main() {
	fmt.Println(quickSort([]int{34, 2, 4, 1, 5}))
}
