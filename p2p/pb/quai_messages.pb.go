// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: p2p/pb/quai_messages.proto

package pb

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

type QuaiProtocolMessage_ActionType int32

const (
	QuaiProtocolMessage_REQUEST_BLOCK       QuaiProtocolMessage_ActionType = 0
	QuaiProtocolMessage_REQUEST_TRANSACTION QuaiProtocolMessage_ActionType = 1
)

// Enum value maps for QuaiProtocolMessage_ActionType.
var (
	QuaiProtocolMessage_ActionType_name = map[int32]string{
		0: "REQUEST_BLOCK",
		1: "REQUEST_TRANSACTION",
	}
	QuaiProtocolMessage_ActionType_value = map[string]int32{
		"REQUEST_BLOCK":       0,
		"REQUEST_TRANSACTION": 1,
	}
)

func (x QuaiProtocolMessage_ActionType) Enum() *QuaiProtocolMessage_ActionType {
	p := new(QuaiProtocolMessage_ActionType)
	*p = x
	return p
}

func (x QuaiProtocolMessage_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuaiProtocolMessage_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_p2p_pb_quai_messages_proto_enumTypes[0].Descriptor()
}

func (QuaiProtocolMessage_ActionType) Type() protoreflect.EnumType {
	return &file_p2p_pb_quai_messages_proto_enumTypes[0]
}

func (x QuaiProtocolMessage_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuaiProtocolMessage_ActionType.Descriptor instead.
func (QuaiProtocolMessage_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{2, 0}
}

// GossipSub messages for broadcasting blocks and transactions
type GossipBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GossipBlock) Reset() {
	*x = GossipBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipBlock) ProtoMessage() {}

func (x *GossipBlock) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipBlock.ProtoReflect.Descriptor instead.
func (*GossipBlock) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GossipBlock) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type GossipTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GossipTransaction) Reset() {
	*x = GossipTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipTransaction) ProtoMessage() {}

func (x *GossipTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipTransaction.ProtoReflect.Descriptor instead.
func (*GossipTransaction) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GossipTransaction) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// Stream-based request-response messages
type QuaiProtocolMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action QuaiProtocolMessage_ActionType `protobuf:"varint,1,opt,name=action,proto3,enum=quaiprotocol.QuaiProtocolMessage_ActionType" json:"action,omitempty"`
	Data   []byte                         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // Encoded payload specific to the action
}

func (x *QuaiProtocolMessage) Reset() {
	*x = QuaiProtocolMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuaiProtocolMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuaiProtocolMessage) ProtoMessage() {}

func (x *QuaiProtocolMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuaiProtocolMessage.ProtoReflect.Descriptor instead.
func (*QuaiProtocolMessage) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{2}
}

func (x *QuaiProtocolMessage) GetAction() QuaiProtocolMessage_ActionType {
	if x != nil {
		return x.Action
	}
	return QuaiProtocolMessage_REQUEST_BLOCK
}

func (x *QuaiProtocolMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Define a block structure
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash             string         `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`                         // Unique identifier of the block
	ParentHash       string         `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`             // Hash of the parent block
	Number           uint64         `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`                    // Block number
	StateRoot        string         `protobuf:"bytes,4,opt,name=stateRoot,proto3" json:"stateRoot,omitempty"`               // Root of the state trie
	TransactionsRoot string         `protobuf:"bytes,5,opt,name=transactionsRoot,proto3" json:"transactionsRoot,omitempty"` // Root of the transactions trie
	ReceiptsRoot     string         `protobuf:"bytes,6,opt,name=receiptsRoot,proto3" json:"receiptsRoot,omitempty"`         // Root of the receipts trie
	Transactions     []*Transaction `protobuf:"bytes,7,rep,name=transactions,proto3" json:"transactions,omitempty"`         // Transactions in the block
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{3}
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Block) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *Block) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Block) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *Block) GetTransactionsRoot() string {
	if x != nil {
		return x.TransactionsRoot
	}
	return ""
}

func (x *Block) GetReceiptsRoot() string {
	if x != nil {
		return x.ReceiptsRoot
	}
	return ""
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

// Define a transaction structure
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash     string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`         // Unique identifier of the transaction
	From     string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`         // Sender address
	To       string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`             // Recipient address (empty for contract creation)
	Nonce    uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`      // Nonce of the sender
	Value    string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`       // Value transferred in Wei
	GasPrice string `protobuf:"bytes,6,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"` // Gas price in Wei
	Gas      uint64 `protobuf:"varint,7,opt,name=gas,proto3" json:"gas,omitempty"`          // Gas limit
	Input    []byte `protobuf:"bytes,8,opt,name=input,proto3" json:"input,omitempty"`       // Input data (for contract calls)
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Transaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *Transaction) GetGas() uint64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *Transaction) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

// Request and response messages for block and transaction queries
type BlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *BlockRequest) Reset() {
	*x = BlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequest) ProtoMessage() {}

func (x *BlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequest.ProtoReflect.Descriptor instead.
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{5}
}

func (x *BlockRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool   `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Block *Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{6}
}

func (x *BlockResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *BlockResponse) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{7}
}

func (x *TransactionRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found       bool         `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{8}
}

func (x *TransactionResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *TransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

var File_p2p_pb_quai_messages_proto protoreflect.FileDescriptor

var file_p2p_pb_quai_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x61, 0x69, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x71, 0x75,
	0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x38, 0x0a, 0x0b, 0x47, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x29, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x50, 0x0a, 0x11, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa9, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x61, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x44,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x51, 0x75,
	0x61, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x01, 0x22, 0x80, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x52, 0x6f,
	0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71,
	0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x22, 0x0a,
	0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x22, 0x50, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x22, 0x28, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x68, 0x0a,
	0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x71, 0x75, 0x61,
	0x69, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_pb_quai_messages_proto_rawDescOnce sync.Once
	file_p2p_pb_quai_messages_proto_rawDescData = file_p2p_pb_quai_messages_proto_rawDesc
)

func file_p2p_pb_quai_messages_proto_rawDescGZIP() []byte {
	file_p2p_pb_quai_messages_proto_rawDescOnce.Do(func() {
		file_p2p_pb_quai_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_pb_quai_messages_proto_rawDescData)
	})
	return file_p2p_pb_quai_messages_proto_rawDescData
}

var file_p2p_pb_quai_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_p2p_pb_quai_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_p2p_pb_quai_messages_proto_goTypes = []interface{}{
	(QuaiProtocolMessage_ActionType)(0), // 0: quaiprotocol.QuaiProtocolMessage.ActionType
	(*GossipBlock)(nil),                 // 1: quaiprotocol.GossipBlock
	(*GossipTransaction)(nil),           // 2: quaiprotocol.GossipTransaction
	(*QuaiProtocolMessage)(nil),         // 3: quaiprotocol.QuaiProtocolMessage
	(*Block)(nil),                       // 4: quaiprotocol.Block
	(*Transaction)(nil),                 // 5: quaiprotocol.Transaction
	(*BlockRequest)(nil),                // 6: quaiprotocol.BlockRequest
	(*BlockResponse)(nil),               // 7: quaiprotocol.BlockResponse
	(*TransactionRequest)(nil),          // 8: quaiprotocol.TransactionRequest
	(*TransactionResponse)(nil),         // 9: quaiprotocol.TransactionResponse
}
var file_p2p_pb_quai_messages_proto_depIdxs = []int32{
	4, // 0: quaiprotocol.GossipBlock.block:type_name -> quaiprotocol.Block
	5, // 1: quaiprotocol.GossipTransaction.transaction:type_name -> quaiprotocol.Transaction
	0, // 2: quaiprotocol.QuaiProtocolMessage.action:type_name -> quaiprotocol.QuaiProtocolMessage.ActionType
	5, // 3: quaiprotocol.Block.transactions:type_name -> quaiprotocol.Transaction
	4, // 4: quaiprotocol.BlockResponse.block:type_name -> quaiprotocol.Block
	5, // 5: quaiprotocol.TransactionResponse.transaction:type_name -> quaiprotocol.Transaction
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_p2p_pb_quai_messages_proto_init() }
func file_p2p_pb_quai_messages_proto_init() {
	if File_p2p_pb_quai_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_pb_quai_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipBlock); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipTransaction); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuaiProtocolMessage); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequest); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
			RawDescriptor: file_p2p_pb_quai_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_pb_quai_messages_proto_goTypes,
		DependencyIndexes: file_p2p_pb_quai_messages_proto_depIdxs,
		EnumInfos:         file_p2p_pb_quai_messages_proto_enumTypes,
		MessageInfos:      file_p2p_pb_quai_messages_proto_msgTypes,
	}.Build()
	File_p2p_pb_quai_messages_proto = out.File
	file_p2p_pb_quai_messages_proto_rawDesc = nil
	file_p2p_pb_quai_messages_proto_goTypes = nil
	file_p2p_pb_quai_messages_proto_depIdxs = nil
}
