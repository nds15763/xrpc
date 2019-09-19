package xrpc

type XServer struct {
}

//NewServer 建立服务端
func NewServer(network, address string) (XServer, error) {
	//从上层开始经过 net/rpc(可能需要重写) 然后再下层传输的地方使用yamux

	//所以说需要返回一个结构体，其中包含对于RPC的所有方法

	//如 xrpc.NewClient(xxxx)

	return XServer{}, nil
}
