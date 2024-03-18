package main

import (
	"fmt"
	"strings"
)

func add(a, b interface{}) interface{} {
	//ai,ok := a.(int)
	//if !ok {
	//	panic("not int type")
	//}
	//bi,_ := b.(int)
	//return ai+bi

	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := a.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := a.(int32)
		return ai + bi
	case int64:
		ai, _ := a.(int64)
		bi, _ := a.(int64)
		return ai + bi
	case string:
		as, _ := a.(int32)
		bs, _ := a.(int32)
		return as + bs
	default:
		panic("not supported type")
	}
}

//func addint32(a,b int32) int32 {
//	return a+b
//}
//
//func addint64(a,b int64) int64 {
//	return a+b
//}

func main() {
	//a := 1.0
	//b := 2.0
	re := add("hello", "bobby")
	res, _ := re.(string)
	strings.Split(res, "")
	fmt.Println(re)
}
