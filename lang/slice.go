package main

import "fmt"

func updateSlice(s []int)  {
	s[0] = 100
}

func main()  {
	arr := [...]int{0, 1, 2, 3, 4 , 5, 6, 7} //slice
	//代码中一般用半开半闭区间
	//下面的slice都是arr的view
	//slice都是值传递
	fmt.Println("arr[2:6] =", arr[2:6]) //[2 3 4 5]
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])

	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("after updateSlice s1")
	updateSlice(s1)
	fmt.Println("s1 =", s1)
	fmt.Println("arr =", arr)//arr本身也改变，说明s1是arr的一个view

	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println("s2 =", s2)

	fmt.Println("Extending slice")
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

}