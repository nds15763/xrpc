package xrpc

import ( 
	"net"
	"net/rpc"
)

type RPCServiceMap map[string]interface{}

type RPCService struct {
	server interface{}
}

func NewRPCService(serviceMap map[string]interface{}) *RPCService {
	//1.服务注册
	server = NewServer()
	return &RPCService{}
}

func (s *RPCService) InitRPCServer(addr string, secrKey string) {
	//密码初始化

	//服务初始化
	client, err := rpc.DialTCP("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }
}

func (s *RPCService) Run() {

}
