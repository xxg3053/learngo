package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	s := "hello world,你好" //UTF-8
	fmt.Println(len(s))
	for i, b := range []byte(s){
		fmt.Printf("(%d %X) ", i, b)
	}

	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0{
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s){
		fmt.Printf("(%d %c)", i, ch)
	}




}
