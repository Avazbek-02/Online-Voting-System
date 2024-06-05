// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: candidate.proto

package voting

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

type CreateCandidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	PartyId    string `protobuf:"bytes,4,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *CreateCandidateRequest) Reset() {
	*x = CreateCandidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCandidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCandidateRequest) ProtoMessage() {}

func (x *CreateCandidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCandidateRequest.ProtoReflect.Descriptor instead.
func (*CreateCandidateRequest) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCandidateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateCandidateRequest) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *CreateCandidateRequest) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *CreateCandidateRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type GetCandidateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCandidateInfoRequest) Reset() {
	*x = GetCandidateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCandidateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandidateInfoRequest) ProtoMessage() {}

func (x *GetCandidateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandidateInfoRequest.ProtoReflect.Descriptor instead.
func (*GetCandidateInfoRequest) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{1}
}

func (x *GetCandidateInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateCandidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateCandidateRequest) Reset() {
	*x = UpdateCandidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCandidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCandidateRequest) ProtoMessage() {}

func (x *UpdateCandidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCandidateRequest.ProtoReflect.Descriptor instead.
func (*UpdateCandidateRequest) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCandidateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCandidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCandidateRequest) Reset() {
	*x = DeleteCandidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCandidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCandidateRequest) ProtoMessage() {}

func (x *DeleteCandidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCandidateRequest.ProtoReflect.Descriptor instead.
func (*DeleteCandidateRequest) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCandidateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CandidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	PartyId    string `protobuf:"bytes,4,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *CandidateResponse) Reset() {
	*x = CandidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateResponse) ProtoMessage() {}

func (x *CandidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateResponse.ProtoReflect.Descriptor instead.
func (*CandidateResponse) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{4}
}

func (x *CandidateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CandidateResponse) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *CandidateResponse) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *CandidateResponse) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type Void1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void1) Reset() {
	*x = Void1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void1) ProtoMessage() {}

func (x *Void1) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void1.ProtoReflect.Descriptor instead.
func (*Void1) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{5}
}

var File_candidate_proto protoreflect.FileDescriptor

var file_candidate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x81, 0x01, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x22, 0x29, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7c, 0x0a, 0x11,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x56, 0x6f,
	0x69, 0x64, 0x31, 0x32, 0xc8, 0x02, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x76,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x31, 0x22, 0x00, 0x42, 0x12,
	0x5a, 0x10, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_candidate_proto_rawDescOnce sync.Once
	file_candidate_proto_rawDescData = file_candidate_proto_rawDesc
)

func file_candidate_proto_rawDescGZIP() []byte {
	file_candidate_proto_rawDescOnce.Do(func() {
		file_candidate_proto_rawDescData = protoimpl.X.CompressGZIP(file_candidate_proto_rawDescData)
	})
	return file_candidate_proto_rawDescData
}

var file_candidate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_candidate_proto_goTypes = []interface{}{
	(*CreateCandidateRequest)(nil),  // 0: voting.CreateCandidateRequest
	(*GetCandidateInfoRequest)(nil), // 1: voting.GetCandidateInfoRequest
	(*UpdateCandidateRequest)(nil),  // 2: voting.UpdateCandidateRequest
	(*DeleteCandidateRequest)(nil),  // 3: voting.DeleteCandidateRequest
	(*CandidateResponse)(nil),       // 4: voting.CandidateResponse
	(*Void1)(nil),                   // 5: voting.Void1
}
var file_candidate_proto_depIdxs = []int32{
	0, // 0: voting.CandidateService.CreateCandidate:input_type -> voting.CreateCandidateRequest
	1, // 1: voting.CandidateService.GetCandidateInfo:input_type -> voting.GetCandidateInfoRequest
	2, // 2: voting.CandidateService.UpdateCandidate:input_type -> voting.UpdateCandidateRequest
	3, // 3: voting.CandidateService.DeleteCandidate:input_type -> voting.DeleteCandidateRequest
	4, // 4: voting.CandidateService.CreateCandidate:output_type -> voting.CandidateResponse
	4, // 5: voting.CandidateService.GetCandidateInfo:output_type -> voting.CandidateResponse
	4, // 6: voting.CandidateService.UpdateCandidate:output_type -> voting.CandidateResponse
	5, // 7: voting.CandidateService.DeleteCandidate:output_type -> voting.Void1
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_candidate_proto_init() }
func file_candidate_proto_init() {
	if File_candidate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_candidate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCandidateRequest); i {
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
		file_candidate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCandidateInfoRequest); i {
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
		file_candidate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCandidateRequest); i {
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
		file_candidate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCandidateRequest); i {
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
		file_candidate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateResponse); i {
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
		file_candidate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void1); i {
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
			RawDescriptor: file_candidate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_candidate_proto_goTypes,
		DependencyIndexes: file_candidate_proto_depIdxs,
		MessageInfos:      file_candidate_proto_msgTypes,
	}.Build()
	File_candidate_proto = out.File
	file_candidate_proto_rawDesc = nil
	file_candidate_proto_goTypes = nil
	file_candidate_proto_depIdxs = nil
}