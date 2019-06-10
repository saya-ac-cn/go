package main

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
)

/**
捕获错误，不影响后续执行
 */

func readConf_2(neme string) (err error) {
	if neme != ""{
		return nil
	}else {
		return errors.New("读取文件错误，文件名不能为空")
	}
}

func upload_2()  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	err := readConf_2("")
	if err != nil{
		// 如果读取异常则发送错误
		panic(err)
	}
	fmt.Println("继续执行后续的代码")
}

func main() {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	upload_2()
	fmt.Println("执行main后续代码")
}
