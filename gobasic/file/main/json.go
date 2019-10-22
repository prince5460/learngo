package main

import (
	"encoding/json"
	"fmt"
)

//json的序列化

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Salary   float64 `json:"salary"`
	Skill    string  `json:"skill"`
}

//结构体
func testStruct() {
	monster := Monster{
		Name:     "牛魔",
		Age:      500,
		Birthday: "1000-10-10",
		Salary:   10000,
		Skill:    "牛魔拳",
	}

	//	将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	fmt.Printf("monster:%v\n", string(data))
}

//map
func testMap() {
	var person map[string]interface{}
	person = make(map[string]interface{})
	person["name"] = "铁扇公主"
	person["age"] = 300
	person["address"] = "火焰山"

	//将people序列化
	data, err := json.Marshal(&person)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	fmt.Printf("people:%v\n", string(data))
}

//slice
func testSlice() {
	var animal []map[string]interface{}
	var a1 map[string]interface{}
	a1 = make(map[string]interface{})
	a1["name"] = "玉面狐狸"
	a1["age"] = 200
	a1["address"] = "火云洞"
	animal = append(animal, a1)

	var a2 map[string]interface{}
	a2 = make(map[string]interface{})
	a2["name"] = "蜘蛛精"
	a2["age"] = 100
	a2["address"] = "盘丝洞"
	animal = append(animal, a2)

	//将animal序列化
	data, err := json.Marshal(&animal)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	fmt.Printf("animal:%v\n", string(data))
}

func main() {
	testStruct()

	testMap()

	testSlice()
}
