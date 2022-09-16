// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: service/shell/v1/shell.proto

package v1

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

type ShellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DataIn string `protobuf:"bytes,2,opt,name=data_in,json=dataIn,proto3" json:"data_in,omitempty"`
}

func (x *ShellRequest) Reset() {
	*x = ShellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_shell_v1_shell_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellRequest) ProtoMessage() {}

func (x *ShellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_shell_v1_shell_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellRequest.ProtoReflect.Descriptor instead.
func (*ShellRequest) Descriptor() ([]byte, []int) {
	return file_service_shell_v1_shell_proto_rawDescGZIP(), []int{0}
}

func (x *ShellRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShellRequest) GetDataIn() string {
	if x != nil {
		return x.DataIn
	}
	return ""
}

type ShellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataOut string `protobuf:"bytes,1,opt,name=data_out,json=dataOut,proto3" json:"data_out,omitempty"`
	DataErr string `protobuf:"bytes,2,opt,name=data_err,json=dataErr,proto3" json:"data_err,omitempty"`
	Eof     bool   `protobuf:"varint,3,opt,name=eof,proto3" json:"eof,omitempty"`
}

func (x *ShellResponse) Reset() {
	*x = ShellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_shell_v1_shell_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellResponse) ProtoMessage() {}

func (x *ShellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_shell_v1_shell_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellResponse.ProtoReflect.Descriptor instead.
func (*ShellResponse) Descriptor() ([]byte, []int) {
	return file_service_shell_v1_shell_proto_rawDescGZIP(), []int{1}
}

func (x *ShellResponse) GetDataOut() string {
	if x != nil {
		return x.DataOut
	}
	return ""
}

func (x *ShellResponse) GetDataErr() string {
	if x != nil {
		return x.DataErr
	}
	return ""
}

func (x *ShellResponse) GetEof() bool {
	if x != nil {
		return x.Eof
	}
	return false
}

var File_service_shell_v1_shell_proto protoreflect.FileDescriptor

var file_service_shell_v1_shell_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x76, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x68, 0x65,
	0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x3b, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x61,
	0x49, 0x6e, 0x22, 0x57, 0x0a, 0x0d, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x4f, 0x75, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6f, 0x66, 0x32, 0x66, 0x0a, 0x0c, 0x53,
	0x68, 0x65, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x05, 0x53,
	0x68, 0x65, 0x6c, 0x6c, 0x12, 0x23, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x69, 0x61, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x3d, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x76, 0x31,
	0x5a, 0x20, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_shell_v1_shell_proto_rawDescOnce sync.Once
	file_service_shell_v1_shell_proto_rawDescData = file_service_shell_v1_shell_proto_rawDesc
)

func file_service_shell_v1_shell_proto_rawDescGZIP() []byte {
	file_service_shell_v1_shell_proto_rawDescOnce.Do(func() {
		file_service_shell_v1_shell_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_shell_v1_shell_proto_rawDescData)
	})
	return file_service_shell_v1_shell_proto_rawDescData
}

var file_service_shell_v1_shell_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_shell_v1_shell_proto_goTypes = []interface{}{
	(*ShellRequest)(nil),  // 0: viam.service.shell.v1.ShellRequest
	(*ShellResponse)(nil), // 1: viam.service.shell.v1.ShellResponse
}
var file_service_shell_v1_shell_proto_depIdxs = []int32{
	0, // 0: viam.service.shell.v1.ShellService.Shell:input_type -> viam.service.shell.v1.ShellRequest
	1, // 1: viam.service.shell.v1.ShellService.Shell:output_type -> viam.service.shell.v1.ShellResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_shell_v1_shell_proto_init() }
func file_service_shell_v1_shell_proto_init() {
	if File_service_shell_v1_shell_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_shell_v1_shell_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellRequest); i {
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
		file_service_shell_v1_shell_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellResponse); i {
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
			RawDescriptor: file_service_shell_v1_shell_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_shell_v1_shell_proto_goTypes,
		DependencyIndexes: file_service_shell_v1_shell_proto_depIdxs,
		MessageInfos:      file_service_shell_v1_shell_proto_msgTypes,
	}.Build()
	File_service_shell_v1_shell_proto = out.File
	file_service_shell_v1_shell_proto_rawDesc = nil
	file_service_shell_v1_shell_proto_goTypes = nil
	file_service_shell_v1_shell_proto_depIdxs = nil
}