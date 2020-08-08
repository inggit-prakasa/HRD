// Service

//INGGIT
// 1. CreateEmployee (Employee) string
// 2. ReadEmployee (nama/id) Employee
// 3. UpdateEmployee (Employee) Employee
// 4. Delete (nama/id) string

// DIAN
// 5. Laporan (nama/id) Employee
// 6. LaporanAll (kosong) []Employee

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: hrd.proto

package employee

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetEmpByUsername struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *GetEmpByUsername) Reset() {
	*x = GetEmpByUsername{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmpByUsername) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmpByUsername) ProtoMessage() {}

func (x *GetEmpByUsername) ProtoReflect() protoreflect.Message {
	mi := &file_hrd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmpByUsername.ProtoReflect.Descriptor instead.
func (*GetEmpByUsername) Descriptor() ([]byte, []int) {
	return file_hrd_proto_rawDescGZIP(), []int{0}
}

func (x *GetEmpByUsername) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type LaporanEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Nama      string `protobuf:"bytes,2,opt,name=Nama,proto3" json:"Nama,omitempty"`
	TotalGaji string `protobuf:"bytes,3,opt,name=TotalGaji,proto3" json:"TotalGaji,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *LaporanEmployee) Reset() {
	*x = LaporanEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaporanEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaporanEmployee) ProtoMessage() {}

func (x *LaporanEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_hrd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaporanEmployee.ProtoReflect.Descriptor instead.
func (*LaporanEmployee) Descriptor() ([]byte, []int) {
	return file_hrd_proto_rawDescGZIP(), []int{1}
}

func (x *LaporanEmployee) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LaporanEmployee) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *LaporanEmployee) GetTotalGaji() string {
	if x != nil {
		return x.TotalGaji
	}
	return ""
}

func (x *LaporanEmployee) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ArrEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArrLaporan []*LaporanEmployee `protobuf:"bytes,1,rep,name=ArrLaporan,proto3" json:"ArrLaporan,omitempty"`
}

func (x *ArrEmployee) Reset() {
	*x = ArrEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrEmployee) ProtoMessage() {}

func (x *ArrEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_hrd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrEmployee.ProtoReflect.Descriptor instead.
func (*ArrEmployee) Descriptor() ([]byte, []int) {
	return file_hrd_proto_rawDescGZIP(), []int{2}
}

func (x *ArrEmployee) GetArrLaporan() []*LaporanEmployee {
	if x != nil {
		return x.ArrLaporan
	}
	return nil
}

var File_hrd_proto protoreflect.FileDescriptor

var file_hrd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x68, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x0f, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x6a, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x6a, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x48, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x41, 0x72, 0x72, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x52, 0x0a, 0x41, 0x72, 0x72, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x32, 0x53, 0x0a, 0x03,
	0x48, 0x52, 0x44, 0x12, 0x4c, 0x0a, 0x11, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hrd_proto_rawDescOnce sync.Once
	file_hrd_proto_rawDescData = file_hrd_proto_rawDesc
)

func file_hrd_proto_rawDescGZIP() []byte {
	file_hrd_proto_rawDescOnce.Do(func() {
		file_hrd_proto_rawDescData = protoimpl.X.CompressGZIP(file_hrd_proto_rawDescData)
	})
	return file_hrd_proto_rawDescData
}

var file_hrd_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hrd_proto_goTypes = []interface{}{
	(*GetEmpByUsername)(nil), // 0: employee.GetEmpByUsername
	(*LaporanEmployee)(nil),  // 1: employee.LaporanEmployee
	(*ArrEmployee)(nil),      // 2: employee.ArrEmployee
}
var file_hrd_proto_depIdxs = []int32{
	1, // 0: employee.ArrEmployee.ArrLaporan:type_name -> employee.LaporanEmployee
	0, // 1: employee.HRD.LaporanByUsername:input_type -> employee.GetEmpByUsername
	1, // 2: employee.HRD.LaporanByUsername:output_type -> employee.LaporanEmployee
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hrd_proto_init() }
func file_hrd_proto_init() {
	if File_hrd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hrd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmpByUsername); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hrd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaporanEmployee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hrd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrEmployee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hrd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hrd_proto_goTypes,
		DependencyIndexes: file_hrd_proto_depIdxs,
		MessageInfos:      file_hrd_proto_msgTypes,
	}.Build()
	File_hrd_proto = out.File
	file_hrd_proto_rawDesc = nil
	file_hrd_proto_goTypes = nil
	file_hrd_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HRDClient is the client API for HRD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HRDClient interface {
	LaporanByUsername(ctx context.Context, in *GetEmpByUsername, opts ...grpc.CallOption) (*LaporanEmployee, error)
}

type hRDClient struct {
	cc grpc.ClientConnInterface
}

func NewHRDClient(cc grpc.ClientConnInterface) HRDClient {
	return &hRDClient{cc}
}

func (c *hRDClient) LaporanByUsername(ctx context.Context, in *GetEmpByUsername, opts ...grpc.CallOption) (*LaporanEmployee, error) {
	out := new(LaporanEmployee)
	err := c.cc.Invoke(ctx, "/employee.HRD/LaporanByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HRDServer is the server API for HRD service.
type HRDServer interface {
	LaporanByUsername(context.Context, *GetEmpByUsername) (*LaporanEmployee, error)
}

// UnimplementedHRDServer can be embedded to have forward compatible implementations.
type UnimplementedHRDServer struct {
}

func (*UnimplementedHRDServer) LaporanByUsername(context.Context, *GetEmpByUsername) (*LaporanEmployee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaporanByUsername not implemented")
}

func RegisterHRDServer(s *grpc.Server, srv HRDServer) {
	s.RegisterService(&_HRD_serviceDesc, srv)
}

func _HRD_LaporanByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmpByUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HRDServer).LaporanByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.HRD/LaporanByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HRDServer).LaporanByUsername(ctx, req.(*GetEmpByUsername))
	}
	return interceptor(ctx, in, info, handler)
}

var _HRD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "employee.HRD",
	HandlerType: (*HRDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LaporanByUsername",
			Handler:    _HRD_LaporanByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hrd.proto",
}
