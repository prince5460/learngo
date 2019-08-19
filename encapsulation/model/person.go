package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄不对！")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary < 3000 && salary > 50000 {
		fmt.Println("金额不对！")
	} else {
		p.salary = salary
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
