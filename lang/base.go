package main

import (
	"math"
	"fmt"
)

func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//const 数值可以作为各种类型来用
func consts()  {
	//const filename = "abc.txt"
	//const a, b = 3, 4
	const(
		filename = "abc.txt"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举类型
func enums()  {
	const (
		cpp = iota
		_
		java
		python
		golang
		javascript
	)
	fmt.Println(cpp, java, python, golang, javascript)

	//b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main()  {
	fmt.Printf("Hello, world")

	triangle()
	consts()
	enums()
}
