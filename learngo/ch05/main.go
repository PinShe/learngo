package main

import (
	"fmt"
)

func main() {
	//go 语言提供了那些集合类型的数据结构，数组、切片(slice)、map、list
	//数组 定义：var name [count]int
	/*	var courses1 [] string //courses类型，数组 只有3个元素的数组
		//var courses2 [4] string
		//切片：[]strings 和 数组：[3]string //这是两种不同的类型

		courses1[0] = "go"
		courses1[1] = "grpc"
		courses1[2] = "gin"*/

	//fmt.Println(courses1)

	//fmt.Printf("%T\r\n",courses1)
	//fmt.Printf("%Tn",courses2)

	//数组的初始化
	//var courses1 [3]string = [3]string{"go","grpc","gin"}
	//for _ , value := range courses1{
	//	fmt.Println(value)
	//}

	//courses1 := [3]string{"go","grpc","gin"}
	//for _ , value := range courses1{
	//	fmt.Println(value)
	//}
	//数组初始化 2
	//courses2 := [3]string{2:"gin"}
	//for _ , value := range courses2{
	//	fmt.Println(value)
	//}
	//数组初始化 3
	//courses3 := [...]string{"go","grpc"}
	//
	//for i := 0;i < len(courses3);i++{
	//	fmt.Println(courses3[i])
	//}

	//for _ , value := range courses3{
	//	fmt.Println(value)
	//}

	//courses4 := [...]string{"go","grpc"}
	//
	//if courses3 == courses4 {
	//	fmt.Println("equal")
	//}

	//多维数组
	var courseInfo [3][4]string
	courseInfo[0] = [4]string{"go", "1h", "bobby"}
	courseInfo[1] = [4]string{"grpc", "2h", "bobby2"}
	courseInfo[2] = [4]string{"gin", "1.5h", "bobby3"}

	fmt.Println(len(courseInfo))

	for i := 0; i < len(courseInfo); i++ {
		for j := 0; j < len(courseInfo[i]); j++ {
			fmt.Print(courseInfo[i][j] + " ")
		}
		fmt.Println()
	}
	for _, row := range courseInfo {
		fmt.Println(row)
	}
}
