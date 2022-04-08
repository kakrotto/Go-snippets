package main

import (
	"fmt"
	"strings"
)

// 自定义的移除前缀的处理函数
func removePrefix(s string) string {
	return strings.TrimPrefix(s, "go")
}

func MyStringProcess(list []*string, chain []func(s string) string) {
	for _, s := range list {
		for _, f := range chain {
			*s = f(*s)
		}
	}
}

func main() {
	a := "go scanner"
	b := "go go parser"
	list := []*string{
		&a,
		&b,
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	MyStringProcess(list, chain)
	for _, v := range list {
		fmt.Println(*v)
	}
}
