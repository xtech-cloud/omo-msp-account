// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/shared.proto

package omo_msa_account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 策略
type Strategy int32

const (
	Strategy_STRATEGY_NONE Strategy = 0
	Strategy_STRATEGY_JWT  Strategy = 1
)

var Strategy_name = map[int32]string{
	0: "STRATEGY_NONE",
	1: "STRATEGY_JWT",
}

var Strategy_value = map[string]int32{
	"STRATEGY_NONE": 0,
	"STRATEGY_JWT":  1,
}

func (x Strategy) String() string {
	return proto.EnumName(Strategy_name, int32(x))
}

func (Strategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49158a9607562c1b, []int{0}
}

// 查询字段
type QueryField int32

const (
	QueryField_QUERY_FIELD_NONE     QueryField = 0
	QueryField_QUERY_FIELD_UUID     QueryField = 1
	QueryField_QUERY_FIELD_USERNAME QueryField = 2
)

var QueryField_name = map[int32]string{
	0: "QUERY_FIELD_NONE",
	1: "QUERY_FIELD_UUID",
	2: "QUERY_FIELD_USERNAME",
}

var QueryField_value = map[string]int32{
	"QUERY_FIELD_NONE":     0,
	"QUERY_FIELD_UUID":     1,
	"QUERY_FIELD_USERNAME": 2,
}

func (x QueryField) String() string {
	return proto.EnumName(QueryField_name, int32(x))
}

func (QueryField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49158a9607562c1b, []int{1}
}

// 状态
type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_49158a9607562c1b, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 账号实体
type AccountEntity struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Profile              string   `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	CreatedAt            int64    `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountEntity) Reset()         { *m = AccountEntity{} }
func (m *AccountEntity) String() string { return proto.CompactTextString(m) }
func (*AccountEntity) ProtoMessage()    {}
func (*AccountEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_49158a9607562c1b, []int{1}
}

func (m *AccountEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEntity.Unmarshal(m, b)
}
func (m *AccountEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEntity.Marshal(b, m, deterministic)
}
func (m *AccountEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEntity.Merge(m, src)
}
func (m *AccountEntity) XXX_Size() int {
	return xxx_messageInfo_AccountEntity.Size(m)
}
func (m *AccountEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEntity.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEntity proto.InternalMessageInfo

func (m *AccountEntity) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *AccountEntity) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *AccountEntity) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *AccountEntity) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("omo.msa.account.Strategy", Strategy_name, Strategy_value)
	proto.RegisterEnum("omo.msa.account.QueryField", QueryField_name, QueryField_value)
	proto.RegisterType((*Status)(nil), "omo.msa.account.Status")
	proto.RegisterType((*AccountEntity)(nil), "omo.msa.account.AccountEntity")
}

func init() {
	proto.RegisterFile("proto/account/shared.proto", fileDescriptor_49158a9607562c1b)
}

var fileDescriptor_49158a9607562c1b = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcb, 0x6a, 0x02, 0x31,
	0x18, 0x46, 0x1b, 0x6f, 0xd5, 0x9f, 0x4a, 0xd3, 0xe0, 0x22, 0x48, 0x17, 0xe2, 0x4a, 0x5c, 0xe8,
	0xa2, 0xd0, 0xfd, 0x80, 0xb1, 0x58, 0x5a, 0x8b, 0x19, 0xa5, 0xb8, 0x92, 0x74, 0x92, 0xda, 0x01,
	0xc7, 0x0c, 0xb9, 0x2c, 0xe6, 0x41, 0xfa, 0xbe, 0x65, 0x2e, 0x4e, 0x69, 0x77, 0xff, 0x77, 0x0e,
	0x39, 0x8b, 0xc0, 0x30, 0x35, 0xda, 0xe9, 0xb9, 0x88, 0x22, 0xed, 0xcf, 0x6e, 0x6e, 0xbf, 0x84,
	0x51, 0x72, 0x56, 0x40, 0x72, 0xab, 0x13, 0x3d, 0x4b, 0xac, 0x98, 0x55, 0x76, 0xfc, 0x08, 0x9d,
	0xd0, 0x09, 0xe7, 0x2d, 0x21, 0xd0, 0x8a, 0xb4, 0x54, 0x14, 0x8d, 0xd0, 0xa4, 0xcd, 0x8b, 0x9b,
	0x50, 0xb8, 0x4e, 0x94, 0xb5, 0xe2, 0xa8, 0x68, 0x63, 0x84, 0x26, 0x3d, 0x7e, 0x99, 0xe3, 0x6f,
	0x04, 0xfd, 0xa0, 0x6c, 0xb0, 0xb3, 0x8b, 0x5d, 0x46, 0x86, 0xd0, 0xf5, 0x56, 0x99, 0xb3, 0x48,
	0xca, 0x46, 0x8f, 0xd7, 0x3b, 0x6f, 0x7b, 0x1f, 0xcb, 0x2a, 0x52, 0xdc, 0x79, 0x3b, 0x35, 0xfa,
	0x33, 0x3e, 0x29, 0xda, 0x2c, 0xdb, 0xd5, 0x24, 0xf7, 0xd0, 0x8b, 0x8c, 0x12, 0x4e, 0xc9, 0xc0,
	0xd1, 0xd6, 0x08, 0x4d, 0x9a, 0xfc, 0x17, 0xe4, 0xd6, 0xa7, 0xb2, 0xb2, 0xed, 0xd2, 0xd6, 0x60,
	0x3a, 0x87, 0x6e, 0xe8, 0x8c, 0x70, 0xea, 0x98, 0x91, 0x3b, 0xe8, 0x87, 0x5b, 0x1e, 0x6c, 0xd9,
	0xd3, 0xfe, 0xb0, 0x7e, 0x5b, 0x33, 0x7c, 0x45, 0x30, 0xdc, 0xd4, 0xe8, 0xf9, 0x7d, 0x8b, 0xd1,
	0x94, 0x03, 0x6c, 0xbc, 0x32, 0xd9, 0x32, 0x56, 0x27, 0x49, 0x06, 0x80, 0x37, 0x3b, 0xc6, 0xf7,
	0x87, 0xe5, 0x8a, 0xbd, 0x2c, 0x2e, 0xaf, 0xfe, 0xd1, 0xdd, 0x6e, 0xb5, 0xc0, 0x88, 0x50, 0x18,
	0xfc, 0xa1, 0x21, 0xe3, 0xeb, 0xe0, 0x95, 0xe1, 0xc6, 0x47, 0xa7, 0xf8, 0xec, 0x87, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0xae, 0x5d, 0x79, 0x8a, 0x01, 0x00, 0x00,
}
