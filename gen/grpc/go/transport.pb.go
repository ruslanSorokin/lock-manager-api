// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: transport.proto

package proto

import (
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

// Error is a enum error.
type LockRes_Error int32

const (
	LockRes_ERR_UNSPECIFIED LockRes_Error = 0
	// ERR_INVALID_RESOURCE_ID means resource_id violates one of its constraint
	LockRes_ERR_INVALID_RESOURCE_ID LockRes_Error = 1
	// ERR_RESOURCE_ID_ALREADY_EXISTS means resource_id is already locked
	LockRes_ERR_RESOURCE_ID_ALREADY_EXISTS LockRes_Error = 2
)

// Enum value maps for LockRes_Error.
var (
	LockRes_Error_name = map[int32]string{
		0: "ERR_UNSPECIFIED",
		1: "ERR_INVALID_RESOURCE_ID",
		2: "ERR_RESOURCE_ID_ALREADY_EXISTS",
	}
	LockRes_Error_value = map[string]int32{
		"ERR_UNSPECIFIED":                0,
		"ERR_INVALID_RESOURCE_ID":        1,
		"ERR_RESOURCE_ID_ALREADY_EXISTS": 2,
	}
)

func (x LockRes_Error) Enum() *LockRes_Error {
	p := new(LockRes_Error)
	*p = x
	return p
}

func (x LockRes_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockRes_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_proto_enumTypes[0].Descriptor()
}

func (LockRes_Error) Type() protoreflect.EnumType {
	return &file_transport_proto_enumTypes[0]
}

func (x LockRes_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LockRes_Error.Descriptor instead.
func (LockRes_Error) EnumDescriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1, 0}
}

// Error is a enum error.
type UnlockRes_Error int32

const (
	UnlockRes_ERR_UNSPECIFIED UnlockRes_Error = 0
	// ERR_INVALID_TOKEN means token violates one of its constraint
	UnlockRes_ERR_INVALID_TOKEN UnlockRes_Error = 1
	// ERR_WRONG_TOKEN means token doesn't fit
	UnlockRes_ERR_WRONG_TOKEN UnlockRes_Error = 2
	// ERR_INVALID_RESOURCE_ID means resource_id violates one of its constraint
	UnlockRes_ERR_INVALID_RESOURCE_ID UnlockRes_Error = 3
	// ERR_RESOURCE_ID_NOT_FOUND means resource_id was not found
	UnlockRes_ERR_RESOURCE_ID_NOT_FOUND UnlockRes_Error = 4
)

// Enum value maps for UnlockRes_Error.
var (
	UnlockRes_Error_name = map[int32]string{
		0: "ERR_UNSPECIFIED",
		1: "ERR_INVALID_TOKEN",
		2: "ERR_WRONG_TOKEN",
		3: "ERR_INVALID_RESOURCE_ID",
		4: "ERR_RESOURCE_ID_NOT_FOUND",
	}
	UnlockRes_Error_value = map[string]int32{
		"ERR_UNSPECIFIED":           0,
		"ERR_INVALID_TOKEN":         1,
		"ERR_WRONG_TOKEN":           2,
		"ERR_INVALID_RESOURCE_ID":   3,
		"ERR_RESOURCE_ID_NOT_FOUND": 4,
	}
)

func (x UnlockRes_Error) Enum() *UnlockRes_Error {
	p := new(UnlockRes_Error)
	*p = x
	return p
}

func (x UnlockRes_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnlockRes_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_proto_enumTypes[1].Descriptor()
}

func (UnlockRes_Error) Type() protoreflect.EnumType {
	return &file_transport_proto_enumTypes[1]
}

func (x UnlockRes_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UnlockRes_Error.Descriptor instead.
func (UnlockRes_Error) EnumDescriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{3, 0}
}

// LockReq is a request for Lock call.
type LockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	ResourceId *string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3,oneof" json:"resource_id,omitempty"`
}

func (x *LockReq) Reset() {
	*x = LockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockReq) ProtoMessage() {}

func (x *LockReq) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockReq.ProtoReflect.Descriptor instead.
func (*LockReq) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{0}
}

func (x *LockReq) GetResourceId() string {
	if x != nil && x.ResourceId != nil {
		return *x.ResourceId
	}
	return ""
}

// LockRes is a response for LockReq.
type LockRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []LockRes_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=lock_manager.LockRes_Error" json:"errors,omitempty"`
	// Required.
	Token *string `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *LockRes) Reset() {
	*x = LockRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockRes) ProtoMessage() {}

func (x *LockRes) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockRes.ProtoReflect.Descriptor instead.
func (*LockRes) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1}
}

func (x *LockRes) GetErrors() []LockRes_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *LockRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

// UnlockReq is a request for Unlock call.
type UnlockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	ResourceId *string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3,oneof" json:"resource_id,omitempty"`
	// Required.
	Token *string `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *UnlockReq) Reset() {
	*x = UnlockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockReq) ProtoMessage() {}

func (x *UnlockReq) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockReq.ProtoReflect.Descriptor instead.
func (*UnlockReq) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{2}
}

func (x *UnlockReq) GetResourceId() string {
	if x != nil && x.ResourceId != nil {
		return *x.ResourceId
	}
	return ""
}

func (x *UnlockReq) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

// UnlockRes is a response for UnlockReq.
type UnlockRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []UnlockRes_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=lock_manager.UnlockRes_Error" json:"errors,omitempty"`
	// Required.
	Token *string `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *UnlockRes) Reset() {
	*x = UnlockRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockRes) ProtoMessage() {}

func (x *UnlockRes) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockRes.ProtoReflect.Descriptor instead.
func (*UnlockRes) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{3}
}

func (x *UnlockRes) GetErrors() []UnlockRes_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *UnlockRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

var File_transport_proto protoreflect.FileDescriptor

var file_transport_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22,
	0x3f, 0x0a, 0x07, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x22, 0xc2, 0x01, 0x0a, 0x07, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x22, 0x5d, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52,
	0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x52, 0x52, 0x5f, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x02, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x66, 0x0a, 0x09, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xee, 0x01,
	0x0a, 0x09, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x22, 0x84, 0x01,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x04, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x43,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x75, 0x73,
	0x6c, 0x61, 0x6e, 0x53, 0x6f, 0x72, 0x6f, 0x6b, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_proto_rawDescOnce sync.Once
	file_transport_proto_rawDescData = file_transport_proto_rawDesc
)

func file_transport_proto_rawDescGZIP() []byte {
	file_transport_proto_rawDescOnce.Do(func() {
		file_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_proto_rawDescData)
	})
	return file_transport_proto_rawDescData
}

var file_transport_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transport_proto_goTypes = []interface{}{
	(LockRes_Error)(0),   // 0: lock_manager.LockRes.Error
	(UnlockRes_Error)(0), // 1: lock_manager.UnlockRes.Error
	(*LockReq)(nil),      // 2: lock_manager.LockReq
	(*LockRes)(nil),      // 3: lock_manager.LockRes
	(*UnlockReq)(nil),    // 4: lock_manager.UnlockReq
	(*UnlockRes)(nil),    // 5: lock_manager.UnlockRes
}
var file_transport_proto_depIdxs = []int32{
	0, // 0: lock_manager.LockRes.errors:type_name -> lock_manager.LockRes.Error
	1, // 1: lock_manager.UnlockRes.errors:type_name -> lock_manager.UnlockRes.Error
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transport_proto_init() }
func file_transport_proto_init() {
	if File_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockReq); i {
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
		file_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockRes); i {
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
		file_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockReq); i {
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
		file_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockRes); i {
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
	file_transport_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_transport_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_transport_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_transport_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transport_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_proto_goTypes,
		DependencyIndexes: file_transport_proto_depIdxs,
		EnumInfos:         file_transport_proto_enumTypes,
		MessageInfos:      file_transport_proto_msgTypes,
	}.Build()
	File_transport_proto = out.File
	file_transport_proto_rawDesc = nil
	file_transport_proto_goTypes = nil
	file_transport_proto_depIdxs = nil
}
