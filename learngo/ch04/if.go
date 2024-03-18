package main

import "fmt"

/*if 布尔表达式{
	逻辑
}
*/

func main() {
	//if条件表达式
	age := 16
	//courtry := "美国"
	/*	if (age < 18)&&(courtry == "中国"){
		fmt.Println("未成年")
	}*/

	if age < 18 {
		//if courtry == "中国"{
		//	fmt.Println("未成年")
		//}
	} else if age == 18 {
		fmt.Println("刚好成年")
	} else {
		fmt.Println("已成年")
	}
}
