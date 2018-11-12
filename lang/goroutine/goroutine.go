package main

import (
	"fmt"
	"time"
	"runtime"
)

func main()  {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
				runtime.Gosched()//交出协程控制权
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
}
