// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloWorld.proto

/*
Package helloWorld is a generated protocol buffer package.

It is generated from these files:
	helloWorld.proto

It has these top-level messages:
	WeekDoRequest
	WeekDoReply
	Bar
	Baz
	Bab
	Profile
*/
package helloWorld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// 定义枚举类型
type SportsName int32

const (
	SportsName_CHINESE_BASKETBALL SportsName = 0
	SportsName_JAPAN_FOOTBALL     SportsName = 1
	SportsName_CHINESE_HAHA       SportsName = 2
)

var SportsName_name = map[int32]string{
	0: "CHINESE_BASKETBALL",
	1: "JAPAN_FOOTBALL",
	2: "CHINESE_HAHA",
}
var SportsName_value = map[string]int32{
	"CHINESE_BASKETBALL": 0,
	"JAPAN_FOOTBALL":     1,
	"CHINESE_HAHA":       2,
}

func (x SportsName) String() string {
	return proto.EnumName(SportsName_name, int32(x))
}
func (SportsName) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 测试--枚举类型
// 将同一个数值,赋值给不同的常量
type EnumAllowingAlias int32

const (
	EnumAllowingAlias_UNKNOWN EnumAllowingAlias = 0
	EnumAllowingAlias_STARTED EnumAllowingAlias = 1
	EnumAllowingAlias_RUNNING EnumAllowingAlias = 1
)

var EnumAllowingAlias_name = map[int32]string{
	0: "UNKNOWN",
	1: "STARTED",
	// Duplicate value: 1: "RUNNING",
}
var EnumAllowingAlias_value = map[string]int32{
	"UNKNOWN": 0,
	"STARTED": 1,
	"RUNNING": 1,
}

func (x EnumAllowingAlias) String() string {
	return proto.EnumName(EnumAllowingAlias_name, int32(x))
}
func (EnumAllowingAlias) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 定义请求结构
type WeekDoRequest struct {
	SportsName     string   `protobuf:"bytes,1,opt,name=sports_name,json=sportsName" json:"sports_name,omitempty"`
	SportsAge      int32    `protobuf:"varint,2,opt,name=sports_age,json=sportsAge" json:"sports_age,omitempty"`
	SportsKinds    []string `protobuf:"bytes,3,rep,name=sports_kinds,json=sportsKinds" json:"sports_kinds,omitempty"`
	XTestCase_1    string   `protobuf:"bytes,4,opt,name=_test_case_1,json=TestCase1" json:"_test_case_1,omitempty"`
	TestCaseDemo_2 string   `protobuf:"bytes,5,opt,name=test_case_demo_2,json=testCaseDemo2" json:"test_case_demo_2,omitempty"`
}

func (m *WeekDoRequest) Reset()                    { *m = WeekDoRequest{} }
func (m *WeekDoRequest) String() string            { return proto.CompactTextString(m) }
func (*WeekDoRequest) ProtoMessage()               {}
func (*WeekDoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WeekDoRequest) GetSportsName() string {
	if m != nil {
		return m.SportsName
	}
	return ""
}

func (m *WeekDoRequest) GetSportsAge() int32 {
	if m != nil {
		return m.SportsAge
	}
	return 0
}

func (m *WeekDoRequest) GetSportsKinds() []string {
	if m != nil {
		return m.SportsKinds
	}
	return nil
}

func (m *WeekDoRequest) GetXTestCase_1() string {
	if m != nil {
		return m.XTestCase_1
	}
	return ""
}

func (m *WeekDoRequest) GetTestCaseDemo_2() string {
	if m != nil {
		return m.TestCaseDemo_2
	}
	return ""
}

// 定义响应结构
type WeekDoReply struct {
	ActivityNames []string `protobuf:"bytes,1,rep,name=activity_names,json=activityNames" json:"activity_names,omitempty"`
}

func (m *WeekDoReply) Reset()                    { *m = WeekDoReply{} }
func (m *WeekDoReply) String() string            { return proto.CompactTextString(m) }
func (*WeekDoReply) ProtoMessage()               {}
func (*WeekDoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WeekDoReply) GetActivityNames() []string {
	if m != nil {
		return m.ActivityNames
	}
	return nil
}

// 测试message的嵌套--也就是Go中结构体的使用
type Bar struct {
}

func (m *Bar) Reset()                    { *m = Bar{} }
func (m *Bar) String() string            { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()               {}
func (*Bar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Baz struct {
	Foo *Bar `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
}

func (m *Baz) Reset()                    { *m = Baz{} }
func (m *Baz) String() string            { return proto.CompactTextString(m) }
func (*Baz) ProtoMessage()               {}
func (*Baz) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Baz) GetFoo() *Bar {
	if m != nil {
		return m.Foo
	}
	return nil
}

type Bab struct {
	Foo map[string]*Bar `protobuf:"bytes,1,rep,name=foo" json:"foo,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Bab) Reset()                    { *m = Bab{} }
func (m *Bab) String() string            { return proto.CompactTextString(m) }
func (*Bab) ProtoMessage()               {}
func (*Bab) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Bab) GetFoo() map[string]*Bar {
	if m != nil {
		return m.Foo
	}
	return nil
}

type Profile struct {
	// Types that are valid to be assigned to Avatar:
	//	*Profile_ImageUrl
	//	*Profile_ImageData
	Avatar isProfile_Avatar `protobuf_oneof:"avatar"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isProfile_Avatar interface {
	isProfile_Avatar()
}

type Profile_ImageUrl struct {
	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,oneof"`
}
type Profile_ImageData struct {
	ImageData []byte `protobuf:"bytes,2,opt,name=image_data,json=imageData,proto3,oneof"`
}

func (*Profile_ImageUrl) isProfile_Avatar()  {}
func (*Profile_ImageData) isProfile_Avatar() {}

func (m *Profile) GetAvatar() isProfile_Avatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *Profile) GetImageUrl() string {
	if x, ok := m.GetAvatar().(*Profile_ImageUrl); ok {
		return x.ImageUrl
	}
	return ""
}

func (m *Profile) GetImageData() []byte {
	if x, ok := m.GetAvatar().(*Profile_ImageData); ok {
		return x.ImageData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Profile) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Profile_OneofMarshaler, _Profile_OneofUnmarshaler, _Profile_OneofSizer, []interface{}{
		(*Profile_ImageUrl)(nil),
		(*Profile_ImageData)(nil),
	}
}

func _Profile_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Profile)
	// avatar
	switch x := m.Avatar.(type) {
	case *Profile_ImageUrl:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ImageUrl)
	case *Profile_ImageData:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.ImageData)
	case nil:
	default:
		return fmt.Errorf("Profile.Avatar has unexpected type %T", x)
	}
	return nil
}

func _Profile_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Profile)
	switch tag {
	case 1: // avatar.image_url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Avatar = &Profile_ImageUrl{x}
		return true, err
	case 2: // avatar.image_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Avatar = &Profile_ImageData{x}
		return true, err
	default:
		return false, nil
	}
}

func _Profile_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Profile)
	// avatar
	switch x := m.Avatar.(type) {
	case *Profile_ImageUrl:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ImageUrl)))
		n += len(x.ImageUrl)
	case *Profile_ImageData:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ImageData)))
		n += len(x.ImageData)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*WeekDoRequest)(nil), "helloWorld.WeekDoRequest")
	proto.RegisterType((*WeekDoReply)(nil), "helloWorld.WeekDoReply")
	proto.RegisterType((*Bar)(nil), "helloWorld.Bar")
	proto.RegisterType((*Baz)(nil), "helloWorld.Baz")
	proto.RegisterType((*Bab)(nil), "helloWorld.Bab")
	proto.RegisterType((*Profile)(nil), "helloWorld.Profile")
	proto.RegisterEnum("helloWorld.SportsName", SportsName_name, SportsName_value)
	proto.RegisterEnum("helloWorld.EnumAllowingAlias", EnumAllowingAlias_name, EnumAllowingAlias_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WeekDo service

type WeekDoClient interface {
	PlayGames(ctx context.Context, in *WeekDoRequest, opts ...grpc.CallOption) (*WeekDoReply, error)
	WatchMoves(ctx context.Context, in *WeekDoRequest, opts ...grpc.CallOption) (*WeekDoReply, error)
}

type weekDoClient struct {
	cc *grpc.ClientConn
}

func NewWeekDoClient(cc *grpc.ClientConn) WeekDoClient {
	return &weekDoClient{cc}
}

func (c *weekDoClient) PlayGames(ctx context.Context, in *WeekDoRequest, opts ...grpc.CallOption) (*WeekDoReply, error) {
	out := new(WeekDoReply)
	err := grpc.Invoke(ctx, "/helloWorld.WeekDo/PlayGames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weekDoClient) WatchMoves(ctx context.Context, in *WeekDoRequest, opts ...grpc.CallOption) (*WeekDoReply, error) {
	out := new(WeekDoReply)
	err := grpc.Invoke(ctx, "/helloWorld.WeekDo/WatchMoves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WeekDo service

type WeekDoServer interface {
	PlayGames(context.Context, *WeekDoRequest) (*WeekDoReply, error)
	WatchMoves(context.Context, *WeekDoRequest) (*WeekDoReply, error)
}

func RegisterWeekDoServer(s *grpc.Server, srv WeekDoServer) {
	s.RegisterService(&_WeekDo_serviceDesc, srv)
}

func _WeekDo_PlayGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeekDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeekDoServer).PlayGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloWorld.WeekDo/PlayGames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeekDoServer).PlayGames(ctx, req.(*WeekDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeekDo_WatchMoves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeekDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeekDoServer).WatchMoves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloWorld.WeekDo/WatchMoves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeekDoServer).WatchMoves(ctx, req.(*WeekDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WeekDo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloWorld.WeekDo",
	HandlerType: (*WeekDoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlayGames",
			Handler:    _WeekDo_PlayGames_Handler,
		},
		{
			MethodName: "WatchMoves",
			Handler:    _WeekDo_WatchMoves_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloWorld.proto",
}

func init() { proto.RegisterFile("helloWorld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x65, 0x71, 0x21, 0x78, 0x80, 0xd4, 0xdd, 0x43, 0xeb, 0x22, 0x55, 0x10, 0x4b, 0x51, 0x11,
	0x07, 0xa4, 0xd0, 0x1e, 0xaa, 0x1e, 0x2a, 0xad, 0x83, 0x03, 0x09, 0xa9, 0x41, 0x06, 0xc4, 0xd1,
	0x5a, 0x60, 0x43, 0x2c, 0xd6, 0x2c, 0xf5, 0x2e, 0x54, 0xce, 0x4f, 0xf4, 0x87, 0xfa, 0x71, 0x95,
	0x6d, 0x10, 0x89, 0x94, 0x53, 0x4f, 0xf6, 0xbc, 0xf7, 0xf6, 0xcd, 0xce, 0x1b, 0x1b, 0x8c, 0x47,
	0xc6, 0xb9, 0x98, 0x89, 0x88, 0x2f, 0xdb, 0xdb, 0x48, 0x28, 0x81, 0xe1, 0x84, 0x58, 0x7f, 0x11,
	0x54, 0x67, 0x8c, 0xad, 0xbb, 0xc2, 0x63, 0xbf, 0x76, 0x4c, 0x2a, 0x5c, 0x87, 0xb2, 0xdc, 0x8a,
	0x48, 0x49, 0x7f, 0x43, 0x43, 0x66, 0xa2, 0x06, 0x6a, 0xea, 0x1e, 0x64, 0x90, 0x4b, 0x43, 0x86,
	0x3f, 0xc1, 0xa1, 0xf2, 0xe9, 0x8a, 0x99, 0xf9, 0x06, 0x6a, 0x16, 0x3c, 0x3d, 0x43, 0xc8, 0x8a,
	0xe1, 0x0b, 0xa8, 0x1c, 0xe8, 0x75, 0xb0, 0x59, 0x4a, 0x53, 0x6b, 0x68, 0x4d, 0xdd, 0x3b, 0x78,
	0x0e, 0x12, 0x08, 0xd7, 0xa1, 0xe2, 0x2b, 0x26, 0x95, 0xbf, 0xa0, 0x92, 0xf9, 0x57, 0xe6, 0x9b,
	0xb4, 0x87, 0x3e, 0x61, 0x52, 0x5d, 0x53, 0xc9, 0xae, 0xf0, 0x67, 0x30, 0x4e, 0xfc, 0x92, 0x85,
	0xc2, 0xef, 0x98, 0x85, 0x54, 0x54, 0x55, 0x07, 0x51, 0x97, 0x85, 0xa2, 0x63, 0x7d, 0x85, 0xf2,
	0xf1, 0xf6, 0x5b, 0x1e, 0xe3, 0x4b, 0x38, 0xa7, 0x0b, 0x15, 0xec, 0x03, 0x15, 0xa7, 0xb7, 0x97,
	0x26, 0x4a, 0xbb, 0x57, 0x8f, 0x68, 0x32, 0x80, 0xb4, 0x0a, 0xa0, 0xd9, 0x34, 0xb2, 0x9a, 0xc9,
	0xe3, 0x09, 0x5f, 0x80, 0xf6, 0x20, 0x44, 0x3a, 0x68, 0xb9, 0xf3, 0xb6, 0xfd, 0x2c, 0x2e, 0x9b,
	0x46, 0x5e, 0xc2, 0x59, 0x4f, 0x89, 0x72, 0x8e, 0x5b, 0x47, 0xa5, 0xd6, 0x2c, 0x77, 0xcc, 0x97,
	0xca, 0x79, 0xfb, 0x46, 0x08, 0x67, 0xa3, 0xa2, 0x38, 0x3d, 0x52, 0xeb, 0x41, 0xe9, 0x08, 0x60,
	0x03, 0xb4, 0x35, 0x8b, 0x0f, 0x51, 0x26, 0xaf, 0xf8, 0x12, 0x0a, 0x7b, 0xca, 0x77, 0x59, 0x7c,
	0xaf, 0x74, 0xcd, 0xd8, 0xef, 0xf9, 0x6f, 0xc8, 0x1a, 0xc3, 0xd9, 0x28, 0x12, 0x0f, 0x01, 0x4f,
	0x92, 0xd7, 0x83, 0x90, 0xae, 0x98, 0xbf, 0x8b, 0x78, 0xe6, 0xd6, 0xcf, 0x79, 0xa5, 0x14, 0x9a,
	0x46, 0x1c, 0xd7, 0x01, 0x32, 0x7a, 0x49, 0x15, 0x4d, 0x9d, 0x2b, 0xfd, 0x9c, 0x97, 0x1d, 0xe9,
	0x52, 0x45, 0xed, 0x12, 0x14, 0xe9, 0x9e, 0x2a, 0x1a, 0xb5, 0xee, 0x00, 0xc6, 0xa7, 0x8d, 0xbe,
	0x07, 0x7c, 0xdd, 0xbf, 0x75, 0x9d, 0xb1, 0xe3, 0xdb, 0x64, 0x3c, 0x70, 0x26, 0x36, 0xb9, 0xbf,
	0x37, 0x72, 0x18, 0xc3, 0xf9, 0x1d, 0x19, 0x11, 0xd7, 0xbf, 0x19, 0x0e, 0x33, 0x0c, 0x61, 0x03,
	0x2a, 0x47, 0x6d, 0x9f, 0xf4, 0x89, 0x91, 0x6f, 0xfd, 0x80, 0x77, 0xce, 0x66, 0x17, 0x12, 0xce,
	0xc5, 0xef, 0x60, 0xb3, 0x22, 0x3c, 0xa0, 0x12, 0x97, 0xe1, 0x6c, 0xea, 0x0e, 0xdc, 0xe1, 0xcc,
	0x35, 0x72, 0x49, 0x31, 0x9e, 0x10, 0x6f, 0xe2, 0x74, 0x0d, 0x94, 0x14, 0xde, 0xd4, 0x75, 0x6f,
	0xdd, 0x9e, 0x81, 0x6a, 0x79, 0x03, 0x75, 0xfe, 0x20, 0x28, 0x66, 0x4b, 0xc4, 0x04, 0xf4, 0x11,
	0xa7, 0x71, 0x2f, 0xd9, 0x12, 0xfe, 0xf8, 0x3c, 0x94, 0x17, 0xdf, 0x68, 0xed, 0xc3, 0x6b, 0xd4,
	0x96, 0xc7, 0x56, 0x0e, 0xdb, 0x00, 0x33, 0xaa, 0x16, 0x8f, 0x3f, 0xc5, 0xfe, 0x7f, 0x3d, 0xe6,
	0xc5, 0xf4, 0x3f, 0xf9, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xab, 0xfd, 0xec, 0x59, 0x3b, 0x03,
	0x00, 0x00,
}
