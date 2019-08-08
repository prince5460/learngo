package main

import "fmt"

//二分查找

func main() {
	arr := []int{1, 3, 4, 6, 7, 9, 12, 99}

	fmt.Println(BinarySearch(&arr, 12))
}

func BinarySearch(arr *[]int, item int) int {
	low := 0
	high := len(*arr) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := (*arr)[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
