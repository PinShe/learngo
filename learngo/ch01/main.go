package main

import (
	"fmt"
)

//全局变量和局部变量
/*var name = "boby"
var age = 18
var ok   bool*/

var (
	name = "boby"
	age  = 18
	ok   bool
)

func main() {
	//go为静态语言
	//1.变量必须先定义后使用 2.变量必须有类型 3.类型定下来后不能改变
	//定义变量的方式
	//var name int
	//name = 1

	//var age = 1
	age := 18
	var age2 int

	//go语言中变量定义了不使用是不行的
	fmt.Println(age, age2)

	//2.多变量定义
	var user1, user2, user3 = "boby1", 123, "boby3"
	fmt.Println(user1, user2, user3)
	/*注意：
	    变量必须先定义才能使用
		go语言是静态语言，要求变量的类型和赋值类型一致
		变量名不能冲突
		简洁变量定义不能用全局变量
		变量是有零值，也就是默认值
		定义了变量一定要使用，否则会报错
	*/
}
