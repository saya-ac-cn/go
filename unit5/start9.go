package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 判断文件夹或文件是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 打印到文件
func print(file io.Writer,data string)  {
	write := bufio.NewWriter(file)
	write.WriteString(data)
	write.Flush()
}

// 初始化文件
func initFile(filePath string) (bool, error) {
	exist, err := pathExists(filePath)
	if err != nil {
		fmt.Printf("获取文件或目录时，发生异常[%v]\n", err)
		return false ,err
	}
	if !exist {
		//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666
		file,err:=os.Create(filePath)
		if err!=nil{
			fmt.Println(err)
		}
		file.Close()
	}
	return true,nil
}

// 打开文件并追加内容
func main() {
	filePath := "./unit5/map.txt"
	_, err := initFile(filePath)
	if err != nil {
		fmt.Printf("初始化文件或目录时，发生异常[%v]\n", err)
		return
	}
	file,err := os.OpenFile(filePath,os.O_WRONLY|os.O_APPEND,0666)
	if err != nil{
		fmt.Println("文件打开异常：",err)
		return
	}
	// 及时关闭文件
	defer file.Close()
	str := "测试数据\n"
	print(file,str)
}
