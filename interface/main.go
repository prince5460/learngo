package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("Camera start working...")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop working...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	camera := Camera{}

	computer.Working(camera)
}
