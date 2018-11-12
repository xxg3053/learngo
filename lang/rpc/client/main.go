package main

import (
	"net"
	"net/rpc/jsonrpc"
	"github.com/xxg3053/learngo/lang/rpc"
	"fmt"
)

func main()  {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil{
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{3,4}, &result)
	if err != nil{
		panic(err)
	}else{
		fmt.Println(result)
	}
}
