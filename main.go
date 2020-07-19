package main

import (
	"fmt"
	"golang_study/myreflect"
)

func main() {
	fmt.Println("haha this is a modules")

	go myreflect.PrintStructInfo()

}
