package main

import (
	"fmt"
	"strconv"
	"time"
)
// 每秒打印hello,world
func test(){
	for i :=1; i <= 10; i++{
		fmt.Println("test () hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	// 开启一个协程
	go test()
	for i := 1; i <= 10; i++{
		fmt.Println("main () hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
