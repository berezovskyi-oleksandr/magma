// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/sms_orc8r.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// section 8.4
type SMODownlinkUnitdata struct {
	Imsi                 string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	NasMessageContainer  []byte   `protobuf:"bytes,2,opt,name=nas_message_container,json=nasMessageContainer,proto3" json:"nas_message_container,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMODownlinkUnitdata) Reset()         { *m = SMODownlinkUnitdata{} }
func (m *SMODownlinkUnitdata) String() string { return proto.CompactTextString(m) }
func (*SMODownlinkUnitdata) ProtoMessage()    {}
func (*SMODownlinkUnitdata) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e3e558366760a7d, []int{0}
}

func (m *SMODownlinkUnitdata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMODownlinkUnitdata.Unmarshal(m, b)
}
func (m *SMODownlinkUnitdata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMODownlinkUnitdata.Marshal(b, m, deterministic)
}
func (m *SMODownlinkUnitdata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMODownlinkUnitdata.Merge(m, src)
}
func (m *SMODownlinkUnitdata) XXX_Size() int {
	return xxx_messageInfo_SMODownlinkUnitdata.Size(m)
}
func (m *SMODownlinkUnitdata) XXX_DiscardUnknown() {
	xxx_messageInfo_SMODownlinkUnitdata.DiscardUnknown(m)
}

var xxx_messageInfo_SMODownlinkUnitdata proto.InternalMessageInfo

func (m *SMODownlinkUnitdata) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *SMODownlinkUnitdata) GetNasMessageContainer() []byte {
	if m != nil {
		return m.NasMessageContainer
	}
	return nil
}

// section 8.14
type SMOPagingRequest struct {
	Imsi                   string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	VlrName                string   `protobuf:"bytes,2,opt,name=vlr_name,json=vlrName,proto3" json:"vlr_name,omitempty"`
	ServiceIndicator       []byte   `protobuf:"bytes,3,opt,name=service_indicator,json=serviceIndicator,proto3" json:"service_indicator,omitempty"`
	Tmsi                   []byte   `protobuf:"bytes,4,opt,name=tmsi,proto3" json:"tmsi,omitempty"`
	Cli                    []byte   `protobuf:"bytes,5,opt,name=cli,proto3" json:"cli,omitempty"`
	LocationAreaIdentifier []byte   `protobuf:"bytes,6,opt,name=location_area_identifier,json=locationAreaIdentifier,proto3" json:"location_area_identifier,omitempty"`
	GlobalCnId             []byte   `protobuf:"bytes,7,opt,name=global_cn_id,json=globalCnId,proto3" json:"global_cn_id,omitempty"`
	SsCode                 []byte   `protobuf:"bytes,8,opt,name=ss_code,json=ssCode,proto3" json:"ss_code,omitempty"`
	LcsIndicator           []byte   `protobuf:"bytes,9,opt,name=lcs_indicator,json=lcsIndicator,proto3" json:"lcs_indicator,omitempty"`
	LcsClientIdentity      []byte   `protobuf:"bytes,10,opt,name=lcs_client_identity,json=lcsClientIdentity,proto3" json:"lcs_client_identity,omitempty"`
	ChannelNeeded          []byte   `protobuf:"bytes,11,opt,name=channel_needed,json=channelNeeded,proto3" json:"channel_needed,omitempty"`
	EmlppPriority          []byte   `protobuf:"bytes,12,opt,name=emlpp_priority,json=emlppPriority,proto3" json:"emlpp_priority,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *SMOPagingRequest) Reset()         { *m = SMOPagingRequest{} }
func (m *SMOPagingRequest) String() string { return proto.CompactTextString(m) }
func (*SMOPagingRequest) ProtoMessage()    {}
func (*SMOPagingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e3e558366760a7d, []int{1}
}

func (m *SMOPagingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMOPagingRequest.Unmarshal(m, b)
}
func (m *SMOPagingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMOPagingRequest.Marshal(b, m, deterministic)
}
func (m *SMOPagingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMOPagingRequest.Merge(m, src)
}
func (m *SMOPagingRequest) XXX_Size() int {
	return xxx_messageInfo_SMOPagingRequest.Size(m)
}
func (m *SMOPagingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SMOPagingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SMOPagingRequest proto.InternalMessageInfo

func (m *SMOPagingRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *SMOPagingRequest) GetVlrName() string {
	if m != nil {
		return m.VlrName
	}
	return ""
}

func (m *SMOPagingRequest) GetServiceIndicator() []byte {
	if m != nil {
		return m.ServiceIndicator
	}
	return nil
}

func (m *SMOPagingRequest) GetTmsi() []byte {
	if m != nil {
		return m.Tmsi
	}
	return nil
}

func (m *SMOPagingRequest) GetCli() []byte {
	if m != nil {
		return m.Cli
	}
	return nil
}

func (m *SMOPagingRequest) GetLocationAreaIdentifier() []byte {
	if m != nil {
		return m.LocationAreaIdentifier
	}
	return nil
}

func (m *SMOPagingRequest) GetGlobalCnId() []byte {
	if m != nil {
		return m.GlobalCnId
	}
	return nil
}

func (m *SMOPagingRequest) GetSsCode() []byte {
	if m != nil {
		return m.SsCode
	}
	return nil
}

func (m *SMOPagingRequest) GetLcsIndicator() []byte {
	if m != nil {
		return m.LcsIndicator
	}
	return nil
}

func (m *SMOPagingRequest) GetLcsClientIdentity() []byte {
	if m != nil {
		return m.LcsClientIdentity
	}
	return nil
}

func (m *SMOPagingRequest) GetChannelNeeded() []byte {
	if m != nil {
		return m.ChannelNeeded
	}
	return nil
}

func (m *SMOPagingRequest) GetEmlppPriority() []byte {
	if m != nil {
		return m.EmlppPriority
	}
	return nil
}

// section 8.22
type SMOUplinkUnitdata struct {
	Imsi                    string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	NasMessageContainer     []byte   `protobuf:"bytes,2,opt,name=nas_message_container,json=nasMessageContainer,proto3" json:"nas_message_container,omitempty"`
	Imeisv                  []byte   `protobuf:"bytes,3,opt,name=imeisv,proto3" json:"imeisv,omitempty"`
	UeTimeZone              []byte   `protobuf:"bytes,4,opt,name=ue_time_zone,json=ueTimeZone,proto3" json:"ue_time_zone,omitempty"`
	MobileStationClassmark2 []byte   `protobuf:"bytes,5,opt,name=mobile_station_classmark2,json=mobileStationClassmark2,proto3" json:"mobile_station_classmark2,omitempty"`
	Tai                     []byte   `protobuf:"bytes,6,opt,name=tai,proto3" json:"tai,omitempty"`
	ECgi                    []byte   `protobuf:"bytes,7,opt,name=e_cgi,json=eCgi,proto3" json:"e_cgi,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *SMOUplinkUnitdata) Reset()         { *m = SMOUplinkUnitdata{} }
func (m *SMOUplinkUnitdata) String() string { return proto.CompactTextString(m) }
func (*SMOUplinkUnitdata) ProtoMessage()    {}
func (*SMOUplinkUnitdata) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e3e558366760a7d, []int{2}
}

func (m *SMOUplinkUnitdata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMOUplinkUnitdata.Unmarshal(m, b)
}
func (m *SMOUplinkUnitdata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMOUplinkUnitdata.Marshal(b, m, deterministic)
}
func (m *SMOUplinkUnitdata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMOUplinkUnitdata.Merge(m, src)
}
func (m *SMOUplinkUnitdata) XXX_Size() int {
	return xxx_messageInfo_SMOUplinkUnitdata.Size(m)
}
func (m *SMOUplinkUnitdata) XXX_DiscardUnknown() {
	xxx_messageInfo_SMOUplinkUnitdata.DiscardUnknown(m)
}

var xxx_messageInfo_SMOUplinkUnitdata proto.InternalMessageInfo

func (m *SMOUplinkUnitdata) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *SMOUplinkUnitdata) GetNasMessageContainer() []byte {
	if m != nil {
		return m.NasMessageContainer
	}
	return nil
}

func (m *SMOUplinkUnitdata) GetImeisv() []byte {
	if m != nil {
		return m.Imeisv
	}
	return nil
}

func (m *SMOUplinkUnitdata) GetUeTimeZone() []byte {
	if m != nil {
		return m.UeTimeZone
	}
	return nil
}

func (m *SMOUplinkUnitdata) GetMobileStationClassmark2() []byte {
	if m != nil {
		return m.MobileStationClassmark2
	}
	return nil
}

func (m *SMOUplinkUnitdata) GetTai() []byte {
	if m != nil {
		return m.Tai
	}
	return nil
}

func (m *SMOUplinkUnitdata) GetECgi() []byte {
	if m != nil {
		return m.ECgi
	}
	return nil
}

func init() {
	proto.RegisterType((*SMODownlinkUnitdata)(nil), "magma.lte.SMODownlinkUnitdata")
	proto.RegisterType((*SMOPagingRequest)(nil), "magma.lte.SMOPagingRequest")
	proto.RegisterType((*SMOUplinkUnitdata)(nil), "magma.lte.SMOUplinkUnitdata")
}

func init() { proto.RegisterFile("lte/protos/sms_orc8r.proto", fileDescriptor_5e3e558366760a7d) }

var fileDescriptor_5e3e558366760a7d = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0x13, 0x3b,
	0x14, 0x7d, 0x79, 0x6d, 0xd3, 0xe6, 0x36, 0x7d, 0xaf, 0x71, 0xf4, 0xda, 0x49, 0xfb, 0x84, 0xaa,
	0x20, 0xa4, 0x4a, 0x48, 0x89, 0x14, 0x36, 0x15, 0x0b, 0x04, 0x0d, 0x12, 0xca, 0x22, 0x4d, 0xc9,
	0x50, 0x16, 0x95, 0x90, 0xe5, 0x7a, 0x2e, 0xc3, 0x55, 0xfd, 0x11, 0x6c, 0x27, 0x55, 0xf9, 0x31,
	0xfc, 0x02, 0x7e, 0x20, 0x4b, 0x34, 0x9e, 0x49, 0x68, 0xa1, 0x5b, 0x56, 0x73, 0xe7, 0x9c, 0xa3,
	0x33, 0x1e, 0x1f, 0x1f, 0xc3, 0x81, 0x0a, 0xd8, 0x9f, 0x39, 0x1b, 0xac, 0xef, 0x7b, 0xed, 0xb9,
	0x75, 0xf2, 0xc4, 0xf5, 0x22, 0xc0, 0x1a, 0x5a, 0xe4, 0x5a, 0xf4, 0x54, 0xc0, 0x83, 0x4e, 0xc4,
	0x97, 0x42, 0x69, 0xb5, 0xb6, 0xa6, 0x54, 0x75, 0x3f, 0x40, 0x3b, 0x1d, 0x4f, 0x5e, 0xdb, 0x1b,
	0xa3, 0xc8, 0x5c, 0x5f, 0x18, 0x0a, 0x99, 0x08, 0x82, 0x31, 0x58, 0x27, 0xed, 0x29, 0xa9, 0x1d,
	0xd5, 0x8e, 0x1b, 0xd3, 0x38, 0xb3, 0x01, 0xfc, 0x67, 0x84, 0xe7, 0x1a, 0xbd, 0x17, 0x39, 0x72,
	0x69, 0x4d, 0x10, 0x64, 0xd0, 0x25, 0x7f, 0x1f, 0xd5, 0x8e, 0x9b, 0xd3, 0xb6, 0x11, 0x7e, 0x5c,
	0x72, 0xc3, 0x25, 0xd5, 0xfd, 0xb6, 0x06, 0xbb, 0xe9, 0x78, 0x72, 0x2e, 0x72, 0x32, 0xf9, 0x14,
	0x3f, 0xcf, 0xd1, 0x87, 0x07, 0xcd, 0x3b, 0xb0, 0xb5, 0x50, 0x8e, 0x1b, 0xa1, 0x31, 0xfa, 0x35,
	0xa6, 0x9b, 0x0b, 0xe5, 0xce, 0x84, 0x46, 0xf6, 0x14, 0x5a, 0x1e, 0xdd, 0x82, 0x24, 0x72, 0x32,
	0x19, 0x49, 0x11, 0xac, 0x4b, 0xd6, 0xe2, 0x37, 0x77, 0x2b, 0x62, 0xb4, 0xc4, 0x0b, 0xef, 0x50,
	0x78, 0xaf, 0x47, 0x3e, 0xce, 0x6c, 0x17, 0xd6, 0xa4, 0xa2, 0x64, 0x23, 0x42, 0xc5, 0xc8, 0x4e,
	0x20, 0x51, 0x56, 0x8a, 0x40, 0xd6, 0x70, 0xe1, 0x50, 0x70, 0xca, 0xd0, 0x04, 0xfa, 0x48, 0xe8,
	0x92, 0x7a, 0x94, 0xed, 0x2d, 0xf9, 0x57, 0x0e, 0xc5, 0x68, 0xc5, 0xb2, 0x23, 0x68, 0xe6, 0xca,
	0x5e, 0x09, 0xc5, 0xa5, 0xe1, 0x94, 0x25, 0x9b, 0x51, 0x0d, 0x25, 0x36, 0x34, 0xa3, 0x8c, 0xed,
	0xc3, 0xa6, 0xf7, 0x5c, 0xda, 0x0c, 0x93, 0xad, 0x48, 0xd6, 0xbd, 0x1f, 0xda, 0x0c, 0xd9, 0x63,
	0xd8, 0x51, 0xd2, 0xdf, 0xf9, 0x87, 0x46, 0xa4, 0x9b, 0x4a, 0xfa, 0x9f, 0xeb, 0xef, 0x41, 0xbb,
	0x10, 0x49, 0x45, 0x68, 0x42, 0xb5, 0xac, 0x70, 0x9b, 0x40, 0x94, 0xb6, 0x94, 0xf4, 0xc3, 0xc8,
	0x8c, 0x2a, 0x82, 0x3d, 0x81, 0x7f, 0xe4, 0x27, 0x61, 0x0c, 0x2a, 0x6e, 0x10, 0x33, 0xcc, 0x92,
	0xed, 0x28, 0xdd, 0xa9, 0xd0, 0xb3, 0x08, 0x16, 0x32, 0xd4, 0x6a, 0x36, 0xe3, 0x33, 0x47, 0xd6,
	0x15, 0x8e, 0xcd, 0x52, 0x16, 0xd1, 0xf3, 0x0a, 0xec, 0x7e, 0xaf, 0x41, 0x2b, 0x1d, 0x4f, 0x2e,
	0x66, 0x7f, 0xe2, 0x30, 0xb0, 0x3d, 0xa8, 0x93, 0x46, 0xf2, 0x8b, 0x2a, 0xbd, 0xea, 0xad, 0xd8,
	0xd3, 0x39, 0xf2, 0x40, 0x1a, 0xf9, 0x17, 0x6b, 0xb0, 0xca, 0x0e, 0xe6, 0xf8, 0x8e, 0x34, 0x5e,
	0x5a, 0x83, 0xec, 0x39, 0x74, 0xb4, 0xbd, 0x22, 0x85, 0xdc, 0x87, 0x32, 0x35, 0xa9, 0x84, 0xf7,
	0x5a, 0xb8, 0xeb, 0x41, 0x95, 0xeb, 0x7e, 0x29, 0x48, 0x4b, 0x7e, 0xb8, 0xa2, 0x8b, 0xf4, 0x83,
	0xa0, 0x2a, 0xd6, 0x62, 0x64, 0x6d, 0xd8, 0x40, 0x2e, 0x73, 0xaa, 0xc2, 0x5b, 0xc7, 0x61, 0x4e,
	0x83, 0xb7, 0xf0, 0x6f, 0x3a, 0x4e, 0x27, 0x45, 0x51, 0xd2, 0xf2, 0x50, 0xb1, 0x17, 0xd0, 0x58,
	0x6d, 0x06, 0xfb, 0xbf, 0xb7, 0xea, 0x53, 0xef, 0xb7, 0x2d, 0x3a, 0x68, 0x55, 0x6c, 0x59, 0xc0,
	0xf7, 0x96, 0xb2, 0xee, 0x5f, 0x83, 0xaf, 0x35, 0xd8, 0x5b, 0x7a, 0xbe, 0x11, 0x01, 0x6f, 0xc4,
	0xed, 0xd2, 0xfa, 0x14, 0xb6, 0xef, 0xd4, 0x8e, 0x3d, 0xba, 0x6f, 0xfe, 0x6b, 0x1d, 0x1f, 0xb4,
	0x67, 0x2f, 0xa1, 0x79, 0xb7, 0x5a, 0xec, 0xf0, 0xbe, 0xc9, 0xbd, 0xce, 0x3d, 0xe8, 0x70, 0x7a,
	0x78, 0xd9, 0x89, 0x68, 0xbf, 0xb8, 0x46, 0xa4, 0xb2, 0xf3, 0xac, 0x9f, 0xdb, 0xea, 0x9a, 0xb8,
	0xaa, 0xc7, 0xe7, 0xb3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x01, 0xd1, 0xc7, 0x64, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SMSOrc8RServiceClient is the client API for SMSOrc8RService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SMSOrc8RServiceClient interface {
	SMOUplink(ctx context.Context, in *SMOUplinkUnitdata, opts ...grpc.CallOption) (*protos.Void, error)
}

type sMSOrc8RServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSOrc8RServiceClient(cc grpc.ClientConnInterface) SMSOrc8RServiceClient {
	return &sMSOrc8RServiceClient{cc}
}

func (c *sMSOrc8RServiceClient) SMOUplink(ctx context.Context, in *SMOUplinkUnitdata, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.SMSOrc8rService/SMOUplink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSOrc8RServiceServer is the server API for SMSOrc8RService service.
type SMSOrc8RServiceServer interface {
	SMOUplink(context.Context, *SMOUplinkUnitdata) (*protos.Void, error)
}

// UnimplementedSMSOrc8RServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSMSOrc8RServiceServer struct {
}

func (*UnimplementedSMSOrc8RServiceServer) SMOUplink(ctx context.Context, req *SMOUplinkUnitdata) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SMOUplink not implemented")
}

func RegisterSMSOrc8RServiceServer(s *grpc.Server, srv SMSOrc8RServiceServer) {
	s.RegisterService(&_SMSOrc8RService_serviceDesc, srv)
}

func _SMSOrc8RService_SMOUplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMOUplinkUnitdata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSOrc8RServiceServer).SMOUplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.SMSOrc8rService/SMOUplink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSOrc8RServiceServer).SMOUplink(ctx, req.(*SMOUplinkUnitdata))
	}
	return interceptor(ctx, in, info, handler)
}

var _SMSOrc8RService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.SMSOrc8rService",
	HandlerType: (*SMSOrc8RServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SMOUplink",
			Handler:    _SMSOrc8RService_SMOUplink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/sms_orc8r.proto",
}

// SMSOrc8RGatewayServiceClient is the client API for SMSOrc8RGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SMSOrc8RGatewayServiceClient interface {
	SMODownlink(ctx context.Context, in *SMODownlinkUnitdata, opts ...grpc.CallOption) (*protos.Void, error)
	SMOPagingReq(ctx context.Context, in *SMOPagingRequest, opts ...grpc.CallOption) (*protos.Void, error)
}

type sMSOrc8RGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSOrc8RGatewayServiceClient(cc grpc.ClientConnInterface) SMSOrc8RGatewayServiceClient {
	return &sMSOrc8RGatewayServiceClient{cc}
}

func (c *sMSOrc8RGatewayServiceClient) SMODownlink(ctx context.Context, in *SMODownlinkUnitdata, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.SMSOrc8rGatewayService/SMODownlink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sMSOrc8RGatewayServiceClient) SMOPagingReq(ctx context.Context, in *SMOPagingRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.SMSOrc8rGatewayService/SMOPagingReq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSOrc8RGatewayServiceServer is the server API for SMSOrc8RGatewayService service.
type SMSOrc8RGatewayServiceServer interface {
	SMODownlink(context.Context, *SMODownlinkUnitdata) (*protos.Void, error)
	SMOPagingReq(context.Context, *SMOPagingRequest) (*protos.Void, error)
}

// UnimplementedSMSOrc8RGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSMSOrc8RGatewayServiceServer struct {
}

func (*UnimplementedSMSOrc8RGatewayServiceServer) SMODownlink(ctx context.Context, req *SMODownlinkUnitdata) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SMODownlink not implemented")
}
func (*UnimplementedSMSOrc8RGatewayServiceServer) SMOPagingReq(ctx context.Context, req *SMOPagingRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SMOPagingReq not implemented")
}

func RegisterSMSOrc8RGatewayServiceServer(s *grpc.Server, srv SMSOrc8RGatewayServiceServer) {
	s.RegisterService(&_SMSOrc8RGatewayService_serviceDesc, srv)
}

func _SMSOrc8RGatewayService_SMODownlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMODownlinkUnitdata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSOrc8RGatewayServiceServer).SMODownlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.SMSOrc8rGatewayService/SMODownlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSOrc8RGatewayServiceServer).SMODownlink(ctx, req.(*SMODownlinkUnitdata))
	}
	return interceptor(ctx, in, info, handler)
}

func _SMSOrc8RGatewayService_SMOPagingReq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMOPagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSOrc8RGatewayServiceServer).SMOPagingReq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.SMSOrc8rGatewayService/SMOPagingReq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSOrc8RGatewayServiceServer).SMOPagingReq(ctx, req.(*SMOPagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SMSOrc8RGatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.SMSOrc8rGatewayService",
	HandlerType: (*SMSOrc8RGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SMODownlink",
			Handler:    _SMSOrc8RGatewayService_SMODownlink_Handler,
		},
		{
			MethodName: "SMOPagingReq",
			Handler:    _SMSOrc8RGatewayService_SMOPagingReq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/sms_orc8r.proto",
}
