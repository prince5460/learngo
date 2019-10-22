package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计文件中各种字符串数量

type CharCount struct {
	ZhCount    int //中文字符个数
	EnCount    int //英文字符个数
	NumCount   int // 数字个数
	SpaceCount int //空格个数
	OtherCount int // 其他字符个数
}

func main() {
	fileName := "test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}
	defer file.Close()

	//定义CharCount实例
	var count CharCount
	//	创建一个reader
	reader := bufio.NewReader(file)

	//	循环读取fileName中的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读到文件末尾
			break
		}

		//兼容中文
		str = string([]rune(str))
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.EnCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++

			}
		}

	}
	fmt.Printf("en:%v,num:%v,space:%v,other:%v\n", count.EnCount, count.NumCount,
		count.SpaceCount, count.OtherCount)
}
