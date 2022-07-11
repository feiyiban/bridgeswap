// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/contract/exchange_contract.proto

package core

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

type ExchangeCreateContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	FirstTokenId         []byte   `protobuf:"bytes,2,opt,name=first_token_id,json=firstTokenId,proto3" json:"first_token_id,omitempty"`
	FirstTokenBalance    int64    `protobuf:"varint,3,opt,name=first_token_balance,json=firstTokenBalance,proto3" json:"first_token_balance,omitempty"`
	SecondTokenId        []byte   `protobuf:"bytes,4,opt,name=second_token_id,json=secondTokenId,proto3" json:"second_token_id,omitempty"`
	SecondTokenBalance   int64    `protobuf:"varint,5,opt,name=second_token_balance,json=secondTokenBalance,proto3" json:"second_token_balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeCreateContract) Reset()         { *m = ExchangeCreateContract{} }
func (m *ExchangeCreateContract) String() string { return proto.CompactTextString(m) }
func (*ExchangeCreateContract) ProtoMessage()    {}
func (*ExchangeCreateContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c7f776bef66f61, []int{0}
}

func (m *ExchangeCreateContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeCreateContract.Unmarshal(m, b)
}
func (m *ExchangeCreateContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeCreateContract.Marshal(b, m, deterministic)
}
func (m *ExchangeCreateContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeCreateContract.Merge(m, src)
}
func (m *ExchangeCreateContract) XXX_Size() int {
	return xxx_messageInfo_ExchangeCreateContract.Size(m)
}
func (m *ExchangeCreateContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeCreateContract.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeCreateContract proto.InternalMessageInfo

func (m *ExchangeCreateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ExchangeCreateContract) GetFirstTokenId() []byte {
	if m != nil {
		return m.FirstTokenId
	}
	return nil
}

func (m *ExchangeCreateContract) GetFirstTokenBalance() int64 {
	if m != nil {
		return m.FirstTokenBalance
	}
	return 0
}

func (m *ExchangeCreateContract) GetSecondTokenId() []byte {
	if m != nil {
		return m.SecondTokenId
	}
	return nil
}

func (m *ExchangeCreateContract) GetSecondTokenBalance() int64 {
	if m != nil {
		return m.SecondTokenBalance
	}
	return 0
}

type ExchangeInjectContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ExchangeId           int64    `protobuf:"varint,2,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	TokenId              []byte   `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Quant                int64    `protobuf:"varint,4,opt,name=quant,proto3" json:"quant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeInjectContract) Reset()         { *m = ExchangeInjectContract{} }
func (m *ExchangeInjectContract) String() string { return proto.CompactTextString(m) }
func (*ExchangeInjectContract) ProtoMessage()    {}
func (*ExchangeInjectContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c7f776bef66f61, []int{1}
}

func (m *ExchangeInjectContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeInjectContract.Unmarshal(m, b)
}
func (m *ExchangeInjectContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeInjectContract.Marshal(b, m, deterministic)
}
func (m *ExchangeInjectContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeInjectContract.Merge(m, src)
}
func (m *ExchangeInjectContract) XXX_Size() int {
	return xxx_messageInfo_ExchangeInjectContract.Size(m)
}
func (m *ExchangeInjectContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeInjectContract.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeInjectContract proto.InternalMessageInfo

func (m *ExchangeInjectContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ExchangeInjectContract) GetExchangeId() int64 {
	if m != nil {
		return m.ExchangeId
	}
	return 0
}

func (m *ExchangeInjectContract) GetTokenId() []byte {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *ExchangeInjectContract) GetQuant() int64 {
	if m != nil {
		return m.Quant
	}
	return 0
}

type ExchangeWithdrawContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ExchangeId           int64    `protobuf:"varint,2,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	TokenId              []byte   `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Quant                int64    `protobuf:"varint,4,opt,name=quant,proto3" json:"quant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeWithdrawContract) Reset()         { *m = ExchangeWithdrawContract{} }
func (m *ExchangeWithdrawContract) String() string { return proto.CompactTextString(m) }
func (*ExchangeWithdrawContract) ProtoMessage()    {}
func (*ExchangeWithdrawContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c7f776bef66f61, []int{2}
}

func (m *ExchangeWithdrawContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeWithdrawContract.Unmarshal(m, b)
}
func (m *ExchangeWithdrawContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeWithdrawContract.Marshal(b, m, deterministic)
}
func (m *ExchangeWithdrawContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeWithdrawContract.Merge(m, src)
}
func (m *ExchangeWithdrawContract) XXX_Size() int {
	return xxx_messageInfo_ExchangeWithdrawContract.Size(m)
}
func (m *ExchangeWithdrawContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeWithdrawContract.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeWithdrawContract proto.InternalMessageInfo

func (m *ExchangeWithdrawContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ExchangeWithdrawContract) GetExchangeId() int64 {
	if m != nil {
		return m.ExchangeId
	}
	return 0
}

func (m *ExchangeWithdrawContract) GetTokenId() []byte {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *ExchangeWithdrawContract) GetQuant() int64 {
	if m != nil {
		return m.Quant
	}
	return 0
}

type ExchangeTransactionContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ExchangeId           int64    `protobuf:"varint,2,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	TokenId              []byte   `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Quant                int64    `protobuf:"varint,4,opt,name=quant,proto3" json:"quant,omitempty"`
	Expected             int64    `protobuf:"varint,5,opt,name=expected,proto3" json:"expected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeTransactionContract) Reset()         { *m = ExchangeTransactionContract{} }
func (m *ExchangeTransactionContract) String() string { return proto.CompactTextString(m) }
func (*ExchangeTransactionContract) ProtoMessage()    {}
func (*ExchangeTransactionContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c7f776bef66f61, []int{3}
}

func (m *ExchangeTransactionContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeTransactionContract.Unmarshal(m, b)
}
func (m *ExchangeTransactionContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeTransactionContract.Marshal(b, m, deterministic)
}
func (m *ExchangeTransactionContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeTransactionContract.Merge(m, src)
}
func (m *ExchangeTransactionContract) XXX_Size() int {
	return xxx_messageInfo_ExchangeTransactionContract.Size(m)
}
func (m *ExchangeTransactionContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeTransactionContract.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeTransactionContract proto.InternalMessageInfo

func (m *ExchangeTransactionContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ExchangeTransactionContract) GetExchangeId() int64 {
	if m != nil {
		return m.ExchangeId
	}
	return 0
}

func (m *ExchangeTransactionContract) GetTokenId() []byte {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *ExchangeTransactionContract) GetQuant() int64 {
	if m != nil {
		return m.Quant
	}
	return 0
}

func (m *ExchangeTransactionContract) GetExpected() int64 {
	if m != nil {
		return m.Expected
	}
	return 0
}

func init() {
	proto.RegisterType((*ExchangeCreateContract)(nil), "protocol.ExchangeCreateContract")
	proto.RegisterType((*ExchangeInjectContract)(nil), "protocol.ExchangeInjectContract")
	proto.RegisterType((*ExchangeWithdrawContract)(nil), "protocol.ExchangeWithdrawContract")
	proto.RegisterType((*ExchangeTransactionContract)(nil), "protocol.ExchangeTransactionContract")
}

func init() {
	proto.RegisterFile("core/contract/exchange_contract.proto", fileDescriptor_10c7f776bef66f61)
}

var fileDescriptor_10c7f776bef66f61 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0xd3, 0xaf, 0x1f, 0x4a, 0xae, 0xa0, 0x71, 0x24, 0xa6, 0xea, 0x42, 0xc4, 0x3f, 0x61,
	0xd5, 0x9a, 0xb8, 0x76, 0x21, 0xc4, 0x05, 0xdb, 0x86, 0xc4, 0xc4, 0x4d, 0x33, 0x9d, 0xb9, 0x94,
	0x2a, 0xce, 0xe0, 0xf4, 0x12, 0x78, 0x0b, 0xe3, 0x9b, 0xf8, 0x62, 0xbe, 0x83, 0x61, 0xda, 0x01,
	0x5c, 0xba, 0x62, 0xd5, 0xdc, 0xd3, 0x5f, 0x4f, 0xce, 0xe9, 0xdc, 0x81, 0x6b, 0xa1, 0x0d, 0x46,
	0x42, 0x2b, 0x32, 0x5c, 0x50, 0x84, 0x0b, 0x31, 0xe6, 0x2a, 0xc3, 0xc4, 0x29, 0xe1, 0xd4, 0x68,
	0xd2, 0xac, 0x6e, 0x1f, 0x42, 0x4f, 0x3a, 0xdf, 0x1e, 0x1c, 0x3f, 0x56, 0x54, 0xdf, 0x20, 0x27,
	0xec, 0x57, 0x28, 0xbb, 0x84, 0xa6, 0x9e, 0x2b, 0x34, 0x09, 0x97, 0xd2, 0x60, 0x51, 0x04, 0x5e,
	0xdb, 0xeb, 0x36, 0xe2, 0x86, 0x15, 0x1f, 0x4a, 0x8d, 0x5d, 0xc1, 0xfe, 0x28, 0x37, 0x05, 0x25,
	0xa4, 0x5f, 0x51, 0x25, 0xb9, 0x0c, 0xfe, 0x95, 0x94, 0x55, 0x87, 0x4b, 0x71, 0x20, 0x59, 0x08,
	0x47, 0x9b, 0x54, 0xca, 0x27, 0x5c, 0x09, 0x0c, 0xfc, 0xb6, 0xd7, 0xf5, 0xe3, 0xc3, 0x35, 0xda,
	0x2b, 0x5f, 0xb0, 0x1b, 0x38, 0x28, 0x50, 0x68, 0x25, 0xd7, 0xb6, 0xff, 0xad, 0x6d, 0xb3, 0x94,
	0x9d, 0xef, 0x2d, 0xb4, 0x7e, 0x71, 0xce, 0xb8, 0x66, 0x8d, 0xd9, 0x06, 0x5c, 0x39, 0x77, 0x3e,
	0x36, 0xfa, 0x0e, 0xd4, 0x0b, 0x0a, 0xfa, 0x5b, 0xdf, 0x73, 0xd8, 0x5b, 0xfd, 0xd4, 0xaa, 0xac,
	0x1f, 0x83, 0x93, 0x06, 0x92, 0x9d, 0x40, 0x7d, 0x95, 0xd9, 0xb7, 0x06, 0xbb, 0x54, 0xa5, 0x6d,
	0x41, 0xed, 0x7d, 0xc6, 0x15, 0xd9, 0x2e, 0x7e, 0x5c, 0x0e, 0x9d, 0x4f, 0x0f, 0x02, 0x97, 0xe8,
	0x29, 0xa7, 0xb1, 0x34, 0x7c, 0xbe, 0xed, 0x4c, 0x5f, 0x1e, 0x9c, 0xb9, 0x4c, 0x43, 0xc3, 0x55,
	0xc1, 0x05, 0xe5, 0x5a, 0x6d, 0x39, 0x16, 0x3b, 0x85, 0x3a, 0x2e, 0xa6, 0x28, 0x08, 0x65, 0x75,
	0xc4, 0xab, 0xb9, 0x77, 0x0f, 0x81, 0x36, 0x59, 0x48, 0x46, 0xab, 0x72, 0xc9, 0x8b, 0xd0, 0x2d,
	0xfd, 0xf3, 0x45, 0x96, 0xd3, 0x78, 0x96, 0x86, 0x42, 0xbf, 0x45, 0xa3, 0xb4, 0xd0, 0xa9, 0xc1,
	0xdc, 0xf0, 0x28, 0xd3, 0x4b, 0x3a, 0x5a, 0x5e, 0x99, 0x74, 0xc7, 0x7e, 0x73, 0xf7, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x52, 0x53, 0x7f, 0x44, 0x41, 0x03, 0x00, 0x00,
}
