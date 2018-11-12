package main

import (
	"fmt"
	"time"
)

func chanDemo()  {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2

}

func worker(id int, c chan int)  {
	for{
		//n := <-c
		//fmt.Println(n)
		fmt.Printf("Worker %d received %d\n", id, <-c)
	}
}

func bufferedChannel()  {
	c := make(chan int, 3)//三个缓存
    go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	
	time.Sleep(time.Millisecond)
}

func channelClose()  {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	time.Sleep(time.Millisecond)
}

func main()  {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}