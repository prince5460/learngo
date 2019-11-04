package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//百度贴吧并发爬虫

func main() {
	var start, end int

	fmt.Printf("请输入起始页(>=1):")
	fmt.Scan(&start)
	fmt.Printf("请输入起始页(>=起始页):")
	fmt.Scan(&end)

	DoWork(start, end)

}

//http://tieba.baidu.com/f?kw=%E4%B8%AD%E5%9B%BD&ie=utf-8&pn=
func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func SpiderPage(i int, page chan<- int) {
	url := "http://tieba.baidu.com/f?kw=%E4%B8%AD%E5%9B%BD&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬第 %d 个网页:%s \n", i, url)
	//爬取网页
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}
	//把内容写入文件
	fileName := strconv.Itoa(i) + ".html"
	f, e := os.Create(fileName)
	if e != nil {
		fmt.Println("os.Create err = ", e)
		return
	}

	f.WriteString(result) //写内容

	f.Close() //关闭文件

	page <- i

}

//爬去网页内容
func HttpGet(url string) (result string, err error) {
	resp, e := http.Get(url)
	if e != nil {
		err = e
		return
	}

	defer resp.Body.Close()

	//读取网页body内容
	bytes := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(bytes)
		if n == 0 { //读取结束或出问题
			fmt.Println("resp.Body.Read err= ", err)
			break
		}

		result += string(bytes[:n])
	}
	return
}
