package main

import (
	"fmt"
	"path"
	"strings"
)
//https://www.cnblogs.com/lyxy/p/5629151.html
func main() {

	fullFilename := "test.tar.gz"
	fmt.Println("fullFilename =", fullFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	fmt.Println("fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)//获取文件名
	fmt.Println("filenameOnly =", filenameOnly)
	
}
