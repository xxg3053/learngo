package main

import (
	"io/ioutil"
	"fmt"
)

//if
func ifopen(){
	const filename = ".gitignore"
	contents, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}
	//简写
	if contents, err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n", contents)
	}
}

//switch 不需要break, 如果需要继续使用fallthrough
func grade(score int) string  {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("error score")
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <=100:
		g = "A"
	default:
		panic("error score")
	}
	fmt.Println(g)
	return g
}

func main()  {
	ifopen()
	grade(30)
}
