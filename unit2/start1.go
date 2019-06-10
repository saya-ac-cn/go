package main

import "fmt"

//冒泡排序
func BubbleSort(arr *[5]int){
	fmt.Println("排序前arr=",*arr)
	temp := 0//临时变量用于交换
	for i := 0; i < len(*arr); i++{
		for j := 0; j < len(*arr) - i - 1; j++{
			if (*arr)[j] > (*arr)[j + 1]{
				// 交换
				temp  = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}
}

func main(){
	arr := [5]int{12,56,89,24,80}
	BubbleSort(&arr)
	fmt.Println("排序后：",arr)
}
