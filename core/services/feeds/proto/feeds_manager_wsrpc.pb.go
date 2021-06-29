// Code generated by protoc-gen-go-wsrpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-wsrpc v0.0.1
// - protoc             v3.15.8

package proto

import (
	context "context"
	wsrpc "github.com/smartcontractkit/wsrpc"
)

// FeedsManagerClient is the client API for FeedsManager service.
//
type FeedsManagerClient interface {
	ApproveJob(ctx context.Context, in *ApproveJobRequest) (*ApproveJobResponse, error)
}

type feedsManagerClient struct {
	cc wsrpc.ClientInterface
}

func NewFeedsManagerClient(cc wsrpc.ClientInterface) FeedsManagerClient {
	return &feedsManagerClient{cc}
}

func (c *feedsManagerClient) ApproveJob(ctx context.Context, in *ApproveJobRequest) (*ApproveJobResponse, error) {
	out := new(ApproveJobResponse)
	err := c.cc.Invoke(ctx, "ApproveJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedsManagerServer is the server API for FeedsManager service.
type FeedsManagerServer interface {
	ApproveJob(context.Context, *ApproveJobRequest) (*ApproveJobResponse, error)
}

func RegisterFeedsManagerServer(s wsrpc.ServiceRegistrar, srv FeedsManagerServer) {
	s.RegisterService(&FeedsManager_ServiceDesc, srv)
}

func _FeedsManager_ApproveJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ApproveJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(FeedsManagerServer).ApproveJob(ctx, in)
}

// FeedsManager_ServiceDesc is the wsrpc.ServiceDesc for FeedsManager service.
// It's only intended for direct use with wsrpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedsManager_ServiceDesc = wsrpc.ServiceDesc{
	ServiceName: "cfm.FeedsManager",
	HandlerType: (*FeedsManagerServer)(nil),
	Methods: []wsrpc.MethodDesc{
		{
			MethodName: "ApproveJob",
			Handler:    _FeedsManager_ApproveJob_Handler,
		},
	},
}

// NodeServiceClient is the client API for NodeService service.
//
type NodeServiceClient interface {
	ProposeJob(ctx context.Context, in *ProposeJobRequest) (*ProposeJobResponse, error)
}

type nodeServiceClient struct {
	cc wsrpc.ClientInterface
}

func NewNodeServiceClient(cc wsrpc.ClientInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) ProposeJob(ctx context.Context, in *ProposeJobRequest) (*ProposeJobResponse, error) {
	out := new(ProposeJobResponse)
	err := c.cc.Invoke(ctx, "ProposeJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	ProposeJob(context.Context, *ProposeJobRequest) (*ProposeJobResponse, error)
}

func RegisterNodeServiceServer(s wsrpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_ProposeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProposeJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(NodeServiceServer).ProposeJob(ctx, in)
}

// NodeService_ServiceDesc is the wsrpc.ServiceDesc for NodeService service.
// It's only intended for direct use with wsrpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = wsrpc.ServiceDesc{
	ServiceName: "cfm.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []wsrpc.MethodDesc{
		{
			MethodName: "ProposeJob",
			Handler:    _NodeService_ProposeJob_Handler,
		},
	},
}
