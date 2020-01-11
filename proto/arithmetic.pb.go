// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arithmetic.proto

package arithmetic

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

type Input struct {
	X                    float64  `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_dae331f0499ff811, []int{0}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Input) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type InputMultiple struct {
	X                    []float64 `protobuf:"fixed64,1,rep,packed,name=x,proto3" json:"x,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InputMultiple) Reset()         { *m = InputMultiple{} }
func (m *InputMultiple) String() string { return proto.CompactTextString(m) }
func (*InputMultiple) ProtoMessage()    {}
func (*InputMultiple) Descriptor() ([]byte, []int) {
	return fileDescriptor_dae331f0499ff811, []int{1}
}

func (m *InputMultiple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputMultiple.Unmarshal(m, b)
}
func (m *InputMultiple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputMultiple.Marshal(b, m, deterministic)
}
func (m *InputMultiple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputMultiple.Merge(m, src)
}
func (m *InputMultiple) XXX_Size() int {
	return xxx_messageInfo_InputMultiple.Size(m)
}
func (m *InputMultiple) XXX_DiscardUnknown() {
	xxx_messageInfo_InputMultiple.DiscardUnknown(m)
}

var xxx_messageInfo_InputMultiple proto.InternalMessageInfo

func (m *InputMultiple) GetX() []float64 {
	if m != nil {
		return m.X
	}
	return nil
}

type Result struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_dae331f0499ff811, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Input)(nil), "arithmetic.Input")
	proto.RegisterType((*InputMultiple)(nil), "arithmetic.InputMultiple")
	proto.RegisterType((*Result)(nil), "arithmetic.Result")
}

func init() { proto.RegisterFile("arithmetic.proto", fileDescriptor_dae331f0499ff811) }

var fileDescriptor_dae331f0499ff811 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0xca, 0x2c,
	0xc9, 0xc8, 0x4d, 0x2d, 0xc9, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x73, 0xb1, 0x7a, 0xe6, 0x15, 0x94, 0x96, 0x08, 0xf1, 0x70, 0x31, 0x56, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x30, 0x06, 0x31, 0x56, 0x80, 0x78, 0x95, 0x12, 0x4c, 0x10, 0x5e, 0xa5, 0x92,
	0x2c, 0x17, 0x2f, 0x58, 0x91, 0x6f, 0x69, 0x4e, 0x49, 0x66, 0x41, 0x4e, 0x2a, 0x4c, 0x31, 0x33,
	0x58, 0xb1, 0x92, 0x02, 0x17, 0x5b, 0x50, 0x6a, 0x71, 0x69, 0x4e, 0x89, 0x90, 0x18, 0x17, 0x5b,
	0x11, 0x98, 0x05, 0x35, 0x09, 0xca, 0x33, 0x7a, 0xc6, 0xcc, 0xc5, 0xe5, 0x08, 0xb7, 0x54, 0x48,
	0x8f, 0x8b, 0xd9, 0x31, 0x25, 0x45, 0x48, 0x50, 0x0f, 0xc9, 0x69, 0x60, 0x0b, 0xa4, 0x84, 0x90,
	0x85, 0x20, 0x86, 0x2a, 0x31, 0x08, 0xd9, 0x71, 0x71, 0x3b, 0xa6, 0xa4, 0xc0, 0x6d, 0x97, 0xc4,
	0xd0, 0x07, 0x93, 0xc2, 0xa1, 0xdf, 0x98, 0x8b, 0x23, 0xb8, 0x34, 0xa9, 0xa4, 0x28, 0x31, 0xb9,
	0x84, 0x78, 0x4b, 0x9d, 0xb9, 0x04, 0x60, 0x9a, 0x28, 0xb2, 0x19, 0xaa, 0xa2, 0x92, 0x24, 0x9b,
	0x61, 0x9a, 0xc8, 0xb7, 0xd9, 0x90, 0x8b, 0xcd, 0x25, 0xb3, 0x2c, 0x33, 0x25, 0x95, 0x78, 0x7b,
	0x1d, 0xb9, 0xf8, 0x20, 0x5a, 0xc8, 0xb6, 0x35, 0x89, 0x0d, 0x9c, 0xc2, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x4b, 0x26, 0xf8, 0x26, 0x75, 0x02, 0x00, 0x00,
}