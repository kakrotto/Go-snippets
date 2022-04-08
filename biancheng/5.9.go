package main

import (
	"bytes"
	"fmt"
	"strings"
)

// MyFunc 从内部实现机理上来说，类型 ...type 本质上是一个数组切片，也就是 []type
func MyFunc(s string, i int, list ...string) {
	fmt.Printf("%+v %+v %+v\n", s, i, list)
}

func joinStrings(list ...string) string {
	return strings.Join(list, ",")
}

func MyJoinStrings(list ...string) string {
	var b bytes.Buffer
	for _, s := range list {
		b.WriteString(s)
	}
	return b.String()
}

func MyPrint(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is int")
		case string:
			fmt.Println(arg, "is string")
		case bool:
			fmt.Println(arg, "is bool")
		default:
			fmt.Println(arg, "is unknown")
		}
	}
}

// 获得可变参数类型，获得每一个参数的类型
func printTypeValue(list ...interface{}) string {

	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer

	// 遍历参数
	for _, s := range list {

		// 将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)

		// 类型的字符串描述
		var typeString string

		// 对s进行类型断言
		switch s.(type) {
		case bool: // 当s为布尔类型时
			typeString = "bool"
		case string: // 当s为字符串类型时
			typeString = "string"
		case int: // 当s为整型类型时
			typeString = "int"
		default:
			typeString = "unknown"
		}

		// 写字符串前缀
		b.WriteString("value: ")

		// 写入值
		b.WriteString(str)

		// 写类型前缀
		b.WriteString(" type: ")

		// 写类型字符串
		b.WriteString(typeString)

		// 写入换行符
		b.WriteString("\n")

	}
	return b.String()
}

// 在多个可变参数函数中传递参数
// 实际打印的函数
func rawPrint(rawList ...interface{}) {

	// 遍历可变参数切片
	for _, a := range rawList {

		// 打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func prints(list ...interface{}) {
	// 将 list 可变参数切片完整传递给下一个函数
	rawPrint(list...)
}

func main() {
	MyFunc("hello", 1, "world")
	MyFunc("hello", 1, "world", "go")
	MyPrint(1, "hello", 1.2, true)
	fmt.Println(joinStrings("hello", "world"))
	fmt.Println(MyJoinStrings("你好", "golang"))
	fmt.Println(printTypeValue(1, "hello", 1.2, true))
	prints(1, 2, 3)
}
