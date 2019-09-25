package xrpc

import (
	"errors"
	"io"
	"log"
	"net"
	oRpc "net/rpc"

	reg "xrpc/register"
	"github.com/hashicorp/yamux"
)

//XClient 客户端
type XClient struct {
	rpc     *oRpc.Client
	consul *reg.Consul
	session yamux.Session
	funcMap map[string]*XFunc
}

//XFunc 同net/rpc的call结构，用于储存方法
type XFunc struct {
	ServiceMethod string      // The name of the service and method to call.
	Args          interface{} // The argument to the function (*struct).
	Reply         interface{} // The reply from the function (*struct).
	Error         error       // After completion, the error status.
	Done          chan *XFunc // Strobes when call is complete.
}


//NewClient 建立客户端
func NewClient(consulAddr, serviceName string) (*XClient, error) {
	//从上层开始经过 net/rpc(可能需要重写) 然后再下层传输的地方使用yamux

	//所以说需要返回一个结构体，其中包含对于RPC的所有方法

	//如 xrpc.NewClient(xxxx)

	//一堆源码，我在想舍弃他这些东西，直接调用input方法。而且input方法也只是单纯的发送用，也不涉及一些反射

	//这里需要考虑一下分层
	//第一层是链接层，向外暴露接口
	//NEW()
	//Call():普通 Client / Server 调用(北向)接口
	//C2S():Client 向 Server 的流式(北向)流式接口
	//S2C():Server 向 Cinet 调用(南向)流式接口
	//C2C():Server / Client 双向流式接口
	//SetFunc():保存方法及参数
	

	//0.5 实例化链接是否需要先去consul中注册服务和查找
	//注册服务应该是微服务启动时准备的
	//这个我要再看一下，是通过什么能直接访问到负载均衡，是consul+nginx还是怎么
	//如果是nginx的话就需要在consul中存一个路由地址
	reg.DoDiscover(consulAddr, serviceName)

	if reg.ServicsMap[serviceName] == nil{
		return nil,err.Errors("Server Not Found")
	}
	
	
	//1 实例化链接，先调用net/rpc建立实体链接
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	//2 实例化链接之后 可能需要封装源码的调用，中间加一层yamux
	//session, _ := yamux.Client(conn, nil)

	//3 建立连接之后去注册
	x := &XClient{session: session}


	return x, nil
}


//SetFuncMap 保存方法及参数
func (x *XClient) SetFuncMap(xfm map[string]*XFunc) {

}

//Call 普通 Client / Server 调用(北向)接口
func (x *XClient) Call() {

	//todo

}

//C2S() Client 向 Server 的流式(北向)流式接口
func (x *XClient) C2S() {

	//todo

}

//Call 普通 Client / Server 调用(北向)接口
func (x *XClient) C2C() {

	//todo

}