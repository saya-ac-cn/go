package main

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
)

/**
发生异常直接中断退出，不予捕获
 */

func readConf_1(neme string) (err error) {
	if neme != ""{
		return nil
	}else {
		return errors.New("读取文件错误，文件名不能为空")
	}
}

func upload_1()  {
	err := readConf_1("")
	if err != nil{
		// 如果读取异常则发送错误
		panic(err)
	}
	fmt.Println("继续执行后续的代码")
}

func main() {
	upload_1()
	fmt.Println("执行main后续代码")
}
