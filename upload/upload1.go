package main

import (
	"fmt"
	"go-start/tools"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

var (
	// 可执行文件禁止上传
	executFile = [...]string{".exe",".bat",".com",".sys"}
)

// 最大文件上传大小
//const maxUploadSize  = 10 * 1024 * 1024

func uploadOne(w http.ResponseWriter, r *http.Request) {
	//判断请求方式
	if r.Method == "POST" {
		//设置内存大小
		r.ParseMultipartForm(32 << 20);
		//获取上传的第一个文件
		file, header, err := r.FormFile("file");
		// header 取到的主要信息：Filename（文件名），Size（文件大小，单位字节）
		fmt.Println("文件大小：",header.Size)
		fmt.Println("文件名：",header.Filename)
		fileSuffix := checkFileType(header.Filename)
		defer file.Close();
		if err != nil {
			log.Fatal(err);
		}
		//创建上传目录
		os.Mkdir("./upload", os.ModePerm);
		//创建上传文件
		cur, err := os.Create("./upload/" + generateRandomFilename()+fileSuffix);
		defer cur.Close();
		if err != nil {
			log.Fatal(err);
		}
		//把上传文件数据拷贝到我们新建的文件
		io.Copy(cur, file);
	} else {
		//解析模板文件
		//谨慎对待文件
		t, err := template.ParseFiles("./upload/upload.html");
		if err != nil {
			log.Fatal(err);
		}
		//输出文件数据
		t.Execute(w, nil)
	}
}

// 文件扩展验证，不符合返回Ox001
func checkFileType(fileName string) string{
	// 映射成File路径字符串
	filenameWithSuffix := path.Base(fileName)
	//获取文件后缀
	fileSuffix := path.Ext(filenameWithSuffix)
	// 针对复合扩展的处理
	switch fileSuffix {
	case ".gz":
		fileSuffix = getFileSubSuffix(fileName,fileSuffix)
	case ".bz2":
		fileSuffix = getFileSubSuffix(fileName,fileSuffix)
	case ".xz" :
		fileSuffix = getFileSubSuffix(fileName,fileSuffix)
	default:

	}
	if fileSuffix == ""{
		return ""
	}else{
		for _,v := range executFile{
			if v == strings.ToLower(fileSuffix){
				return "Ox001"
			}
		}
	}
	return fileSuffix
}

// 取出文件副扩展名
func getFileSubSuffix(fileName string,fileSuffix string) string {
	// 取出最后一次.出现的位置
	last1PointIndex := strings.LastIndex(fileName, ".")
	// 取出最后最后一个.之前的字符
	subSuffix := fileName[0 : last1PointIndex]
	// 取出倒数第二次.出现的位置
	last2PointIndex := strings.LastIndex(subSuffix, ".")
	if last2PointIndex != -1{
		return fileName[last2PointIndex : last1PointIndex] + fileSuffix
	}else {
		return fileSuffix
	}
}

// 生成文件名（yyyyMMdd加6位随机数）
func generateRandomFilename() string {
	return tools.FormatDate(time.Now(), tools.YYYYMMDD) + fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func main() {
	http.HandleFunc("/upload", uploadOne);
	err := http.ListenAndServe(":9090", nil);
	if err != nil {
		log.Fatal(err)
	}
}