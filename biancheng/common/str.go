/*
字符串相关操作示例
*/
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"unicode/utf8"
	"unsafe"
)

// Go 语言的内建函数 len()，可以用来获取切片、字符串、通道（channel）等的长度
// len() 函数的返回值的类型为 int，表示字符串的 ASCII 字符个数或字节长度。
// Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节

func LenStr(s string) int {
	/* ASCII 字符串长度使用 len() 函数。
	   Unicode 字符串长度使用 utf8.RuneCountInString() 函数。*/
	return utf8.RuneCountInString(s)
}

func TraversalStr(s string) {
	// ASCII 字符串遍历直接使用下标。
	// Unicode 字符串遍历用 for range。
	for _, v := range s {
		fmt.Printf("%c, %d \n", v, v)
	}
}

func SpliceStr(str1, str2 string) string {
	// 声明字节缓冲
	var stringBuilder bytes.Buffer
	// 把字符串写入缓冲
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)
	// 将缓冲以字符串形式输出
	return stringBuilder.String()
}

func base64Str(){

	// 需要处理的字符串
	message := "Away from keyboard. https://golang.org/"

	// 编码消息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte (message))

	// 输出编码完成的消息
	fmt.Println(encodedMessage)

	// 解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMessage)

	// 出错处理
	if err != nil {
		fmt.Println(err)
	} else {
		// 打印解码完成的数据
		fmt.Println(string(data))
	}

}
func main(){
	var str = "去打游戏不要加班了 go"
	fmt.Println("原字符串:", str)
	fmt.Println(LenStr(str))
	TraversalStr(str)
	var s1, s2 = "allen", "赵"
	fmt.Println(SpliceStr(s1, s2))
	base64Str()
	var a uint8
	fmt.Printf("%T, %d \n", a, unsafe.Sizeof(a))
}
