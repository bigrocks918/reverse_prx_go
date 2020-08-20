// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.4.0
// source: message.proto

package message

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//服务器类型
type SERVICE int32

const (
	SERVICE_NONE          SERVICE = 0
	SERVICE_CLIENT        SERVICE = 1
	SERVICE_GATESERVER    SERVICE = 2 //网关,转发服务
	SERVICE_ACCOUNTSERVER SERVICE = 3 //账号
	SERVICE_WORLDSERVER   SERVICE = 4 //世界
	SERVICE_ZONESERVER    SERVICE = 5 //地图
	SERVICE_WORLDDBSERVER SERVICE = 6 //db
)

// Enum value maps for SERVICE.
var (
	SERVICE_name = map[int32]string{
		0: "NONE",
		1: "CLIENT",
		2: "GATESERVER",
		3: "ACCOUNTSERVER",
		4: "WORLDSERVER",
		5: "ZONESERVER",
		6: "WORLDDBSERVER",
	}
	SERVICE_value = map[string]int32{
		"NONE":          0,
		"CLIENT":        1,
		"GATESERVER":    2,
		"ACCOUNTSERVER": 3,
		"WORLDSERVER":   4,
		"ZONESERVER":    5,
		"WORLDDBSERVER": 6,
	}
)

func (x SERVICE) Enum() *SERVICE {
	p := new(SERVICE)
	*p = x
	return p
}

func (x SERVICE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SERVICE) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (SERVICE) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x SERVICE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SERVICE.Descriptor instead.
func (SERVICE) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type CHAT int32

const (
	CHAT_MSG_TYPE_WORLD   CHAT = 0
	CHAT_MSG_TYPE_PRIVATE CHAT = 1
	CHAT_MSG_TYPE_ORG     CHAT = 2
	CHAT_MSG_TYPE_COUNT   CHAT = 3
)

// Enum value maps for CHAT.
var (
	CHAT_name = map[int32]string{
		0: "MSG_TYPE_WORLD",
		1: "MSG_TYPE_PRIVATE",
		2: "MSG_TYPE_ORG",
		3: "MSG_TYPE_COUNT",
	}
	CHAT_value = map[string]int32{
		"MSG_TYPE_WORLD":   0,
		"MSG_TYPE_PRIVATE": 1,
		"MSG_TYPE_ORG":     2,
		"MSG_TYPE_COUNT":   3,
	}
)

func (x CHAT) Enum() *CHAT {
	p := new(CHAT)
	*p = x
	return p
}

func (x CHAT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CHAT) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (CHAT) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x CHAT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CHAT.Descriptor instead.
func (CHAT) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

//发送标志
type SEND int32

const (
	SEND_POINT      SEND = 0 //指定集群id
	SEND_BALANCE    SEND = 1 //负载
	SEND_BOARD_CAST SEND = 2 //广播
)

// Enum value maps for SEND.
var (
	SEND_name = map[int32]string{
		0: "POINT",
		1: "BALANCE",
		2: "BOARD_CAST",
	}
	SEND_value = map[string]int32{
		"POINT":      0,
		"BALANCE":    1,
		"BOARD_CAST": 2,
	}
)

func (x SEND) Enum() *SEND {
	p := new(SEND)
	*p = x
	return p
}

func (x SEND) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SEND) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[2].Descriptor()
}

func (SEND) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[2]
}

func (x SEND) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SEND.Descriptor instead.
func (SEND) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

//客户端包头
type Ipacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stx            int32   `protobuf:"varint,1,opt,name=Stx,proto3" json:"Stx,omitempty"`
	DestServerType SERVICE `protobuf:"varint,2,opt,name=DestServerType,proto3,enum=message.SERVICE" json:"DestServerType,omitempty"`
	Ckx            int32   `protobuf:"varint,3,opt,name=Ckx,proto3" json:"Ckx,omitempty"`
	Id             int64   `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Ipacket) Reset() {
	*x = Ipacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipacket) ProtoMessage() {}

func (x *Ipacket) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipacket.ProtoReflect.Descriptor instead.
func (*Ipacket) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Ipacket) GetStx() int32 {
	if x != nil {
		return x.Stx
	}
	return 0
}

func (x *Ipacket) GetDestServerType() SERVICE {
	if x != nil {
		return x.DestServerType
	}
	return SERVICE_NONE
}

func (x *Ipacket) GetCkx() int32 {
	if x != nil {
		return x.Ckx
	}
	return 0
}

func (x *Ipacket) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PlayerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID   int64  `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	PlayerName string `protobuf:"bytes,2,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	PlayerGold int32  `protobuf:"varint,3,opt,name=PlayerGold,proto3" json:"PlayerGold,omitempty"`
}

func (x *PlayerData) Reset() {
	*x = PlayerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerData) ProtoMessage() {}

func (x *PlayerData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerData.ProtoReflect.Descriptor instead.
func (*PlayerData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerData) GetPlayerID() int64 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *PlayerData) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PlayerData) GetPlayerGold() int32 {
	if x != nil {
		return x.PlayerGold
	}
	return 0
}

//rpc 包头
type RpcHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"` //token
	SocketId       uint32  `protobuf:"varint,2,opt,name=SocketId,proto3" json:"SocketId,omitempty"`
	SrcClusterId   uint32  `protobuf:"varint,3,opt,name=SrcClusterId,proto3" json:"SrcClusterId,omitempty"`                          //源集群id
	ClusterId      uint32  `protobuf:"varint,4,opt,name=ClusterId,proto3" json:"ClusterId,omitempty"`                                //目标集群id
	DestServerType SERVICE `protobuf:"varint,5,opt,name=DestServerType,proto3,enum=message.SERVICE" json:"DestServerType,omitempty"` //目标集群
	SendType       SEND    `protobuf:"varint,6,opt,name=SendType,proto3,enum=message.SEND" json:"SendType,omitempty"`
	ActorName      string  `protobuf:"bytes,7,opt,name=ActorName,proto3" json:"ActorName,omitempty"`
	SeqId          int64   `protobuf:"varint,8,opt,name=SeqId,proto3" json:"SeqId,omitempty"`
}

func (x *RpcHead) Reset() {
	*x = RpcHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcHead) ProtoMessage() {}

func (x *RpcHead) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcHead.ProtoReflect.Descriptor instead.
func (*RpcHead) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *RpcHead) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RpcHead) GetSocketId() uint32 {
	if x != nil {
		return x.SocketId
	}
	return 0
}

func (x *RpcHead) GetSrcClusterId() uint32 {
	if x != nil {
		return x.SrcClusterId
	}
	return 0
}

func (x *RpcHead) GetClusterId() uint32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *RpcHead) GetDestServerType() SERVICE {
	if x != nil {
		return x.DestServerType
	}
	return SERVICE_NONE
}

func (x *RpcHead) GetSendType() SEND {
	if x != nil {
		return x.SendType
	}
	return SEND_POINT
}

func (x *RpcHead) GetActorName() string {
	if x != nil {
		return x.ActorName
	}
	return ""
}

func (x *RpcHead) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

//rpc 包
type RpcPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FuncName string   `protobuf:"bytes,1,opt,name=FuncName,proto3" json:"FuncName,omitempty"`
	ArgLen   int32    `protobuf:"varint,2,opt,name=ArgLen,proto3" json:"ArgLen,omitempty"`
	RpcHead  *RpcHead `protobuf:"bytes,3,opt,name=RpcHead,proto3" json:"RpcHead,omitempty"`
	RpcBody  []byte   `protobuf:"bytes,4,opt,name=RpcBody,proto3" json:"RpcBody,omitempty"`
}

func (x *RpcPacket) Reset() {
	*x = RpcPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcPacket) ProtoMessage() {}

func (x *RpcPacket) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcPacket.ProtoReflect.Descriptor instead.
func (*RpcPacket) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *RpcPacket) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *RpcPacket) GetArgLen() int32 {
	if x != nil {
		return x.ArgLen
	}
	return 0
}

func (x *RpcPacket) GetRpcHead() *RpcHead {
	if x != nil {
		return x.RpcHead
	}
	return nil
}

func (x *RpcPacket) GetRpcBody() []byte {
	if x != nil {
		return x.RpcBody
	}
	return nil
}

//集群信息
type ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     SERVICE `protobuf:"varint,1,opt,name=Type,proto3,enum=message.SERVICE" json:"Type,omitempty"`
	Ip       string  `protobuf:"bytes,2,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Port     int32   `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
	Weight   int32   `protobuf:"varint,4,opt,name=Weight,proto3" json:"Weight,omitempty"`
	SocketId uint32  `protobuf:"varint,5,opt,name=SocketId,proto3" json:"SocketId,omitempty"`
}

func (x *ClusterInfo) Reset() {
	*x = ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInfo.ProtoReflect.Descriptor instead.
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *ClusterInfo) GetType() SERVICE {
	if x != nil {
		return x.Type
	}
	return SERVICE_NONE
}

func (x *ClusterInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ClusterInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ClusterInfo) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ClusterInfo) GetSocketId() uint32 {
	if x != nil {
		return x.SocketId
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x77, 0x0a, 0x07, 0x49, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x53, 0x74, 0x78, 0x12, 0x38, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x52,
	0x0e, 0x44, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x43, 0x6b, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x43, 0x6b,
	0x78, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x64, 0x22, 0x68, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x6f, 0x6c, 0x64, 0x22, 0x90, 0x02, 0x0a, 0x07,
	0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x72, 0x63, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x53, 0x72, 0x63, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x52,
	0x0e, 0x44, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x45, 0x4e, 0x44,
	0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63,
	0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x65, 0x71, 0x49,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x65, 0x71, 0x49, 0x64, 0x22, 0x85,
	0x01, 0x0a, 0x09, 0x52, 0x70, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x72, 0x67, 0x4c,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x72, 0x67, 0x4c, 0x65, 0x6e,
	0x12, 0x2a, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x70, 0x63, 0x48,
	0x65, 0x61, 0x64, 0x52, 0x07, 0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x70, 0x63, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x52,
	0x70, 0x63, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x2a, 0x76, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x41, 0x54, 0x45, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x4f, 0x52, 0x4c,
	0x44, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x5a, 0x4f, 0x4e,
	0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x4f, 0x52,
	0x4c, 0x44, 0x44, 0x42, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x06, 0x2a, 0x56, 0x0a, 0x04,
	0x43, 0x48, 0x41, 0x54, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x53, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52, 0x47, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x10, 0x03, 0x2a, 0x2e, 0x0a, 0x04, 0x53, 0x45, 0x4e, 0x44, 0x12, 0x09, 0x0a, 0x05,
	0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x41, 0x4c, 0x41, 0x4e,
	0x43, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x5f, 0x43, 0x41,
	0x53, 0x54, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(SERVICE)(0),        // 0: message.SERVICE
	(CHAT)(0),           // 1: message.CHAT
	(SEND)(0),           // 2: message.SEND
	(*Ipacket)(nil),     // 3: message.Ipacket
	(*PlayerData)(nil),  // 4: message.PlayerData
	(*RpcHead)(nil),     // 5: message.RpcHead
	(*RpcPacket)(nil),   // 6: message.RpcPacket
	(*ClusterInfo)(nil), // 7: message.ClusterInfo
}
var file_message_proto_depIdxs = []int32{
	0, // 0: message.Ipacket.DestServerType:type_name -> message.SERVICE
	0, // 1: message.RpcHead.DestServerType:type_name -> message.SERVICE
	2, // 2: message.RpcHead.SendType:type_name -> message.SEND
	5, // 3: message.RpcPacket.RpcHead:type_name -> message.RpcHead
	0, // 4: message.ClusterInfo.Type:type_name -> message.SERVICE
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipacket); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerData); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcHead); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcPacket); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInfo); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
