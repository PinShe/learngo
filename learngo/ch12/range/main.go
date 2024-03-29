package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int
	//有缓冲和无缓冲的channel
	msg = make(chan int, 2) //channel的初始化值，如果为0的话，你放的值会阻塞
	go func(msg chan int) { //go有一种happen-before的机制，可以保障
		//data := <-msg
		//fmt.Println(data)
		//data = <-msg
		//fmt.Println(data)
		//
		//data = <-msg
		for date := range msg {
			fmt.Println(date)
		}
		fmt.Println("all done")
	}(msg)

	msg <- 1 //放值到channel中
	msg <- 2 //放值到channel中

	close(msg) //其他编程语言有很大的区别
	d := <-msg //已经关闭的channel可以继续取值，但是不能再放值了
	fmt.Println(d)

	//msg <- 3//放值到channel中，已关闭的channel不能再放值了

	//waitgroup 如果少了done调用，容易出现deadlock，无缓冲的channel也容易出现
	time.Sleep(time.Second * 10)
}
