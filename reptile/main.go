package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-start/tools"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// https://www.jianshu.com/p/757d133021de
/**
*打印每个省下的市区县
*flog 如果为真，则自动去除市名前的+ eg：+北京市
*     如果为假，则原样输出
**/
func filter(selection *goquery.Selection, flog bool) string {
	var result bytes.Buffer
	selection.Find("td").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			if flog {
				result.WriteString(strings.Replace(replace(tools.GbkToUtf8(selection.Text())), "+", "", -1))
			} else {
				result.WriteString(replace(tools.GbkToUtf8(selection.Text())))
			}
		} else {
			//fmt.Println("now",replace(tools.GbkToUtf8(selection.Text())))
			result.WriteString("|")
			result.WriteString(replace(tools.GbkToUtf8(selection.Text())))
		}
	})
	return result.String()
}

// 去除空白字符
func replace(str string) (reslut string) {
	if str == "" {
		return "nil"
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	reslut = reg.ReplaceAllString(str, "")
	return
}

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
func print(file io.Writer, data string) {
	fmt.Print("正在写入：", data)
	write := bufio.NewWriter(file)
	write.WriteString(data)
	write.Flush()
}

// 初始化文件
func initFile(filePath string) (bool, error) {
	exist, err := pathExists(filePath)
	if err != nil {
		fmt.Printf("获取文件或目录时，发生异常[%v]\n", err)
		return false, err
	}
	if !exist {
		//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
	return true, nil
}

// 爬去数据省级开始
func province(url string) {
	// 准备文件写入环境
	filePath := "./reptile/map.txt"
	_, err := initFile(filePath)
	if err != nil {
		fmt.Printf("初始化文件或目录时，发生异常[%v]\n", err)
		return
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开异常：", err)
		return
	}
	// 及时关闭文件
	defer file.Close()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("发生异常")
		log.Fatal(err)
	}

	fmt.Println(replace(tools.GbkToUtf8(doc.Find(".info_table #cp").Text())))

	// 第一层 找出所有的.info_table tbody 集合
	doc.Find(".info_table tbody").Each(func(i int, selection *goquery.Selection) {
		// 遍历.info_table tbody 集合
		// 找出所有的.info_table tbody tr集合
		selection.Find("tr").Each(func(j int, selection *goquery.Selection) {
			// 依次遍历每个tr
			if j == 0 {
				// 表格行首 放空操作
			} else {
				// 判断当前文档对象是否为 市级
				if selection.Is(".shi_nub") {
					//fmt.Println("子节点数量",selection.Find("td").Length())
					print(file, filter(selection, true)+"\n")
					//fmt.Println("当前值为：",)
				} else {
					// 为县区
					print(file, filter(selection, false)+"\n")
					//fmt.Println("当前值为：",filter(selection,false))
				}
			}
		})
	})
}

func main() {
	province("http://xzqh.mca.gov.cn/defaultQuery?shengji=%B0%C4%C3%C5%CC%D8%B1%F0%D0%D0%D5%FE%C7%F8%28%B0%C4%29&diji=-1&xianji=-1")
}
