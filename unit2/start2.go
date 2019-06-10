package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string,10)
	a["key1"] = "A"
	a["key2"] = "B"
	fmt.Println(a)

	code := make(map[string]string)
	code["succ"] = "200"
	code["erro"] = "500"
	fmt.Println(code)

	enum := map[string]string{
		"key1" : "a",
		"key2" : "b",
	}
	fmt.Println(enum)
}
