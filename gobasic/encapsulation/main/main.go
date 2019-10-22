package main

import (
	"fmt"
	"learngo/encapsulation/model"
)

func main() {
	p := model.NewPerson("Tom")
	p.SetAge(18)
	p.SetSalary(50000.0)

	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
	fmt.Println(p.GetSalary())

}
