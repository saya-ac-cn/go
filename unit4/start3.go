package main

import "fmt"

func main() {

	a := [...]int{1,2,3,4,0,5,6,7,8}
	b := a[0:4]

	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Println("-----------")
	//c := a[5:7]
	b = append(b,99)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Println("-----------")
	c := a[5:7]
	b = append(b,c...)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	
}
