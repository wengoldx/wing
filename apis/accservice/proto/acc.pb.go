// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/acc.proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{0}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // account request token
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // account unique id
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[2]
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
	return file_proto_acc_proto_rawDescGZIP(), []int{2}
}

func (x *UUID) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type UIDS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uids []string `protobuf:"bytes,1,rep,name=uids,proto3" json:"uids,omitempty"`
}

func (x *UIDS) Reset() {
	*x = UIDS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UIDS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UIDS) ProtoMessage() {}

func (x *UIDS) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UIDS.ProtoReflect.Descriptor instead.
func (*UIDS) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{3}
}

func (x *UIDS) GetUids() []string {
	if x != nil {
		return x.Uids
	}
	return nil
}

type AccPwd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // account unique id
	Pwd  string `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`   // account login password enripted by RSA + Base64
}

func (x *AccPwd) Reset() {
	*x = AccPwd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccPwd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccPwd) ProtoMessage() {}

func (x *AccPwd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccPwd.ProtoReflect.Descriptor instead.
func (*AccPwd) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{4}
}

func (x *AccPwd) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AccPwd) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`          // account unique id
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`  // account nickname
	Heardurl string `protobuf:"bytes,3,opt,name=heardurl,proto3" json:"heardurl,omitempty"`  // account heardurl, the value set from wechat avatar or manual update by userself
	Sex      int64  `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`           // account sex, 0:none, 1:male, 2:female, 3:neutral
	Birthday string `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`  // account birthday, the value set from real-name authentication
	Isverify bool   `protobuf:"varint,6,opt,name=isverify,proto3" json:"isverify,omitempty"` // status to indicate account if pass real-name autnenticated
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{5}
}

func (x *Profile) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Profile) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Profile) GetHeardurl() string {
	if x != nil {
		return x.Heardurl
	}
	return ""
}

func (x *Profile) GetSex() int64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *Profile) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *Profile) GetIsverify() bool {
	if x != nil {
		return x.Isverify
	}
	return false
}

type ProfSumm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`         // account unique id
	Sex      int64  `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`          // account sex, 0:none, 1:male, 2:female, 3:neutral
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"` // account nickname
	Heardurl string `protobuf:"bytes,4,opt,name=heardurl,proto3" json:"heardurl,omitempty"` // account heardurl, the value set from wechat avatar
}

func (x *ProfSumm) Reset() {
	*x = ProfSumm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfSumm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfSumm) ProtoMessage() {}

func (x *ProfSumm) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfSumm.ProtoReflect.Descriptor instead.
func (*ProfSumm) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{6}
}

func (x *ProfSumm) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ProfSumm) GetSex() int64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *ProfSumm) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ProfSumm) GetHeardurl() string {
	if x != nil {
		return x.Heardurl
	}
	return ""
}

type ProfSumms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profsumms []*ProfSumm `protobuf:"bytes,1,rep,name=profsumms,proto3" json:"profsumms,omitempty"`
}

func (x *ProfSumms) Reset() {
	*x = ProfSumms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfSumms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfSumms) ProtoMessage() {}

func (x *ProfSumms) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfSumms.ProtoReflect.Descriptor instead.
func (*ProfSumms) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{7}
}

func (x *ProfSumms) GetProfsumms() []*ProfSumm {
	if x != nil {
		return x.Profsumms
	}
	return nil
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`       // account unique id
	Contact string `protobuf:"bytes,2,opt,name=contact,proto3" json:"contact,omitempty"` // account nickname
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`     // account email address
	Phone   string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`     // account email phone
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{8}
}

func (x *Contact) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Contact) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Contact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ProfStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logourl  string `protobuf:"bytes,1,opt,name=logourl,proto3" json:"logourl,omitempty"`   // store logo image url from heardurl field value of database
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"` // store name frist perfix string
	Unionid  string `protobuf:"bytes,5,opt,name=unionid,proto3" json:"unionid,omitempty"`   // wechat unionid bind with store account
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`       // store account email address, set when account generate
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`       // store account phone number, set when account generate
}

func (x *ProfStore) Reset() {
	*x = ProfStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfStore) ProtoMessage() {}

func (x *ProfStore) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfStore.ProtoReflect.Descriptor instead.
func (*ProfStore) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{9}
}

func (x *ProfStore) GetLogourl() string {
	if x != nil {
		return x.Logourl
	}
	return ""
}

func (x *ProfStore) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ProfStore) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

func (x *ProfStore) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ProfStore) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Secures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`       // account unique id
	Unionid string `protobuf:"bytes,2,opt,name=unionid,proto3" json:"unionid,omitempty"` // wechat unionid bind with account
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`     // account email address
	Pwd     string `protobuf:"bytes,4,opt,name=pwd,proto3" json:"pwd,omitempty"`         // account login password enripted by RSA + Base64
}

func (x *Secures) Reset() {
	*x = Secures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secures) ProtoMessage() {}

func (x *Secures) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secures.ProtoReflect.Descriptor instead.
func (*Secures) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{10}
}

func (x *Secures) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Secures) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

func (x *Secures) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Secures) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`       // account unique id
	Md5Pwd  string `protobuf:"bytes,2,opt,name=md5pwd,proto3" json:"md5pwd,omitempty"`   // admin account password encode as MD5
	Unionid string `protobuf:"bytes,3,opt,name=unionid,proto3" json:"unionid,omitempty"` // wechat unionid bind with account
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_proto_acc_proto_rawDescGZIP(), []int{11}
}

func (x *Admin) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Admin) GetMd5Pwd() string {
	if x != nil {
		return x.Md5Pwd
	}
	return ""
}

func (x *Admin) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

var File_proto_acc_proto protoreflect.FileDescriptor

var file_proto_acc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x1a, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x04,
	0x55, 0x49, 0x44, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x69, 0x64, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x50,
	0x77, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x72, 0x64, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x72, 0x64, 0x75, 0x72, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x73, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0x68, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x66, 0x53, 0x75, 0x6d, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x72,
	0x64, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x72,
	0x64, 0x75, 0x72, 0x6c, 0x22, 0x3a, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x53, 0x75, 0x6d, 0x6d,
	0x73, 0x12, 0x2d, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x73, 0x75, 0x6d, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x53, 0x75, 0x6d, 0x6d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x73, 0x75, 0x6d, 0x6d, 0x73,
	0x22, 0x63, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69,
	0x6f, 0x6e, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f,
	0x6e, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0x5f, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64,
	0x22, 0x4d, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x64, 0x35, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x64, 0x35, 0x70, 0x77, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x32,
	0x8c, 0x03, 0x0a, 0x03, 0x41, 0x63, 0x63, 0x12, 0x27, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x50,
	0x77, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x27, 0x0a, 0x08, 0x56, 0x69, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x50, 0x77, 0x64, 0x12, 0x26, 0x0a, 0x08, 0x56, 0x69, 0x61,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x29, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x53, 0x75, 0x6d, 0x6d, 0x73, 0x12, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x49, 0x44, 0x53, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x53, 0x75, 0x6d, 0x6d, 0x73, 0x12, 0x29, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x2b, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x73,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_acc_proto_rawDescOnce sync.Once
	file_proto_acc_proto_rawDescData = file_proto_acc_proto_rawDesc
)

func file_proto_acc_proto_rawDescGZIP() []byte {
	file_proto_acc_proto_rawDescOnce.Do(func() {
		file_proto_acc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_acc_proto_rawDescData)
	})
	return file_proto_acc_proto_rawDescData
}

var file_proto_acc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_acc_proto_goTypes = []interface{}{
	(*Empty)(nil),     // 0: proto.Empty
	(*Token)(nil),     // 1: proto.Token
	(*UUID)(nil),      // 2: proto.UUID
	(*UIDS)(nil),      // 3: proto.UIDS
	(*AccPwd)(nil),    // 4: proto.AccPwd
	(*Profile)(nil),   // 5: proto.Profile
	(*ProfSumm)(nil),  // 6: proto.ProfSumm
	(*ProfSumms)(nil), // 7: proto.ProfSumms
	(*Contact)(nil),   // 8: proto.Contact
	(*ProfStore)(nil), // 9: proto.ProfStore
	(*Secures)(nil),   // 10: proto.Secures
	(*Admin)(nil),     // 11: proto.Admin
}
var file_proto_acc_proto_depIdxs = []int32{
	6,  // 0: proto.ProfSumms.profsumms:type_name -> proto.ProfSumm
	4,  // 1: proto.Acc.AccLogin:input_type -> proto.AccPwd
	1,  // 2: proto.Acc.ViaToken:input_type -> proto.Token
	11, // 3: proto.Acc.ViaAdmin:input_type -> proto.Admin
	2,  // 4: proto.Acc.GetProfile:input_type -> proto.UUID
	3,  // 5: proto.Acc.GetProfSumms:input_type -> proto.UIDS
	2,  // 6: proto.Acc.GetContact:input_type -> proto.UUID
	2,  // 7: proto.Acc.StoreProfile:input_type -> proto.UUID
	8,  // 8: proto.Acc.SetContact:input_type -> proto.Contact
	10, // 9: proto.Acc.BindAccount:input_type -> proto.Secures
	1,  // 10: proto.Acc.AccLogin:output_type -> proto.Token
	4,  // 11: proto.Acc.ViaToken:output_type -> proto.AccPwd
	0,  // 12: proto.Acc.ViaAdmin:output_type -> proto.Empty
	5,  // 13: proto.Acc.GetProfile:output_type -> proto.Profile
	7,  // 14: proto.Acc.GetProfSumms:output_type -> proto.ProfSumms
	8,  // 15: proto.Acc.GetContact:output_type -> proto.Contact
	9,  // 16: proto.Acc.StoreProfile:output_type -> proto.ProfStore
	0,  // 17: proto.Acc.SetContact:output_type -> proto.Empty
	1,  // 18: proto.Acc.BindAccount:output_type -> proto.Token
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_acc_proto_init() }
func file_proto_acc_proto_init() {
	if File_proto_acc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_acc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_acc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_proto_acc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_acc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UIDS); i {
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
		file_proto_acc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccPwd); i {
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
		file_proto_acc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_proto_acc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfSumm); i {
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
		file_proto_acc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfSumms); i {
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
		file_proto_acc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_proto_acc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfStore); i {
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
		file_proto_acc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secures); i {
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
		file_proto_acc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
			RawDescriptor: file_proto_acc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_acc_proto_goTypes,
		DependencyIndexes: file_proto_acc_proto_depIdxs,
		MessageInfos:      file_proto_acc_proto_msgTypes,
	}.Build()
	File_proto_acc_proto = out.File
	file_proto_acc_proto_rawDesc = nil
	file_proto_acc_proto_goTypes = nil
	file_proto_acc_proto_depIdxs = nil
}
