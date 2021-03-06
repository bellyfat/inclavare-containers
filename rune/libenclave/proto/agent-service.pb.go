// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent-service.proto

package libenclave_proto

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

type AgentServiceRequest struct {
	Exec                 *AgentServiceRequest_Execute `protobuf:"bytes,1,opt,name=exec,proto3" json:"exec,omitempty"`
	Kill                 *AgentServiceRequest_Kill    `protobuf:"bytes,2,opt,name=kill,proto3" json:"kill,omitempty"`
	Attest               *AgentServiceRequest_Attest  `protobuf:"bytes,3,opt,name=attest,proto3" json:"attest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AgentServiceRequest) Reset()         { *m = AgentServiceRequest{} }
func (m *AgentServiceRequest) String() string { return proto.CompactTextString(m) }
func (*AgentServiceRequest) ProtoMessage()    {}
func (*AgentServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_145c985d4df28d67, []int{0}
}

func (m *AgentServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentServiceRequest.Unmarshal(m, b)
}
func (m *AgentServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentServiceRequest.Marshal(b, m, deterministic)
}
func (m *AgentServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentServiceRequest.Merge(m, src)
}
func (m *AgentServiceRequest) XXX_Size() int {
	return xxx_messageInfo_AgentServiceRequest.Size(m)
}
func (m *AgentServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentServiceRequest proto.InternalMessageInfo

func (m *AgentServiceRequest) GetExec() *AgentServiceRequest_Execute {
	if m != nil {
		return m.Exec
	}
	return nil
}

func (m *AgentServiceRequest) GetKill() *AgentServiceRequest_Kill {
	if m != nil {
		return m.Kill
	}
	return nil
}

func (m *AgentServiceRequest) GetAttest() *AgentServiceRequest_Attest {
	if m != nil {
		return m.Attest
	}
	return nil
}

type AgentServiceRequest_Execute struct {
	Argv                 string   `protobuf:"bytes,1,opt,name=argv,proto3" json:"argv,omitempty"`
	Envp                 string   `protobuf:"bytes,2,opt,name=envp,proto3" json:"envp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentServiceRequest_Execute) Reset()         { *m = AgentServiceRequest_Execute{} }
func (m *AgentServiceRequest_Execute) String() string { return proto.CompactTextString(m) }
func (*AgentServiceRequest_Execute) ProtoMessage()    {}
func (*AgentServiceRequest_Execute) Descriptor() ([]byte, []int) {
	return fileDescriptor_145c985d4df28d67, []int{0, 0}
}

func (m *AgentServiceRequest_Execute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentServiceRequest_Execute.Unmarshal(m, b)
}
func (m *AgentServiceRequest_Execute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentServiceRequest_Execute.Marshal(b, m, deterministic)
}
func (m *AgentServiceRequest_Execute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentServiceRequest_Execute.Merge(m, src)
}
func (m *AgentServiceRequest_Execute) XXX_Size() int {
	return xxx_messageInfo_AgentServiceRequest_Execute.Size(m)
}
func (m *AgentServiceRequest_Execute) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentServiceRequest_Execute.DiscardUnknown(m)
}

var xxx_messageInfo_AgentServiceRequest_Execute proto.InternalMessageInfo

func (m *AgentServiceRequest_Execute) GetArgv() string {
	if m != nil {
		return m.Argv
	}
	return ""
}

func (m *AgentServiceRequest_Execute) GetEnvp() string {
	if m != nil {
		return m.Envp
	}
	return ""
}

type AgentServiceRequest_Kill struct {
	Sig                  int32    `protobuf:"varint,1,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentServiceRequest_Kill) Reset()         { *m = AgentServiceRequest_Kill{} }
func (m *AgentServiceRequest_Kill) String() string { return proto.CompactTextString(m) }
func (*AgentServiceRequest_Kill) ProtoMessage()    {}
func (*AgentServiceRequest_Kill) Descriptor() ([]byte, []int) {
	return fileDescriptor_145c985d4df28d67, []int{0, 1}
}

func (m *AgentServiceRequest_Kill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentServiceRequest_Kill.Unmarshal(m, b)
}
func (m *AgentServiceRequest_Kill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentServiceRequest_Kill.Marshal(b, m, deterministic)
}
func (m *AgentServiceRequest_Kill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentServiceRequest_Kill.Merge(m, src)
}
func (m *AgentServiceRequest_Kill) XXX_Size() int {
	return xxx_messageInfo_AgentServiceRequest_Kill.Size(m)
}
func (m *AgentServiceRequest_Kill) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentServiceRequest_Kill.DiscardUnknown(m)
}

var xxx_messageInfo_AgentServiceRequest_Kill proto.InternalMessageInfo

func (m *AgentServiceRequest_Kill) GetSig() int32 {
	if m != nil {
		return m.Sig
	}
	return 0
}

type AgentServiceRequest_Attest struct {
	Spid                 string   `protobuf:"bytes,1,opt,name=spid,proto3" json:"spid,omitempty"`
	SubscriptionKey      string   `protobuf:"bytes,2,opt,name=subscriptionKey,proto3" json:"subscriptionKey,omitempty"`
	Product              uint32   `protobuf:"varint,3,opt,name=product,proto3" json:"product,omitempty"`
	QuoteType            uint32   `protobuf:"varint,4,opt,name=quoteType,proto3" json:"quoteType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentServiceRequest_Attest) Reset()         { *m = AgentServiceRequest_Attest{} }
func (m *AgentServiceRequest_Attest) String() string { return proto.CompactTextString(m) }
func (*AgentServiceRequest_Attest) ProtoMessage()    {}
func (*AgentServiceRequest_Attest) Descriptor() ([]byte, []int) {
	return fileDescriptor_145c985d4df28d67, []int{0, 2}
}

func (m *AgentServiceRequest_Attest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentServiceRequest_Attest.Unmarshal(m, b)
}
func (m *AgentServiceRequest_Attest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentServiceRequest_Attest.Marshal(b, m, deterministic)
}
func (m *AgentServiceRequest_Attest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentServiceRequest_Attest.Merge(m, src)
}
func (m *AgentServiceRequest_Attest) XXX_Size() int {
	return xxx_messageInfo_AgentServiceRequest_Attest.Size(m)
}
func (m *AgentServiceRequest_Attest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentServiceRequest_Attest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentServiceRequest_Attest proto.InternalMessageInfo

func (m *AgentServiceRequest_Attest) GetSpid() string {
	if m != nil {
		return m.Spid
	}
	return ""
}

func (m *AgentServiceRequest_Attest) GetSubscriptionKey() string {
	if m != nil {
		return m.SubscriptionKey
	}
	return ""
}

func (m *AgentServiceRequest_Attest) GetProduct() uint32 {
	if m != nil {
		return m.Product
	}
	return 0
}

func (m *AgentServiceRequest_Attest) GetQuoteType() uint32 {
	if m != nil {
		return m.QuoteType
	}
	return 0
}

type AgentServiceResponse struct {
	Exec                 *AgentServiceResponse_Execute `protobuf:"bytes,1,opt,name=exec,proto3" json:"exec,omitempty"`
	Attest               *AgentServiceResponse_Attest  `protobuf:"bytes,2,opt,name=attest,proto3" json:"attest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AgentServiceResponse) Reset()         { *m = AgentServiceResponse{} }
func (m *AgentServiceResponse) String() string { return proto.CompactTextString(m) }
func (*AgentServiceResponse) ProtoMessage()    {}
func (*AgentServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_145c985d4df28d67, []int{1}
}

func (m *AgentServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentServiceResponse.Unmarshal(m, b)
}
func (m *AgentServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentServiceResponse.Marshal(b, m, deterministic)
}
func (m *AgentServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentServiceResponse.Merge(m, src)
}
func (m *AgentServiceResponse) XXX_Size() int {
	return xxx_messageInfo_AgentServiceResponse.Size(m)
}
func (m *AgentServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AgentServiceResponse proto.InternalMessageInfo

func (m *AgentServiceResponse) GetExec() *AgentServiceResponse_Execute {
	if m != nil {
		return m.Exec
	}
	return nil
}

func (m *AgentServiceResponse) GetAttest() *AgentServiceResponse_Attest {
	if m != nil {
		return m.Attest
	}
	return nil
}

type AgentServiceResponse_Execute struct {
	ExitCode             int32    `protobuf:"varint,1,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentServiceResponse_Execute) Reset()         { *m = AgentServiceResponse_Execute{} }
func (m *AgentServiceResponse_Execute) String() string { return proto.CompactTextString(m) }
func (*AgentServiceResponse_Execute) ProtoMessage()    {}
func (*AgentServiceResponse_Execute) Descriptor() ([]byte, []int) {
	return fileDescriptor_145c985d4df28d67, []int{1, 0}
}

func (m *AgentServiceResponse_Execute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentServiceResponse_Execute.Unmarshal(m, b)
}
func (m *AgentServiceResponse_Execute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentServiceResponse_Execute.Marshal(b, m, deterministic)
}
func (m *AgentServiceResponse_Execute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentServiceResponse_Execute.Merge(m, src)
}
func (m *AgentServiceResponse_Execute) XXX_Size() int {
	return xxx_messageInfo_AgentServiceResponse_Execute.Size(m)
}
func (m *AgentServiceResponse_Execute) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentServiceResponse_Execute.DiscardUnknown(m)
}

var xxx_messageInfo_AgentServiceResponse_Execute proto.InternalMessageInfo

func (m *AgentServiceResponse_Execute) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *AgentServiceResponse_Execute) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type AgentServiceResponse_Attest struct {
	ExitCode             int32    `protobuf:"varint,1,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	LocalReport          []byte   `protobuf:"bytes,3,opt,name=localReport,proto3" json:"localReport,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentServiceResponse_Attest) Reset()         { *m = AgentServiceResponse_Attest{} }
func (m *AgentServiceResponse_Attest) String() string { return proto.CompactTextString(m) }
func (*AgentServiceResponse_Attest) ProtoMessage()    {}
func (*AgentServiceResponse_Attest) Descriptor() ([]byte, []int) {
	return fileDescriptor_145c985d4df28d67, []int{1, 1}
}

func (m *AgentServiceResponse_Attest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentServiceResponse_Attest.Unmarshal(m, b)
}
func (m *AgentServiceResponse_Attest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentServiceResponse_Attest.Marshal(b, m, deterministic)
}
func (m *AgentServiceResponse_Attest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentServiceResponse_Attest.Merge(m, src)
}
func (m *AgentServiceResponse_Attest) XXX_Size() int {
	return xxx_messageInfo_AgentServiceResponse_Attest.Size(m)
}
func (m *AgentServiceResponse_Attest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentServiceResponse_Attest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentServiceResponse_Attest proto.InternalMessageInfo

func (m *AgentServiceResponse_Attest) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *AgentServiceResponse_Attest) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AgentServiceResponse_Attest) GetLocalReport() []byte {
	if m != nil {
		return m.LocalReport
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentServiceRequest)(nil), "libenclave_proto.AgentServiceRequest")
	proto.RegisterType((*AgentServiceRequest_Execute)(nil), "libenclave_proto.AgentServiceRequest.Execute")
	proto.RegisterType((*AgentServiceRequest_Kill)(nil), "libenclave_proto.AgentServiceRequest.Kill")
	proto.RegisterType((*AgentServiceRequest_Attest)(nil), "libenclave_proto.AgentServiceRequest.Attest")
	proto.RegisterType((*AgentServiceResponse)(nil), "libenclave_proto.AgentServiceResponse")
	proto.RegisterType((*AgentServiceResponse_Execute)(nil), "libenclave_proto.AgentServiceResponse.Execute")
	proto.RegisterType((*AgentServiceResponse_Attest)(nil), "libenclave_proto.AgentServiceResponse.Attest")
}

func init() {
	proto.RegisterFile("agent-service.proto", fileDescriptor_145c985d4df28d67)
}

var fileDescriptor_145c985d4df28d67 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x51, 0x57, 0xcd, 0x67, 0x91, 0x8c, 0x1e, 0x96, 0xa5, 0x83, 0x78, 0x92, 0xc8, 0x85,
	0xea, 0x18, 0x04, 0x56, 0x9e, 0xbc, 0x4d, 0x1d, 0x83, 0x58, 0xc7, 0x87, 0x0c, 0x0d, 0x3b, 0xe3,
	0xcc, 0xec, 0xa2, 0x97, 0xbe, 0x50, 0x5f, 0xa1, 0x0f, 0x17, 0x3b, 0xbb, 0x6a, 0x2e, 0x41, 0xdb,
	0xed, 0xcd, 0x7f, 0xde, 0xef, 0xcf, 0x7b, 0xff, 0x07, 0xfd, 0x68, 0x85, 0xb1, 0x9d, 0x18, 0xd4,
	0x29, 0x67, 0x18, 0x2a, 0x2d, 0xad, 0x24, 0x3d, 0xc1, 0x17, 0x18, 0x33, 0x11, 0xa5, 0xf8, 0xe6,
	0x94, 0xd1, 0x67, 0x03, 0xfa, 0xd3, 0xac, 0xf3, 0x39, 0x6f, 0xa4, 0xb8, 0x4e, 0xd0, 0x58, 0x32,
	0x05, 0x0f, 0x37, 0xc8, 0xfc, 0xda, 0xb0, 0x36, 0xee, 0xde, 0x4c, 0xc2, 0x32, 0x18, 0xfe, 0x02,
	0x85, 0xb3, 0x0d, 0xb2, 0xc4, 0x22, 0x75, 0x28, 0xb9, 0x07, 0xef, 0x9d, 0x0b, 0xe1, 0xd7, 0x9d,
	0xc5, 0x65, 0x35, 0x8b, 0x39, 0x17, 0x82, 0x3a, 0x8e, 0x3c, 0x41, 0x2b, 0xb2, 0x16, 0x8d, 0xf5,
	0x1b, 0xce, 0xe1, 0xaa, 0x9a, 0xc3, 0xd4, 0x31, 0xb4, 0x60, 0x83, 0x6b, 0x68, 0x17, 0x63, 0x11,
	0x02, 0x5e, 0xa4, 0x57, 0xa9, 0xdb, 0xa9, 0x43, 0x5d, 0x9d, 0x69, 0x18, 0xa7, 0xca, 0x0d, 0xd9,
	0xa1, 0xae, 0x0e, 0x7c, 0xf0, 0xb2, 0x31, 0x48, 0x0f, 0x1a, 0x86, 0xaf, 0x5c, 0x7b, 0x93, 0x66,
	0x65, 0xf0, 0x01, 0xad, 0xdc, 0x3e, 0xe3, 0x8c, 0xe2, 0xcb, 0x9d, 0x57, 0x56, 0x93, 0x31, 0x9c,
	0x9b, 0x64, 0x61, 0x98, 0xe6, 0xca, 0x72, 0x19, 0xcf, 0x71, 0x5b, 0xd8, 0x96, 0x65, 0xe2, 0x43,
	0x5b, 0x69, 0xb9, 0x4c, 0x58, 0xbe, 0xdb, 0x19, 0xdd, 0x3d, 0xc9, 0x05, 0x74, 0xd6, 0x89, 0xb4,
	0xf8, 0xb2, 0x55, 0xe8, 0x7b, 0xee, 0xef, 0x20, 0x8c, 0xbe, 0xea, 0x30, 0x38, 0xde, 0xd9, 0x28,
	0x19, 0x1b, 0x24, 0x0f, 0x47, 0xe7, 0x0a, 0xff, 0x4a, 0x2a, 0xa7, 0x4a, 0xf7, 0x9a, 0xed, 0xf3,
	0xae, 0x57, 0x3b, 0x7a, 0xe1, 0x52, 0x0a, 0xfc, 0xee, 0x10, 0x78, 0x00, 0x27, 0xb8, 0xe1, 0xf6,
	0x51, 0x2e, 0xb1, 0x48, 0x71, 0xff, 0x26, 0x03, 0x68, 0xa2, 0xd6, 0x52, 0x17, 0x11, 0xe5, 0x8f,
	0xe0, 0x75, 0x1f, 0xf0, 0xbf, 0x59, 0x32, 0x84, 0xae, 0x90, 0x2c, 0x12, 0x14, 0x95, 0xd4, 0x79,
	0xb0, 0xa7, 0xf4, 0xa7, 0xb4, 0x68, 0xb9, 0x2d, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x63,
	0x98, 0xf5, 0x2d, 0x1c, 0x03, 0x00, 0x00,
}
