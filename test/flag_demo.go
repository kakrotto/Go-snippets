// flag 包常用方法
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	// &user 就是接收用户命令行中输入的 -u 后面的参数值 value：默认值，usage：用户提示信息
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为 localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为 3306")
	flag.Parse()
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%d\n", user, pwd, host, port)
	// 启动命令：go run flag_demo.go  -u=root -pwd=123456 -h=127.0.0.1 -port=3306
}
