// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: events/v1/events.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ReceiveEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mid      uint32                        `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Type     string                        `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Created  string                        `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Image    []byte                        `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Finished bool                          `protobuf:"varint,6,opt,name=finished,proto3" json:"finished,omitempty"`
	Objects  []*ReceiveEventRequest_Object `protobuf:"bytes,7,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *ReceiveEventRequest) Reset() {
	*x = ReceiveEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveEventRequest) ProtoMessage() {}

func (x *ReceiveEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveEventRequest.ProtoReflect.Descriptor instead.
func (*ReceiveEventRequest) Descriptor() ([]byte, []int) {
	return file_events_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReceiveEventRequest) GetMid() uint32 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *ReceiveEventRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ReceiveEventRequest) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *ReceiveEventRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ReceiveEventRequest) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

func (x *ReceiveEventRequest) GetObjects() []*ReceiveEventRequest_Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

type ReceiveEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveEventReply) Reset() {
	*x = ReceiveEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_v1_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveEventReply) ProtoMessage() {}

func (x *ReceiveEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_events_v1_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveEventReply.ProtoReflect.Descriptor instead.
func (*ReceiveEventReply) Descriptor() ([]byte, []int) {
	return file_events_v1_events_proto_rawDescGZIP(), []int{1}
}

type ReceiveEventRequest_Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X uint32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y uint32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *ReceiveEventRequest_Point) Reset() {
	*x = ReceiveEventRequest_Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_v1_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveEventRequest_Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveEventRequest_Point) ProtoMessage() {}

func (x *ReceiveEventRequest_Point) ProtoReflect() protoreflect.Message {
	mi := &file_events_v1_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveEventRequest_Point.ProtoReflect.Descriptor instead.
func (*ReceiveEventRequest_Point) Descriptor() ([]byte, []int) {
	return file_events_v1_events_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ReceiveEventRequest_Point) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ReceiveEventRequest_Point) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type ReceiveEventRequest_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ReceiveEventRequest_Property) Reset() {
	*x = ReceiveEventRequest_Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_v1_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveEventRequest_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveEventRequest_Property) ProtoMessage() {}

func (x *ReceiveEventRequest_Property) ProtoReflect() protoreflect.Message {
	mi := &file_events_v1_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveEventRequest_Property.ProtoReflect.Descriptor instead.
func (*ReceiveEventRequest_Property) Descriptor() ([]byte, []int) {
	return file_events_v1_events_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ReceiveEventRequest_Property) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReceiveEventRequest_Property) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ReceiveEventRequest_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Aid        uint32                          `protobuf:"varint,2,opt,name=aid,proto3" json:"aid,omitempty"`
	Points     []*ReceiveEventRequest_Point    `protobuf:"bytes,3,rep,name=points,proto3" json:"points,omitempty"`
	Properties []*ReceiveEventRequest_Property `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *ReceiveEventRequest_Object) Reset() {
	*x = ReceiveEventRequest_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_v1_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveEventRequest_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveEventRequest_Object) ProtoMessage() {}

func (x *ReceiveEventRequest_Object) ProtoReflect() protoreflect.Message {
	mi := &file_events_v1_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveEventRequest_Object.ProtoReflect.Descriptor instead.
func (*ReceiveEventRequest_Object) Descriptor() ([]byte, []int) {
	return file_events_v1_events_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ReceiveEventRequest_Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReceiveEventRequest_Object) GetAid() uint32 {
	if x != nil {
		return x.Aid
	}
	return 0
}

func (x *ReceiveEventRequest_Object) GetPoints() []*ReceiveEventRequest_Point {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *ReceiveEventRequest_Object) GetProperties() []*ReceiveEventRequest_Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

var File_events_v1_events_proto protoreflect.FileDescriptor

var file_events_v1_events_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0,
	0x04, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xa0, 0x01, 0x24, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a,
	0x02, 0x10, 0x01, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a, 0x3b, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x2a, 0x05,
	0x18, 0x90, 0x4e, 0x28, 0x00, 0x52, 0x01, 0x78, 0x12, 0x18, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x2a, 0x05, 0x18, 0x90, 0x4e, 0x28, 0x00, 0x52,
	0x01, 0x79, 0x1a, 0x44, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x19,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0xd6, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x69, 0x64, 0x12, 0x4c,
	0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x92, 0x01, 0x04,
	0x08, 0x02, 0x10, 0x02, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x53, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x06, 0xfa,
	0x42, 0x03, 0x92, 0x01, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x79, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x6f, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01,
	0x2a, 0x42, 0x19, 0x5a, 0x17, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_v1_events_proto_rawDescOnce sync.Once
	file_events_v1_events_proto_rawDescData = file_events_v1_events_proto_rawDesc
)

func file_events_v1_events_proto_rawDescGZIP() []byte {
	file_events_v1_events_proto_rawDescOnce.Do(func() {
		file_events_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_v1_events_proto_rawDescData)
	})
	return file_events_v1_events_proto_rawDescData
}

var file_events_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_events_v1_events_proto_goTypes = []interface{}{
	(*ReceiveEventRequest)(nil),          // 0: api.events.v1.ReceiveEventRequest
	(*ReceiveEventReply)(nil),            // 1: api.events.v1.ReceiveEventReply
	(*ReceiveEventRequest_Point)(nil),    // 2: api.events.v1.ReceiveEventRequest.Point
	(*ReceiveEventRequest_Property)(nil), // 3: api.events.v1.ReceiveEventRequest.Property
	(*ReceiveEventRequest_Object)(nil),   // 4: api.events.v1.ReceiveEventRequest.Object
}
var file_events_v1_events_proto_depIdxs = []int32{
	4, // 0: api.events.v1.ReceiveEventRequest.objects:type_name -> api.events.v1.ReceiveEventRequest.Object
	2, // 1: api.events.v1.ReceiveEventRequest.Object.points:type_name -> api.events.v1.ReceiveEventRequest.Point
	3, // 2: api.events.v1.ReceiveEventRequest.Object.properties:type_name -> api.events.v1.ReceiveEventRequest.Property
	0, // 3: api.events.v1.Events.ReceiveEvent:input_type -> api.events.v1.ReceiveEventRequest
	1, // 4: api.events.v1.Events.ReceiveEvent:output_type -> api.events.v1.ReceiveEventReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_events_v1_events_proto_init() }
func file_events_v1_events_proto_init() {
	if File_events_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveEventRequest); i {
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
		file_events_v1_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveEventReply); i {
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
		file_events_v1_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveEventRequest_Point); i {
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
		file_events_v1_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveEventRequest_Property); i {
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
		file_events_v1_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveEventRequest_Object); i {
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
			RawDescriptor: file_events_v1_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_events_v1_events_proto_goTypes,
		DependencyIndexes: file_events_v1_events_proto_depIdxs,
		MessageInfos:      file_events_v1_events_proto_msgTypes,
	}.Build()
	File_events_v1_events_proto = out.File
	file_events_v1_events_proto_rawDesc = nil
	file_events_v1_events_proto_goTypes = nil
	file_events_v1_events_proto_depIdxs = nil
}
