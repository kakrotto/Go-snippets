// 介绍 strconv 包的常用方法
// strconv 提供了字符串和基本数据类型之间的转换功能。

/*
Itoa()  func Itoa(i int) string
Atoi()  func Atoi(s string) (i int, err error)
 */
package main

import (
	"fmt"
	"strconv"
)

func intToStr(i int) string {
	return strconv.Itoa(i)
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil{
		fmt.Printf("%v 转换失败！", s)
		return -1
	}
	return i
}

func main() {
	var i int64 = 100
	fmt.Println(strconv.Itoa(int(i))+"all")
}
