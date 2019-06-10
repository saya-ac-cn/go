package main

import "fmt"

/*有两个变量，现要求不使用中间变量的情况下对其交换*/

func main() {
	a := 10
	b := 20
	fmt.Printf("交换前：%d，%d",a,b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("交换后：%d，%d",a,b)
}
