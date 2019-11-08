package main

import (
	"github.com/astaxie/beego"
	_ "learngo/class/models"
	_ "learngo/class/routers"
)

func main() {
	beego.AddFuncMap("ShowPrePage", HandlePrePage)
	beego.AddFuncMap("ShowNextPage", HandleNextPage)
	beego.Run()
}

//上一页
func HandlePrePage(data int) int {
	pageIndex := data - 1
	return pageIndex
}

//下一页
func HandleNextPage(data int) int {
	pageIndex := data + 1
	return pageIndex
}
