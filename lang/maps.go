package main

import "fmt"

func main()  {

	//hashmap 无序
	m := map[string]string{
		"name": "abc",
		"age": "124",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)
	fmt.Println("Traversing map")
	for k, v := range m{
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	name := m["name"]
	age, ok := m["agesssss"] //没有key的拿到空串
	fmt.Println(name, age, ok)
	if ok{
		fmt.Println("key exist")
	}

	fmt.Println("Deleteing valuse")
	delete(m, "age")
	age, ok = m["age"]
	if !ok{
		fmt.Println("key delete")
	}
	fmt.Println(m)
}
