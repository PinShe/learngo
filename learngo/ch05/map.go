package main

import "fmt"

func main() {
	//map是一个key(索引)和value(值)的无序集合，主要是查询方便o(1)
	var courseMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}
	//取值
	//fmt.Println(courseMap["grpc"])

	//放值
	//coursesMap["mysql"] = "mysql的原理"

	//
	//var courseMap map[string]string//nil,map类型想要设置值必须要先初始化
	//var courseMap = make(map[string]string,3)//make是内置函数，主要用于初始化slice、map、channel
	courseMap["mysql"] = "mysql的原理"
	fmt.Println(courseMap)

	//map必须初始化才能使用	1.map[string]string{}	2.map(map[string]string,3)但是slice可以不初始化
	//var m []string
	//if m == nil {
	//	fmt.Println("yes")
	//}
	//赋值，遍历
	//for _ ,value := range courseMap{
	//	fmt.Println(value)
	//}
	for key := range courseMap {
		fmt.Println(key, courseMap[key])
	}

	//map是无序的，而且不保证每次打印都是相同的顺序

	if _, ok := courseMap["java"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("in")
	}

	//删除一个元素
	delete(courseMap, "grpc")
	fmt.Println(courseMap)

	//重要提示，map不是线程安全的
	fmt.Println(courseMap)
}
