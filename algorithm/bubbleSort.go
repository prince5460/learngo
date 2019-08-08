package main

import "fmt"

//冒泡排序

func main() {
	arr := [5]int{34, 6, 1, 23, 8}

	BubbleSort(&arr)

	fmt.Println("main arr:", arr)
}

func BubbleSort(arr *[5]int) {
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
	fmt.Println(*arr)
}
