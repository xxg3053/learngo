package main

import (
	"net/rpc"
	"github.com/xxg3053/learngo/lang/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

// telnet localhost 1234
// {"method":"DemoService.Div", "params":[{"A":3,"B":4}], "id":1}
func main()  {
	rpc.Register(rpcdemo.DemoService{})

	listerer, err := net.Listen("tcp", ":1234")
	if err != nil{
		panic(err)
	}
	for{
		conn, err := listerer.Accept()
		if err != nil{
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
