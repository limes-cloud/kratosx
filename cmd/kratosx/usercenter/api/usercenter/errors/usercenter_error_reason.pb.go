// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: usercenter_error_reason.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_SystemError      ErrorReason = 0
	ErrorReason_ParamsError      ErrorReason = 1
	ErrorReason_DatabaseError    ErrorReason = 2
	ErrorReason_TransformError   ErrorReason = 3
	ErrorReason_GetError         ErrorReason = 4
	ErrorReason_ListError        ErrorReason = 5
	ErrorReason_CreateError      ErrorReason = 6
	ErrorReason_ImportError      ErrorReason = 7
	ErrorReason_ExportError      ErrorReason = 8
	ErrorReason_UpdateError      ErrorReason = 9
	ErrorReason_DeleteError      ErrorReason = 10
	ErrorReason_GetTrashError    ErrorReason = 11
	ErrorReason_ListTrashError   ErrorReason = 12
	ErrorReason_DeleteTrashError ErrorReason = 13
	ErrorReason_RevertTrashError ErrorReason = 14
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "SystemError",
		1:  "ParamsError",
		2:  "DatabaseError",
		3:  "TransformError",
		4:  "GetError",
		5:  "ListError",
		6:  "CreateError",
		7:  "ImportError",
		8:  "ExportError",
		9:  "UpdateError",
		10: "DeleteError",
		11: "GetTrashError",
		12: "ListTrashError",
		13: "DeleteTrashError",
		14: "RevertTrashError",
	}
	ErrorReason_value = map[string]int32{
		"SystemError":      0,
		"ParamsError":      1,
		"DatabaseError":    2,
		"TransformError":   3,
		"GetError":         4,
		"ListError":        5,
		"CreateError":      6,
		"ImportError":      7,
		"ExportError":      8,
		"UpdateError":      9,
		"DeleteError":      10,
		"GetTrashError":    11,
		"ListTrashError":   12,
		"DeleteTrashError": 13,
		"RevertTrashError": 14,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_usercenter_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_usercenter_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_usercenter_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_usercenter_error_reason_proto protoreflect.FileDescriptor

var file_usercenter_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9b, 0x05, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x1a, 0x0f, 0xb2,
	0x45, 0x0c, 0xe7, 0xb3, 0xbb, 0xe7, 0xbb, 0x9f, 0xe5, 0xbc, 0x82, 0xe5, 0xb8, 0xb8, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x1a,
	0x0f, 0xb2, 0x45, 0x0c, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf,
	0x12, 0x25, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x02, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xba,
	0x93, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x29, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x03, 0x1a, 0x15, 0xb2, 0x45, 0x12,
	0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe8, 0xbd, 0xac, 0xe6, 0x8d, 0xa2, 0xe5, 0xa4, 0xb1, 0xe8,
	0xb4, 0xa5, 0x12, 0x23, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x04,
	0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe6, 0x95, 0xb0, 0xe6, 0x8d,
	0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x2a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x05, 0x1a, 0x1b, 0xb2, 0x45, 0x18, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f,
	0x96, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1,
	0xe8, 0xb4, 0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x06, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe6,
	0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x07, 0x1a, 0x15, 0xb2, 0x45,
	0x12, 0xe5, 0xaf, 0xbc, 0xe5, 0x85, 0xa5, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1,
	0xe8, 0xb4, 0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x08, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe5, 0xaf, 0xbc, 0xe5, 0x87, 0xba, 0xe6,
	0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x09, 0x1a, 0x15, 0xb2, 0x45,
	0x12, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1,
	0xe8, 0xb4, 0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x0a, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe6,
	0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x31, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x73, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0b, 0x1a, 0x1e,
	0xb2, 0x45, 0x1b, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe5, 0x9b, 0x9e, 0xe6, 0x94, 0xb6, 0xe7,
	0xab, 0x99, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x38,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x73, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x0c, 0x1a, 0x24, 0xb2, 0x45, 0x21, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe5, 0x9b, 0x9e,
	0xe6, 0x94, 0xb6, 0xe7, 0xab, 0x99, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0xe6, 0x95, 0xb0, 0xe6,
	0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x34, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x73, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0d, 0x1a, 0x1e,
	0xb2, 0x45, 0x1b, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe5, 0x9b, 0x9e, 0xe6, 0x94, 0xb6, 0xe7,
	0xab, 0x99, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x34,
	0x0a, 0x10, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x54, 0x72, 0x61, 0x73, 0x68, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x0e, 0x1a, 0x1e, 0xb2, 0x45, 0x1b, 0xe8, 0xbf, 0x98, 0xe5, 0x8e, 0x9f, 0xe5,
	0x9b, 0x9e, 0xe6, 0x94, 0xb6, 0xe7, 0xab, 0x99, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4,
	0xb1, 0xe8, 0xb4, 0xa5, 0x1a, 0x04, 0xa0, 0x45, 0x90, 0x03, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f,
	0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usercenter_error_reason_proto_rawDescOnce sync.Once
	file_usercenter_error_reason_proto_rawDescData = file_usercenter_error_reason_proto_rawDesc
)

func file_usercenter_error_reason_proto_rawDescGZIP() []byte {
	file_usercenter_error_reason_proto_rawDescOnce.Do(func() {
		file_usercenter_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_usercenter_error_reason_proto_rawDescData)
	})
	return file_usercenter_error_reason_proto_rawDescData
}

var file_usercenter_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_usercenter_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: errors.ErrorReason
}
var file_usercenter_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_usercenter_error_reason_proto_init() }
func file_usercenter_error_reason_proto_init() {
	if File_usercenter_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_usercenter_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_usercenter_error_reason_proto_goTypes,
		DependencyIndexes: file_usercenter_error_reason_proto_depIdxs,
		EnumInfos:         file_usercenter_error_reason_proto_enumTypes,
	}.Build()
	File_usercenter_error_reason_proto = out.File
	file_usercenter_error_reason_proto_rawDesc = nil
	file_usercenter_error_reason_proto_goTypes = nil
	file_usercenter_error_reason_proto_depIdxs = nil
}
