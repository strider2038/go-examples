// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.2
// source: messaging.proto

package messaging

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DispatchResult_Status int32

const (
	DispatchResult_Success DispatchResult_Status = 0
	DispatchResult_Fail    DispatchResult_Status = 1
)

// Enum value maps for DispatchResult_Status.
var (
	DispatchResult_Status_name = map[int32]string{
		0: "Success",
		1: "Fail",
	}
	DispatchResult_Status_value = map[string]int32{
		"Success": 0,
		"Fail":    1,
	}
)

func (x DispatchResult_Status) Enum() *DispatchResult_Status {
	p := new(DispatchResult_Status)
	*p = x
	return p
}

func (x DispatchResult_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DispatchResult_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_messaging_proto_enumTypes[0].Descriptor()
}

func (DispatchResult_Status) Type() protoreflect.EnumType {
	return &file_messaging_proto_enumTypes[0]
}

func (x DispatchResult_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DispatchResult_Status.Descriptor instead.
func (DispatchResult_Status) EnumDescriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{4, 0}
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hi uint64 `protobuf:"varint,1,opt,name=hi,proto3" json:"hi,omitempty"`
	Lo uint64 `protobuf:"varint,2,opt,name=lo,proto3" json:"lo,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetHi() uint64 {
	if x != nil {
		return x.Hi
	}
	return 0
}

func (x *UUID) GetLo() uint64 {
	if x != nil {
		return x.Lo
	}
	return 0
}

type Headers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Headers) Reset() {
	*x = Headers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Headers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Headers) ProtoMessage() {}

func (x *Headers) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Headers.ProtoReflect.Descriptor instead.
func (*Headers) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{1}
}

func (x *Headers) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers    *Headers        `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	Body       *structpb.Value `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Properties *structpb.Value `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{2}
}

func (x *Envelope) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Envelope) GetBody() *structpb.Value {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Envelope) GetProperties() *structpb.Value {
	if x != nil {
		return x.Properties
	}
	return nil
}

type EnvelopeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelopes []*Envelope `protobuf:"bytes,1,rep,name=envelopes,proto3" json:"envelopes,omitempty"`
}

func (x *EnvelopeList) Reset() {
	*x = EnvelopeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvelopeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvelopeList) ProtoMessage() {}

func (x *EnvelopeList) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvelopeList.ProtoReflect.Descriptor instead.
func (*EnvelopeList) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{3}
}

func (x *EnvelopeList) GetEnvelopes() []*Envelope {
	if x != nil {
		return x.Envelopes
	}
	return nil
}

type DispatchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status DispatchResult_Status `protobuf:"varint,1,opt,name=status,proto3,enum=messaging.DispatchResult_Status" json:"status,omitempty"`
	Id     *UUID                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Errors []string              `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *DispatchResult) Reset() {
	*x = DispatchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchResult) ProtoMessage() {}

func (x *DispatchResult) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchResult.ProtoReflect.Descriptor instead.
func (*DispatchResult) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{4}
}

func (x *DispatchResult) GetStatus() DispatchResult_Status {
	if x != nil {
		return x.Status
	}
	return DispatchResult_Success
}

func (x *DispatchResult) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DispatchResult) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_messaging_proto protoreflect.FileDescriptor

var file_messaging_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x12, 0x0e, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x68, 0x69,
	0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x6c, 0x6f,
	0x22, 0x2a, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a,
	0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0c, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x22, 0xa4,
	0x01, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x61, 0x69, 0x6c, 0x10, 0x01, 0x32, 0xbd, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messaging_proto_rawDescOnce sync.Once
	file_messaging_proto_rawDescData = file_messaging_proto_rawDesc
)

func file_messaging_proto_rawDescGZIP() []byte {
	file_messaging_proto_rawDescOnce.Do(func() {
		file_messaging_proto_rawDescData = protoimpl.X.CompressGZIP(file_messaging_proto_rawDescData)
	})
	return file_messaging_proto_rawDescData
}

var file_messaging_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messaging_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_messaging_proto_goTypes = []interface{}{
	(DispatchResult_Status)(0), // 0: messaging.DispatchResult.Status
	(*UUID)(nil),               // 1: messaging.UUID
	(*Headers)(nil),            // 2: messaging.Headers
	(*Envelope)(nil),           // 3: messaging.Envelope
	(*EnvelopeList)(nil),       // 4: messaging.EnvelopeList
	(*DispatchResult)(nil),     // 5: messaging.DispatchResult
	(*structpb.Value)(nil),     // 6: google.protobuf.Value
}
var file_messaging_proto_depIdxs = []int32{
	1, // 0: messaging.Headers.id:type_name -> messaging.UUID
	2, // 1: messaging.Envelope.headers:type_name -> messaging.Headers
	6, // 2: messaging.Envelope.body:type_name -> google.protobuf.Value
	6, // 3: messaging.Envelope.properties:type_name -> google.protobuf.Value
	3, // 4: messaging.EnvelopeList.envelopes:type_name -> messaging.Envelope
	0, // 5: messaging.DispatchResult.status:type_name -> messaging.DispatchResult.Status
	1, // 6: messaging.DispatchResult.id:type_name -> messaging.UUID
	3, // 7: messaging.Messenger.Dispatch:input_type -> messaging.Envelope
	4, // 8: messaging.Messenger.DispatchQueue:input_type -> messaging.EnvelopeList
	5, // 9: messaging.Messenger.Dispatch:output_type -> messaging.DispatchResult
	5, // 10: messaging.Messenger.DispatchQueue:output_type -> messaging.DispatchResult
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_messaging_proto_init() }
func file_messaging_proto_init() {
	if File_messaging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messaging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_messaging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Headers); i {
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
		file_messaging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_messaging_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvelopeList); i {
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
		file_messaging_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchResult); i {
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
			RawDescriptor: file_messaging_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messaging_proto_goTypes,
		DependencyIndexes: file_messaging_proto_depIdxs,
		EnumInfos:         file_messaging_proto_enumTypes,
		MessageInfos:      file_messaging_proto_msgTypes,
	}.Build()
	File_messaging_proto = out.File
	file_messaging_proto_rawDesc = nil
	file_messaging_proto_goTypes = nil
	file_messaging_proto_depIdxs = nil
}
