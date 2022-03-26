package main

import (
	"fmt"
	"net/rpc"
)

func main()  {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil{
		panic("链接失败")
	}
	// var reply *string = new(string)  默认值 nil，所以需要 new(string) 分配内存
	var reply string
	err = client.Call("HelloService.Hello", "allen", &reply)
	if err != nil{
		panic("调用失败")
	}
	fmt.Println(reply)
}