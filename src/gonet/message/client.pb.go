// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type C_A_LoginRequest struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	AccountName          *string  `protobuf:"bytes,2,req,name=AccountName" json:"AccountName,omitempty"`
	BuildNo              *string  `protobuf:"bytes,3,req,name=BuildNo" json:"BuildNo,omitempty"`
	SocketId             *int32   `protobuf:"varint,4,req,name=SocketId" json:"SocketId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_A_LoginRequest) Reset()         { *m = C_A_LoginRequest{} }
func (m *C_A_LoginRequest) String() string { return proto.CompactTextString(m) }
func (*C_A_LoginRequest) ProtoMessage()    {}
func (*C_A_LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{0}
}
func (m *C_A_LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_A_LoginRequest.Unmarshal(m, b)
}
func (m *C_A_LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_A_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *C_A_LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_A_LoginRequest.Merge(dst, src)
}
func (m *C_A_LoginRequest) XXX_Size() int {
	return xxx_messageInfo_C_A_LoginRequest.Size(m)
}
func (m *C_A_LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_C_A_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_C_A_LoginRequest proto.InternalMessageInfo

func (m *C_A_LoginRequest) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_A_LoginRequest) GetAccountName() string {
	if m != nil && m.AccountName != nil {
		return *m.AccountName
	}
	return ""
}

func (m *C_A_LoginRequest) GetBuildNo() string {
	if m != nil && m.BuildNo != nil {
		return *m.BuildNo
	}
	return ""
}

func (m *C_A_LoginRequest) GetSocketId() int32 {
	if m != nil && m.SocketId != nil {
		return *m.SocketId
	}
	return 0
}

type C_A_LoginRequest1 struct {
	Login                *C_A_LoginRequest `protobuf:"bytes,1,req,name=Login" json:"Login,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *C_A_LoginRequest1) Reset()         { *m = C_A_LoginRequest1{} }
func (m *C_A_LoginRequest1) String() string { return proto.CompactTextString(m) }
func (*C_A_LoginRequest1) ProtoMessage()    {}
func (*C_A_LoginRequest1) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{1}
}
func (m *C_A_LoginRequest1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_A_LoginRequest1.Unmarshal(m, b)
}
func (m *C_A_LoginRequest1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_A_LoginRequest1.Marshal(b, m, deterministic)
}
func (dst *C_A_LoginRequest1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_A_LoginRequest1.Merge(dst, src)
}
func (m *C_A_LoginRequest1) XXX_Size() int {
	return xxx_messageInfo_C_A_LoginRequest1.Size(m)
}
func (m *C_A_LoginRequest1) XXX_DiscardUnknown() {
	xxx_messageInfo_C_A_LoginRequest1.DiscardUnknown(m)
}

var xxx_messageInfo_C_A_LoginRequest1 proto.InternalMessageInfo

func (m *C_A_LoginRequest1) GetLogin() *C_A_LoginRequest {
	if m != nil {
		return m.Login
	}
	return nil
}

type C_A_RegisterRequest struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	AccountName          *string  `protobuf:"bytes,2,req,name=AccountName" json:"AccountName,omitempty"`
	SocketId             *int32   `protobuf:"varint,3,req,name=SocketId" json:"SocketId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_A_RegisterRequest) Reset()         { *m = C_A_RegisterRequest{} }
func (m *C_A_RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*C_A_RegisterRequest) ProtoMessage()    {}
func (*C_A_RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{2}
}
func (m *C_A_RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_A_RegisterRequest.Unmarshal(m, b)
}
func (m *C_A_RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_A_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *C_A_RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_A_RegisterRequest.Merge(dst, src)
}
func (m *C_A_RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_C_A_RegisterRequest.Size(m)
}
func (m *C_A_RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_C_A_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_C_A_RegisterRequest proto.InternalMessageInfo

func (m *C_A_RegisterRequest) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_A_RegisterRequest) GetAccountName() string {
	if m != nil && m.AccountName != nil {
		return *m.AccountName
	}
	return ""
}

func (m *C_A_RegisterRequest) GetSocketId() int32 {
	if m != nil && m.SocketId != nil {
		return *m.SocketId
	}
	return 0
}

type C_W_CreatePlayerRequest struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	PlayerName           *string  `protobuf:"bytes,2,req,name=PlayerName" json:"PlayerName,omitempty"`
	Sex                  *int32   `protobuf:"varint,3,req,name=Sex" json:"Sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_W_CreatePlayerRequest) Reset()         { *m = C_W_CreatePlayerRequest{} }
func (m *C_W_CreatePlayerRequest) String() string { return proto.CompactTextString(m) }
func (*C_W_CreatePlayerRequest) ProtoMessage()    {}
func (*C_W_CreatePlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{3}
}
func (m *C_W_CreatePlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_W_CreatePlayerRequest.Unmarshal(m, b)
}
func (m *C_W_CreatePlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_W_CreatePlayerRequest.Marshal(b, m, deterministic)
}
func (dst *C_W_CreatePlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_W_CreatePlayerRequest.Merge(dst, src)
}
func (m *C_W_CreatePlayerRequest) XXX_Size() int {
	return xxx_messageInfo_C_W_CreatePlayerRequest.Size(m)
}
func (m *C_W_CreatePlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_C_W_CreatePlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_C_W_CreatePlayerRequest proto.InternalMessageInfo

func (m *C_W_CreatePlayerRequest) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_W_CreatePlayerRequest) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *C_W_CreatePlayerRequest) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

type C_W_Game_LoginRequset struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	PlayerId             *int64   `protobuf:"varint,2,req,name=PlayerId" json:"PlayerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_W_Game_LoginRequset) Reset()         { *m = C_W_Game_LoginRequset{} }
func (m *C_W_Game_LoginRequset) String() string { return proto.CompactTextString(m) }
func (*C_W_Game_LoginRequset) ProtoMessage()    {}
func (*C_W_Game_LoginRequset) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{4}
}
func (m *C_W_Game_LoginRequset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_W_Game_LoginRequset.Unmarshal(m, b)
}
func (m *C_W_Game_LoginRequset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_W_Game_LoginRequset.Marshal(b, m, deterministic)
}
func (dst *C_W_Game_LoginRequset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_W_Game_LoginRequset.Merge(dst, src)
}
func (m *C_W_Game_LoginRequset) XXX_Size() int {
	return xxx_messageInfo_C_W_Game_LoginRequset.Size(m)
}
func (m *C_W_Game_LoginRequset) XXX_DiscardUnknown() {
	xxx_messageInfo_C_W_Game_LoginRequset.DiscardUnknown(m)
}

var xxx_messageInfo_C_W_Game_LoginRequset proto.InternalMessageInfo

func (m *C_W_Game_LoginRequset) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_W_Game_LoginRequset) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type C_G_LogoutResponse struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_G_LogoutResponse) Reset()         { *m = C_G_LogoutResponse{} }
func (m *C_G_LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*C_G_LogoutResponse) ProtoMessage()    {}
func (*C_G_LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{5}
}
func (m *C_G_LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_G_LogoutResponse.Unmarshal(m, b)
}
func (m *C_G_LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_G_LogoutResponse.Marshal(b, m, deterministic)
}
func (dst *C_G_LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_G_LogoutResponse.Merge(dst, src)
}
func (m *C_G_LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_C_G_LogoutResponse.Size(m)
}
func (m *C_G_LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_C_G_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_C_G_LogoutResponse proto.InternalMessageInfo

func (m *C_G_LogoutResponse) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

type C_W_ChatMessage struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	Sender               *int64   `protobuf:"varint,2,req,name=Sender" json:"Sender,omitempty"`
	Recver               *int64   `protobuf:"varint,3,req,name=Recver" json:"Recver,omitempty"`
	MessageType          *int32   `protobuf:"varint,4,req,name=MessageType" json:"MessageType,omitempty"`
	Message              *string  `protobuf:"bytes,5,req,name=Message" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_W_ChatMessage) Reset()         { *m = C_W_ChatMessage{} }
func (m *C_W_ChatMessage) String() string { return proto.CompactTextString(m) }
func (*C_W_ChatMessage) ProtoMessage()    {}
func (*C_W_ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{6}
}
func (m *C_W_ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_W_ChatMessage.Unmarshal(m, b)
}
func (m *C_W_ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_W_ChatMessage.Marshal(b, m, deterministic)
}
func (dst *C_W_ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_W_ChatMessage.Merge(dst, src)
}
func (m *C_W_ChatMessage) XXX_Size() int {
	return xxx_messageInfo_C_W_ChatMessage.Size(m)
}
func (m *C_W_ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_C_W_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_C_W_ChatMessage proto.InternalMessageInfo

func (m *C_W_ChatMessage) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_W_ChatMessage) GetSender() int64 {
	if m != nil && m.Sender != nil {
		return *m.Sender
	}
	return 0
}

func (m *C_W_ChatMessage) GetRecver() int64 {
	if m != nil && m.Recver != nil {
		return *m.Recver
	}
	return 0
}

func (m *C_W_ChatMessage) GetMessageType() int32 {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return 0
}

func (m *C_W_ChatMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type W_C_ChatMessage struct {
	PacketHead           *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	Sender               *int64   `protobuf:"varint,2,req,name=Sender" json:"Sender,omitempty"`
	SenderName           *string  `protobuf:"bytes,3,req,name=SenderName" json:"SenderName,omitempty"`
	Recver               *int64   `protobuf:"varint,4,req,name=Recver" json:"Recver,omitempty"`
	RecverName           *string  `protobuf:"bytes,5,req,name=RecverName" json:"RecverName,omitempty"`
	MessageType          *int32   `protobuf:"varint,6,req,name=MessageType" json:"MessageType,omitempty"`
	Message              *string  `protobuf:"bytes,7,req,name=Message" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *W_C_ChatMessage) Reset()         { *m = W_C_ChatMessage{} }
func (m *W_C_ChatMessage) String() string { return proto.CompactTextString(m) }
func (*W_C_ChatMessage) ProtoMessage()    {}
func (*W_C_ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{7}
}
func (m *W_C_ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_W_C_ChatMessage.Unmarshal(m, b)
}
func (m *W_C_ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_W_C_ChatMessage.Marshal(b, m, deterministic)
}
func (dst *W_C_ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_W_C_ChatMessage.Merge(dst, src)
}
func (m *W_C_ChatMessage) XXX_Size() int {
	return xxx_messageInfo_W_C_ChatMessage.Size(m)
}
func (m *W_C_ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_W_C_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_W_C_ChatMessage proto.InternalMessageInfo

func (m *W_C_ChatMessage) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *W_C_ChatMessage) GetSender() int64 {
	if m != nil && m.Sender != nil {
		return *m.Sender
	}
	return 0
}

func (m *W_C_ChatMessage) GetSenderName() string {
	if m != nil && m.SenderName != nil {
		return *m.SenderName
	}
	return ""
}

func (m *W_C_ChatMessage) GetRecver() int64 {
	if m != nil && m.Recver != nil {
		return *m.Recver
	}
	return 0
}

func (m *W_C_ChatMessage) GetRecverName() string {
	if m != nil && m.RecverName != nil {
		return *m.RecverName
	}
	return ""
}

func (m *W_C_ChatMessage) GetMessageType() int32 {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return 0
}

func (m *W_C_ChatMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type W_C_Test struct {
	Recv                 []int32  `protobuf:"varint,1,rep,name=Recv" json:"Recv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *W_C_Test) Reset()         { *m = W_C_Test{} }
func (m *W_C_Test) String() string { return proto.CompactTextString(m) }
func (*W_C_Test) ProtoMessage()    {}
func (*W_C_Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_847e00738f02fccc, []int{8}
}
func (m *W_C_Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_W_C_Test.Unmarshal(m, b)
}
func (m *W_C_Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_W_C_Test.Marshal(b, m, deterministic)
}
func (dst *W_C_Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_W_C_Test.Merge(dst, src)
}
func (m *W_C_Test) XXX_Size() int {
	return xxx_messageInfo_W_C_Test.Size(m)
}
func (m *W_C_Test) XXX_DiscardUnknown() {
	xxx_messageInfo_W_C_Test.DiscardUnknown(m)
}

var xxx_messageInfo_W_C_Test proto.InternalMessageInfo

func (m *W_C_Test) GetRecv() []int32 {
	if m != nil {
		return m.Recv
	}
	return nil
}

func init() {
	proto.RegisterType((*C_A_LoginRequest)(nil), "message.C_A_LoginRequest")
	proto.RegisterType((*C_A_LoginRequest1)(nil), "message.C_A_LoginRequest1")
	proto.RegisterType((*C_A_RegisterRequest)(nil), "message.C_A_RegisterRequest")
	proto.RegisterType((*C_W_CreatePlayerRequest)(nil), "message.C_W_CreatePlayerRequest")
	proto.RegisterType((*C_W_Game_LoginRequset)(nil), "message.C_W_Game_LoginRequset")
	proto.RegisterType((*C_G_LogoutResponse)(nil), "message.C_G_LogoutResponse")
	proto.RegisterType((*C_W_ChatMessage)(nil), "message.C_W_ChatMessage")
	proto.RegisterType((*W_C_ChatMessage)(nil), "message.W_C_ChatMessage")
	proto.RegisterType((*W_C_Test)(nil), "message.W_C_Test")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_client_847e00738f02fccc) }

var fileDescriptor_client_847e00738f02fccc = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcb, 0xae, 0xd3, 0x30,
	0x10, 0x55, 0x92, 0xe6, 0xb6, 0xcc, 0x05, 0xdd, 0x62, 0x04, 0x84, 0x2e, 0xa2, 0x28, 0xab, 0xac,
	0x2e, 0x8f, 0x3f, 0xb8, 0x04, 0x51, 0x2a, 0x41, 0x55, 0xb9, 0x95, 0xba, 0x8c, 0xac, 0x64, 0x54,
	0x22, 0xda, 0x38, 0x24, 0x0e, 0xa2, 0x0b, 0x96, 0xfc, 0x05, 0xff, 0xc0, 0xa7, 0xf1, 0x0b, 0xc8,
	0x8f, 0x16, 0x37, 0x0b, 0x24, 0x8a, 0xba, 0x9b, 0x73, 0x3c, 0xf1, 0x39, 0x67, 0x3c, 0x81, 0xfb,
	0xf9, 0xb6, 0xc4, 0x4a, 0xdc, 0xd6, 0x0d, 0x17, 0x9c, 0x0c, 0x77, 0xd8, 0xb6, 0x6c, 0x83, 0x93,
	0x07, 0xa6, 0xd0, 0x7c, 0xfc, 0xc3, 0x81, 0x71, 0x9a, 0xdd, 0x65, 0xef, 0xf9, 0xa6, 0xac, 0x28,
	0x7e, 0xee, 0xb0, 0x15, 0xe4, 0x05, 0xc0, 0x82, 0xe5, 0x9f, 0x50, 0xbc, 0x43, 0x56, 0x04, 0x4e,
	0xe4, 0x26, 0xd7, 0xaf, 0xc6, 0xb7, 0x87, 0x0f, 0x67, 0xb5, 0x3a, 0xa3, 0x56, 0x0f, 0x89, 0xe0,
	0xfa, 0x2e, 0xcf, 0x79, 0x57, 0x89, 0x39, 0xdb, 0x61, 0xe0, 0x46, 0x6e, 0x72, 0x8f, 0xda, 0x14,
	0x09, 0x60, 0xf8, 0xba, 0x2b, 0xb7, 0xc5, 0x9c, 0x07, 0x9e, 0x3a, 0x3d, 0x40, 0x32, 0x81, 0xd1,
	0x92, 0xcb, 0x9b, 0x66, 0x45, 0x30, 0x88, 0xdc, 0xc4, 0xa7, 0x47, 0x1c, 0xbf, 0x81, 0x87, 0x7d,
	0x77, 0x2f, 0xc9, 0x73, 0xf0, 0x15, 0x61, 0x9c, 0x3d, 0x3b, 0x3a, 0xeb, 0xb7, 0x52, 0xdd, 0x17,
	0x7f, 0x77, 0xe0, 0x91, 0x3c, 0xa3, 0xb8, 0x29, 0x5b, 0x81, 0xcd, 0x25, 0x73, 0xda, 0x69, 0xbc,
	0x5e, 0x9a, 0x6f, 0xf0, 0x34, 0xcd, 0xd6, 0x59, 0xda, 0x20, 0x13, 0xb8, 0xd8, 0xb2, 0xfd, 0xff,
	0x58, 0x09, 0x01, 0xf4, 0x15, 0x96, 0x13, 0x8b, 0x21, 0x63, 0xf0, 0x96, 0xf8, 0xd5, 0x78, 0x90,
	0x65, 0x8c, 0xf0, 0x58, 0xca, 0x4f, 0xd9, 0x0e, 0xff, 0x8c, 0xa9, 0xc5, 0x73, 0xc4, 0x27, 0x30,
	0xd2, 0x52, 0xb3, 0x42, 0x49, 0x7b, 0xf4, 0x88, 0xe3, 0xb7, 0x40, 0xd2, 0x6c, 0x2a, 0x15, 0x78,
	0x27, 0x28, 0xb6, 0x35, 0xaf, 0x5a, 0xfc, 0x77, 0x8d, 0xf8, 0xa7, 0x03, 0x37, 0x6a, 0x5c, 0x1f,
	0x99, 0xf8, 0xa0, 0xfb, 0xce, 0x70, 0xfa, 0x04, 0xae, 0x96, 0x58, 0x15, 0xd8, 0x18, 0x9f, 0x06,
	0x49, 0x9e, 0x62, 0xfe, 0x05, 0x1b, 0x35, 0x21, 0x8f, 0x1a, 0x24, 0x5f, 0xd8, 0x88, 0xad, 0xf6,
	0x35, 0x9a, 0x85, 0xb4, 0x29, 0xb9, 0xc9, 0x06, 0x06, 0xbe, 0xde, 0x64, 0x03, 0xe3, 0x5f, 0x0e,
	0xdc, 0xac, 0xb3, 0xf4, 0x42, 0x8e, 0x43, 0x00, 0x5d, 0xa9, 0x07, 0xd7, 0x3f, 0x91, 0xc5, 0x58,
	0x89, 0x06, 0x27, 0x89, 0x42, 0x00, 0x5d, 0xa9, 0xef, 0xb4, 0x65, 0x8b, 0xe9, 0x27, 0xbe, 0xfa,
	0x6b, 0xe2, 0xe1, 0x69, 0xe2, 0x10, 0x46, 0x32, 0xf0, 0x4a, 0xae, 0x30, 0x81, 0x81, 0xbc, 0x35,
	0x70, 0x22, 0x2f, 0xf1, 0xa9, 0xaa, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x8c, 0x00, 0x6d,
	0x85, 0x04, 0x00, 0x00,
}
