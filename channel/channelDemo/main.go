package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {

	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	cat1 := Cat{
		Name: "tom",
		Age:  5,
	}
	cat2 := Cat{
		Name: "jerry",
		Age:  4,
	}

	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"

	newCat := <-allChan
	fmt.Printf("newCat=%T,newCat=%v\n", newCat, newCat)
	a := newCat.(Cat) //断言
	fmt.Println(a.Name)

}
