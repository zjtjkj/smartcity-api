// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: api/home-config/v1/home-config.proto

package v1

import (
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

type HomeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Conf string `protobuf:"bytes,2,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (x *HomeConf) Reset() {
	*x = HomeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeConf) ProtoMessage() {}

func (x *HomeConf) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeConf.ProtoReflect.Descriptor instead.
func (*HomeConf) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{0}
}

func (x *HomeConf) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HomeConf) GetConf() string {
	if x != nil {
		return x.Conf
	}
	return ""
}

type SaveHomeConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HomeConf *HomeConf `protobuf:"bytes,1,opt,name=home_conf,json=homeConf,proto3" json:"home_conf,omitempty"`
}

func (x *SaveHomeConfigRequest) Reset() {
	*x = SaveHomeConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveHomeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveHomeConfigRequest) ProtoMessage() {}

func (x *SaveHomeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveHomeConfigRequest.ProtoReflect.Descriptor instead.
func (*SaveHomeConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{1}
}

func (x *SaveHomeConfigRequest) GetHomeConf() *HomeConf {
	if x != nil {
		return x.HomeConf
	}
	return nil
}

type SaveHomeConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveHomeConfigReply) Reset() {
	*x = SaveHomeConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveHomeConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveHomeConfigReply) ProtoMessage() {}

func (x *SaveHomeConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveHomeConfigReply.ProtoReflect.Descriptor instead.
func (*SaveHomeConfigReply) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{2}
}

type GetHomeConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHomeConfigRequest) Reset() {
	*x = GetHomeConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeConfigRequest) ProtoMessage() {}

func (x *GetHomeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeConfigRequest.ProtoReflect.Descriptor instead.
func (*GetHomeConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{3}
}

type GetHomeConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HomeConf *HomeConf `protobuf:"bytes,1,opt,name=home_conf,json=homeConf,proto3" json:"home_conf,omitempty"`
}

func (x *GetHomeConfigReply) Reset() {
	*x = GetHomeConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeConfigReply) ProtoMessage() {}

func (x *GetHomeConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeConfigReply.ProtoReflect.Descriptor instead.
func (*GetHomeConfigReply) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetHomeConfigReply) GetHomeConf() *HomeConf {
	if x != nil {
		return x.HomeConf
	}
	return nil
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data int32  `protobuf:"varint,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{5}
}

func (x *Option) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Option) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Option) GetData() int32 {
	if x != nil {
		return x.Data
	}
	return 0
}

type GetOptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOptionsRequest) Reset() {
	*x = GetOptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptionsRequest) ProtoMessage() {}

func (x *GetOptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptionsRequest.ProtoReflect.Descriptor instead.
func (*GetOptionsRequest) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{6}
}

type GetOptionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options []*Option `protobuf:"bytes,1,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *GetOptionsReply) Reset() {
	*x = GetOptionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptionsReply) ProtoMessage() {}

func (x *GetOptionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptionsReply.ProtoReflect.Descriptor instead.
func (*GetOptionsReply) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{7}
}

func (x *GetOptionsReply) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

type GetOptionByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOptionByIdRequest) Reset() {
	*x = GetOptionByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptionByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptionByIdRequest) ProtoMessage() {}

func (x *GetOptionByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptionByIdRequest.ProtoReflect.Descriptor instead.
func (*GetOptionByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{8}
}

func (x *GetOptionByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOptionByIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Option *Option `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *GetOptionByIdReply) Reset() {
	*x = GetOptionByIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_home_config_v1_home_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptionByIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptionByIdReply) ProtoMessage() {}

func (x *GetOptionByIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_home_config_v1_home_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptionByIdReply.ProtoReflect.Descriptor instead.
func (*GetOptionByIdReply) Descriptor() ([]byte, []int) {
	return file_api_home_config_v1_home_config_proto_rawDescGZIP(), []int{9}
}

func (x *GetOptionByIdReply) GetOption() *Option {
	if x != nil {
		return x.Option
	}
	return nil
}

var File_api_home_config_v1_home_config_proto protoreflect.FileDescriptor

var file_api_home_config_v1_home_config_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x6e,
	0x66, 0x22, 0x4b, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x68, 0x6f,
	0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x22, 0x15,
	0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x08, 0x68,
	0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x22, 0x40, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2d, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b,
	0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xeb, 0x03, 0x0a, 0x0a,
	0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x7b, 0x0a, 0x0e, 0x53, 0x61,
	0x76, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x48,
	0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x73, 0x61, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f,
	0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x6e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f,
	0x6d, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x77, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x32, 0x0a, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x21, 0x68, 0x6f, 0x6d, 0x65,
	0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x6f, 0x6d, 0x65,
	0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_home_config_v1_home_config_proto_rawDescOnce sync.Once
	file_api_home_config_v1_home_config_proto_rawDescData = file_api_home_config_v1_home_config_proto_rawDesc
)

func file_api_home_config_v1_home_config_proto_rawDescGZIP() []byte {
	file_api_home_config_v1_home_config_proto_rawDescOnce.Do(func() {
		file_api_home_config_v1_home_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_home_config_v1_home_config_proto_rawDescData)
	})
	return file_api_home_config_v1_home_config_proto_rawDescData
}

var file_api_home_config_v1_home_config_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_home_config_v1_home_config_proto_goTypes = []interface{}{
	(*HomeConf)(nil),              // 0: api.home.v1.HomeConf
	(*SaveHomeConfigRequest)(nil), // 1: api.home.v1.SaveHomeConfigRequest
	(*SaveHomeConfigReply)(nil),   // 2: api.home.v1.SaveHomeConfigReply
	(*GetHomeConfigRequest)(nil),  // 3: api.home.v1.GetHomeConfigRequest
	(*GetHomeConfigReply)(nil),    // 4: api.home.v1.GetHomeConfigReply
	(*Option)(nil),                // 5: api.home.v1.Option
	(*GetOptionsRequest)(nil),     // 6: api.home.v1.GetOptionsRequest
	(*GetOptionsReply)(nil),       // 7: api.home.v1.GetOptionsReply
	(*GetOptionByIdRequest)(nil),  // 8: api.home.v1.GetOptionByIdRequest
	(*GetOptionByIdReply)(nil),    // 9: api.home.v1.GetOptionByIdReply
}
var file_api_home_config_v1_home_config_proto_depIdxs = []int32{
	0, // 0: api.home.v1.SaveHomeConfigRequest.home_conf:type_name -> api.home.v1.HomeConf
	0, // 1: api.home.v1.GetHomeConfigReply.home_conf:type_name -> api.home.v1.HomeConf
	5, // 2: api.home.v1.GetOptionsReply.options:type_name -> api.home.v1.Option
	5, // 3: api.home.v1.GetOptionByIdReply.option:type_name -> api.home.v1.Option
	1, // 4: api.home.v1.HomeConfig.SaveHomeConfig:input_type -> api.home.v1.SaveHomeConfigRequest
	3, // 5: api.home.v1.HomeConfig.GetHomeConfig:input_type -> api.home.v1.GetHomeConfigRequest
	6, // 6: api.home.v1.HomeConfig.GetOptions:input_type -> api.home.v1.GetOptionsRequest
	8, // 7: api.home.v1.HomeConfig.GetOptionById:input_type -> api.home.v1.GetOptionByIdRequest
	2, // 8: api.home.v1.HomeConfig.SaveHomeConfig:output_type -> api.home.v1.SaveHomeConfigReply
	4, // 9: api.home.v1.HomeConfig.GetHomeConfig:output_type -> api.home.v1.GetHomeConfigReply
	7, // 10: api.home.v1.HomeConfig.GetOptions:output_type -> api.home.v1.GetOptionsReply
	9, // 11: api.home.v1.HomeConfig.GetOptionById:output_type -> api.home.v1.GetOptionByIdReply
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_home_config_v1_home_config_proto_init() }
func file_api_home_config_v1_home_config_proto_init() {
	if File_api_home_config_v1_home_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_home_config_v1_home_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeConf); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveHomeConfigRequest); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveHomeConfigReply); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeConfigRequest); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeConfigReply); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptionsRequest); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptionsReply); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptionByIdRequest); i {
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
		file_api_home_config_v1_home_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptionByIdReply); i {
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
			RawDescriptor: file_api_home_config_v1_home_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_home_config_v1_home_config_proto_goTypes,
		DependencyIndexes: file_api_home_config_v1_home_config_proto_depIdxs,
		MessageInfos:      file_api_home_config_v1_home_config_proto_msgTypes,
	}.Build()
	File_api_home_config_v1_home_config_proto = out.File
	file_api_home_config_v1_home_config_proto_rawDesc = nil
	file_api_home_config_v1_home_config_proto_goTypes = nil
	file_api_home_config_v1_home_config_proto_depIdxs = nil
}
