package xrpc

import (
	"errors"
	"io"
	"log"
	"net"
	oRpc "net/rpc"

	"github.com/hashicorp/yamux"
)

//XClient 客户端
type XClient struct {
	rpc     *oRpc.Client
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

// A ClientCodec implements writing of RPC requests and
// reading of RPC responses for the client side of an RPC session.
// The client calls WriteRequest to write a request to the connection
// and calls ReadResponseHeader and ReadResponseBody in pairs
// to read responses. The client calls Close when finished with the
// connection. ReadResponseBody may be called with a nil
// argument to force the body of the response to be read and then
// discarded.
// See NewClient's comment for information about concurrent access.
type ClientCodec interface {
	WriteRequest(*orpc.Request, interface{}) error
	ReadResponseHeader(*orpc.Response) error
	ReadResponseBody(interface{}) error

	Close() error
}

//NewClient 建立客户端
func NewClient(network, address string) (*XClient, error) {
	//从上层开始经过 net/rpc(可能需要重写) 然后再下层传输的地方使用yamux

	//所以说需要返回一个结构体，其中包含对于RPC的所有方法

	//如 xrpc.NewClient(xxxx)

	//一堆源码，我在想舍弃他这些东西，直接调用input方法。而且input方法也只是单纯的发送用，也不涉及一些反射

	//所以第一步还是建立连接

	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	session, _ := yamux.Client(conn, nil)

	x := NewClientWithCodec{}
	x.input()

	return x, nil
}

// NewClientWithCodec is like NewClient but uses the specified
// codec to encode requests and decode responses.
func NewClientWithCodec(codec ClientCodec) *XClient {
	x := &XClient{
		codec:   codec,
		pending: make(map[uint64]*XFunc),
	}
	go x.input()
	return x
}

func (x *XClient) input() {
	var err error
	var response Response
	for err == nil {
		response = Response{}
		err = client.codec.ReadResponseHeader(&response)
		if err != nil {
			break
		}
		seq := response.Seq
		client.mutex.Lock()
		call := client.pending[seq]
		delete(client.pending, seq)
		client.mutex.Unlock()

		switch {
		case call == nil:
			// We've got no pending call. That usually means that
			// WriteRequest partially failed, and call was already
			// removed; response is a server telling us about an
			// error reading request body. We should still attempt
			// to read error body, but there's no one to give it to.
			err = client.codec.ReadResponseBody(nil)
			if err != nil {
				err = errors.New("reading error body: " + err.Error())
			}
		case response.Error != "":
			// We've got an error response. Give this to the request;
			// any subsequent requests will get the ReadResponseBody
			// error if there is one.
			call.Error = ServerError(response.Error)
			err = client.codec.ReadResponseBody(nil)
			if err != nil {
				err = errors.New("reading error body: " + err.Error())
			}
			call.done()
		default:
			err = client.codec.ReadResponseBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	// Terminate pending calls.
	client.reqMutex.Lock()
	client.mutex.Lock()
	client.shutdown = true
	closing := client.closing
	if err == io.EOF {
		if closing {
			err = ErrShutdown
		} else {
			err = io.ErrUnexpectedEOF
		}
	}
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
	client.mutex.Unlock()
	client.reqMutex.Unlock()
	if debugLog && err != io.EOF && !closing {
		log.Println("rpc: client protocol error:", err)
	}
}

//SetFuncMap 保存方法及参数
func (x *XClient) SetFuncMap(xfm map[string]*XFunc) {

}

//Call 单次调用
func (x *XClient) Call() {

	//如果我想写成grpc的那种形式怎么做比较好

}

//Stream 建立stream链接
func (x *XClient) Stream() {

}
