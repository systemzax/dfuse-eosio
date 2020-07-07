// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/eosio/trxdb/v1/trxdb.proto

package pbtrxdb

import (
	fmt "fmt"
	v1 "github.com/dfuse-io/dfuse-eosio/pb/dfuse/eosio/codec/v1"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// IndexableCategory represents the valid catgory that can be read/write by
// the TrxDB instance.
//
// Those enum are used mainly throughout the code to determine what needs to
// be written to the database as well as where to read each component.
type IndexableCategory int32

const (
	IndexableCategory_INDEXABLE_CATEGORY_ACCOUNT     IndexableCategory = 0
	IndexableCategory_INDEXABLE_CATEGORY_BLOCK       IndexableCategory = 1
	IndexableCategory_INDEXABLE_CATEGORY_TRANSACTION IndexableCategory = 2
	IndexableCategory_INDEXABLE_CATEGORY_TIMELINE    IndexableCategory = 3
)

var IndexableCategory_name = map[int32]string{
	0: "INDEXABLE_CATEGORY_ACCOUNT",
	1: "INDEXABLE_CATEGORY_BLOCK",
	2: "INDEXABLE_CATEGORY_TRANSACTION",
	3: "INDEXABLE_CATEGORY_TIMELINE",
}

var IndexableCategory_value = map[string]int32{
	"INDEXABLE_CATEGORY_ACCOUNT":     0,
	"INDEXABLE_CATEGORY_BLOCK":       1,
	"INDEXABLE_CATEGORY_TRANSACTION": 2,
	"INDEXABLE_CATEGORY_TIMELINE":    3,
}

func (x IndexableCategory) String() string {
	return proto.EnumName(IndexableCategory_name, int32(x))
}

func (IndexableCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_014579ff7d302fea, []int{0}
}

type ImplicitTrxRow struct {
	// Primary key to include: transaction_id and block_id
	SignedTrx            *v1.SignedTransaction `protobuf:"bytes,1,opt,name=signed_trx,json=signedTrx,proto3" json:"signed_trx,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ImplicitTrxRow) Reset()         { *m = ImplicitTrxRow{} }
func (m *ImplicitTrxRow) String() string { return proto.CompactTextString(m) }
func (*ImplicitTrxRow) ProtoMessage()    {}
func (*ImplicitTrxRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_014579ff7d302fea, []int{0}
}

func (m *ImplicitTrxRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImplicitTrxRow.Unmarshal(m, b)
}
func (m *ImplicitTrxRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImplicitTrxRow.Marshal(b, m, deterministic)
}
func (m *ImplicitTrxRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImplicitTrxRow.Merge(m, src)
}
func (m *ImplicitTrxRow) XXX_Size() int {
	return xxx_messageInfo_ImplicitTrxRow.Size(m)
}
func (m *ImplicitTrxRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ImplicitTrxRow.DiscardUnknown(m)
}

var xxx_messageInfo_ImplicitTrxRow proto.InternalMessageInfo

func (m *ImplicitTrxRow) GetSignedTrx() *v1.SignedTransaction {
	if m != nil {
		return m.SignedTrx
	}
	return nil
}

func (m *ImplicitTrxRow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TrxRow struct {
	SignedTrx            *v1.SignedTransaction  `protobuf:"bytes,1,opt,name=signed_trx,json=signedTrx,proto3" json:"signed_trx,omitempty"`
	Receipt              *v1.TransactionReceipt `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	PublicKeys           *v1.PublicKeys         `protobuf:"bytes,3,opt,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TrxRow) Reset()         { *m = TrxRow{} }
func (m *TrxRow) String() string { return proto.CompactTextString(m) }
func (*TrxRow) ProtoMessage()    {}
func (*TrxRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_014579ff7d302fea, []int{1}
}

func (m *TrxRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrxRow.Unmarshal(m, b)
}
func (m *TrxRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrxRow.Marshal(b, m, deterministic)
}
func (m *TrxRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrxRow.Merge(m, src)
}
func (m *TrxRow) XXX_Size() int {
	return xxx_messageInfo_TrxRow.Size(m)
}
func (m *TrxRow) XXX_DiscardUnknown() {
	xxx_messageInfo_TrxRow.DiscardUnknown(m)
}

var xxx_messageInfo_TrxRow proto.InternalMessageInfo

func (m *TrxRow) GetSignedTrx() *v1.SignedTransaction {
	if m != nil {
		return m.SignedTrx
	}
	return nil
}

func (m *TrxRow) GetReceipt() *v1.TransactionReceipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *TrxRow) GetPublicKeys() *v1.PublicKeys {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type TrxTraceRow struct {
	TrxTrace             *v1.TransactionTrace `protobuf:"bytes,1,opt,name=trx_trace,json=trxTrace,proto3" json:"trx_trace,omitempty"`
	BlockHeader          *v1.BlockHeader      `protobuf:"bytes,2,opt,name=block_header,json=blockHeader,proto3" json:"block_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TrxTraceRow) Reset()         { *m = TrxTraceRow{} }
func (m *TrxTraceRow) String() string { return proto.CompactTextString(m) }
func (*TrxTraceRow) ProtoMessage()    {}
func (*TrxTraceRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_014579ff7d302fea, []int{2}
}

func (m *TrxTraceRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrxTraceRow.Unmarshal(m, b)
}
func (m *TrxTraceRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrxTraceRow.Marshal(b, m, deterministic)
}
func (m *TrxTraceRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrxTraceRow.Merge(m, src)
}
func (m *TrxTraceRow) XXX_Size() int {
	return xxx_messageInfo_TrxTraceRow.Size(m)
}
func (m *TrxTraceRow) XXX_DiscardUnknown() {
	xxx_messageInfo_TrxTraceRow.DiscardUnknown(m)
}

var xxx_messageInfo_TrxTraceRow proto.InternalMessageInfo

func (m *TrxTraceRow) GetTrxTrace() *v1.TransactionTrace {
	if m != nil {
		return m.TrxTrace
	}
	return nil
}

func (m *TrxTraceRow) GetBlockHeader() *v1.BlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

type DtrxRow struct {
	SignedTrx            *v1.SignedTransaction `protobuf:"bytes,1,opt,name=signed_trx,json=signedTrx,proto3" json:"signed_trx,omitempty"`
	CreatedBy            *v1.ExtDTrxOp         `protobuf:"bytes,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CanceledBy           *v1.ExtDTrxOp         `protobuf:"bytes,3,opt,name=canceled_by,json=canceledBy,proto3" json:"canceled_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DtrxRow) Reset()         { *m = DtrxRow{} }
func (m *DtrxRow) String() string { return proto.CompactTextString(m) }
func (*DtrxRow) ProtoMessage()    {}
func (*DtrxRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_014579ff7d302fea, []int{3}
}

func (m *DtrxRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DtrxRow.Unmarshal(m, b)
}
func (m *DtrxRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DtrxRow.Marshal(b, m, deterministic)
}
func (m *DtrxRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DtrxRow.Merge(m, src)
}
func (m *DtrxRow) XXX_Size() int {
	return xxx_messageInfo_DtrxRow.Size(m)
}
func (m *DtrxRow) XXX_DiscardUnknown() {
	xxx_messageInfo_DtrxRow.DiscardUnknown(m)
}

var xxx_messageInfo_DtrxRow proto.InternalMessageInfo

func (m *DtrxRow) GetSignedTrx() *v1.SignedTransaction {
	if m != nil {
		return m.SignedTrx
	}
	return nil
}

func (m *DtrxRow) GetCreatedBy() *v1.ExtDTrxOp {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *DtrxRow) GetCanceledBy() *v1.ExtDTrxOp {
	if m != nil {
		return m.CanceledBy
	}
	return nil
}

type BlockRow struct {
	Block                *v1.Block           `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	ImplicitTrxRefs      *v1.TransactionRefs `protobuf:"bytes,2,opt,name=implicit_trx_refs,json=implicitTrxRefs,proto3" json:"implicit_trx_refs,omitempty"`
	TrxRefs              *v1.TransactionRefs `protobuf:"bytes,3,opt,name=trx_refs,json=trxRefs,proto3" json:"trx_refs,omitempty"`
	TraceRefs            *v1.TransactionRefs `protobuf:"bytes,4,opt,name=trace_refs,json=traceRefs,proto3" json:"trace_refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BlockRow) Reset()         { *m = BlockRow{} }
func (m *BlockRow) String() string { return proto.CompactTextString(m) }
func (*BlockRow) ProtoMessage()    {}
func (*BlockRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_014579ff7d302fea, []int{4}
}

func (m *BlockRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRow.Unmarshal(m, b)
}
func (m *BlockRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRow.Marshal(b, m, deterministic)
}
func (m *BlockRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRow.Merge(m, src)
}
func (m *BlockRow) XXX_Size() int {
	return xxx_messageInfo_BlockRow.Size(m)
}
func (m *BlockRow) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRow.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRow proto.InternalMessageInfo

func (m *BlockRow) GetBlock() *v1.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockRow) GetImplicitTrxRefs() *v1.TransactionRefs {
	if m != nil {
		return m.ImplicitTrxRefs
	}
	return nil
}

func (m *BlockRow) GetTrxRefs() *v1.TransactionRefs {
	if m != nil {
		return m.TrxRefs
	}
	return nil
}

func (m *BlockRow) GetTraceRefs() *v1.TransactionRefs {
	if m != nil {
		return m.TraceRefs
	}
	return nil
}

type AccountRow struct {
	// Key will be account:block_id    // block_id, block_num and account implicit in key
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Creator              string               `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	BlockId              string               `protobuf:"bytes,3,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	TrxId                string               `protobuf:"bytes,4,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	BlockTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AccountRow) Reset()         { *m = AccountRow{} }
func (m *AccountRow) String() string { return proto.CompactTextString(m) }
func (*AccountRow) ProtoMessage()    {}
func (*AccountRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_014579ff7d302fea, []int{5}
}

func (m *AccountRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRow.Unmarshal(m, b)
}
func (m *AccountRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRow.Marshal(b, m, deterministic)
}
func (m *AccountRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRow.Merge(m, src)
}
func (m *AccountRow) XXX_Size() int {
	return xxx_messageInfo_AccountRow.Size(m)
}
func (m *AccountRow) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRow.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRow proto.InternalMessageInfo

func (m *AccountRow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountRow) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AccountRow) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func (m *AccountRow) GetTrxId() string {
	if m != nil {
		return m.TrxId
	}
	return ""
}

func (m *AccountRow) GetBlockTime() *timestamp.Timestamp {
	if m != nil {
		return m.BlockTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("dfuse.eosio.trxdb.v1.IndexableCategory", IndexableCategory_name, IndexableCategory_value)
	proto.RegisterType((*ImplicitTrxRow)(nil), "dfuse.eosio.trxdb.v1.ImplicitTrxRow")
	proto.RegisterType((*TrxRow)(nil), "dfuse.eosio.trxdb.v1.TrxRow")
	proto.RegisterType((*TrxTraceRow)(nil), "dfuse.eosio.trxdb.v1.TrxTraceRow")
	proto.RegisterType((*DtrxRow)(nil), "dfuse.eosio.trxdb.v1.DtrxRow")
	proto.RegisterType((*BlockRow)(nil), "dfuse.eosio.trxdb.v1.BlockRow")
	proto.RegisterType((*AccountRow)(nil), "dfuse.eosio.trxdb.v1.AccountRow")
}

func init() { proto.RegisterFile("dfuse/eosio/trxdb/v1/trxdb.proto", fileDescriptor_014579ff7d302fea) }

var fileDescriptor_014579ff7d302fea = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x6e, 0xda, 0x26, 0x9e, 0x1c, 0x9d, 0x93, 0xae, 0x40, 0x0a, 0x29, 0x6a, 0x43, 0x24,
	0xa0, 0x42, 0xc2, 0x51, 0xca, 0x15, 0x42, 0x82, 0x3a, 0x3f, 0x80, 0xd5, 0x92, 0xc0, 0xd6, 0x48,
	0xc0, 0x8d, 0x65, 0xaf, 0x37, 0xe9, 0xaa, 0x8e, 0xd7, 0x5a, 0x6f, 0x8a, 0xf3, 0x22, 0xa8, 0x6f,
	0xc0, 0x03, 0x71, 0xc1, 0xeb, 0x20, 0xef, 0x3a, 0x6d, 0x91, 0x2c, 0x68, 0xa5, 0xde, 0xcd, 0x78,
	0xbf, 0xef, 0xdb, 0xfd, 0xc6, 0x33, 0x03, 0xed, 0x70, 0xba, 0x48, 0x69, 0x97, 0xf2, 0x94, 0xf1,
	0xae, 0x14, 0x59, 0x18, 0x74, 0xcf, 0x7a, 0x3a, 0xb0, 0x12, 0xc1, 0x25, 0x47, 0x77, 0x14, 0xc2,
	0x52, 0x08, 0x4b, 0x1f, 0x9c, 0xf5, 0x5a, 0xbf, 0xf1, 0x08, 0x0f, 0x29, 0xc9, 0x79, 0x2a, 0xd0,
	0xbc, 0xd6, 0xee, 0x8c, 0xf3, 0x59, 0x44, 0xbb, 0x2a, 0x0b, 0x16, 0xd3, 0xae, 0x64, 0x73, 0x9a,
	0x4a, 0x7f, 0x9e, 0x68, 0x40, 0x27, 0x82, 0xff, 0x9c, 0x79, 0x12, 0x31, 0xc2, 0xa4, 0x2b, 0x32,
	0xcc, 0xbf, 0xa2, 0xd7, 0x00, 0x29, 0x9b, 0xc5, 0x34, 0xf4, 0xa4, 0xc8, 0x9a, 0x46, 0xdb, 0xd8,
	0xab, 0xef, 0x3f, 0xb6, 0xae, 0xde, 0xaf, 0x2f, 0x38, 0xeb, 0x59, 0xc7, 0x0a, 0xe7, 0x0a, 0x3f,
	0x4e, 0x7d, 0x22, 0x19, 0x8f, 0xb1, 0x99, 0x16, 0x9f, 0x32, 0x84, 0x60, 0x3d, 0xf6, 0xe7, 0xb4,
	0xb9, 0xd6, 0x36, 0xf6, 0x4c, 0xac, 0xe2, 0xce, 0x4f, 0x03, 0x36, 0x6f, 0xf9, 0x9a, 0x3e, 0x54,
	0x05, 0x25, 0x94, 0x25, 0x52, 0xdd, 0x54, 0xdf, 0xdf, 0x2b, 0x17, 0xb9, 0x4a, 0xd7, 0x78, 0xbc,
	0x22, 0x22, 0x1b, 0xea, 0xc9, 0x22, 0x88, 0x18, 0xf1, 0x4e, 0xe9, 0x32, 0x6d, 0x56, 0x94, 0x4e,
	0xbb, 0x5c, 0xe7, 0xbd, 0x02, 0x1e, 0xd2, 0x65, 0x8a, 0x21, 0xb9, 0x88, 0x3b, 0xe7, 0x06, 0xd4,
	0x5d, 0x91, 0xb9, 0xc2, 0x27, 0x34, 0xb7, 0x37, 0x00, 0x53, 0x8a, 0xcc, 0x93, 0x79, 0x5e, 0xb8,
	0x7b, 0xf4, 0xd7, 0x87, 0x69, 0x76, 0x4d, 0x16, 0x3a, 0x68, 0x08, 0xff, 0x06, 0x11, 0x27, 0xa7,
	0xde, 0x09, 0xf5, 0x43, 0x2a, 0x0a, 0x83, 0x0f, 0xca, 0x75, 0xfa, 0x39, 0xf2, 0xad, 0x02, 0xe2,
	0x7a, 0x70, 0x99, 0x74, 0x7e, 0x18, 0x50, 0x1d, 0xca, 0xdb, 0xad, 0xfa, 0x4b, 0x00, 0x22, 0xa8,
	0x2f, 0x69, 0xe8, 0x05, 0xcb, 0xe2, 0x5d, 0xbb, 0xe5, 0x3a, 0xa3, 0x4c, 0x0e, 0x5d, 0x91, 0x4d,
	0x12, 0x6c, 0x16, 0x94, 0xfe, 0x12, 0x1d, 0x40, 0x9d, 0xf8, 0x31, 0xa1, 0x91, 0x16, 0xa8, 0x5c,
	0x4f, 0x00, 0x56, 0x9c, 0xfe, 0xb2, 0x73, 0xbe, 0x06, 0x35, 0x65, 0x39, 0xb7, 0xd5, 0x83, 0x0d,
	0xe5, 0xb8, 0x70, 0xb4, 0xfd, 0x87, 0x0a, 0x61, 0x8d, 0x44, 0x1f, 0x60, 0x8b, 0x15, 0x8d, 0x9f,
	0xd7, 0xc2, 0x13, 0x74, 0x9a, 0x16, 0x46, 0x1e, 0x5e, 0xa3, 0x83, 0xa6, 0x29, 0xfe, 0x9f, 0x5d,
	0x19, 0x1c, 0x3a, 0x4d, 0xd1, 0x01, 0xd4, 0x2e, 0x94, 0x2a, 0x37, 0x51, 0xaa, 0xca, 0x42, 0x61,
	0x08, 0xa0, 0x3a, 0x46, 0x6b, 0xac, 0xdf, 0x44, 0xc3, 0x54, 0xc4, 0x3c, 0xec, 0x7c, 0x37, 0x00,
	0x6c, 0x42, 0xf8, 0x22, 0x96, 0x79, 0x71, 0x56, 0x83, 0x68, 0x5c, 0x0e, 0x22, 0x6a, 0x42, 0x55,
	0xfd, 0x0c, 0x2e, 0x8a, 0xf9, 0x5c, 0xa5, 0xe8, 0x1e, 0xd4, 0x74, 0xcf, 0xb1, 0x50, 0x99, 0x30,
	0x71, 0x55, 0xe5, 0x4e, 0x88, 0xee, 0xc2, 0x66, 0xee, 0x8f, 0x85, 0xea, 0x65, 0x26, 0xde, 0x90,
	0x22, 0x73, 0x42, 0xf4, 0x1c, 0x40, 0x33, 0xf2, 0xdd, 0xd2, 0xdc, 0x50, 0x8f, 0x6e, 0x59, 0x7a,
	0xf1, 0x58, 0xab, 0xc5, 0x63, 0xb9, 0xab, 0xc5, 0x83, 0x4d, 0x85, 0xce, 0xf3, 0x27, 0xdf, 0x0c,
	0xd8, 0x72, 0xe2, 0x90, 0x66, 0x7e, 0x10, 0xd1, 0x81, 0x2f, 0xe9, 0x8c, 0x8b, 0x25, 0xda, 0x81,
	0x96, 0x33, 0x1e, 0x8e, 0x3e, 0xd9, 0xfd, 0xa3, 0x91, 0x37, 0xb0, 0xdd, 0xd1, 0x9b, 0x09, 0xfe,
	0xec, 0xd9, 0x83, 0xc1, 0xe4, 0xe3, 0xd8, 0x6d, 0xfc, 0x83, 0xee, 0x43, 0xb3, 0xe4, 0xbc, 0x7f,
	0x34, 0x19, 0x1c, 0x36, 0x0c, 0xd4, 0x81, 0x9d, 0x92, 0x53, 0x17, 0xdb, 0xe3, 0x63, 0x7b, 0xe0,
	0x3a, 0x93, 0x71, 0x63, 0x0d, 0xed, 0xc2, 0x76, 0x19, 0xc6, 0x79, 0x37, 0x3a, 0x72, 0xc6, 0xa3,
	0x46, 0xa5, 0x6f, 0x7f, 0x79, 0x35, 0x63, 0xf2, 0x64, 0x11, 0x58, 0x84, 0xcf, 0xbb, 0xea, 0x07,
	0x3c, 0x65, 0xbc, 0x08, 0xf4, 0xbe, 0x4d, 0x82, 0x6e, 0xd9, 0xda, 0x7e, 0x91, 0x04, 0x2a, 0x0c,
	0x36, 0x95, 0xf5, 0x67, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xdb, 0x2a, 0x36, 0xdd, 0x05,
	0x00, 0x00,
}
