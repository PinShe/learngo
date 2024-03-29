package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var mylist list.List

	mylist := list.New()

	//放在尾部
	mylist.PushBack("go")
	mylist.PushBack("grpc")
	mylist.PushBack("mysql")

	//头部放数据
	mylist.PushFront("gin")

	//遍历打印值,正序
	i := mylist.Front()
	for ; i != nil; i = i.Next() {
		if i.Next().Value.(string) == "grpc" {
			break
		}
	}
	mylist.Remove(i)

	//反向遍历
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	//集合类型四种 1.数组 - 不同长度的数组类型不一样 2.切片- 动态数组，用起来很方便而且性能高，尽量使用 3.map 4.list
}
