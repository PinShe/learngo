package main

import "fmt"

//我想在想要放多个person的信息到一个集合中
//var persons [][]string

type Person struct {
	name    string
	age     int
	address string
	height  float32
}

//个人信息：name、age、address、height

func main() {
	//	//麻烦
	//	//persons := append(persons,[]interface{}{"bobby",18,"慕课网",1.80})//断言
	//	//如何初始化结构体
	//	p1 := Person{"bobby1",18,"慕课网",1.80}
	//	p2 := Person{
	//		name: "bobby2",
	//		age: 22,
	//		address: "青海",
	//		height: 1.88,
	//	}
	//	var persons []Person
	//	persons = append(persons,p1)
	//	persons = append(persons,Person{
	//		name: "bobby3",
	//	})
	//	persons2 := []Person{
	//		{"bobby1",18,"慕课网",1.80},
	//		{
	//			age: 22,
	//		},
	//	}

	var p Person
	p.age = 20

	fmt.Println(p.height)

	//匿名结构体，匿名函数
	address := struct {
		province string
		city     string
		address  string
	}{
		"北京市",
		"海淀区",
		"xxx",
	}
	fmt.Println(address.city)
}
