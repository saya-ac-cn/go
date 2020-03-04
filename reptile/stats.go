package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"regexp"
)

/**
 * 爬取国家统计局行政区数据，
 * 注意：由于统计局网页开发采用的gb2312编码，请注意编码的转换
**/

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

// 爬取数据省级开始
func province(url string, file *os.File) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("一级错误,一级错误,一级错误,一级错误,一级错误\t" + url)
		print(file, "一级错误,一级错误,一级错误,一级错误,一级错误\t"+url+"\n")
	}
	fmt.Println(doc)
	// 第一层 找出所有的省级行政单位
	doc.Find(".provincetable a").Each(func(i int, selection *goquery.Selection) {
		province_url, _ := selection.Attr("href")
		fmt.Println(province_url)
		// 递归查询市
		//city(url,province_url,ConvertToString(selection.Text(), "gbk", "utf-8"),file)
	})
}

// 爬取市级单位
func city(home_url string, url string, province_name string, file *os.File) {
	doc, err := goquery.NewDocument(home_url + url)
	if err != nil {
		fmt.Println("二级错误,二级错误,二级错误,二级错误,二级错误\t" + url)
		print(file, "二级错误,二级错误,二级错误,二级错误,二级错误\t"+url+"\n")
		fmt.Println("致命异常：", err)
		return
	}
	doc.Find(".citytr").Each(func(i int, selection *goquery.Selection) {
		city := selection.Find("td").Last().Find("a")
		city_url, _ := city.Attr("href")
		city_name := ConvertToString(city.Text(), "gbk", "utf-8")
		county(home_url, city_url, province_name, city_name, file)
	})
}

// 爬取县级单位
func county(home_url string, url string, province_name string, city_name string, file *os.File) {
	doc, err := goquery.NewDocument(home_url + url)
	if err != nil {
		fmt.Println("三级错误,三级错误,三级错误,三级错误,三级错误\t" + home_url + url)
		print(file, "三级错误,三级错误,三级错误,三级错误,三级错误\t"+home_url+url+"\n")
		fmt.Println("致命异常：", err)
		return
	}
	// 有些地级市为特殊情况，市直接到县,比如广东省清远市
	countytr := doc.Find(".countytr")
	towntr := doc.Find(".towntr")
	if countytr.Size() != 0 {
		// 常规走法，查县级
		countytr.Each(func(i int, selection *goquery.Selection) {
			county := selection.Find("td").Last().Find("a")
			if 0 == county.Size() {
				// 为 市本级
				county_code := ConvertToString(selection.Find("td").First().Text(), "gbk", "utf-8")
				county_name := ConvertToString(selection.Find("td").Last().Text(), "gbk", "utf-8")
				//selection.Find("td").Last().Text();
				fmt.Println(province_name + "," + city_name + "," + county_name + "," + county_name + "." + county_name + "\t" + county_code)
				print(file, province_name+","+city_name+","+county_name+","+county_name+"."+county_name+"\t"+county_code+"\n")
			} else {
				county_url, _ := county.Attr("href")
				county_name := ConvertToString(county.Text(), "gbk", "utf-8")
				// 寻找下一级 县
				//fmt.Println(substring(home_url+url,57,county_url))
				town(substring(home_url+url, 57, county_url), province_name, city_name, county_name, file)
				//fmt.Println((home_url+url)+county_url+"\t"+county_name)
			}
		})
	} else if towntr.Size() != 0 {
		// 非正常走位，越级查镇级
		towntr.Each(func(i int, selection *goquery.Selection) {
			town := selection.Find("td").Last().Find("a")
			if 0 == town.Size() {
				// 为 县级
				town_code := ConvertToString(selection.Find("td").First().Text(), "gbk", "utf-8")
				town_name := ConvertToString(selection.Find("td").Last().Text(), "gbk", "utf-8")
				//selection.Find("td").Last().Text();
				fmt.Println(province_name + "," + city_name + "," + city_name + "," + town_name + "." + town_name + "\t" + town_code)
				print(file, province_name+","+city_name+","+city_name+","+town_name+"."+town_name+"\t"+town_code+"\n")
			} else {
				town_url, _ := town.Attr("href")
				town_name := ConvertToString(town.Text(), "gbk", "utf-8")
				// 寻找下一级 镇
				//fmt.Println(url+town_name+"\t"+town_url+"\t"+substring(home_url+url,57,town_url))
				village(substring(home_url+url, 57, town_url), province_name, city_name, city_name, town_name, file)
			}
		})
	}
}

// 爬取镇级单位
func town(url string, province_name string, city_name string, county_name string, file *os.File) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("四级错误,四级错误,四级错误,四级错误,四级错误\t" + url)
		print(file, "四级错误,四级错误,四级错误,四级错误,四级错误\t"+url+"\n")
		fmt.Println("致命异常：", err)
		return
	}
	doc.Find(".towntr").Each(func(i int, selection *goquery.Selection) {
		town := selection.Find("td").Last().Find("a")
		if 0 == town.Size() {
			// 为 县级
			town_code := ConvertToString(selection.Find("td").First().Text(), "gbk", "utf-8")
			town_name := ConvertToString(selection.Find("td").Last().Text(), "gbk", "utf-8")
			//selection.Find("td").Last().Text();
			fmt.Println(province_name + "," + city_name + "," + county_name + "," + town_name + "." + town_name + "\t" + town_code)
			print(file, province_name+","+city_name+","+county_name+","+town_name+"."+town_name+"\t"+town_code+"\n")
		} else {
			town_url, _ := town.Attr("href")
			town_name := ConvertToString(town.Text(), "gbk", "utf-8")
			// 寻找下一级 镇
			//fmt.Println(url+town_name+"\t"+town_url+"\t"+substring(url,60,town_url))
			village(substring(url, 60, town_url), province_name, city_name, county_name, town_name, file)
		}
	})
}

// 爬取村级单位
func village(url string, province_name string, city_name string, county_name string, town_name string, file *os.File) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("五级错误,五级错误,五级错误,五级错误,五级错误\t" + url)
		print(file, "五级错误,五级错误,五级错误,五级错误,五级错误\t"+url+"\n")
		fmt.Println("致命异常：", err)
		return
	}
	doc.Find(".villagetr").Each(func(i int, selection *goquery.Selection) {
		village_code := selection.Find("td").First().Text()
		village_name := ConvertToString(selection.Find("td").Last().Text(), "gbk", "utf-8")
		fmt.Println(province_name + "," + city_name + "," + county_name + "," + town_name + "," + village_name + "\t" + village_code)
		print(file, province_name+","+city_name+","+county_name+","+town_name+","+village_name+"\t"+village_code+"\n")
	})
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result

}

func substring(homeurl string, index int, joinurl string) string {
	content := homeurl[0:index]
	return content + joinurl
}

func main() {
	// 统计局省级页面
	//province("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/")
	// 准备文件写入环境
	filePath := "./reptile/China-32.txt"
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
	city("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/", "32.html", "江苏省", file)
	//county("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/","45/4514.html","广西自治区","崇左市",file)
	//town("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/65/31/653131.html","新疆维吾尔自治区","喀什地区","塔什库尔干塔吉克自治县",file)
	//village("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/51/03/22/510322001.html","111")
}

// 任务执行节点：
//41、河南省；42、湖北省；43、湖南省；44、广东省；45、广西自治区；46、海南省
//50、重庆市；51、四川省；52、贵州省；53、云南省；54、西藏自治区
//61、陕西省；62、甘肃省；63、青海省；64、宁夏自治区；65、新疆自治区
