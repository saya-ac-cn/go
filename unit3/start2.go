package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int,10)
	lock sync.Mutex
)

func test1(n int){
	res := 1
	for i := 1; i <= n; i++{
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i:= 1; i <= 200; i++ {
		go test1(i)
	}
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock()
}
