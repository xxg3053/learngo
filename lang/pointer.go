package main

import "fmt"

//值传递，不改变值本身，交换无效
func swap1(a, b int){
	a, b = b, a
}

func swap2(a, b *int){
	*a, *b = *b, *a
}

func main()  {
	a, b := 3, 4
	swap1(a, b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
}
