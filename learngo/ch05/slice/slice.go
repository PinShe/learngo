package main

import "fmt"

func main() {
	//go 折中, slice 切片 - 数组
	var courses []string
	fmt.Printf("%T", courses)

	//这个方法很特别，第一次头疼，原理
	//courses = append(courses,"go")
	//courses = append(courses,"grpc")
	//courses = append(courses,"gin")
	//
	//fmt.Println(courses[1])
	//
	////切片的初始化 3种：1：从数组直接创建 2：使用String{} 3：make
	//allCourses := [5]string{"go","grpc","mysql","mysql","elasticsearch"}
	//coursesSlice := allCourses[0:len(allCourses)]	//左闭右开
	//fmt.Println(coursesSlice)

	//coursesSlice := []string{"go","grpc","mysql","mysql","elasticsearch"}

	//make	函数
	//allCourses2 := make([]string,3)
	//allCourses2[0] = "c"
	//allCourses2[1] = "c"
	//allCourses2[2] = "c"
	//allCourses2[3] = "c"
	//fmt.Println(allCourses2)
	//
	//var allCourses3 []string
	//allCourses3= append(allCourses3,"c")
	//fmt.Println(allCourses3)

	//访问切片的元素，访问单个、多个
	//fmt.Println(coursesSlice[1])

	//[start:end]
	/*
		1.如果只有start 没有end 就表示从start开始到结尾的所有数据
		2.如果没有start 有end 就表示从0到end的所有数据
		3.如果有start 没有有end
		3.如果有start 有end
	*/
	//fmt.Println(coursesSlice[1:4])
	//fmt.Println(coursesSlice[1:])
	//fmt.Println(coursesSlice[:])

	//coursesSlice := []string{"go","grpc"}
	//coursesSlice2 := []string{"mysql","gin","elasticsearch"}
	//coursesSlice = append(coursesSlice,coursesSlice2[1:]...)

	//如何删除slice中的元素：比较麻烦
	//fmt.Println(coursesSlice)

	coursesSlice := []string{"go", "grpc", "mysql", "gin", "elasticsearch"}
	myslice := append(coursesSlice[:2], coursesSlice[3:]...)
	fmt.Println(myslice)

	coursesSlice = coursesSlice[:3]
	fmt.Println(coursesSlice)

	//复制 slice
	//coursesSliceCopy := coursesSlice
	//coursesSliceCopy2 := coursesSlice[:]
	//fmt.Println(coursesSliceCopy,coursesSliceCopy2)

	coursesSliceCopy2 := coursesSlice[:]
	fmt.Println(coursesSliceCopy2)

	var coursesSliceCopy = make([]string, len(coursesSlice))
	copy(coursesSliceCopy, coursesSlice)
	fmt.Println(coursesSliceCopy)

	fmt.Println("--------------------------------")
	coursesSlice[0] = "java"
	coursesSlice[2] = "php"
	fmt.Println(coursesSliceCopy2)
	fmt.Println(coursesSliceCopy)
}
