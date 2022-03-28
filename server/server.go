package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello," + request
	return nil
}

func main()  {
	// 实例化一个 server
	// 注册 handler
	// 启动服务
	listener, _ := net.Listen("tcp", ":1234")
	_ = rpc.RegisterName("HelloService", &HelloService{})
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)

	// 替换 rpc 序列化协议为 json
	//rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
}

