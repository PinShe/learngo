package main

import "fmt"

type Person struct {
	name string
	age  int
	//adress struct{
	//	city string
	//}
}
type Student struct {
	//第一种嵌套方式
	//p Person

	//第二种嵌套方式，匿名嵌套
	Person
	score float32
	name  string
}

// 接收器有两种形态
func (p *Person) print() {
	//有可能该函数中想要改变p的值，就是person对象很大，数据较大的时候考虑使用指针方式
	p.age = 23
	fmt.Printf("name:%s,age:%d\r\n", p.name, p.age)
}
func (p *Person) print2() {
	fmt.Printf("name:%s,age:%d\r\n", p.name, p.age)
}

//func

func main() {
	p := &Person{
		"bobby", 18,
	}
	p.print()
	fmt.Println(p.age)
	//s := Student{
	//	Person{
	//		"bobby",18,
	//	},
	//	98.5,
	//	"bobby1",
	//}
	//s.print()

	//s := Student{}
	//s.p.name = "bobby"
	//s.name = "bobby3"
	//fmt.Println(s.name)
}
