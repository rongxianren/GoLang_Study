package coroutinenest

import (
	"fmt"
	"time"
)

// OuterCoroutine 父协程结束后 子协程是不会继续执行的
func OuterCoroutine() {
	time.Sleep(1 * time.Second)
	fmt.Println("OuterCoroutine is executed")
	go nestCoroutine()
}

func nestCoroutine() {
	time.Sleep(3 * time.Second)
	fmt.Println("nestCoroutine is executed")
}
