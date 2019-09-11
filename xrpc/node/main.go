package main

import (
	"fmt"
	"reflect"
)

// import (
// 	"stest/server"
// )

// func main() {

// 	//1.注册service
// 	server := server.NewRPCService(server.RPCServiceMap{"stest": ""})

// 	//2.加密方式

// 	//3.服务初始化
// 	server.InitRPCServer(":8081")

// 	//4.启动服务
// 	server.Run()

// 	//5.长连接注册

// 	//6.查找注册链接的节点，并进行通讯
// }

func hello(i int, s string) {
	fmt.Println("Hello world!")
}

func main() {
	hl := hello
	fv := reflect.ValueOf(hl)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(1, "2")
}
