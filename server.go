package xrpc

type XServer struct {
}

//NewServer 建立服务端
func NewServer(network, address string) (XServer, error) {
	//从上层开始经过 net/rpc(可能需要重写) 然后再下层传输的地方使用yamux

	//所以说需要返回一个结构体，其中包含对于RPC的所有方法

	//第一层是链接层，向外暴露接口
	//NEW()
	//Call():普通 Client / Server 调用(北向)接口
	//C2S():Client 向 Server 的流式(北向)流式接口
	//S2C():Server 向 Cinet 调用(南向)流式接口
	//S2S():Server / Client 双向流式接口
	//SetFunc():保存方法及参数

	return XServer{}, nil
}

//S2C Server向Client调用(南向)流式接口
func (x *XServer) S2C() {

	//todo

}
