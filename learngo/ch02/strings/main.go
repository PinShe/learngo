package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	S := "Yes我爱慕课网!!!"
	fmt.Println(S)
	for _, b := range []byte(S) {
		fmt.Printf("%X", b)
	}
	fmt.Println()
	for i, ch := range S {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(S))
	bytes := []byte(S)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Println("%c", ch)
	}
	fmt.Println()

	for i, ch := range []rune(S) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	//长度计算

	//转义符
	courseName := "hello\r\ngo\"体系课\""
	fmt.Println(courseName)
	fmt.Print("hello\r\n")
	fmt.Print("world\n")

	//格式化输出
	username := "bobby"
	age := 18
	address := "青海"
	mobile := "123456789321"

	var ages = []int{1, 2, 3}

	fmt.Println("用户名:"+username, ",年龄:"+strconv.Itoa(age)+",地址:"+address, ",电视:"+mobile) //极难维护
	fmt.Printf("用户名:%s,年龄:%d,地址:%s,电话:%s\r\n", username, age, address, mobile)           //常用，性能一般
	userMsg := fmt.Sprintf("用户名:%T,年龄:%d,地址:%s,电话:%s\r\n", ages, age, address, mobile)
	fmt.Println(userMsg)

	//通过String的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(username)
	builder.WriteString("年龄:")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString("地址:")
	builder.WriteString(address)
	builder.WriteString("电话:")
	builder.WriteString(mobile)

	re := builder.String()
	fmt.Println(re)
}
