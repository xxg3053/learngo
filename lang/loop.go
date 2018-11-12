package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

//整数转二进制,对2取模
func convertTobin(n int) string {
	result := ""
	for ; n > 0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func pringFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

//死循环  go中经常使用死循环
func forever()  {
	for{
		fmt.Println("forverver")
	}
}

func main()  {
	fmt.Println(
		convertTobin(5), //101
		convertTobin(13), //1101
	)

	pringFile(".gitignore")
}
