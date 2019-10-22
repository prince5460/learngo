package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v,rValue=%T\n", rValue, rValue)

	n1 := 2 + rValue.Int()
	fmt.Println("n1=", n1)

	rV := rValue.Interface()
	fmt.Println("rV=", rV)

}

func main() {
	var num int = 100
	reflectTest(num)
}
