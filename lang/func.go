package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) int  {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupperted operaton: " + op)
	}
}


//带余除法
func div(a, b int) (q, r int) {
	return a/b, a%b
}

//参数是函数
func apply(op func(int, int) int, a,b int) int  {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers{
		s = s + numbers[i]
	}
	return s
}

func main()  {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(div(3,4))

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1,2,3,4))
}
