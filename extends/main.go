package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Run() {
	fmt.Printf("%s,%d Runing\n", a.Name, a.Age)
}

type Dog struct {
	Animal
}

func main() {
	dog := Dog{
		Animal{"大黄", 3},
	}
	dog.Run()
}
