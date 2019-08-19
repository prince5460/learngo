package main

/**
循环
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func add() {
	//for的条件里不需要括号
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//for的条件里可以省略初始条件，结束条件，递增表达式
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}

}

func main() {
	add()

	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(101),
		convertToBin(0),
	)

	printFile("abc.txt")

	forever()
}
