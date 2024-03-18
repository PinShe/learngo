package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	/*
		不同类型的数据零值不一样
		bool		false
		numbers		0
		String		""
		pointer		nil
		slice 		nil
		map			nil
		channel、interface、function、nil
	*/
	p1 := Person{
		name: "bobby",
		age:  29,
	}

	p2 := Person{
		name: "bobby",
		age:  19,
	}
	if p1 == p2 {
		fmt.Println("yes")
	}

	//slice的默认值
	//var ps []Person		//nil slice
	//var ps2 = make([]Person,0)		//empty slice
	//if ps2 != nil {
	//	fmt.Println("nil slice")
	//}

	//var m map[string]string//nil map and empty map
	var m map[string]string
	var m2 = make(map[string]string, 0)

	for key, value := range m {
		fmt.Println(key, value)
	}

	fmt.Println(m["name"])
	//m["name"] = "bobby"
	m2["name"] = "bobby"

	for key, value := range m2 {
		fmt.Println(key, value)
	}

	//if m2 == nil {
	//	fmt.Println("nil map")
	//}
}
