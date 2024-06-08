// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: election.proto

package genproto

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

type CreateElectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *CreateElectionRequest) Reset() {
	*x = CreateElectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateElectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateElectionRequest) ProtoMessage() {}

func (x *CreateElectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_election_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateElectionRequest.ProtoReflect.Descriptor instead.
func (*CreateElectionRequest) Descriptor() ([]byte, []int) {
	return file_election_proto_rawDescGZIP(), []int{0}
}

func (x *CreateElectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateElectionRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetElectionInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetElectionInfoRequest) Reset() {
	*x = GetElectionInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetElectionInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetElectionInfoRequest) ProtoMessage() {}

func (x *GetElectionInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_election_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetElectionInfoRequest.ProtoReflect.Descriptor instead.
func (*GetElectionInfoRequest) Descriptor() ([]byte, []int) {
	return file_election_proto_rawDescGZIP(), []int{1}
}

func (x *GetElectionInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateElectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateElectionRequest) Reset() {
	*x = UpdateElectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateElectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateElectionRequest) ProtoMessage() {}

func (x *UpdateElectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_election_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateElectionRequest.ProtoReflect.Descriptor instead.
func (*UpdateElectionRequest) Descriptor() ([]byte, []int) {
	return file_election_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateElectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteElectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteElectionRequest) Reset() {
	*x = DeleteElectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteElectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteElectionRequest) ProtoMessage() {}

func (x *DeleteElectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_election_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteElectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteElectionRequest) Descriptor() ([]byte, []int) {
	return file_election_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteElectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ElectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ElectionResponse) Reset() {
	*x = ElectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionResponse) ProtoMessage() {}

func (x *ElectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_election_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionResponse.ProtoReflect.Descriptor instead.
func (*ElectionResponse) Descriptor() ([]byte, []int) {
	return file_election_proto_rawDescGZIP(), []int{4}
}

func (x *ElectionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ElectionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ElectionResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type Void2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void2) Reset() {
	*x = Void2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void2) ProtoMessage() {}

func (x *Void2) ProtoReflect() protoreflect.Message {
	mi := &file_election_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void2.ProtoReflect.Descriptor instead.
func (*Void2) Descriptor() ([]byte, []int) {
	return file_election_proto_rawDescGZIP(), []int{5}
}

var File_election_proto protoreflect.FileDescriptor

var file_election_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x10, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x32, 0xbc, 0x02, 0x0a, 0x0f, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e,
	0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x76,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_election_proto_rawDescOnce sync.Once
	file_election_proto_rawDescData = file_election_proto_rawDesc
)

func file_election_proto_rawDescGZIP() []byte {
	file_election_proto_rawDescOnce.Do(func() {
		file_election_proto_rawDescData = protoimpl.X.CompressGZIP(file_election_proto_rawDescData)
	})
	return file_election_proto_rawDescData
}

var file_election_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_election_proto_goTypes = []interface{}{
	(*CreateElectionRequest)(nil),  // 0: voting.CreateElectionRequest
	(*GetElectionInfoRequest)(nil), // 1: voting.GetElectionInfoRequest
	(*UpdateElectionRequest)(nil),  // 2: voting.UpdateElectionRequest
	(*DeleteElectionRequest)(nil),  // 3: voting.DeleteElectionRequest
	(*ElectionResponse)(nil),       // 4: voting.ElectionResponse
	(*Void2)(nil),                  // 5: voting.Void2
}
var file_election_proto_depIdxs = []int32{
	0, // 0: voting.ElectionService.CreateElection:input_type -> voting.CreateElectionRequest
	1, // 1: voting.ElectionService.GetElectionInfo:input_type -> voting.GetElectionInfoRequest
	2, // 2: voting.ElectionService.UpdateElection:input_type -> voting.UpdateElectionRequest
	3, // 3: voting.ElectionService.DeleteElection:input_type -> voting.DeleteElectionRequest
	4, // 4: voting.ElectionService.CreateElection:output_type -> voting.ElectionResponse
	4, // 5: voting.ElectionService.GetElectionInfo:output_type -> voting.ElectionResponse
	4, // 6: voting.ElectionService.UpdateElection:output_type -> voting.ElectionResponse
	5, // 7: voting.ElectionService.DeleteElection:output_type -> voting.Void2
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_election_proto_init() }
func file_election_proto_init() {
	if File_election_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_election_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateElectionRequest); i {
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
		file_election_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetElectionInfoRequest); i {
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
		file_election_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateElectionRequest); i {
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
		file_election_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteElectionRequest); i {
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
		file_election_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionResponse); i {
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
		file_election_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void2); i {
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
			RawDescriptor: file_election_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_election_proto_goTypes,
		DependencyIndexes: file_election_proto_depIdxs,
		MessageInfos:      file_election_proto_msgTypes,
	}.Build()
	File_election_proto = out.File
	file_election_proto_rawDesc = nil
	file_election_proto_goTypes = nil
	file_election_proto_depIdxs = nil
}
