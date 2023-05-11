// Code generated by protoc-gen-rpcx. DO NOT EDIT.
// versions:
// - protoc-gen-rpcx v0.3.0
// - protoc          (unknown)
// source: http/rest/channel/channel.proto

package channel

import (
	context "context"
	client "github.com/smallnest/rpcx/client"
	protocol "github.com/smallnest/rpcx/protocol"
	server "github.com/smallnest/rpcx/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO
var _ = server.NewServer
var _ = client.NewClient
var _ = protocol.NewMessage

//================== interface skeleton ===================
type ChannelAble interface {
	// ChannelAble can be used for interface verification.

	// Create is server rpc method as defined
	Create(ctx context.Context, args *CreateReq, reply *CreateRsp) (err error)

	// Add is server rpc method as defined
	Add(ctx context.Context, args *AddReq, reply *AddRsp) (err error)
}

//================== server skeleton ===================
type ChannelImpl struct{}

// ServeForChannel starts a server only registers one service.
// You can register more services and only start one server.
// It blocks until the application exits.
func ServeForChannel(addr string) error {
	s := server.NewServer()
	s.RegisterName("Channel", new(ChannelImpl), "")
	return s.Serve("tcp", addr)
}

// Create is server rpc method as defined
func (s *ChannelImpl) Create(ctx context.Context, args *CreateReq, reply *CreateRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = CreateRsp{}

	return nil
}

// Add is server rpc method as defined
func (s *ChannelImpl) Add(ctx context.Context, args *AddReq, reply *AddRsp) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = AddRsp{}

	return nil
}

//================== client stub ===================
// Channel is a client wrapped XClient.
type ChannelClient struct {
	xclient client.XClient
}

// NewChannelClient wraps a XClient as ChannelClient.
// You can pass a shared XClient object created by NewXClientForChannel.
func NewChannelClient(xclient client.XClient) *ChannelClient {
	return &ChannelClient{xclient: xclient}
}

// NewXClientForChannel creates a XClient.
// You can configure this client with more options such as etcd registry, serialize type, select algorithm and fail mode.
func NewXClientForChannel(addr string) (client.XClient, error) {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	if err != nil {
		return nil, err
	}

	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient("Channel", client.Failtry, client.RoundRobin, d, opt)

	return xclient, nil
}

// Create is client rpc method as defined
func (c *ChannelClient) Create(ctx context.Context, args *CreateReq) (reply *CreateRsp, err error) {
	reply = &CreateRsp{}
	err = c.xclient.Call(ctx, "Create", args, reply)
	return reply, err
}

// Add is client rpc method as defined
func (c *ChannelClient) Add(ctx context.Context, args *AddReq) (reply *AddRsp, err error) {
	reply = &AddRsp{}
	err = c.xclient.Call(ctx, "Add", args, reply)
	return reply, err
}

//================== oneclient stub ===================
// ChannelOneClient is a client wrapped oneClient.
type ChannelOneClient struct {
	serviceName string
	oneclient   *client.OneClient
}

// NewChannelOneClient wraps a OneClient as ChannelOneClient.
// You can pass a shared OneClient object created by NewOneClientForChannel.
func NewChannelOneClient(oneclient *client.OneClient) *ChannelOneClient {
	return &ChannelOneClient{
		serviceName: "Channel",
		oneclient:   oneclient,
	}
}

// ======================================================

// Create is client rpc method as defined
func (c *ChannelOneClient) Create(ctx context.Context, args *CreateReq) (reply *CreateRsp, err error) {
	reply = &CreateRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "Create", args, reply)
	return reply, err
}

// Add is client rpc method as defined
func (c *ChannelOneClient) Add(ctx context.Context, args *AddReq) (reply *AddRsp, err error) {
	reply = &AddRsp{}
	err = c.oneclient.Call(ctx, c.serviceName, "Add", args, reply)
	return reply, err
}