package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	//panic("this is an panic")//panic会导致程序的退出，平时开发中不要随便使用
	//如果我们的服务启动检查中发现了这些任何一个不满足那就调用panic 主动调用
	//recover 这个函数能捕获到panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recoverd if A:", r)
		}
	}()
	panic("this is an panic")
	var names map[string]string
	names["go"] = "go工程师"

	//fmt.Println("this is a func")
	return 0, errors.New("this is an error")

	//1.defer 需要放在panic之前定义，另外 recover 只有在 defer 调用的函数中才会生效
	//2.recover	处理异常后，逻辑并不会恢复到	panic 的那个点去
	//3.多个defer会形成栈，后定义的defer会先执行
}

func main() {
	//panic	这个函数会导致你的程序退出，不推荐随便使用panic

	if _, err := A(); err != nil {
		fmt.Println(err)
	}
}
