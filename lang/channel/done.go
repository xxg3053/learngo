

package main

import (
"fmt"
)



type worker struct {
	in chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func doWorker(id int, c chan int, done chan bool)  {
	for n := range c{
		fmt.Printf("Worker %d received %c\n", id, n)
		go func() {
			done <- true
		}()

	}
}

func chanDemo()  {
	var workers [10]worker
	for i :=0; i<10; i++{
		workers[i] = createWorker(i)
	}

	for i,worker := range workers{
		worker.in <- 'a' + i
		//<-workers[i].done
	}
	for i,worker := range workers{
		worker.in <- 'A' + i
		//<-workers[i].done
	}

	for _,worker := range workers{
		<-worker.done
		<-worker.done
	}

}



func main()  {
	chanDemo()
}