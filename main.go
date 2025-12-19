package main

import (
	"hexlet-go/greeting"

	"github.com/fatih/color"
)

func main() {
	//fmt.Println(greeting.Hello())
	color.Red(greeting.Hello())
	color.Blue(greeting.Hello())
	color.Cyan(greeting.Hello())
}
