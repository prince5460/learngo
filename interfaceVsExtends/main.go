package main

import "fmt"

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "can climbing!")
}

type LittleMonkey struct {
	Monkey
}

type Bird interface {
	flying()
}

type Fish interface {
	swimming()
}

func (littleMonkey *LittleMonkey) flying() {
	fmt.Println(littleMonkey.Name, "can flying!")
}

func (littleMonkey *LittleMonkey) swimming() {
	fmt.Println(littleMonkey.Name, "can swiming!")
}

func main() {
	littleMonkey := LittleMonkey{
		Monkey{
			Name: "WuKong",
		},
	}

	littleMonkey.climbing()

	littleMonkey.flying()

	littleMonkey.swimming()
}
