package main

import (
	"encoding/json"
	"fmt"
)

//json的反序列化

type Monster1 struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Salary   float64 `json:"salary"`
	Skill    string  `json:"skill"`
}

func unMarshalStrcut() {
	str := "{\"name\":\"牛魔\",\"age\":500,\"birthday\":\"1000-10-10\",\"salary\":10000," +
		"\"skill\":\"牛魔拳\"}"

	var monster Monster1

	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err= %v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v\n", monster, monster.Name)
}

func unMarshalMap() {
	str := "{\"address\":\"火焰山\",\"age\":300,\"name\":\"铁扇公主\"}"

	var people map[string]interface{}
	//反序列化不需要make
	err := json.Unmarshal([]byte(str), &people)
	if err != nil {
		fmt.Printf("unmarshal err= %v\n", err)
	}
	fmt.Printf("反序列化后 people=%v\n", people)

}

func unMarshalSlice() {
	str := "[{\"address\":\"火云洞\",\"age\":200,\"name\":\"玉面狐狸\"}," +
		"{\"address\":\"盘丝洞\",\"age\":100,\"name\":\"蜘蛛精\"}]"

	var slice []map[string]interface{}
	//反序列化不需要make
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err= %v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)

}

func main() {
	unMarshalStrcut()

	unMarshalMap()

	unMarshalSlice()
}
