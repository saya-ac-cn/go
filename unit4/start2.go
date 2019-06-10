package main

import (
	"fmt"
	"go-start/dl"
)

func main() {
	for i := 0; i < 20; i++ {
		model := dl.GetInstance()
		fmt.Println(&model)
	}
}
