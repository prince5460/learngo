package main

import "fmt"

//冒泡排序

func main() {
	arr := [5]int{34, 6, 1, 23, 8}

	BubbleSort(&arr)

	sl := []int{34, 1, 9, 4, 6, 2}
	BubbleSort2(sl)
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

func BubbleSort2(sl []int) {
	for i := 0; i < len(sl)-1; i++ {
		for j := 0; j < len(sl)-i-1; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
	fmt.Println(sl)
}
