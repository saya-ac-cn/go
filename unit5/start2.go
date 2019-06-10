package main

import "fmt"

func suma(arr ...int) (sum int) {
	for _,v:=range arr{
		sum += v
	}
	return
}

func sumb(arr []int) (sum int) {
	for _,v:=range arr{
		sum += v
	}
	return
}

func main() {
	slice := []int{1,2,3,4}
	//array := [...]int{1,2,3,4}
	fmt.Println(suma(slice ...))
	fmt.Printf("%t\n",suma)
	fmt.Printf("%t\n",sumb)

}
