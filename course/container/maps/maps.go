package main

import "fmt"

//map
//map的key：
//map使用哈希表，必须可以比较相等
//除了slice,map,function的内建类型都可以作为key
//Struct类型不包含上述字段，也可作为key

func main() {
	m := map[string]string{
		"name":    "zhou",
		"course":  "golang",
		"site":    "baidu.com",
		"quality": "normal",
	}

	m2 := make(map[string]int) // m2 == empty

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	//cname,ok := m["cause"]
	//fmt.Println(cname,ok)
	if cname, ok := m["cause"]; ok {
		fmt.Println(cname)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
