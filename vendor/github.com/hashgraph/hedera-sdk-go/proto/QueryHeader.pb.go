// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/QueryHeader.proto

package proto

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

// The client uses the ResponseType to request that the node send it just the answer, or both the answer and a state proof. It can also ask for just the cost for getting the answer or both. If the payment in the query fails the precheck, then the response may have some fields blank. The state proof is only available for some types of information. It is available for a Record, but not a receipt. It is available for the information in each kind of *GetInfo request.
type ResponseType int32

const (
	ResponseType_ANSWER_ONLY             ResponseType = 0
	ResponseType_ANSWER_STATE_PROOF      ResponseType = 1
	ResponseType_COST_ANSWER             ResponseType = 2
	ResponseType_COST_ANSWER_STATE_PROOF ResponseType = 3
)

var ResponseType_name = map[int32]string{
	0: "ANSWER_ONLY",
	1: "ANSWER_STATE_PROOF",
	2: "COST_ANSWER",
	3: "COST_ANSWER_STATE_PROOF",
}

var ResponseType_value = map[string]int32{
	"ANSWER_ONLY":             0,
	"ANSWER_STATE_PROOF":      1,
	"COST_ANSWER":             2,
	"COST_ANSWER_STATE_PROOF": 3,
}

func (x ResponseType) String() string {
	return proto.EnumName(ResponseType_name, int32(x))
}

func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0a7b838bafa1873, []int{0}
}

// Each query from the client to the node will contain the QueryHeader, which gives the requested response type, and includes a payment for the response. It will sometimes leave payment blank: it is blank for TransactionGetReceiptQuery. It can also be left blank when the responseType is costAnswer or costAnswerStateProof. But it needs to be filled in for all other cases. The idea is that an answer that is only a few bytes (or that was paid for earlier) can be given for free. But if the answer is something that requires many bytes or much computation (like a state proof), then it should be paid for.
type QueryHeader struct {
	Payment              *Transaction `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	ResponseType         ResponseType `protobuf:"varint,2,opt,name=responseType,proto3,enum=proto.ResponseType" json:"responseType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *QueryHeader) Reset()         { *m = QueryHeader{} }
func (m *QueryHeader) String() string { return proto.CompactTextString(m) }
func (*QueryHeader) ProtoMessage()    {}
func (*QueryHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0a7b838bafa1873, []int{0}
}

func (m *QueryHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryHeader.Unmarshal(m, b)
}
func (m *QueryHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryHeader.Marshal(b, m, deterministic)
}
func (m *QueryHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHeader.Merge(m, src)
}
func (m *QueryHeader) XXX_Size() int {
	return xxx_messageInfo_QueryHeader.Size(m)
}
func (m *QueryHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHeader.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHeader proto.InternalMessageInfo

func (m *QueryHeader) GetPayment() *Transaction {
	if m != nil {
		return m.Payment
	}
	return nil
}

func (m *QueryHeader) GetResponseType() ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return ResponseType_ANSWER_ONLY
}

func init() {
	proto.RegisterEnum("proto.ResponseType", ResponseType_name, ResponseType_value)
	proto.RegisterType((*QueryHeader)(nil), "proto.QueryHeader")
}

func init() { proto.RegisterFile("proto/QueryHeader.proto", fileDescriptor_d0a7b838bafa1873) }

var fileDescriptor_d0a7b838bafa1873 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0x87, 0xcd, 0x44, 0x85, 0x6c, 0x68, 0x89, 0xe0, 0xc6, 0xbc, 0x0c, 0x4f, 0x45, 0x5c, 0x0a,
	0xf3, 0xe0, 0x79, 0xca, 0x64, 0x07, 0x59, 0x67, 0x1a, 0x10, 0xbd, 0x94, 0xac, 0x7d, 0x34, 0x43,
	0xda, 0x84, 0x24, 0x3b, 0xf4, 0xbf, 0x17, 0x9b, 0x2a, 0xd9, 0xe9, 0xc1, 0xef, 0xfb, 0x3e, 0x78,
	0x78, 0xac, 0x8d, 0x72, 0x2a, 0x79, 0x3f, 0x80, 0x69, 0xd7, 0x20, 0x4a, 0x30, 0xb4, 0x5b, 0xc8,
	0x59, 0x77, 0xa6, 0x3d, 0xe7, 0x46, 0x34, 0x56, 0x14, 0x6e, 0xaf, 0x1a, 0xcf, 0xef, 0x1c, 0x1e,
	0x06, 0x11, 0x79, 0xc0, 0x17, 0x5a, 0xb4, 0x35, 0x34, 0x6e, 0x82, 0x66, 0x28, 0x1e, 0x2e, 0x88,
	0xf7, 0x68, 0x50, 0xb2, 0x3f, 0x85, 0x3c, 0xe1, 0x91, 0x01, 0xab, 0x55, 0x63, 0x81, 0xb7, 0x1a,
	0x26, 0x83, 0x19, 0x8a, 0x2f, 0x17, 0xd7, 0x7d, 0xc2, 0x02, 0xc4, 0x8e, 0xc4, 0x7b, 0xc0, 0xa3,
	0x90, 0x92, 0x2b, 0x3c, 0x5c, 0x6e, 0xb2, 0x8f, 0x15, 0xcb, 0xd3, 0xcd, 0xdb, 0x67, 0x74, 0x42,
	0x6e, 0x30, 0xe9, 0x87, 0x8c, 0x2f, 0xf9, 0x2a, 0xdf, 0xb2, 0x34, 0x7d, 0x8d, 0xd0, 0xaf, 0xf8,
	0x92, 0x66, 0x3c, 0xf7, 0x30, 0x1a, 0x90, 0x5b, 0x3c, 0x0e, 0x86, 0x23, 0xfb, 0xf4, 0x79, 0x8d,
	0xa7, 0x85, 0xaa, 0xa9, 0x84, 0x12, 0x8c, 0xa0, 0x52, 0x58, 0x59, 0x19, 0xa1, 0xa5, 0xff, 0x6f,
	0x8b, 0xbe, 0xe2, 0x6a, 0xef, 0xe4, 0x61, 0x47, 0x0b, 0x55, 0x27, 0xff, 0x34, 0xf1, 0xfa, 0xdc,
	0x96, 0xdf, 0xf3, 0x4a, 0x25, 0x9d, 0xbb, 0x3b, 0xef, 0xce, 0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x09, 0xdf, 0xb2, 0x6e, 0x68, 0x01, 0x00, 0x00,
}
