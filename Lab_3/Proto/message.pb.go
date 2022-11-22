// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: Proto/message.proto

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

type BaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base   string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Valor  string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *BaseMessage) Reset() {
	*x = BaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseMessage) ProtoMessage() {}

func (x *BaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseMessage.ProtoReflect.Descriptor instead.
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return file_Proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *BaseMessage) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *BaseMessage) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *BaseMessage) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

type RenameMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector  string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base    string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Newbase string `protobuf:"bytes,3,opt,name=newbase,proto3" json:"newbase,omitempty"`
}

func (x *RenameMessage) Reset() {
	*x = RenameMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameMessage) ProtoMessage() {}

func (x *RenameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameMessage.ProtoReflect.Descriptor instead.
func (*RenameMessage) Descriptor() ([]byte, []int) {
	return file_Proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *RenameMessage) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *RenameMessage) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *RenameMessage) GetNewbase() string {
	if x != nil {
		return x.Newbase
	}
	return ""
}

type QueryMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base   string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *QueryMessage) Reset() {
	*x = QueryMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMessage) ProtoMessage() {}

func (x *QueryMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMessage.ProtoReflect.Descriptor instead.
func (*QueryMessage) Descriptor() ([]byte, []int) {
	return file_Proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMessage) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *QueryMessage) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

type ReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valor string `protobuf:"bytes,1,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *ReplyMessage) Reset() {
	*x = ReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMessage) ProtoMessage() {}

func (x *ReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMessage.ProtoReflect.Descriptor instead.
func (*ReplyMessage) Descriptor() ([]byte, []int) {
	return file_Proto_message_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyMessage) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

var File_Proto_message_proto protoreflect.FileDescriptor

var file_Proto_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x4f, 0x0a, 0x0b, 0x42,
	0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x55, 0x0a, 0x0d,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x77,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x62,
	0x61, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22,
	0x24, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x32, 0xd5, 0x01, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3b, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x53, 0x44,
	0x2f, 0x4c, 0x41, 0x42, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Proto_message_proto_rawDescOnce sync.Once
	file_Proto_message_proto_rawDescData = file_Proto_message_proto_rawDesc
)

func file_Proto_message_proto_rawDescGZIP() []byte {
	file_Proto_message_proto_rawDescOnce.Do(func() {
		file_Proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_message_proto_rawDescData)
	})
	return file_Proto_message_proto_rawDescData
}

var file_Proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Proto_message_proto_goTypes = []interface{}{
	(*BaseMessage)(nil),   // 0: grpc.BaseMessage
	(*RenameMessage)(nil), // 1: grpc.RenameMessage
	(*QueryMessage)(nil),  // 2: grpc.QueryMessage
	(*ReplyMessage)(nil),  // 3: grpc.ReplyMessage
}
var file_Proto_message_proto_depIdxs = []int32{
	0, // 0: grpc.PlanetaryService.Add:input_type -> grpc.BaseMessage
	1, // 1: grpc.PlanetaryService.Rename:input_type -> grpc.RenameMessage
	0, // 2: grpc.PlanetaryService.Update:input_type -> grpc.BaseMessage
	0, // 3: grpc.PlanetaryService.Delete:input_type -> grpc.BaseMessage
	2, // 4: grpc.GetService.Get:input_type -> grpc.QueryMessage
	3, // 5: grpc.PlanetaryService.Add:output_type -> grpc.ReplyMessage
	3, // 6: grpc.PlanetaryService.Rename:output_type -> grpc.ReplyMessage
	3, // 7: grpc.PlanetaryService.Update:output_type -> grpc.ReplyMessage
	3, // 8: grpc.PlanetaryService.Delete:output_type -> grpc.ReplyMessage
	3, // 9: grpc.GetService.Get:output_type -> grpc.ReplyMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Proto_message_proto_init() }
func file_Proto_message_proto_init() {
	if File_Proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseMessage); i {
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
		file_Proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameMessage); i {
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
		file_Proto_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMessage); i {
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
		file_Proto_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMessage); i {
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
			RawDescriptor: file_Proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_Proto_message_proto_goTypes,
		DependencyIndexes: file_Proto_message_proto_depIdxs,
		MessageInfos:      file_Proto_message_proto_msgTypes,
	}.Build()
	File_Proto_message_proto = out.File
	file_Proto_message_proto_rawDesc = nil
	file_Proto_message_proto_goTypes = nil
	file_Proto_message_proto_depIdxs = nil
}
