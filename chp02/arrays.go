package main

import "fmt"

/**
数组
*/

func printArray(arr []int) {
	//数组是值类型
	//[10]int和[20]int是不同类型
	//调用func f(arr [10]int)会拷贝数组
	arr[0] = 100
	//遍历数组
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func maxValue(arr [5]int) {
	//找出数组最大值
	maxi := -1
	maxValue := -1
	for i, v := range arr {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println(maxi, maxValue)
}

func sumArr(arr [5]int) {
	//数组求和
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)

}

func main() {
	//定义数组的几种方法
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 14, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	fmt.Println("------------")
	printArray(arr1[:])
	fmt.Println("------------")
	printArray(arr3[:])
	fmt.Println("------------")
	fmt.Println(arr1, arr3)
	fmt.Println("------------")
	maxValue(arr3)
	fmt.Println("------------")
	sumArr(arr3)

}
