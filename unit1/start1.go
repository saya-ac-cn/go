package main

import "fmt"

func main() {
	var days int = 97
	var week int = days / 7
	var day = days % 7
	fmt.Printf("%d个星期零%d天\n",week,day)

	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("%v 对应的摄氏温度=%v",huashi,sheshi)
}
