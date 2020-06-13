package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

// CoordinateWithWaitGroup 测试WaitGroup的用法
func CoordinateWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(5)
	num := int32(0)
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	go addNum(&num, 3, wg.Done)
	go addNum(&num, 4, wg.Done)
	go addNum(&num, 5, wg.Done)
	go addNum(&num, 6, wg.Done)
	go addNum(&num, 7, wg.Done)
	fmt.Println("wait the result")
	wg.Wait()
	fmt.Printf("num after addNum() this value is %d\n", num)
}

func addNum(src *int32, target int32, cancel cancelFunc) {
	*src = *src + target
	time.Sleep(1 * time.Second)
	cancel()
}

type cancelFunc func()
