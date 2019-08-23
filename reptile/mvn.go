package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func download(dir string, url string) {
	fileName := path.Base(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	file, err := os.Create(dir + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d\n", written)
}

func getFileList(dir string, url string) {
	saveDir := "./reptile/" + dir + "/"
	err := os.Mkdir(saveDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("发生异常")
		os.RemoveAll(saveDir)
		log.Fatal(err)
	}
	doc.Find("pre a").Each(func(i int, selection *goquery.Selection) {
		href, IsExist := selection.Attr("href")
		if i != 0 && IsExist == true {
			fmt.Println(url + href)
			download(saveDir, url+href)
		}
	})
}

func main() {
	getFileList("0.9.0.RELEASE", "https://repo1.maven.org/maven2/org/springframework/cloud/spring-cloud-alibaba-nacos-discovery/0.9.0.RELEASE/")
}
