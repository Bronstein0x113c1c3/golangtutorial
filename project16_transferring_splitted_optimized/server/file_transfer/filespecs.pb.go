// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: filespecs.proto

package file_transfer

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

type File_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *File_Info) Reset() {
	*x = File_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filespecs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File_Info) ProtoMessage() {}

func (x *File_Info) ProtoReflect() protoreflect.Message {
	mi := &file_filespecs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File_Info.ProtoReflect.Descriptor instead.
func (*File_Info) Descriptor() ([]byte, []int) {
	return file_filespecs_proto_rawDescGZIP(), []int{0}
}

func (x *File_Info) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *File_Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filespecs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_filespecs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_filespecs_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_filespecs_proto protoreflect.FileDescriptor

var file_filespecs_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x33, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x32, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65,
	0x5f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x08, 0x53, 0x65, 0x6e,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0a, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x5f, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x28, 0x01, 0x42, 0x15, 0x5a, 0x13,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filespecs_proto_rawDescOnce sync.Once
	file_filespecs_proto_rawDescData = file_filespecs_proto_rawDesc
)

func file_filespecs_proto_rawDescGZIP() []byte {
	file_filespecs_proto_rawDescOnce.Do(func() {
		file_filespecs_proto_rawDescData = protoimpl.X.CompressGZIP(file_filespecs_proto_rawDescData)
	})
	return file_filespecs_proto_rawDescData
}

var file_filespecs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_filespecs_proto_goTypes = []interface{}{
	(*File_Info)(nil), // 0: File_Info
	(*Status)(nil),    // 1: Status
}
var file_filespecs_proto_depIdxs = []int32{
	0, // 0: File_Transfer.SendFile:input_type -> File_Info
	1, // 1: File_Transfer.SendFile:output_type -> Status
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_filespecs_proto_init() }
func file_filespecs_proto_init() {
	if File_filespecs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filespecs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File_Info); i {
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
		file_filespecs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_filespecs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filespecs_proto_goTypes,
		DependencyIndexes: file_filespecs_proto_depIdxs,
		MessageInfos:      file_filespecs_proto_msgTypes,
	}.Build()
	File_filespecs_proto = out.File
	file_filespecs_proto_rawDesc = nil
	file_filespecs_proto_goTypes = nil
	file_filespecs_proto_depIdxs = nil
}
