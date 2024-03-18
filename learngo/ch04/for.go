package main

import (
	"fmt"
	"time"
)

func main() {
	//for循环
	/*
		for init; condition; post{ }
	*/
	//var i int
	//for i<3{
	//	time.Sleep(2*time.Second)
	//	fmt.Println(i)
	//	i ++
	//}

	//1-100 相加的结果

	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	/*
		99乘法表
	*/

	//遍历，处理第几行
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d", x, y, x*y)
		}
	}

	//for 循环还有一种用法，for range，主要是对 字符串、数组、切片、map、channel

	/*
		for key,value := range{
		}
	*/
	//name := "imooc go体系课"
	//nameRune :=[]rune(name)
	//for index := range name{
	//	//fmt.Println(index,value)
	//	//fmt.Println(name[index])
	//	fmt.Printf("%c\r\n",nameRune[index])
	//}
	//
	//for i := 0;i < len(nameRune);i++ {
	//	fmt.Printf("%c\r\n",nameRune[i])
	//}

	//for _,value := range name{
	//	//fmt.Println(index,value)
	//	//fmt.Println(name[index])
	//	fmt.Printf("%c\r\n",value)
	//}

	/*
		字符串	字符串的索引(key)	字符串对应的索引的字符值的拷贝(value)		如果不写key，那么返回的是索引
		数组 	数组的索引		索引对应的值的拷贝						如果不写key，那么返回的是索引
		切片 	切片的索引		索引对应的值的拷贝						如果不写key，那么返回的是索引
		map     map的key		value 返回的是 key 对应的值的拷贝		如果不写key，那么返回的是 map 的值
		channel					value 返回的是 channel 接受的数据
	*/

	//fmt.Println(sum)

	round := 0
	for {
		time.Sleep(1 * time.Second)
		round++
		if round == 5 {
			continue
		}
		fmt.Println(round)
		if round > 10 {
			break
		}
	}
}
