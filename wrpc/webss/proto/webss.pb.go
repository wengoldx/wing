// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/webss.proto

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

type WEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WEmpty) Reset() {
	*x = WEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WEmpty) ProtoMessage() {}

func (x *WEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WEmpty.ProtoReflect.Descriptor instead.
func (*WEmpty) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{0}
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DelFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Files  []string `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *DelFiles) Reset() {
	*x = DelFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelFiles) ProtoMessage() {}

func (x *DelFiles) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelFiles.ProtoReflect.Descriptor instead.
func (*DelFiles) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{2}
}

func (x *DelFiles) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *DelFiles) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type DeleteLifeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lifecycleid []string `protobuf:"bytes,1,rep,name=lifecycleid,proto3" json:"lifecycleid,omitempty"`
	Bucket      string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
}

func (x *DeleteLifeConfig) Reset() {
	*x = DeleteLifeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLifeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLifeConfig) ProtoMessage() {}

func (x *DeleteLifeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLifeConfig.ProtoReflect.Descriptor instead.
func (*DeleteLifeConfig) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteLifeConfig) GetLifecycleid() []string {
	if x != nil {
		return x.Lifecycleid
	}
	return nil
}

func (x *DeleteLifeConfig) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

type LifeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value  string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LifeConfig) Reset() {
	*x = LifeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifeConfig) ProtoMessage() {}

func (x *LifeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifeConfig.ProtoReflect.Descriptor instead.
func (*LifeConfig) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{4}
}

func (x *LifeConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LifeConfig) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *LifeConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LifeConfig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Path   []string `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Key    string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value  string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{5}
}

func (x *Tag) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Tag) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Tag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PerSigUrlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res    string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	Add    string `protobuf:"bytes,2,opt,name=add,proto3" json:"add,omitempty"`
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3" json:"suffix,omitempty"`
}

func (x *PerSigUrlReq) Reset() {
	*x = PerSigUrlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerSigUrlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerSigUrlReq) ProtoMessage() {}

func (x *PerSigUrlReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerSigUrlReq.ProtoReflect.Descriptor instead.
func (*PerSigUrlReq) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{6}
}

func (x *PerSigUrlReq) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

func (x *PerSigUrlReq) GetAdd() string {
	if x != nil {
		return x.Add
	}
	return ""
}

func (x *PerSigUrlReq) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

type PerSigUrlsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res    string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	Add    string   `protobuf:"bytes,2,opt,name=add,proto3" json:"add,omitempty"`
	Suffix []string `protobuf:"bytes,3,rep,name=suffix,proto3" json:"suffix,omitempty"`
}

func (x *PerSigUrlsReq) Reset() {
	*x = PerSigUrlsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerSigUrlsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerSigUrlsReq) ProtoMessage() {}

func (x *PerSigUrlsReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerSigUrlsReq.ProtoReflect.Descriptor instead.
func (*PerSigUrlsReq) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{7}
}

func (x *PerSigUrlsReq) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

func (x *PerSigUrlsReq) GetAdd() string {
	if x != nil {
		return x.Add
	}
	return ""
}

func (x *PerSigUrlsReq) GetSuffix() []string {
	if x != nil {
		return x.Suffix
	}
	return nil
}

type PreSigUrls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreSigUrls []*PreSigUrl `protobuf:"bytes,1,rep,name=PreSigUrls,proto3" json:"PreSigUrls,omitempty"`
}

func (x *PreSigUrls) Reset() {
	*x = PreSigUrls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreSigUrls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreSigUrls) ProtoMessage() {}

func (x *PreSigUrls) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreSigUrls.ProtoReflect.Descriptor instead.
func (*PreSigUrls) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{8}
}

func (x *PreSigUrls) GetPreSigUrls() []*PreSigUrl {
	if x != nil {
		return x.PreSigUrls
	}
	return nil
}

type PreSigUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PreSigUrl) Reset() {
	*x = PreSigUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_webss_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreSigUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreSigUrl) ProtoMessage() {}

func (x *PreSigUrl) ProtoReflect() protoreflect.Message {
	mi := &file_proto_webss_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreSigUrl.ProtoReflect.Descriptor instead.
func (*PreSigUrl) Descriptor() ([]byte, []int) {
	return file_proto_webss_proto_rawDescGZIP(), []int{9}
}

func (x *PreSigUrl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PreSigUrl) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_proto_webss_proto protoreflect.FileDescriptor

var file_proto_webss_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06, 0x57, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x08, 0x44, 0x65,
	0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69,
	0x66, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x22, 0x5c, 0x0a, 0x0a, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x59, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4a, 0x0a, 0x0c, 0x50,
	0x65, 0x72, 0x53, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x64, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x64, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0x4b, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x53, 0x69,
	0x67, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x64,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x64, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75,
	0x66, 0x66, 0x69, 0x78, 0x22, 0x3e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x55, 0x72,
	0x6c, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x65, 0x53, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x0a, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67,
	0x55, 0x72, 0x6c, 0x73, 0x22, 0x31, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x55, 0x72,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0xae, 0x02, 0x0a, 0x05, 0x57, 0x65, 0x62, 0x73,
	0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x36, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x69, 0x66, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69,
	0x66, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x54, 0x61, 0x67, 0x12, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x65, 0x53, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x65, 0x72, 0x53, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x12,
	0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x72, 0x6c, 0x73, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x53, 0x69, 0x67, 0x55, 0x72,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x65, 0x53, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_webss_proto_rawDescOnce sync.Once
	file_proto_webss_proto_rawDescData = file_proto_webss_proto_rawDesc
)

func file_proto_webss_proto_rawDescGZIP() []byte {
	file_proto_webss_proto_rawDescOnce.Do(func() {
		file_proto_webss_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_webss_proto_rawDescData)
	})
	return file_proto_webss_proto_rawDescData
}

var file_proto_webss_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_webss_proto_goTypes = []interface{}{
	(*WEmpty)(nil),           // 0: proto.WEmpty
	(*ID)(nil),               // 1: proto.ID
	(*DelFiles)(nil),         // 2: proto.DelFiles
	(*DeleteLifeConfig)(nil), // 3: proto.DeleteLifeConfig
	(*LifeConfig)(nil),       // 4: proto.LifeConfig
	(*Tag)(nil),              // 5: proto.Tag
	(*PerSigUrlReq)(nil),     // 6: proto.PerSigUrlReq
	(*PerSigUrlsReq)(nil),    // 7: proto.PerSigUrlsReq
	(*PreSigUrls)(nil),       // 8: proto.PreSigUrls
	(*PreSigUrl)(nil),        // 9: proto.PreSigUrl
}
var file_proto_webss_proto_depIdxs = []int32{
	9, // 0: proto.PreSigUrls.PreSigUrls:type_name -> proto.PreSigUrl
	2, // 1: proto.Webss.DeleteFiles:input_type -> proto.DelFiles
	3, // 2: proto.Webss.DeleteConfig:input_type -> proto.DeleteLifeConfig
	4, // 3: proto.Webss.AddConfig:input_type -> proto.LifeConfig
	5, // 4: proto.Webss.AddTag:input_type -> proto.Tag
	6, // 5: proto.Webss.GetPreSigUrl:input_type -> proto.PerSigUrlReq
	7, // 6: proto.Webss.GetMultiUrls:input_type -> proto.PerSigUrlsReq
	0, // 7: proto.Webss.DeleteFiles:output_type -> proto.WEmpty
	0, // 8: proto.Webss.DeleteConfig:output_type -> proto.WEmpty
	1, // 9: proto.Webss.AddConfig:output_type -> proto.ID
	0, // 10: proto.Webss.AddTag:output_type -> proto.WEmpty
	9, // 11: proto.Webss.GetPreSigUrl:output_type -> proto.PreSigUrl
	8, // 12: proto.Webss.GetMultiUrls:output_type -> proto.PreSigUrls
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_webss_proto_init() }
func file_proto_webss_proto_init() {
	if File_proto_webss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_webss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WEmpty); i {
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
		file_proto_webss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_proto_webss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelFiles); i {
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
		file_proto_webss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLifeConfig); i {
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
		file_proto_webss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LifeConfig); i {
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
		file_proto_webss_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_proto_webss_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerSigUrlReq); i {
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
		file_proto_webss_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerSigUrlsReq); i {
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
		file_proto_webss_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreSigUrls); i {
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
		file_proto_webss_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreSigUrl); i {
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
			RawDescriptor: file_proto_webss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_webss_proto_goTypes,
		DependencyIndexes: file_proto_webss_proto_depIdxs,
		MessageInfos:      file_proto_webss_proto_msgTypes,
	}.Build()
	File_proto_webss_proto = out.File
	file_proto_webss_proto_rawDesc = nil
	file_proto_webss_proto_goTypes = nil
	file_proto_webss_proto_depIdxs = nil
}
