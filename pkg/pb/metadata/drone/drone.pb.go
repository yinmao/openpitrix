// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/drone/drone.proto

package pbdrone // import "openpitrix.io/openpitrix/pkg/pb/metadata/drone"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "openpitrix.io/openpitrix/pkg/pb/metadata/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DroneServiceClient is the client API for DroneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DroneServiceClient interface {
	GetDroneConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.DroneConfig, error)
	SetDroneConfig(ctx context.Context, in *types.DroneConfig, opts ...grpc.CallOption) (*types.Empty, error)
	GetConfdConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.ConfdConfig, error)
	SetConfdConfig(ctx context.Context, in *types.ConfdConfig, opts ...grpc.CallOption) (*types.Empty, error)
	GetFrontgateConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.FrontgateConfig, error)
	SetFrontgateConfig(ctx context.Context, in *types.FrontgateConfig, opts ...grpc.CallOption) (*types.Empty, error)
	IsConfdRunning(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Bool, error)
	StartConfd(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	StopConfd(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	GetTemplateFiles(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.StringList, error)
	GetValues(ctx context.Context, in *types.StringList, opts ...grpc.CallOption) (*types.StringMap, error)
	PingPilot(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	PingFrontgate(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	PingDrone(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	RunCommand(ctx context.Context, in *types.RunCommandOnDroneRequest, opts ...grpc.CallOption) (*types.String, error)
}

type droneServiceClient struct {
	cc *grpc.ClientConn
}

func NewDroneServiceClient(cc *grpc.ClientConn) DroneServiceClient {
	return &droneServiceClient{cc}
}

func (c *droneServiceClient) GetDroneConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.DroneConfig, error) {
	out := new(types.DroneConfig)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetDroneConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) SetDroneConfig(ctx context.Context, in *types.DroneConfig, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/SetDroneConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetConfdConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.ConfdConfig, error) {
	out := new(types.ConfdConfig)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetConfdConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) SetConfdConfig(ctx context.Context, in *types.ConfdConfig, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/SetConfdConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetFrontgateConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.FrontgateConfig, error) {
	out := new(types.FrontgateConfig)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetFrontgateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) SetFrontgateConfig(ctx context.Context, in *types.FrontgateConfig, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/SetFrontgateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) IsConfdRunning(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Bool, error) {
	out := new(types.Bool)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/IsConfdRunning", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) StartConfd(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/StartConfd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) StopConfd(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/StopConfd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetTemplateFiles(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.StringList, error) {
	out := new(types.StringList)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetTemplateFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetValues(ctx context.Context, in *types.StringList, opts ...grpc.CallOption) (*types.StringMap, error) {
	out := new(types.StringMap)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/GetValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingPilot(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/PingPilot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingFrontgate(ctx context.Context, in *types.FrontgateEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/PingFrontgate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingDrone(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/PingDrone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) RunCommand(ctx context.Context, in *types.RunCommandOnDroneRequest, opts ...grpc.CallOption) (*types.String, error) {
	out := new(types.String)
	err := c.cc.Invoke(ctx, "/metadata.drone.DroneService/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DroneServiceServer is the server API for DroneService service.
type DroneServiceServer interface {
	GetDroneConfig(context.Context, *types.Empty) (*types.DroneConfig, error)
	SetDroneConfig(context.Context, *types.DroneConfig) (*types.Empty, error)
	GetConfdConfig(context.Context, *types.Empty) (*types.ConfdConfig, error)
	SetConfdConfig(context.Context, *types.ConfdConfig) (*types.Empty, error)
	GetFrontgateConfig(context.Context, *types.Empty) (*types.FrontgateConfig, error)
	SetFrontgateConfig(context.Context, *types.FrontgateConfig) (*types.Empty, error)
	IsConfdRunning(context.Context, *types.Empty) (*types.Bool, error)
	StartConfd(context.Context, *types.Empty) (*types.Empty, error)
	StopConfd(context.Context, *types.Empty) (*types.Empty, error)
	GetTemplateFiles(context.Context, *types.Empty) (*types.StringList, error)
	GetValues(context.Context, *types.StringList) (*types.StringMap, error)
	PingPilot(context.Context, *types.FrontgateEndpoint) (*types.Empty, error)
	PingFrontgate(context.Context, *types.FrontgateEndpoint) (*types.Empty, error)
	PingDrone(context.Context, *types.Empty) (*types.Empty, error)
	RunCommand(context.Context, *types.RunCommandOnDroneRequest) (*types.String, error)
}

func RegisterDroneServiceServer(s *grpc.Server, srv DroneServiceServer) {
	s.RegisterService(&_DroneService_serviceDesc, srv)
}

func _DroneService_GetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetDroneConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_SetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).SetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/SetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).SetDroneConfig(ctx, req.(*types.DroneConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetConfdConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetConfdConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetConfdConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetConfdConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_SetConfdConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ConfdConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).SetConfdConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/SetConfdConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).SetConfdConfig(ctx, req.(*types.ConfdConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetFrontgateConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_SetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).SetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/SetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).SetFrontgateConfig(ctx, req.(*types.FrontgateConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_IsConfdRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).IsConfdRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/IsConfdRunning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).IsConfdRunning(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_StartConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).StartConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/StartConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).StartConfd(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_StopConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).StopConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/StopConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).StopConfd(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetTemplateFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetTemplateFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetTemplateFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetTemplateFiles(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StringList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/GetValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetValues(ctx, req.(*types.StringList))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingPilot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingPilot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/PingPilot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingPilot(ctx, req.(*types.FrontgateEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingFrontgate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingFrontgate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/PingFrontgate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingFrontgate(ctx, req.(*types.FrontgateEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/PingDrone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingDrone(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.RunCommandOnDroneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.drone.DroneService/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).RunCommand(ctx, req.(*types.RunCommandOnDroneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DroneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metadata.drone.DroneService",
	HandlerType: (*DroneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDroneConfig",
			Handler:    _DroneService_GetDroneConfig_Handler,
		},
		{
			MethodName: "SetDroneConfig",
			Handler:    _DroneService_SetDroneConfig_Handler,
		},
		{
			MethodName: "GetConfdConfig",
			Handler:    _DroneService_GetConfdConfig_Handler,
		},
		{
			MethodName: "SetConfdConfig",
			Handler:    _DroneService_SetConfdConfig_Handler,
		},
		{
			MethodName: "GetFrontgateConfig",
			Handler:    _DroneService_GetFrontgateConfig_Handler,
		},
		{
			MethodName: "SetFrontgateConfig",
			Handler:    _DroneService_SetFrontgateConfig_Handler,
		},
		{
			MethodName: "IsConfdRunning",
			Handler:    _DroneService_IsConfdRunning_Handler,
		},
		{
			MethodName: "StartConfd",
			Handler:    _DroneService_StartConfd_Handler,
		},
		{
			MethodName: "StopConfd",
			Handler:    _DroneService_StopConfd_Handler,
		},
		{
			MethodName: "GetTemplateFiles",
			Handler:    _DroneService_GetTemplateFiles_Handler,
		},
		{
			MethodName: "GetValues",
			Handler:    _DroneService_GetValues_Handler,
		},
		{
			MethodName: "PingPilot",
			Handler:    _DroneService_PingPilot_Handler,
		},
		{
			MethodName: "PingFrontgate",
			Handler:    _DroneService_PingFrontgate_Handler,
		},
		{
			MethodName: "PingDrone",
			Handler:    _DroneService_PingDrone_Handler,
		},
		{
			MethodName: "RunCommand",
			Handler:    _DroneService_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata/drone/drone.proto",
}

func init() { proto.RegisterFile("metadata/drone/drone.proto", fileDescriptor_drone_28bf75ca0199208f) }

var fileDescriptor_drone_28bf75ca0199208f = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x6b, 0xdb, 0x30,
	0x14, 0xbf, 0x8d, 0xe5, 0xb1, 0x99, 0x21, 0xb6, 0xc1, 0x34, 0xd8, 0xd8, 0x6d, 0x27, 0x1b, 0x36,
	0x18, 0xcb, 0x46, 0x2f, 0x49, 0x1c, 0x13, 0x48, 0xdb, 0x10, 0x97, 0x1e, 0x7a, 0x53, 0xe2, 0x17,
	0x23, 0x6a, 0x4b, 0xaa, 0xfd, 0x5c, 0x9a, 0x6f, 0xd3, 0x8f, 0x5a, 0x62, 0xe7, 0x8f, 0xe3, 0x44,
	0x2e, 0x69, 0x2f, 0x32, 0xfe, 0xfd, 0xd3, 0x4f, 0xcf, 0x58, 0xc0, 0x53, 0x24, 0x11, 0x09, 0x12,
	0x5e, 0x94, 0x69, 0x85, 0xd5, 0xea, 0x9a, 0x4c, 0x93, 0x66, 0xce, 0x86, 0x73, 0x4b, 0x94, 0xef,
	0xb4, 0xb4, 0x34, 0x98, 0x57, 0x6b, 0xa5, 0x3d, 0xe0, 0xe6, 0x5a, 0x2d, 0x22, 0x0b, 0x57, 0xdb,
	0x83, 0x7f, 0x6b, 0x70, 0x8b, 0x4c, 0x2b, 0x8a, 0x05, 0xad, 0xf9, 0x5f, 0x8f, 0x6f, 0xe1, 0xdd,
	0x60, 0xa5, 0x0f, 0x31, 0xbb, 0x97, 0x73, 0x64, 0x03, 0x70, 0x02, 0xa4, 0x12, 0xea, 0x6b, 0xb5,
	0x90, 0x31, 0xfb, 0xe4, 0x6e, 0x7b, 0x56, 0x8d, 0xfc, 0xd4, 0xd0, 0x92, 0x7f, 0x6d, 0xc2, 0x75,
	0xcf, 0x00, 0x9c, 0x70, 0x3f, 0xa5, 0x4d, 0xce, 0x8f, 0x6f, 0xb1, 0xee, 0xb2, 0xd2, 0x44, 0x27,
	0x76, 0xa9, 0x7b, 0xaa, 0x2e, 0x75, 0xa4, 0x4d, 0x6e, 0xeb, 0x32, 0x06, 0x16, 0x20, 0x0d, 0x37,
	0xe3, 0x6b, 0xef, 0xf3, 0xbd, 0x09, 0x37, 0x7d, 0x63, 0x60, 0xe1, 0x61, 0xda, 0x73, 0x36, 0x5b,
	0xb7, 0x33, 0x70, 0x46, 0x79, 0x79, 0x86, 0x69, 0xa1, 0x94, 0x54, 0xd6, 0x5e, 0x1f, 0x9b, 0x70,
	0x4f, 0xeb, 0x84, 0xfd, 0x03, 0x08, 0x49, 0x64, 0xd5, 0x88, 0x6c, 0x56, 0xcb, 0xd6, 0x5d, 0xe8,
	0x84, 0xa4, 0xcd, 0x4b, 0xac, 0x3e, 0x7c, 0x08, 0x90, 0xae, 0x30, 0x35, 0x89, 0x20, 0x1c, 0xca,
	0x04, 0x73, 0x5b, 0x02, 0x6f, 0xc2, 0x21, 0x65, 0x52, 0xc5, 0x63, 0x99, 0x13, 0xeb, 0x41, 0x27,
	0x40, 0xba, 0x16, 0x49, 0x81, 0x39, 0x6b, 0x11, 0xf2, 0x2f, 0xc7, 0xb9, 0x73, 0x61, 0x98, 0x0f,
	0x9d, 0x89, 0x54, 0xf1, 0x44, 0x26, 0x9a, 0xd8, 0x0f, 0xeb, 0x57, 0xf0, 0x55, 0x64, 0xb4, 0x54,
	0x64, 0x3b, 0xd1, 0x08, 0xde, 0xaf, 0x62, 0xb6, 0xfa, 0x57, 0x44, 0x75, 0xab, 0x46, 0xe5, 0x4f,
	0x72, 0xe2, 0x5c, 0x2f, 0x00, 0xa6, 0x85, 0xea, 0xeb, 0x34, 0x15, 0x2a, 0x62, 0x3f, 0x9b, 0xa2,
	0x1d, 0x77, 0xa9, 0xca, 0xf8, 0x29, 0xde, 0x15, 0x98, 0x13, 0xff, 0x7c, 0x7c, 0x3e, 0xbd, 0xbf,
	0x37, 0x7f, 0xb4, 0x41, 0x65, 0x24, 0x65, 0xf2, 0xc1, 0x95, 0xda, 0xdb, 0xbd, 0x79, 0xe6, 0x36,
	0xf6, 0xcc, 0xcc, 0xdb, 0xbf, 0xe4, 0xfe, 0x9b, 0x59, 0xf9, 0x9c, 0xbd, 0x29, 0xef, 0x98, 0xdf,
	0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x5a, 0xf3, 0x09, 0x05, 0x05, 0x00, 0x00,
}
