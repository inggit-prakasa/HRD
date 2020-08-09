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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Absen              int64   `protobuf:"varint,2,opt,name=absen,proto3" json:"absen,omitempty"`
	Cuti               int64   `protobuf:"varint,3,opt,name=cuti,proto3" json:"cuti,omitempty"`
	Nama               string  `protobuf:"bytes,4,opt,name=nama,proto3" json:"nama,omitempty"`
	Username           string  `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password           string  `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	GajiPokok          float32 `protobuf:"fixed32,7,opt,name=gajiPokok,proto3" json:"gajiPokok,omitempty"`
	TotalGaji          float32 `protobuf:"fixed32,8,opt,name=totalGaji,proto3" json:"totalGaji,omitempty"`
	TunjanganMakan     float32 `protobuf:"fixed32,9,opt,name=tunjanganMakan,proto3" json:"tunjanganMakan,omitempty"`
	TunjanganTransport float32 `protobuf:"fixed32,10,opt,name=tunjanganTransport,proto3" json:"tunjanganTransport,omitempty"`
	Status             string  `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Role               string  `protobuf:"bytes,12,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_hrd_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetAbsen() int64 {
	if x != nil {
		return x.Absen
	}
	return 0
}

func (x *Employee) GetCuti() int64 {
	if x != nil {
		return x.Cuti
	}
	return 0
}

func (x *Employee) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *Employee) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Employee) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Employee) GetGajiPokok() float32 {
	if x != nil {
		return x.GajiPokok
	}
	return 0
}

func (x *Employee) GetTotalGaji() float32 {
	if x != nil {
		return x.TotalGaji
	}
	return 0
}

func (x *Employee) GetTunjanganMakan() float32 {
	if x != nil {
		return x.TunjanganMakan
	}
	return 0
}

func (x *Employee) GetTunjanganTransport() float32 {
	if x != nil {
		return x.TunjanganTransport
	}
	return 0
}

func (x *Employee) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Employee) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type AllEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employes []*Employee `protobuf:"bytes,1,rep,name=employes,proto3" json:"employes,omitempty"`
}

func (x *AllEmployee) Reset() {
	*x = AllEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllEmployee) ProtoMessage() {}

func (x *AllEmployee) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AllEmployee.ProtoReflect.Descriptor instead.
func (*AllEmployee) Descriptor() ([]byte, []int) {
	return file_hrd_proto_rawDescGZIP(), []int{1}
}

func (x *AllEmployee) GetEmployes() []*Employee {
	if x != nil {
		return x.Employes
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_hrd_proto_rawDescGZIP(), []int{2}
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_hrd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_hrd_proto_rawDescGZIP(), []int{3}
}

func (x *Result) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type NameId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameId) Reset() {
	*x = NameId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameId) ProtoMessage() {}

func (x *NameId) ProtoReflect() protoreflect.Message {
	mi := &file_hrd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameId.ProtoReflect.Descriptor instead.
func (*NameId) Descriptor() ([]byte, []int) {
	return file_hrd_proto_rawDescGZIP(), []int{4}
}

func (x *NameId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NameId) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetEmpByUsername struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *GetEmpByUsername) Reset() {
	*x = GetEmpByUsername{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmpByUsername) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmpByUsername) ProtoMessage() {}

func (x *GetEmpByUsername) ProtoReflect() protoreflect.Message {
	mi := &file_hrd_proto_msgTypes[5]
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
	return file_hrd_proto_rawDescGZIP(), []int{5}
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
		mi := &file_hrd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaporanEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaporanEmployee) ProtoMessage() {}

func (x *LaporanEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_hrd_proto_msgTypes[6]
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
	return file_hrd_proto_rawDescGZIP(), []int{6}
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
		mi := &file_hrd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrEmployee) ProtoMessage() {}

func (x *ArrEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_hrd_proto_msgTypes[7]
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
	return file_hrd_proto_rawDescGZIP(), []int{7}
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
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0xd0, 0x02, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x75, 0x74, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x75, 0x74, 0x69, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x6a, 0x69,
	0x50, 0x6f, 0x6b, 0x6f, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x67, 0x61, 0x6a,
	0x69, 0x50, 0x6f, 0x6b, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47,
	0x61, 0x6a, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x47, 0x61, 0x6a, 0x69, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x75, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x61,
	0x6e, 0x4d, 0x61, 0x6b, 0x61, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x74, 0x75,
	0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x4d, 0x61, 0x6b, 0x61, 0x6e, 0x12, 0x2e, 0x0a, 0x12,
	0x74, 0x75, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x74, 0x75, 0x6e, 0x6a, 0x61, 0x6e,
	0x67, 0x61, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x3d, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x24, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x0f, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x6a, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x6a, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x48, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x41, 0x72, 0x72, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x0a, 0x41, 0x72, 0x72, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x32, 0xb6, 0x03, 0x0a, 0x09,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x45, 0x6d, 0x70, 0x12, 0x3a, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x12, 0x2e, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a,
	0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x10, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x0c, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x10, 0x2e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x1a,
	0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0f, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x12, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x12, 0x10, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x49, 0x64, 0x1a, 0x10, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61,
	0x6e, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2e, 0x4c, 0x61, 0x70, 0x6f, 0x72, 0x61, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_hrd_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_hrd_proto_goTypes = []interface{}{
	(*Employee)(nil),         // 0: employee.Employee
	(*AllEmployee)(nil),      // 1: employee.AllEmployee
	(*Empty)(nil),            // 2: employee.Empty
	(*Result)(nil),           // 3: employee.Result
	(*NameId)(nil),           // 4: employee.NameId
	(*GetEmpByUsername)(nil), // 5: employee.GetEmpByUsername
	(*LaporanEmployee)(nil),  // 6: employee.LaporanEmployee
	(*ArrEmployee)(nil),      // 7: employee.ArrEmployee
}
var file_hrd_proto_depIdxs = []int32{
	0, // 0: employee.AllEmployee.employes:type_name -> employee.Employee
	6, // 1: employee.ArrEmployee.ArrLaporan:type_name -> employee.LaporanEmployee
	0, // 2: employee.manageEmp.CreateEmployee:input_type -> employee.Employee
	4, // 3: employee.manageEmp.FindEmployee:input_type -> employee.NameId
	4, // 4: employee.manageEmp.ReadEmployee:input_type -> employee.NameId
	2, // 5: employee.manageEmp.ReadAllEmployee:input_type -> employee.Empty
	0, // 6: employee.manageEmp.UpdateEmployee:input_type -> employee.Employee
	4, // 7: employee.manageEmp.DeleteEmployee:input_type -> employee.NameId
	5, // 8: employee.manageEmp.LaporanByUsername:input_type -> employee.GetEmpByUsername
	0, // 9: employee.manageEmp.CreateEmployee:output_type -> employee.Employee
	0, // 10: employee.manageEmp.FindEmployee:output_type -> employee.Employee
	0, // 11: employee.manageEmp.ReadEmployee:output_type -> employee.Employee
	1, // 12: employee.manageEmp.ReadAllEmployee:output_type -> employee.AllEmployee
	0, // 13: employee.manageEmp.UpdateEmployee:output_type -> employee.Employee
	3, // 14: employee.manageEmp.DeleteEmployee:output_type -> employee.Result
	6, // 15: employee.manageEmp.LaporanByUsername:output_type -> employee.LaporanEmployee
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hrd_proto_init() }
func file_hrd_proto_init() {
	if File_hrd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hrd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
			switch v := v.(*AllEmployee); i {
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
			switch v := v.(*Empty); i {
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
		file_hrd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_hrd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameId); i {
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
		file_hrd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_hrd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_hrd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   8,
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

// ManageEmpClient is the client API for ManageEmp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManageEmpClient interface {
	CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error)
	FindEmployee(ctx context.Context, in *NameId, opts ...grpc.CallOption) (*Employee, error)
	ReadEmployee(ctx context.Context, in *NameId, opts ...grpc.CallOption) (*Employee, error)
	ReadAllEmployee(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllEmployee, error)
	UpdateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error)
	DeleteEmployee(ctx context.Context, in *NameId, opts ...grpc.CallOption) (*Result, error)
	LaporanByUsername(ctx context.Context, in *GetEmpByUsername, opts ...grpc.CallOption) (*LaporanEmployee, error)
}

type manageEmpClient struct {
	cc grpc.ClientConnInterface
}

func NewManageEmpClient(cc grpc.ClientConnInterface) ManageEmpClient {
	return &manageEmpClient{cc}
}

func (c *manageEmpClient) CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/employee.manageEmp/CreateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmpClient) FindEmployee(ctx context.Context, in *NameId, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/employee.manageEmp/FindEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmpClient) ReadEmployee(ctx context.Context, in *NameId, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/employee.manageEmp/ReadEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmpClient) ReadAllEmployee(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllEmployee, error) {
	out := new(AllEmployee)
	err := c.cc.Invoke(ctx, "/employee.manageEmp/ReadAllEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmpClient) UpdateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/employee.manageEmp/UpdateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmpClient) DeleteEmployee(ctx context.Context, in *NameId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/employee.manageEmp/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageEmpClient) LaporanByUsername(ctx context.Context, in *GetEmpByUsername, opts ...grpc.CallOption) (*LaporanEmployee, error) {
	out := new(LaporanEmployee)
	err := c.cc.Invoke(ctx, "/employee.manageEmp/LaporanByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManageEmpServer is the server API for ManageEmp service.
type ManageEmpServer interface {
	CreateEmployee(context.Context, *Employee) (*Employee, error)
	FindEmployee(context.Context, *NameId) (*Employee, error)
	ReadEmployee(context.Context, *NameId) (*Employee, error)
	ReadAllEmployee(context.Context, *Empty) (*AllEmployee, error)
	UpdateEmployee(context.Context, *Employee) (*Employee, error)
	DeleteEmployee(context.Context, *NameId) (*Result, error)
	LaporanByUsername(context.Context, *GetEmpByUsername) (*LaporanEmployee, error)
}

// UnimplementedManageEmpServer can be embedded to have forward compatible implementations.
type UnimplementedManageEmpServer struct {
}

func (*UnimplementedManageEmpServer) CreateEmployee(context.Context, *Employee) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (*UnimplementedManageEmpServer) FindEmployee(context.Context, *NameId) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEmployee not implemented")
}
func (*UnimplementedManageEmpServer) ReadEmployee(context.Context, *NameId) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployee not implemented")
}
func (*UnimplementedManageEmpServer) ReadAllEmployee(context.Context, *Empty) (*AllEmployee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllEmployee not implemented")
}
func (*UnimplementedManageEmpServer) UpdateEmployee(context.Context, *Employee) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (*UnimplementedManageEmpServer) DeleteEmployee(context.Context, *NameId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (*UnimplementedManageEmpServer) LaporanByUsername(context.Context, *GetEmpByUsername) (*LaporanEmployee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaporanByUsername not implemented")
}

func RegisterManageEmpServer(s *grpc.Server, srv ManageEmpServer) {
	s.RegisterService(&_ManageEmp_serviceDesc, srv)
}

func _ManageEmp_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmpServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.manageEmp/CreateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmpServer).CreateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmp_FindEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmpServer).FindEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.manageEmp/FindEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmpServer).FindEmployee(ctx, req.(*NameId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmp_ReadEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmpServer).ReadEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.manageEmp/ReadEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmpServer).ReadEmployee(ctx, req.(*NameId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmp_ReadAllEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmpServer).ReadAllEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.manageEmp/ReadAllEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmpServer).ReadAllEmployee(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmp_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmpServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.manageEmp/UpdateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmpServer).UpdateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmp_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmpServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.manageEmp/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmpServer).DeleteEmployee(ctx, req.(*NameId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageEmp_LaporanByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmpByUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageEmpServer).LaporanByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.manageEmp/LaporanByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageEmpServer).LaporanByUsername(ctx, req.(*GetEmpByUsername))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManageEmp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "employee.manageEmp",
	HandlerType: (*ManageEmpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmployee",
			Handler:    _ManageEmp_CreateEmployee_Handler,
		},
		{
			MethodName: "FindEmployee",
			Handler:    _ManageEmp_FindEmployee_Handler,
		},
		{
			MethodName: "ReadEmployee",
			Handler:    _ManageEmp_ReadEmployee_Handler,
		},
		{
			MethodName: "ReadAllEmployee",
			Handler:    _ManageEmp_ReadAllEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _ManageEmp_UpdateEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _ManageEmp_DeleteEmployee_Handler,
		},
		{
			MethodName: "LaporanByUsername",
			Handler:    _ManageEmp_LaporanByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hrd.proto",
}