package main

import "fmt"

func chavalue(a int) int {
	a = a + 1
	return a
}

func chpointer(a *int) {
	*a = *a + 1
}

func main() {
	a := 10
	chavalue(a)
	fmt.Println(a)

	chpointer(&a)
	fmt.Println(a)
}
