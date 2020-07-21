package utils

import (
	"log"
	"time"
)

//2006-01-02 15:04:05
func DateTimeNowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//处理时间字符串转时间
func StrTimeTotime(strTime string) time.Time {

	const Layout = "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")

	tim, _ := time.ParseInLocation(Layout, strTime /*需要转换的时间类型字符串*/, loc)
	log.Println(tim)
	return tim
}

//处理时间字符串转时间
func StrTimeToNowtime() time.Time {
	strTime := time.Now().Format("2006-01-02 15:04:05")
	const Layout = "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")

	tim, _ := time.ParseInLocation(Layout, strTime /*需要转换的时间类型字符串*/, loc)
	log.Println(tim)
	return tim
}

//处理时间字符串转时间
func StrTimeTodefaultdate() time.Time {
	strTime := "2020-01-01 00:00:00"
	const Layout = "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")

	tim, _ := time.ParseInLocation(Layout, strTime /*需要转换的时间类型字符串*/, loc)
	log.Println(tim)
	return tim
}
