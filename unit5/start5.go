package main

import "fmt"

func addUpper() func(int) int {
	n := 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := addUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
