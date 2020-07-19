package main

import (
	"fmt"
	"golang_study/goroutine/coroutinenest"
	"time"
)

func main() {
	go coroutinenest.OuterCoroutine()
	fmt.Println(" main method is executed")
	time.Sleep(2 * time.Second)
}
