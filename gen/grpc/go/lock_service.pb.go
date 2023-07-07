// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: lock_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_lock_service_proto protoreflect.FileDescriptor

var file_lock_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x1a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x86, 0x01, 0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x4c, 0x6f,
	0x63, 0x6b, 0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x12, 0x3a, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x42, 0x43, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x75, 0x73, 0x6c, 0x61,
	0x6e, 0x53, 0x6f, 0x72, 0x6f, 0x6b, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x2d, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_lock_service_proto_goTypes = []interface{}{
	(*LockReq)(nil),   // 0: lock_manager.LockReq
	(*UnlockReq)(nil), // 1: lock_manager.UnlockReq
	(*LockRes)(nil),   // 2: lock_manager.LockRes
	(*UnlockRes)(nil), // 3: lock_manager.UnlockRes
}
var file_lock_service_proto_depIdxs = []int32{
	0, // 0: lock_manager.LockManagerService.Lock:input_type -> lock_manager.LockReq
	1, // 1: lock_manager.LockManagerService.Unlock:input_type -> lock_manager.UnlockReq
	2, // 2: lock_manager.LockManagerService.Lock:output_type -> lock_manager.LockRes
	3, // 3: lock_manager.LockManagerService.Unlock:output_type -> lock_manager.UnlockRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lock_service_proto_init() }
func file_lock_service_proto_init() {
	if File_lock_service_proto != nil {
		return
	}
	file_transport_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lock_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lock_service_proto_goTypes,
		DependencyIndexes: file_lock_service_proto_depIdxs,
	}.Build()
	File_lock_service_proto = out.File
	file_lock_service_proto_rawDesc = nil
	file_lock_service_proto_goTypes = nil
	file_lock_service_proto_depIdxs = nil
}
