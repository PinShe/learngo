package main

import "fmt"

//class Animal{
//	birthday timestamp
//	legs int
//	mouth int
//}

// duck 必须显示的指定Animal，隐藏的信息：duck也具备了legs字段，也具备了birthday字段、也具备了walk方法
// class Dunk implents Animal{
//
// }
// 鸭子类型强调的事物的外部行为 ，而不是内部的结构，在其他语言中
// 接口的定义
type Dunk interface {
	//方式的申请
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (pd *pskDuck) Swimming() {
	//TODO implement me
	panic("implement me")
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎嘎嘎嘎嘎嘎")
}

func (pd *pskDuck) Walk() {
	fmt.Println("走起来")
}

func (pd *pskDuck) Swimmig() {
	fmt.Println("游起来")
}

func main() {
	//go语言的接口，鸭子类型，php，python
	//go语言中处处都是interface，到处都是鸭子类型 duck typing
	/*
		动词，方法，具备
	*/
	//var duck Duck //多态
	//var an Animal = new Duck//必须的点

	var d Dunk = &pskDuck{}
	d.Walk()
}
