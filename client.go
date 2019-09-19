package xrpc

import (
	oRpc "net/rpc"
)

//XClient 客户端
type XClient struct {
	rpc     *oRpc.Client
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
func NewClient(network, address string) (*XClient, error) {
	//从上层开始经过 net/rpc(可能需要重写) 然后再下层传输的地方使用yamux

	//所以说需要返回一个结构体，其中包含对于RPC的所有方法

	//如 xrpc.NewClient(xxxx)

	orpc, err := oRpc.Dial(network, address)

	if err != nil {
		return nil, err
	}

	return &XClient{rpc: orpc}, nil
}

//SetFuncMap 保存方法及参数
func (x *XClient) SetFuncMap(xfm map[string]*XFunc) {

}
