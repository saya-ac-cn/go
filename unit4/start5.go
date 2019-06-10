package main

import (
	"fmt"
	"go-start/tools"
	"math/rand"
	"time"
)

// 生成文件名（yyyyMMdd加6位随机数）
func generateRandomFilename() string {
	currentTime :=time.Now()
	year := currentTime.Year() //年
	month := currentTime.Month() //月
	day := currentTime.Day() //日
	return fmt.Sprintf("%d%d%d%06v", year, month, day, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func main() {
	fmt.Println(tools.FormatDate(time.Now(), tools.YYYYMMDD))
}
