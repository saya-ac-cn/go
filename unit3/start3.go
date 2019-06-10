package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值= %v 地址=%v.\n",intChan,&intChan)
	intChan<-10
	intChan<-20
	intChan<-30
	fmt.Printf("管道的长度=%v，容量=%v.\n",len(intChan),cap(intChan))
	fmt.Println("退出一个元素：", <-intChan)
	fmt.Printf("管道的长度=%v，容量=%v.\n",len(intChan),cap(intChan))
	close(intChan)
	for v := range intChan{
		fmt.Println("v=",v)
	}
}
