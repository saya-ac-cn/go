package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 发送get请求

const appid = "--------------------" //appid需自己提供
const secret = "-------------------" //secret需自己提供

func main() {
	httpGet()
}

func httpGet() {
	var url bytes.Buffer
	url.WriteString("https://api.weixin.qq.com/sns/jscode2session")
	url.WriteString("?appid=")
	url.WriteString(appid)
	url.WriteString("&secret=")
	url.WriteString(secret)
	url.WriteString("&js_code=")
	url.WriteString("023QNWyY1Cmiq01L8uyY1s8WyY1QNWyn")
	url.WriteString("&grant_type=authorization_code")
	resp, err := http.Get(url.String())
	if err != nil {
		log.Println("请求异常：", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
