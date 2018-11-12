package main

import "fmt"

func printArray(arr [5]int)  {
	for i :=0; i<len(arr); i++{
		fmt.Print(arr[i])
	}
	fmt.Println()
	for i := range arr{
		fmt.Print(arr[i])
	}
	fmt.Println()
	for i, v := range arr{
		fmt.Print(i, v)
	}
}


func main()  {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2, 4, 6, 8 ,10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	//printArray(arr2)
	printArray(arr3)

}