package main

import "fmt"

// go明确规定不支持三目运算
func main() {
	var result int
	var arg1 int = 10
	var arg2 int = 20
	if arg1 > arg2{
		result = arg2
	}else{
		result = arg1
	}
	fmt.Printf("值为：%d",result)
}
