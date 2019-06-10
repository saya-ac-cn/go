package main

import "fmt"

// & *的使用
func main() {
	a := 100
	fmt.Println("a的地址",&a)

	var ptr *int = &a
	fmt.Println("ptr 指向的值",*ptr)
}
