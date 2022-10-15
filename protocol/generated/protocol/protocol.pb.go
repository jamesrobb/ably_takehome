// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: protocol/protocol.proto

package protocol

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

type NumbersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUIDv4 identifying the requesting client.
	ClientId   []byte `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	NumNumbers uint32 `protobuf:"varint,2,opt,name=num_numbers,json=numNumbers,proto3" json:"num_numbers,omitempty"`
	// Used for debugging/testing purposes. Specifies the seed for the server's PRNG.
	Seed uint32 `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (x *NumbersRequest) Reset() {
	*x = NumbersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumbersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumbersRequest) ProtoMessage() {}

func (x *NumbersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumbersRequest.ProtoReflect.Descriptor instead.
func (*NumbersRequest) Descriptor() ([]byte, []int) {
	return file_protocol_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *NumbersRequest) GetClientId() []byte {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *NumbersRequest) GetNumNumbers() uint32 {
	if x != nil {
		return x.NumNumbers
	}
	return 0
}

func (x *NumbersRequest) GetSeed() uint32 {
	if x != nil {
		return x.Seed
	}
	return 0
}

type NumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// When a NumberResponse is the last message for a NumbersRequest the checkum is set, otherwise it is an empty string.
	Checksum string `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *NumberResponse) Reset() {
	*x = NumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberResponse) ProtoMessage() {}

func (x *NumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberResponse.ProtoReflect.Descriptor instead.
func (*NumberResponse) Descriptor() ([]byte, []int) {
	return file_protocol_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *NumberResponse) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *NumberResponse) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

var File_protocol_protocol_proto protoreflect.FileDescriptor

var file_protocol_protocol_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x22, 0x62, 0x0a, 0x0e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x22, 0x44, 0x0a, 0x0e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x32, 0x4d, 0x0a,
	0x07, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6d, 0x65, 0x73,
	0x72, 0x6f, 0x62, 0x62, 0x2f, 0x61, 0x62, 0x6c, 0x79, 0x2d, 0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f,
	0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_protocol_proto_rawDescOnce sync.Once
	file_protocol_protocol_proto_rawDescData = file_protocol_protocol_proto_rawDesc
)

func file_protocol_protocol_proto_rawDescGZIP() []byte {
	file_protocol_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_protocol_proto_rawDescData)
	})
	return file_protocol_protocol_proto_rawDescData
}

var file_protocol_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_protocol_proto_goTypes = []interface{}{
	(*NumbersRequest)(nil), // 0: protocol.NumbersRequest
	(*NumberResponse)(nil), // 1: protocol.NumberResponse
}
var file_protocol_protocol_proto_depIdxs = []int32{
	0, // 0: protocol.Numbers.GetNumbers:input_type -> protocol.NumbersRequest
	1, // 1: protocol.Numbers.GetNumbers:output_type -> protocol.NumberResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_protocol_proto_init() }
func file_protocol_protocol_proto_init() {
	if File_protocol_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumbersRequest); i {
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
		file_protocol_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberResponse); i {
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
			RawDescriptor: file_protocol_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_protocol_proto_depIdxs,
		MessageInfos:      file_protocol_protocol_proto_msgTypes,
	}.Build()
	File_protocol_protocol_proto = out.File
	file_protocol_protocol_proto_rawDesc = nil
	file_protocol_protocol_proto_goTypes = nil
	file_protocol_protocol_proto_depIdxs = nil
}
