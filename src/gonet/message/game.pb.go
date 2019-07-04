// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// point3F
type Point3F struct {
	X float32 `protobuf:"fixed32,1,opt,name=X" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y" json:"Y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=Z" json:"Z,omitempty"`
}

func (m *Point3F) Reset()                    { *m = Point3F{} }
func (m *Point3F) String() string            { return proto.CompactTextString(m) }
func (*Point3F) ProtoMessage()               {}
func (*Point3F) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Point3F) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point3F) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Point3F) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

// 移动包
type C_W_Move struct {
	PacketHead *Ipacket       `protobuf:"bytes,1,opt,name=PacketHead" json:"PacketHead,omitempty"`
	Move       *C_W_Move_Move `protobuf:"bytes,2,opt,name=move" json:"move,omitempty"`
}

func (m *C_W_Move) Reset()                    { *m = C_W_Move{} }
func (m *C_W_Move) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move) ProtoMessage()               {}
func (*C_W_Move) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *C_W_Move) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_W_Move) GetMove() *C_W_Move_Move {
	if m != nil {
		return m.Move
	}
	return nil
}

type C_W_Move_Move struct {
	Mode   int32                 `protobuf:"varint,1,opt,name=Mode" json:"Mode,omitempty"`
	Normal *C_W_Move_Move_Normal `protobuf:"bytes,2,opt,name=normal" json:"normal,omitempty"`
	Path   *C_W_Move_Move_Path   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Link   *C_W_Move_Move_Blink  `protobuf:"bytes,4,opt,name=link" json:"link,omitempty"`
	Jump   *C_W_Move_Move_Jump   `protobuf:"bytes,5,opt,name=jump" json:"jump,omitempty"`
	Line   *C_W_Move_Move_Line   `protobuf:"bytes,6,opt,name=line" json:"line,omitempty"`
}

func (m *C_W_Move_Move) Reset()                    { *m = C_W_Move_Move{} }
func (m *C_W_Move_Move) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move) ProtoMessage()               {}
func (*C_W_Move_Move) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

func (m *C_W_Move_Move) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *C_W_Move_Move) GetNormal() *C_W_Move_Move_Normal {
	if m != nil {
		return m.Normal
	}
	return nil
}

func (m *C_W_Move_Move) GetPath() *C_W_Move_Move_Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *C_W_Move_Move) GetLink() *C_W_Move_Move_Blink {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *C_W_Move_Move) GetJump() *C_W_Move_Move_Jump {
	if m != nil {
		return m.Jump
	}
	return nil
}

func (m *C_W_Move_Move) GetLine() *C_W_Move_Move_Line {
	if m != nil {
		return m.Line
	}
	return nil
}

type C_W_Move_Move_Normal struct {
	Pos      *Point3F `protobuf:"bytes,1,opt,name=Pos" json:"Pos,omitempty"`
	Yaw      float32  `protobuf:"fixed32,2,opt,name=Yaw" json:"Yaw,omitempty"`
	Duration float32  `protobuf:"fixed32,3,opt,name=Duration" json:"Duration,omitempty"`
}

func (m *C_W_Move_Move_Normal) Reset()                    { *m = C_W_Move_Move_Normal{} }
func (m *C_W_Move_Move_Normal) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Normal) ProtoMessage()               {}
func (*C_W_Move_Move_Normal) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 0} }

func (m *C_W_Move_Move_Normal) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *C_W_Move_Move_Normal) GetYaw() float32 {
	if m != nil {
		return m.Yaw
	}
	return 0
}

func (m *C_W_Move_Move_Normal) GetDuration() float32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type C_W_Move_Move_Path struct {
	PathId  int32 `protobuf:"varint,1,opt,name=PathId" json:"PathId,omitempty"`
	TimePos int32 `protobuf:"varint,2,opt,name=TimePos" json:"TimePos,omitempty"`
	MountId int32 `protobuf:"varint,3,opt,name=MountId" json:"MountId,omitempty"`
}

func (m *C_W_Move_Move_Path) Reset()                    { *m = C_W_Move_Move_Path{} }
func (m *C_W_Move_Move_Path) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Path) ProtoMessage()               {}
func (*C_W_Move_Move_Path) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 1} }

func (m *C_W_Move_Move_Path) GetPathId() int32 {
	if m != nil {
		return m.PathId
	}
	return 0
}

func (m *C_W_Move_Move_Path) GetTimePos() int32 {
	if m != nil {
		return m.TimePos
	}
	return 0
}

func (m *C_W_Move_Move_Path) GetMountId() int32 {
	if m != nil {
		return m.MountId
	}
	return 0
}

type C_W_Move_Move_Blink struct {
	Pos  *Point3F `protobuf:"bytes,1,opt,name=Pos" json:"Pos,omitempty"`
	RPos *Point3F `protobuf:"bytes,2,opt,name=RPos" json:"RPos,omitempty"`
}

func (m *C_W_Move_Move_Blink) Reset()                    { *m = C_W_Move_Move_Blink{} }
func (m *C_W_Move_Move_Blink) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Blink) ProtoMessage()               {}
func (*C_W_Move_Move_Blink) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 2} }

func (m *C_W_Move_Move_Blink) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *C_W_Move_Move_Blink) GetRPos() *Point3F {
	if m != nil {
		return m.RPos
	}
	return nil
}

type C_W_Move_Move_Jump struct {
	BPos      *Point3F `protobuf:"bytes,1,opt,name=BPos" json:"BPos,omitempty"`
	EPos      *Point3F `protobuf:"bytes,2,opt,name=EPos" json:"EPos,omitempty"`
	Duration  int32    `protobuf:"varint,3,opt,name=Duration" json:"Duration,omitempty"`
	TimePos   int32    `protobuf:"varint,4,opt,name=TimePos" json:"TimePos,omitempty"`
	UpExDur   int32    `protobuf:"varint,5,opt,name=UpExDur" json:"UpExDur,omitempty"`
	DownExDur int32    `protobuf:"varint,6,opt,name=DownExDur" json:"DownExDur,omitempty"`
	A         int32    `protobuf:"varint,7,opt,name=A" json:"A,omitempty"`
	B         int32    `protobuf:"varint,8,opt,name=B" json:"B,omitempty"`
}

func (m *C_W_Move_Move_Jump) Reset()                    { *m = C_W_Move_Move_Jump{} }
func (m *C_W_Move_Move_Jump) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Jump) ProtoMessage()               {}
func (*C_W_Move_Move_Jump) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 3} }

func (m *C_W_Move_Move_Jump) GetBPos() *Point3F {
	if m != nil {
		return m.BPos
	}
	return nil
}

func (m *C_W_Move_Move_Jump) GetEPos() *Point3F {
	if m != nil {
		return m.EPos
	}
	return nil
}

func (m *C_W_Move_Move_Jump) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetTimePos() int32 {
	if m != nil {
		return m.TimePos
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetUpExDur() int32 {
	if m != nil {
		return m.UpExDur
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetDownExDur() int32 {
	if m != nil {
		return m.DownExDur
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type C_W_Move_Move_Line struct {
	BPos     *Point3F `protobuf:"bytes,1,opt,name=BPos" json:"BPos,omitempty"`
	EPos     *Point3F `protobuf:"bytes,2,opt,name=EPos" json:"EPos,omitempty"`
	Duration int32    `protobuf:"varint,3,opt,name=Duration" json:"Duration,omitempty"`
	TimePos  int32    `protobuf:"varint,4,opt,name=TimePos" json:"TimePos,omitempty"`
}

func (m *C_W_Move_Move_Line) Reset()                    { *m = C_W_Move_Move_Line{} }
func (m *C_W_Move_Move_Line) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Line) ProtoMessage()               {}
func (*C_W_Move_Move_Line) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 4} }

func (m *C_W_Move_Move_Line) GetBPos() *Point3F {
	if m != nil {
		return m.BPos
	}
	return nil
}

func (m *C_W_Move_Move_Line) GetEPos() *Point3F {
	if m != nil {
		return m.EPos
	}
	return nil
}

func (m *C_W_Move_Move_Line) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *C_W_Move_Move_Line) GetTimePos() int32 {
	if m != nil {
		return m.TimePos
	}
	return 0
}

type W_C_LoginMap struct {
	PacketHead *Ipacket `protobuf:"bytes,1,opt,name=PacketHead" json:"PacketHead,omitempty"`
	Id         int64    `protobuf:"varint,2,opt,name=Id" json:"Id,omitempty"`
	Pos        *Point3F `protobuf:"bytes,3,opt,name=Pos" json:"Pos,omitempty"`
	Rotation   float32  `protobuf:"fixed32,4,opt,name=Rotation" json:"Rotation,omitempty"`
}

func (m *W_C_LoginMap) Reset()                    { *m = W_C_LoginMap{} }
func (m *W_C_LoginMap) String() string            { return proto.CompactTextString(m) }
func (*W_C_LoginMap) ProtoMessage()               {}
func (*W_C_LoginMap) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *W_C_LoginMap) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *W_C_LoginMap) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *W_C_LoginMap) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *W_C_LoginMap) GetRotation() float32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

type W_C_Move struct {
	PacketHead *Ipacket `protobuf:"bytes,1,opt,name=PacketHead" json:"PacketHead,omitempty"`
	Id         int64    `protobuf:"varint,2,opt,name=Id" json:"Id,omitempty"`
	Pos        *Point3F `protobuf:"bytes,3,opt,name=Pos" json:"Pos,omitempty"`
	Rotation   float32  `protobuf:"fixed32,4,opt,name=Rotation" json:"Rotation,omitempty"`
}

func (m *W_C_Move) Reset()                    { *m = W_C_Move{} }
func (m *W_C_Move) String() string            { return proto.CompactTextString(m) }
func (*W_C_Move) ProtoMessage()               {}
func (*W_C_Move) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *W_C_Move) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *W_C_Move) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *W_C_Move) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *W_C_Move) GetRotation() float32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

type W_C_ADD_SIMOBJ struct {
	PacketHead *Ipacket `protobuf:"bytes,1,opt,name=PacketHead" json:"PacketHead,omitempty"`
	Id         int64    `protobuf:"varint,2,opt,name=Id" json:"Id,omitempty"`
	Pos        *Point3F `protobuf:"bytes,3,opt,name=Pos" json:"Pos,omitempty"`
	Rotation   float32  `protobuf:"fixed32,4,opt,name=Rotation" json:"Rotation,omitempty"`
}

func (m *W_C_ADD_SIMOBJ) Reset()                    { *m = W_C_ADD_SIMOBJ{} }
func (m *W_C_ADD_SIMOBJ) String() string            { return proto.CompactTextString(m) }
func (*W_C_ADD_SIMOBJ) ProtoMessage()               {}
func (*W_C_ADD_SIMOBJ) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *W_C_ADD_SIMOBJ) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *W_C_ADD_SIMOBJ) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *W_C_ADD_SIMOBJ) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *W_C_ADD_SIMOBJ) GetRotation() float32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

type C_W_LoginCopyMap struct {
	PacketHead *Ipacket `protobuf:"bytes,1,opt,name=PacketHead" json:"PacketHead,omitempty"`
	DataId     int32    `protobuf:"varint,2,opt,name=DataId" json:"DataId,omitempty"`
}

func (m *C_W_LoginCopyMap) Reset()                    { *m = C_W_LoginCopyMap{} }
func (m *C_W_LoginCopyMap) String() string            { return proto.CompactTextString(m) }
func (*C_W_LoginCopyMap) ProtoMessage()               {}
func (*C_W_LoginCopyMap) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *C_W_LoginCopyMap) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_W_LoginCopyMap) GetDataId() int32 {
	if m != nil {
		return m.DataId
	}
	return 0
}

func init() {
	proto.RegisterType((*Point3F)(nil), "message.Point3F")
	proto.RegisterType((*C_W_Move)(nil), "message.C_W_Move")
	proto.RegisterType((*C_W_Move_Move)(nil), "message.C_W_Move.Move")
	proto.RegisterType((*C_W_Move_Move_Normal)(nil), "message.C_W_Move.Move.Normal")
	proto.RegisterType((*C_W_Move_Move_Path)(nil), "message.C_W_Move.Move.Path")
	proto.RegisterType((*C_W_Move_Move_Blink)(nil), "message.C_W_Move.Move.Blink")
	proto.RegisterType((*C_W_Move_Move_Jump)(nil), "message.C_W_Move.Move.Jump")
	proto.RegisterType((*C_W_Move_Move_Line)(nil), "message.C_W_Move.Move.Line")
	proto.RegisterType((*W_C_LoginMap)(nil), "message.W_C_LoginMap")
	proto.RegisterType((*W_C_Move)(nil), "message.W_C_Move")
	proto.RegisterType((*W_C_ADD_SIMOBJ)(nil), "message.W_C_ADD_SIMOBJ")
	proto.RegisterType((*C_W_LoginCopyMap)(nil), "message.C_W_LoginCopyMap")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6a, 0x13, 0x41,
	0x18, 0x67, 0x93, 0xd9, 0xdd, 0xf8, 0xa5, 0x96, 0x30, 0x87, 0xb0, 0xac, 0x15, 0x24, 0x78, 0x10,
	0x0f, 0xb1, 0x24, 0xf8, 0x00, 0x49, 0xb6, 0xe2, 0x96, 0x46, 0xe3, 0xa8, 0xb4, 0x09, 0x42, 0x18,
	0xdd, 0x21, 0x5d, 0x9b, 0xdd, 0x59, 0x92, 0x4d, 0xab, 0x67, 0x2f, 0x1e, 0x3c, 0x88, 0xcf, 0xe7,
	0x33, 0xf8, 0x0c, 0xf2, 0x7d, 0x3b, 0xdb, 0x16, 0xd2, 0x06, 0x09, 0x82, 0x5e, 0xb2, 0xf3, 0x9b,
	0xdf, 0xef, 0xfb, 0x3f, 0x33, 0x01, 0x98, 0xc9, 0x44, 0xb5, 0xb3, 0x85, 0xce, 0x35, 0x77, 0x13,
	0xb5, 0x5c, 0xca, 0x99, 0xf2, 0xef, 0x9a, 0x45, 0xb1, 0xdf, 0xea, 0x82, 0x3b, 0xd2, 0x71, 0x9a,
	0x77, 0x9f, 0xf1, 0x1d, 0xb0, 0x4e, 0x3c, 0xeb, 0x81, 0xf5, 0xa8, 0x22, 0xac, 0x13, 0x44, 0x63,
	0xaf, 0x52, 0xa0, 0x31, 0xa2, 0x89, 0x57, 0x2d, 0xd0, 0xa4, 0xf5, 0xcb, 0x85, 0xda, 0x60, 0x7a,
	0x3c, 0x1d, 0xea, 0x73, 0xc5, 0xf7, 0x01, 0x46, 0xf2, 0xc3, 0x99, 0xca, 0x9f, 0x2b, 0x19, 0x91,
	0x7d, 0xbd, 0xd3, 0x68, 0x97, 0x51, 0xc2, 0x8c, 0x38, 0x71, 0x4d, 0xc3, 0x1f, 0x03, 0x4b, 0xf4,
	0xb9, 0x22, 0xef, 0xf5, 0x4e, 0xf3, 0x52, 0x5b, 0xba, 0x6c, 0xe3, 0x8f, 0x20, 0x8d, 0xff, 0xc5,
	0x05, 0x46, 0x61, 0x38, 0x7e, 0x23, 0x45, 0x01, 0x6c, 0x41, 0x6b, 0xfe, 0x14, 0x9c, 0x54, 0x2f,
	0x12, 0x39, 0x37, 0xae, 0xee, 0xdf, 0xec, 0xaa, 0xfd, 0x82, 0x44, 0xc2, 0x88, 0xf9, 0x13, 0x60,
	0x99, 0xcc, 0x4f, 0xa9, 0x9e, 0x7a, 0xe7, 0xde, 0x2d, 0x46, 0x23, 0x99, 0x9f, 0x0a, 0x12, 0xf2,
	0x7d, 0x60, 0xf3, 0x38, 0x3d, 0xf3, 0x18, 0x19, 0xec, 0xdd, 0x62, 0xd0, 0x47, 0x8d, 0x20, 0x25,
	0x86, 0xf8, 0xb8, 0x4a, 0x32, 0xcf, 0xde, 0x18, 0xe2, 0x70, 0x95, 0x64, 0x82, 0x84, 0x68, 0x30,
	0x8f, 0x53, 0xe5, 0x39, 0x1b, 0x0d, 0x8e, 0xe2, 0x54, 0x51, 0x04, 0xe5, 0x4f, 0xc0, 0x29, 0xca,
	0xe2, 0x2d, 0xa8, 0x8e, 0xf4, 0x72, 0xad, 0xf3, 0x66, 0xac, 0x02, 0x49, 0xde, 0x80, 0xea, 0x58,
	0x5e, 0x98, 0x79, 0xe2, 0x92, 0xfb, 0x50, 0x0b, 0x56, 0x0b, 0x99, 0xc7, 0x3a, 0x35, 0x83, 0xbd,
	0xc4, 0xbe, 0x00, 0x86, 0xd5, 0xf3, 0x26, 0x38, 0xf8, 0x0d, 0x23, 0xd3, 0x75, 0x83, 0xb8, 0x07,
	0xee, 0x9b, 0x38, 0x51, 0x18, 0xb5, 0x42, 0x44, 0x09, 0x91, 0x19, 0xea, 0x55, 0x9a, 0x87, 0x11,
	0x39, 0xb5, 0x45, 0x09, 0xfd, 0x57, 0x60, 0x53, 0x83, 0xfe, 0x28, 0xdd, 0x87, 0xc0, 0x44, 0xe9,
	0xfd, 0x26, 0x11, 0xb1, 0xfe, 0x4f, 0x0b, 0x18, 0xb6, 0x10, 0xe5, 0xfd, 0x4d, 0x3e, 0x89, 0x45,
	0xd5, 0xc1, 0x46, 0xa7, 0xc8, 0xae, 0xf5, 0xc5, 0xbe, 0xea, 0xcb, 0xf5, 0xba, 0xd9, 0x5a, 0xdd,
	0x6f, 0xb3, 0x83, 0x4f, 0xc1, 0x6a, 0x41, 0x23, 0xb7, 0x45, 0x09, 0xf9, 0x1e, 0xdc, 0x09, 0xf4,
	0x45, 0x5a, 0x70, 0x0e, 0x71, 0x57, 0x1b, 0x78, 0xaf, 0x7a, 0x9e, 0x4b, 0xbb, 0x56, 0x0f, 0x51,
	0xdf, 0xab, 0x15, 0xa8, 0xef, 0x7f, 0xb5, 0x80, 0xe1, 0xc0, 0xff, 0x7d, 0x79, 0xad, 0xef, 0x16,
	0xec, 0x1c, 0x4f, 0x07, 0xd3, 0x23, 0x3d, 0x8b, 0xd3, 0xa1, 0xcc, 0xb6, 0xb8, 0xf4, 0xbb, 0x50,
	0x09, 0x23, 0x4a, 0xae, 0x2a, 0x2a, 0x61, 0x54, 0x1e, 0x83, 0xea, 0xa6, 0x63, 0xe0, 0x43, 0x4d,
	0xe8, 0xbc, 0x48, 0x96, 0x15, 0x67, 0xb4, 0xc4, 0xad, 0x6f, 0x16, 0xd4, 0x30, 0xa5, 0x2d, 0xdf,
	0xa0, 0xbf, 0x9d, 0xce, 0x0f, 0x0b, 0x76, 0x31, 0x9d, 0x5e, 0x10, 0x4c, 0x5f, 0x87, 0xc3, 0x97,
	0xfd, 0xc3, 0xff, 0x20, 0xa9, 0x77, 0xd0, 0xc0, 0xf7, 0x83, 0xa6, 0x36, 0xd0, 0xd9, 0xe7, 0xed,
	0x26, 0xd7, 0x04, 0x27, 0x90, 0xb9, 0x34, 0x99, 0xd9, 0xc2, 0xa0, 0xf7, 0x0e, 0xfd, 0x83, 0x74,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x87, 0x66, 0xaa, 0xcd, 0x67, 0x06, 0x00, 0x00,
}
