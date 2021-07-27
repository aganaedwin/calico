// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cnibackend.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	cnibackend.proto

It has these top-level messages:
	AddRequest
	ContainerSettings
	IPConfig
	WorkloadIDs
	Port
	AddReply
	DelRequest
	DelReply
*/
package proto

import (
	fmt "fmt"

	proto1 "github.com/gogo/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AddRequest struct {
	InterfaceName            string             `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Netns                    string             `protobuf:"bytes,2,opt,name=netns,proto3" json:"netns,omitempty"`
	DesiredHostInterfaceName string             `protobuf:"bytes,3,opt,name=desired_host_interface_name,json=desiredHostInterfaceName,proto3" json:"desired_host_interface_name,omitempty"`
	Settings                 *ContainerSettings `protobuf:"bytes,4,opt,name=settings" json:"settings,omitempty"`
	ContainerIps             []*IPConfig        `protobuf:"bytes,5,rep,name=container_ips,json=containerIps" json:"container_ips,omitempty"`
	ContainerRoutes          []string           `protobuf:"bytes,6,rep,name=container_routes,json=containerRoutes" json:"container_routes,omitempty"`
	Workload                 *WorkloadIDs       `protobuf:"bytes,7,opt,name=workload" json:"workload,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto1.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptorCnibackend, []int{0} }

func (m *AddRequest) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *AddRequest) GetNetns() string {
	if m != nil {
		return m.Netns
	}
	return ""
}

func (m *AddRequest) GetDesiredHostInterfaceName() string {
	if m != nil {
		return m.DesiredHostInterfaceName
	}
	return ""
}

func (m *AddRequest) GetSettings() *ContainerSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *AddRequest) GetContainerIps() []*IPConfig {
	if m != nil {
		return m.ContainerIps
	}
	return nil
}

func (m *AddRequest) GetContainerRoutes() []string {
	if m != nil {
		return m.ContainerRoutes
	}
	return nil
}

func (m *AddRequest) GetWorkload() *WorkloadIDs {
	if m != nil {
		return m.Workload
	}
	return nil
}

type ContainerSettings struct {
	AllowIpForwarding bool  `protobuf:"varint,1,opt,name=allow_ip_forwarding,json=allowIpForwarding,proto3" json:"allow_ip_forwarding,omitempty"`
	Mtu               int32 `protobuf:"varint,2,opt,name=mtu,proto3" json:"mtu,omitempty"`
}

func (m *ContainerSettings) Reset()                    { *m = ContainerSettings{} }
func (m *ContainerSettings) String() string            { return proto1.CompactTextString(m) }
func (*ContainerSettings) ProtoMessage()               {}
func (*ContainerSettings) Descriptor() ([]byte, []int) { return fileDescriptorCnibackend, []int{1} }

func (m *ContainerSettings) GetAllowIpForwarding() bool {
	if m != nil {
		return m.AllowIpForwarding
	}
	return false
}

func (m *ContainerSettings) GetMtu() int32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

type IPConfig struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Gateway string `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (m *IPConfig) Reset()                    { *m = IPConfig{} }
func (m *IPConfig) String() string            { return proto1.CompactTextString(m) }
func (*IPConfig) ProtoMessage()               {}
func (*IPConfig) Descriptor() ([]byte, []int) { return fileDescriptorCnibackend, []int{2} }

func (m *IPConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPConfig) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

type WorkloadIDs struct {
	Name         string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace    string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels       map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations  map[string]string `protobuf:"bytes,4,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Endpoint     string            `protobuf:"bytes,5,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Node         string            `protobuf:"bytes,6,opt,name=node,proto3" json:"node,omitempty"`
	Orchestrator string            `protobuf:"bytes,7,opt,name=orchestrator,proto3" json:"orchestrator,omitempty"`
	Pod          string            `protobuf:"bytes,8,opt,name=pod,proto3" json:"pod,omitempty"`
	Ports        []*Port           `protobuf:"bytes,9,rep,name=ports" json:"ports,omitempty"`
}

func (m *WorkloadIDs) Reset()                    { *m = WorkloadIDs{} }
func (m *WorkloadIDs) String() string            { return proto1.CompactTextString(m) }
func (*WorkloadIDs) ProtoMessage()               {}
func (*WorkloadIDs) Descriptor() ([]byte, []int) { return fileDescriptorCnibackend, []int{3} }

func (m *WorkloadIDs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkloadIDs) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WorkloadIDs) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WorkloadIDs) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *WorkloadIDs) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *WorkloadIDs) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *WorkloadIDs) GetOrchestrator() string {
	if m != nil {
		return m.Orchestrator
	}
	return ""
}

func (m *WorkloadIDs) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *WorkloadIDs) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type Port struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Port     uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *Port) Reset()                    { *m = Port{} }
func (m *Port) String() string            { return proto1.CompactTextString(m) }
func (*Port) ProtoMessage()               {}
func (*Port) Descriptor() ([]byte, []int) { return fileDescriptorCnibackend, []int{4} }

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type AddReply struct {
	Successful        bool   `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
	ErrorMessage      string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	HostInterfaceName string `protobuf:"bytes,3,opt,name=host_interface_name,json=hostInterfaceName,proto3" json:"host_interface_name,omitempty"`
	ContainerMac      string `protobuf:"bytes,4,opt,name=container_mac,json=containerMac,proto3" json:"container_mac,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto1.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptorCnibackend, []int{5} }

func (m *AddReply) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *AddReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *AddReply) GetHostInterfaceName() string {
	if m != nil {
		return m.HostInterfaceName
	}
	return ""
}

func (m *AddReply) GetContainerMac() string {
	if m != nil {
		return m.ContainerMac
	}
	return ""
}

type DelRequest struct {
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Netns         string `protobuf:"bytes,2,opt,name=netns,proto3" json:"netns,omitempty"`
}

func (m *DelRequest) Reset()                    { *m = DelRequest{} }
func (m *DelRequest) String() string            { return proto1.CompactTextString(m) }
func (*DelRequest) ProtoMessage()               {}
func (*DelRequest) Descriptor() ([]byte, []int) { return fileDescriptorCnibackend, []int{6} }

func (m *DelRequest) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *DelRequest) GetNetns() string {
	if m != nil {
		return m.Netns
	}
	return ""
}

type DelReply struct {
	Successful   bool   `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (m *DelReply) Reset()                    { *m = DelReply{} }
func (m *DelReply) String() string            { return proto1.CompactTextString(m) }
func (*DelReply) ProtoMessage()               {}
func (*DelReply) Descriptor() ([]byte, []int) { return fileDescriptorCnibackend, []int{7} }

func (m *DelReply) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *DelReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto1.RegisterType((*AddRequest)(nil), "cni.AddRequest")
	proto1.RegisterType((*ContainerSettings)(nil), "cni.ContainerSettings")
	proto1.RegisterType((*IPConfig)(nil), "cni.IPConfig")
	proto1.RegisterType((*WorkloadIDs)(nil), "cni.WorkloadIDs")
	proto1.RegisterType((*Port)(nil), "cni.Port")
	proto1.RegisterType((*AddReply)(nil), "cni.AddReply")
	proto1.RegisterType((*DelRequest)(nil), "cni.DelRequest")
	proto1.RegisterType((*DelReply)(nil), "cni.DelReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CniDataplane service

type CniDataplaneClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelReply, error)
}

type cniDataplaneClient struct {
	cc *grpc.ClientConn
}

func NewCniDataplaneClient(cc *grpc.ClientConn) CniDataplaneClient {
	return &cniDataplaneClient{cc}
}

func (c *cniDataplaneClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/cni.CniDataplane/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cniDataplaneClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelReply, error) {
	out := new(DelReply)
	err := grpc.Invoke(ctx, "/cni.CniDataplane/Del", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CniDataplane service

type CniDataplaneServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
	Del(context.Context, *DelRequest) (*DelReply, error)
}

func RegisterCniDataplaneServer(s *grpc.Server, srv CniDataplaneServer) {
	s.RegisterService(&_CniDataplane_serviceDesc, srv)
}

func _CniDataplane_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CniDataplaneServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cni.CniDataplane/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CniDataplaneServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CniDataplane_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CniDataplaneServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cni.CniDataplane/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CniDataplaneServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CniDataplane_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cni.CniDataplane",
	HandlerType: (*CniDataplaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CniDataplane_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _CniDataplane_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cnibackend.proto",
}

func init() { proto1.RegisterFile("cnibackend.proto", fileDescriptorCnibackend) }

var fileDescriptorCnibackend = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0x6d, 0xea, 0x26, 0x75, 0x26, 0xc9, 0x6d, 0xba, 0xf7, 0xea, 0xca, 0x0a, 0x15, 0x04, 0x57,
	0x88, 0x20, 0xa1, 0x3c, 0x04, 0x1e, 0x00, 0x89, 0x4a, 0x25, 0x01, 0x11, 0x44, 0xa1, 0x32, 0x42,
	0x48, 0xbc, 0x84, 0xad, 0x77, 0x92, 0x5a, 0x75, 0x76, 0xcd, 0xee, 0x86, 0x28, 0xff, 0xc3, 0x67,
	0xf0, 0x03, 0xfc, 0x15, 0xda, 0x89, 0xe3, 0xa4, 0x2d, 0x3c, 0x20, 0xf5, 0xc9, 0x33, 0x67, 0xce,
	0x8c, 0xce, 0xcc, 0xec, 0x18, 0x9a, 0xb1, 0x4c, 0xce, 0x78, 0x7c, 0x81, 0x52, 0x74, 0x33, 0xad,
	0xac, 0x62, 0x5e, 0x2c, 0x93, 0xf0, 0xe7, 0x36, 0xc0, 0xb1, 0x10, 0x11, 0x7e, 0x9d, 0xa1, 0xb1,
	0xec, 0x1e, 0xfc, 0x93, 0x48, 0x8b, 0x7a, 0xcc, 0x63, 0x1c, 0x49, 0x3e, 0xc5, 0xa0, 0xd4, 0x2e,
	0x75, 0xaa, 0x51, 0xa3, 0x40, 0xdf, 0xf1, 0x29, 0xb2, 0xff, 0xa0, 0x2c, 0xd1, 0x4a, 0x13, 0x6c,
	0x53, 0x74, 0xe9, 0xb0, 0xe7, 0x70, 0x4b, 0xa0, 0x49, 0x34, 0x8a, 0xd1, 0xb9, 0x32, 0x76, 0x74,
	0xa5, 0x92, 0x47, 0xdc, 0x20, 0xa7, 0xbc, 0x56, 0xc6, 0x0e, 0x2f, 0x15, 0xed, 0x81, 0x6f, 0xd0,
	0xda, 0x44, 0x4e, 0x4c, 0xb0, 0xd3, 0x2e, 0x75, 0x6a, 0xbd, 0xff, 0xbb, 0xb1, 0x4c, 0xba, 0x7d,
	0x25, 0x2d, 0x4f, 0x24, 0xea, 0x0f, 0x79, 0x34, 0x2a, 0x78, 0xac, 0x07, 0x8d, 0x78, 0x15, 0x1e,
	0x25, 0x99, 0x09, 0xca, 0x6d, 0xaf, 0x53, 0xeb, 0x35, 0x28, 0x71, 0x78, 0xda, 0x57, 0x72, 0x9c,
	0x4c, 0xa2, 0x7a, 0xc1, 0x19, 0x66, 0x86, 0x3d, 0x80, 0xe6, 0x3a, 0x47, 0xab, 0x99, 0x45, 0x13,
	0x54, 0xda, 0x5e, 0xa7, 0x1a, 0xed, 0x15, 0x78, 0x44, 0x30, 0x7b, 0x08, 0xfe, 0x5c, 0xe9, 0x8b,
	0x54, 0x71, 0x11, 0xec, 0x92, 0xa4, 0x26, 0x55, 0xfe, 0x94, 0x83, 0xc3, 0x81, 0x89, 0x0a, 0x46,
	0xf8, 0x11, 0xf6, 0xaf, 0x69, 0x65, 0x5d, 0xf8, 0x97, 0xa7, 0xa9, 0x9a, 0x8f, 0x92, 0x6c, 0x34,
	0x56, 0x7a, 0xce, 0xb5, 0x48, 0xe4, 0x84, 0xc6, 0xea, 0x47, 0xfb, 0x14, 0x1a, 0x66, 0xaf, 0x8a,
	0x00, 0x6b, 0x82, 0x37, 0xb5, 0x33, 0x1a, 0x6c, 0x39, 0x72, 0x66, 0x78, 0x04, 0xfe, 0xaa, 0x13,
	0x16, 0xc0, 0x2e, 0x17, 0x42, 0xa3, 0x31, 0xf9, 0x62, 0x56, 0xae, 0x8b, 0x4c, 0xb8, 0xc5, 0x39,
	0x5f, 0xe4, 0x4b, 0x59, 0xb9, 0xe1, 0x0f, 0x0f, 0x6a, 0x1b, 0x82, 0x19, 0x83, 0x9d, 0x8d, 0xcd,
	0x92, 0xcd, 0x0e, 0xa0, 0xea, 0xbe, 0x26, 0xe3, 0x31, 0xe6, 0xf9, 0x6b, 0x80, 0x3d, 0x86, 0x4a,
	0xca, 0xcf, 0x30, 0x35, 0x81, 0x47, 0xe3, 0x3d, 0xb8, 0x3a, 0x84, 0xee, 0x5b, 0x0a, 0xbf, 0x94,
	0x56, 0x2f, 0xa2, 0x9c, 0xcb, 0xfa, 0x50, 0xe3, 0x52, 0x2a, 0xcb, 0x6d, 0xa2, 0xa4, 0x5b, 0xa9,
	0x4b, 0xbd, 0x7b, 0x2d, 0xf5, 0x78, 0xcd, 0x59, 0xe6, 0x6f, 0x66, 0xb1, 0x16, 0xf8, 0x28, 0x45,
	0xa6, 0x12, 0x69, 0x83, 0x32, 0xe9, 0x2a, 0x7c, 0x6a, 0x44, 0x09, 0x0c, 0x2a, 0x79, 0x23, 0x4a,
	0x20, 0x0b, 0xa1, 0xae, 0x74, 0x7c, 0x8e, 0xc6, 0x6a, 0x6e, 0x95, 0xa6, 0xad, 0x55, 0xa3, 0x4b,
	0x98, 0x1b, 0x71, 0xa6, 0x44, 0xe0, 0x53, 0xc8, 0x99, 0xec, 0x0e, 0x94, 0x33, 0xa5, 0xad, 0x09,
	0xaa, 0x24, 0xb2, 0x4a, 0x22, 0x4f, 0x95, 0xb6, 0xd1, 0x12, 0x6f, 0x3d, 0x85, 0xda, 0x46, 0x8b,
	0xae, 0xc2, 0x05, 0x2e, 0xf2, 0x09, 0x3a, 0xd3, 0x5d, 0xc4, 0x37, 0x9e, 0xce, 0x56, 0xc3, 0x5b,
	0x3a, 0xcf, 0xb6, 0x9f, 0x94, 0x5a, 0x47, 0xd0, 0xbc, 0xda, 0xe2, 0xdf, 0xe4, 0x87, 0x6f, 0x60,
	0xc7, 0x29, 0xf9, 0xed, 0xda, 0x5a, 0xe0, 0xd3, 0x2d, 0xc7, 0x2a, 0xcd, 0x13, 0x0b, 0xdf, 0xf1,
	0x9d, 0x76, 0x3a, 0xbb, 0x46, 0x44, 0x76, 0xf8, 0xbd, 0x04, 0x3e, 0x5d, 0x7b, 0x96, 0x2e, 0xd8,
	0x6d, 0x00, 0x33, 0x8b, 0x63, 0x34, 0x66, 0x3c, 0x4b, 0xf3, 0x07, 0xb9, 0x81, 0xb0, 0x43, 0x68,
	0xa0, 0xd6, 0x4a, 0x8f, 0xa6, 0x68, 0x0c, 0x9f, 0xac, 0xa4, 0xd5, 0x09, 0x3c, 0x59, 0x62, 0xee,
	0x79, 0xff, 0xf9, 0xd6, 0xf7, 0xcf, 0xaf, 0x1d, 0xf9, 0xe1, 0xe6, 0xc1, 0x4e, 0x79, 0x4c, 0x97,
	0x5e, 0xdd, 0xb8, 0xd0, 0x13, 0x1e, 0x87, 0x43, 0x80, 0x01, 0xa6, 0x37, 0xf1, 0x4f, 0x0a, 0xdf,
	0x83, 0x4f, 0xa5, 0x6e, 0xaa, 0xe1, 0xde, 0x17, 0xa8, 0xf7, 0x65, 0x32, 0xe0, 0x96, 0x67, 0x29,
	0x97, 0xc8, 0xee, 0x83, 0x77, 0x2c, 0x04, 0xdb, 0xa3, 0x27, 0xb3, 0xfe, 0x93, 0xb6, 0x1a, 0x6b,
	0x20, 0x4b, 0x17, 0xe1, 0x96, 0x23, 0x0e, 0x30, 0xcd, 0x89, 0xeb, 0xf6, 0x72, 0xe2, 0x4a, 0x64,
	0xb8, 0xf5, 0x62, 0xf7, 0x73, 0x99, 0x96, 0x78, 0x56, 0xa1, 0xcf, 0xa3, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0xc0, 0x8d, 0x8a, 0xbb, 0x05, 0x00, 0x00,
}