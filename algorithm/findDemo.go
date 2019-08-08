package main

import "fmt"

//顺序查找

func main() {
	names := [4]string{"小明", "小强", "小白", "小米"}
	var name = ""
	fmt.Println("请输入名称:")
	fmt.Scanln(&name)

	//方式一
	//for i := 0;i<len(names);i++{
	//	if name == names[i]{
	//		fmt.Printf("name:%v,index:%v",name,i)
	//		break
	//	}else if i == (len(names)-1){
	//		fmt.Println("没有找到",name)
	//	}
	//}

	//方式二
	var index = -1
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("name:%v,index:%v", name, index)
	} else {
		fmt.Println("没有找到", name)
	}

}
