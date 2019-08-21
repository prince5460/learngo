package testcase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (monster *Monster) Store() bool {
	//序列化
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}
	//保存到文件
	filePath := "monster.log"
	ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("WriteFile err=", err)
		return false
	}
	return true
}

func (monster *Monster) ReStore() bool {
	//读取文件
	filePath := "monster.log"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}
	//反序列化
	err = json.Unmarshal(data, monster)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
		return false
	}
	return true
}
