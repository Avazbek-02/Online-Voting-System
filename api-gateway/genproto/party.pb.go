// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: party.proto

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

type CreatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,2,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedAt    string `protobuf:"bytes,3,opt,name=opened_at,json=openedAt,proto3" json:"opened_at,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreatePartyRequest) Reset() {
	*x = CreatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartyRequest) ProtoMessage() {}

func (x *CreatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePartyRequest.ProtoReflect.Descriptor instead.
func (*CreatePartyRequest) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePartyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePartyRequest) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *CreatePartyRequest) GetOpenedAt() string {
	if x != nil {
		return x.OpenedAt
	}
	return ""
}

func (x *CreatePartyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetPartyInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPartyInfoRequest) Reset() {
	*x = GetPartyInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartyInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartyInfoRequest) ProtoMessage() {}

func (x *GetPartyInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartyInfoRequest.ProtoReflect.Descriptor instead.
func (*GetPartyInfoRequest) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{1}
}

func (x *GetPartyInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdatePartyRequest) Reset() {
	*x = UpdatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartyRequest) ProtoMessage() {}

func (x *UpdatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePartyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePartyRequest) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePartyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePartyRequest) Reset() {
	*x = DeletePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePartyRequest) ProtoMessage() {}

func (x *DeletePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePartyRequest.ProtoReflect.Descriptor instead.
func (*DeletePartyRequest) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePartyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedAt    string `protobuf:"bytes,4,opt,name=opened_at,json=openedAt,proto3" json:"opened_at,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PartyResponse) Reset() {
	*x = PartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyResponse) ProtoMessage() {}

func (x *PartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyResponse.ProtoReflect.Descriptor instead.
func (*PartyResponse) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{4}
}

func (x *PartyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PartyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PartyResponse) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *PartyResponse) GetOpenedAt() string {
	if x != nil {
		return x.OpenedAt
	}
	return ""
}

func (x *PartyResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Voidd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Voidd) Reset() {
	*x = Voidd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voidd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voidd) ProtoMessage() {}

func (x *Voidd) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voidd.ProtoReflect.Descriptor instead.
func (*Voidd) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{5}
}

var File_party_proto protoreflect.FileDescriptor

var file_party_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x7f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x56, 0x6f, 0x69, 0x64, 0x64, 0x32,
	0x98, 0x02, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x64, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_proto_rawDescOnce sync.Once
	file_party_proto_rawDescData = file_party_proto_rawDesc
)

func file_party_proto_rawDescGZIP() []byte {
	file_party_proto_rawDescOnce.Do(func() {
		file_party_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_proto_rawDescData)
	})
	return file_party_proto_rawDescData
}

var file_party_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_party_proto_goTypes = []interface{}{
	(*CreatePartyRequest)(nil),  // 0: protos.CreatePartyRequest
	(*GetPartyInfoRequest)(nil), // 1: protos.GetPartyInfoRequest
	(*UpdatePartyRequest)(nil),  // 2: protos.UpdatePartyRequest
	(*DeletePartyRequest)(nil),  // 3: protos.DeletePartyRequest
	(*PartyResponse)(nil),       // 4: protos.PartyResponse
	(*Voidd)(nil),               // 5: protos.Voidd
}
var file_party_proto_depIdxs = []int32{
	0, // 0: protos.PartyService.CreateParty:input_type -> protos.CreatePartyRequest
	1, // 1: protos.PartyService.GetPartyInfo:input_type -> protos.GetPartyInfoRequest
	2, // 2: protos.PartyService.UpdateParty:input_type -> protos.UpdatePartyRequest
	3, // 3: protos.PartyService.DeleteParty:input_type -> protos.DeletePartyRequest
	4, // 4: protos.PartyService.CreateParty:output_type -> protos.PartyResponse
	4, // 5: protos.PartyService.GetPartyInfo:output_type -> protos.PartyResponse
	4, // 6: protos.PartyService.UpdateParty:output_type -> protos.PartyResponse
	5, // 7: protos.PartyService.DeleteParty:output_type -> protos.Voidd
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_party_proto_init() }
func file_party_proto_init() {
	if File_party_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePartyRequest); i {
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
		file_party_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartyInfoRequest); i {
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
		file_party_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePartyRequest); i {
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
		file_party_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePartyRequest); i {
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
		file_party_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyResponse); i {
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
		file_party_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voidd); i {
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
			RawDescriptor: file_party_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_party_proto_goTypes,
		DependencyIndexes: file_party_proto_depIdxs,
		MessageInfos:      file_party_proto_msgTypes,
	}.Build()
	File_party_proto = out.File
	file_party_proto_rawDesc = nil
	file_party_proto_goTypes = nil
	file_party_proto_depIdxs = nil
}
