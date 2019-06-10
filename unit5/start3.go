package main

import (
	"fmt"
)

func first()  {
	fmt.Println("first")
}

func second()  {
	fmt.Println("second")
}

func main() {
	defer first()
	defer second()
	fmt.Println("main method")
}
