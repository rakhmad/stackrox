// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/api_token_service.proto

package v1

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	storage "github.com/stackrox/rox/generated/storage"
	io "io"
	math "math"
	math_bits "math/bits"
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

type GenerateTokenRequest struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Role                 string           `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"` // Deprecated: Do not use.
	Roles                []string         `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Expiration           *types.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GenerateTokenRequest) Reset()         { *m = GenerateTokenRequest{} }
func (m *GenerateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateTokenRequest) ProtoMessage()    {}
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cff1ec89435902, []int{0}
}
func (m *GenerateTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenerateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenerateTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenerateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateTokenRequest.Merge(m, src)
}
func (m *GenerateTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *GenerateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateTokenRequest proto.InternalMessageInfo

func (m *GenerateTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Deprecated: Do not use.
func (m *GenerateTokenRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *GenerateTokenRequest) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *GenerateTokenRequest) GetExpiration() *types.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *GenerateTokenRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GenerateTokenRequest) Clone() *GenerateTokenRequest {
	if m == nil {
		return nil
	}
	cloned := new(GenerateTokenRequest)
	*cloned = *m

	if m.Roles != nil {
		cloned.Roles = make([]string, len(m.Roles))
		copy(cloned.Roles, m.Roles)
	}
	cloned.Expiration = m.Expiration.Clone()
	return cloned
}

type GenerateTokenResponse struct {
	Token                string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Metadata             *storage.TokenMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GenerateTokenResponse) Reset()         { *m = GenerateTokenResponse{} }
func (m *GenerateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateTokenResponse) ProtoMessage()    {}
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cff1ec89435902, []int{1}
}
func (m *GenerateTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenerateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenerateTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenerateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateTokenResponse.Merge(m, src)
}
func (m *GenerateTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *GenerateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateTokenResponse proto.InternalMessageInfo

func (m *GenerateTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GenerateTokenResponse) GetMetadata() *storage.TokenMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *GenerateTokenResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GenerateTokenResponse) Clone() *GenerateTokenResponse {
	if m == nil {
		return nil
	}
	cloned := new(GenerateTokenResponse)
	*cloned = *m

	cloned.Metadata = m.Metadata.Clone()
	return cloned
}

type GetAPITokensRequest struct {
	// Types that are valid to be assigned to RevokedOneof:
	//	*GetAPITokensRequest_Revoked
	RevokedOneof         isGetAPITokensRequest_RevokedOneof `protobuf_oneof:"revoked_oneof"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GetAPITokensRequest) Reset()         { *m = GetAPITokensRequest{} }
func (m *GetAPITokensRequest) String() string { return proto.CompactTextString(m) }
func (*GetAPITokensRequest) ProtoMessage()    {}
func (*GetAPITokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cff1ec89435902, []int{2}
}
func (m *GetAPITokensRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAPITokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAPITokensRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAPITokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAPITokensRequest.Merge(m, src)
}
func (m *GetAPITokensRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAPITokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAPITokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAPITokensRequest proto.InternalMessageInfo

type isGetAPITokensRequest_RevokedOneof interface {
	isGetAPITokensRequest_RevokedOneof()
	MarshalTo([]byte) (int, error)
	Size() int
	Clone() isGetAPITokensRequest_RevokedOneof
}

type GetAPITokensRequest_Revoked struct {
	Revoked bool `protobuf:"varint,1,opt,name=revoked,proto3,oneof" json:"revoked,omitempty"`
}

func (*GetAPITokensRequest_Revoked) isGetAPITokensRequest_RevokedOneof() {}
func (m *GetAPITokensRequest_Revoked) Clone() isGetAPITokensRequest_RevokedOneof {
	if m == nil {
		return nil
	}
	cloned := new(GetAPITokensRequest_Revoked)
	*cloned = *m

	return cloned
}

func (m *GetAPITokensRequest) GetRevokedOneof() isGetAPITokensRequest_RevokedOneof {
	if m != nil {
		return m.RevokedOneof
	}
	return nil
}

func (m *GetAPITokensRequest) GetRevoked() bool {
	if x, ok := m.GetRevokedOneof().(*GetAPITokensRequest_Revoked); ok {
		return x.Revoked
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetAPITokensRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetAPITokensRequest_Revoked)(nil),
	}
}

func (m *GetAPITokensRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetAPITokensRequest) Clone() *GetAPITokensRequest {
	if m == nil {
		return nil
	}
	cloned := new(GetAPITokensRequest)
	*cloned = *m

	if m.RevokedOneof != nil {
		cloned.RevokedOneof = m.RevokedOneof.Clone()
	}
	return cloned
}

type GetAPITokensResponse struct {
	Tokens               []*storage.TokenMetadata `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetAPITokensResponse) Reset()         { *m = GetAPITokensResponse{} }
func (m *GetAPITokensResponse) String() string { return proto.CompactTextString(m) }
func (*GetAPITokensResponse) ProtoMessage()    {}
func (*GetAPITokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cff1ec89435902, []int{3}
}
func (m *GetAPITokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAPITokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAPITokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAPITokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAPITokensResponse.Merge(m, src)
}
func (m *GetAPITokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAPITokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAPITokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAPITokensResponse proto.InternalMessageInfo

func (m *GetAPITokensResponse) GetTokens() []*storage.TokenMetadata {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *GetAPITokensResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetAPITokensResponse) Clone() *GetAPITokensResponse {
	if m == nil {
		return nil
	}
	cloned := new(GetAPITokensResponse)
	*cloned = *m

	if m.Tokens != nil {
		cloned.Tokens = make([]*storage.TokenMetadata, len(m.Tokens))
		for idx, v := range m.Tokens {
			cloned.Tokens[idx] = v.Clone()
		}
	}
	return cloned
}

type ListAllowedTokenRolesResponse struct {
	RoleNames            []string `protobuf:"bytes,1,rep,name=roleNames,proto3" json:"roleNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAllowedTokenRolesResponse) Reset()         { *m = ListAllowedTokenRolesResponse{} }
func (m *ListAllowedTokenRolesResponse) String() string { return proto.CompactTextString(m) }
func (*ListAllowedTokenRolesResponse) ProtoMessage()    {}
func (*ListAllowedTokenRolesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cff1ec89435902, []int{4}
}
func (m *ListAllowedTokenRolesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAllowedTokenRolesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAllowedTokenRolesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAllowedTokenRolesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllowedTokenRolesResponse.Merge(m, src)
}
func (m *ListAllowedTokenRolesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListAllowedTokenRolesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllowedTokenRolesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllowedTokenRolesResponse proto.InternalMessageInfo

func (m *ListAllowedTokenRolesResponse) GetRoleNames() []string {
	if m != nil {
		return m.RoleNames
	}
	return nil
}

func (m *ListAllowedTokenRolesResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ListAllowedTokenRolesResponse) Clone() *ListAllowedTokenRolesResponse {
	if m == nil {
		return nil
	}
	cloned := new(ListAllowedTokenRolesResponse)
	*cloned = *m

	if m.RoleNames != nil {
		cloned.RoleNames = make([]string, len(m.RoleNames))
		copy(cloned.RoleNames, m.RoleNames)
	}
	return cloned
}

func init() {
	proto.RegisterType((*GenerateTokenRequest)(nil), "v1.GenerateTokenRequest")
	proto.RegisterType((*GenerateTokenResponse)(nil), "v1.GenerateTokenResponse")
	proto.RegisterType((*GetAPITokensRequest)(nil), "v1.GetAPITokensRequest")
	proto.RegisterType((*GetAPITokensResponse)(nil), "v1.GetAPITokensResponse")
	proto.RegisterType((*ListAllowedTokenRolesResponse)(nil), "v1.ListAllowedTokenRolesResponse")
}

func init() { proto.RegisterFile("api/v1/api_token_service.proto", fileDescriptor_b5cff1ec89435902) }

var fileDescriptor_b5cff1ec89435902 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdf, 0x6e, 0x12, 0x4f,
	0x14, 0xc7, 0xbb, 0xc0, 0xaf, 0xbf, 0x32, 0x48, 0x30, 0x53, 0xa0, 0xdb, 0xb5, 0x22, 0xdd, 0x18,
	0x43, 0x8c, 0xce, 0x06, 0xbc, 0x30, 0xa9, 0xf1, 0xa2, 0x44, 0xad, 0x4d, 0xac, 0x31, 0xdb, 0x5e,
	0x34, 0xde, 0x90, 0x29, 0x9c, 0x92, 0x0d, 0xec, 0xce, 0xba, 0x33, 0xac, 0x6d, 0x8c, 0x37, 0xbe,
	0x82, 0x5e, 0xf8, 0x48, 0x5e, 0x9a, 0xf8, 0x02, 0x06, 0xbd, 0xf5, 0x1d, 0xcc, 0xfc, 0xd9, 0x02,
	0x96, 0xf6, 0x8a, 0x99, 0xf3, 0x3d, 0xfb, 0x9d, 0xcf, 0x39, 0x73, 0x06, 0xd4, 0xa0, 0x71, 0xe0,
	0xa5, 0x6d, 0x8f, 0xc6, 0x41, 0x4f, 0xb0, 0x11, 0x44, 0x3d, 0x0e, 0x49, 0x1a, 0xf4, 0x81, 0xc4,
	0x09, 0x13, 0x0c, 0xe7, 0xd2, 0xb6, 0xb3, 0x6e, 0x72, 0xfa, 0x2c, 0x0c, 0x59, 0xa4, 0x05, 0x07,
	0x9b, 0x20, 0x84, 0xb1, 0x38, 0x37, 0xb1, 0xad, 0x21, 0x63, 0xc3, 0x31, 0x48, 0x33, 0x8f, 0x46,
	0x11, 0x13, 0x54, 0x04, 0x2c, 0xe2, 0x46, 0xbd, 0x63, 0x54, 0xb5, 0x3b, 0x99, 0x9c, 0x7a, 0x22,
	0x08, 0x81, 0x0b, 0x1a, 0xc6, 0x26, 0x61, 0x83, 0x0b, 0x96, 0xd0, 0x21, 0xcc, 0x60, 0xb4, 0xe0,
	0x7e, 0xb1, 0x50, 0x75, 0x0f, 0x22, 0x48, 0xa8, 0x80, 0x23, 0x19, 0xf7, 0xe1, 0xdd, 0x04, 0xb8,
	0xc0, 0x18, 0x15, 0x22, 0x1a, 0x82, 0x6d, 0x35, 0xad, 0x56, 0xd1, 0x57, 0x6b, 0x5c, 0x47, 0x85,
	0x84, 0x8d, 0xc1, 0xce, 0xc9, 0x58, 0x37, 0x67, 0x5b, 0xbe, 0xda, 0xe3, 0x2a, 0xfa, 0x4f, 0xfe,
	0x72, 0x3b, 0xdf, 0xcc, 0xb7, 0x8a, 0xbe, 0xde, 0xe0, 0x1d, 0x84, 0xe0, 0x2c, 0x0e, 0x12, 0x45,
	0x6a, 0x17, 0x9a, 0x56, 0xab, 0xd4, 0x71, 0x88, 0x26, 0x25, 0x19, 0x29, 0x39, 0xca, 0x48, 0xfd,
	0xb9, 0x6c, 0x97, 0xa2, 0xda, 0x3f, 0x54, 0x3c, 0x66, 0x11, 0x57, 0x47, 0x29, 0x7c, 0xc3, 0xa5,
	0x37, 0xb8, 0x83, 0xd6, 0x42, 0x10, 0x74, 0x40, 0x05, 0x55, 0x70, 0xa5, 0x4e, 0x9d, 0x98, 0x8a,
	0x89, 0xfa, 0xfe, 0xc0, 0xa8, 0xfe, 0x45, 0x9e, 0xdb, 0x45, 0xeb, 0x7b, 0x20, 0x76, 0xdf, 0xec,
	0xab, 0x04, 0x9e, 0xd5, 0xed, 0xa0, 0xff, 0x13, 0x48, 0xd9, 0x08, 0x06, 0xea, 0x88, 0xb5, 0x97,
	0x2b, 0x7e, 0x16, 0xe8, 0x56, 0x50, 0xd9, 0x2c, 0x7b, 0x2c, 0x02, 0x76, 0xea, 0xbe, 0x90, 0xcd,
	0x9b, 0xf7, 0x30, 0x94, 0x04, 0xad, 0x2a, 0x30, 0x6e, 0x5b, 0xcd, 0xfc, 0x35, 0x34, 0x26, 0xcb,
	0x7d, 0x8a, 0x6e, 0xbf, 0x0a, 0xb8, 0xd8, 0x1d, 0x8f, 0xd9, 0x7b, 0x18, 0xe8, 0x8a, 0x65, 0x0f,
	0x2f, 0x0c, 0xb7, 0x50, 0x51, 0x36, 0xf5, 0x35, 0x0d, 0x41, 0x7b, 0x16, 0xfd, 0x59, 0xa0, 0xf3,
	0x27, 0x8f, 0x2a, 0x19, 0xc4, 0xa1, 0x9e, 0x31, 0x7c, 0x88, 0x4a, 0x73, 0x68, 0xf8, 0x26, 0x49,
	0xdb, 0xc4, 0x07, 0xce, 0x26, 0x49, 0x1f, 0xba, 0xe7, 0xfb, 0xcf, 0x9c, 0x2b, 0x98, 0x5c, 0xe7,
	0xd3, 0x8f, 0xdf, 0x9f, 0x73, 0x55, 0x8c, 0xcd, 0xec, 0x6a, 0x44, 0xef, 0x43, 0x30, 0xf8, 0x88,
	0x8f, 0xd1, 0x8d, 0xf9, 0x7a, 0xf1, 0x86, 0x74, 0x5d, 0xd2, 0x45, 0xc7, 0xbe, 0x2c, 0xe8, 0x4a,
	0xdc, 0x9a, 0xb2, 0xaf, 0xe0, 0xf2, 0x82, 0x3d, 0x1e, 0xa2, 0xf2, 0xc2, 0x85, 0x63, 0xe3, 0x70,
	0x79, 0x32, 0x9d, 0xcd, 0x25, 0x8a, 0x31, 0xdf, 0x56, 0xe6, 0xb7, 0xdc, 0xfa, 0x22, 0xfb, 0xd0,
	0x24, 0xef, 0x58, 0xf7, 0xf1, 0x01, 0x2a, 0xf9, 0xea, 0x0e, 0xaf, 0xea, 0x4b, 0x51, 0x46, 0x9e,
	0xcb, 0xa7, 0x97, 0xd9, 0x75, 0x36, 0x17, 0xed, 0xf4, 0x0c, 0xe8, 0x8e, 0x70, 0x54, 0x5b, 0x7a,
	0x73, 0x78, 0x66, 0xe3, 0x6c, 0xcb, 0xe5, 0xb5, 0xf7, 0xeb, 0x3e, 0x50, 0x27, 0xdd, 0xc3, 0x77,
	0x97, 0x83, 0x7b, 0x54, 0x7f, 0xf9, 0x50, 0xbd, 0xac, 0xee, 0xe3, 0x6f, 0xd3, 0x86, 0xf5, 0x7d,
	0xda, 0xb0, 0x7e, 0x4e, 0x1b, 0xd6, 0xd7, 0x5f, 0x8d, 0x15, 0x64, 0x07, 0x8c, 0x70, 0x41, 0xfb,
	0xa3, 0x84, 0x9d, 0xe9, 0xb7, 0x45, 0x68, 0x1c, 0x90, 0xb4, 0xfd, 0xb6, 0x44, 0x3c, 0xfd, 0x77,
	0xf2, 0x24, 0x6d, 0x1f, 0xe7, 0x4e, 0x56, 0x95, 0xf8, 0xe8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2d, 0x00, 0xfc, 0xc3, 0x9d, 0x04, 0x00, 0x00,
}

func (m *GenerateTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenerateTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenerateTokenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Expiration != nil {
		{
			size, err := m.Expiration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApiTokenService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintApiTokenService(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintApiTokenService(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApiTokenService(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenerateTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenerateTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenerateTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApiTokenService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintApiTokenService(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAPITokensRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAPITokensRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAPITokensRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RevokedOneof != nil {
		{
			size := m.RevokedOneof.Size()
			i -= size
			if _, err := m.RevokedOneof.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetAPITokensRequest_Revoked) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAPITokensRequest_Revoked) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.Revoked {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *GetAPITokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAPITokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAPITokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApiTokenService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListAllowedTokenRolesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAllowedTokenRolesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAllowedTokenRolesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RoleNames) > 0 {
		for iNdEx := len(m.RoleNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RoleNames[iNdEx])
			copy(dAtA[i:], m.RoleNames[iNdEx])
			i = encodeVarintApiTokenService(dAtA, i, uint64(len(m.RoleNames[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintApiTokenService(dAtA []byte, offset int, v uint64) int {
	offset -= sovApiTokenService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenerateTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApiTokenService(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovApiTokenService(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovApiTokenService(uint64(l))
		}
	}
	if m.Expiration != nil {
		l = m.Expiration.Size()
		n += 1 + l + sovApiTokenService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GenerateTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovApiTokenService(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovApiTokenService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetAPITokensRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RevokedOneof != nil {
		n += m.RevokedOneof.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetAPITokensRequest_Revoked) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *GetAPITokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovApiTokenService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListAllowedTokenRolesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RoleNames) > 0 {
		for _, s := range m.RoleNames {
			l = len(s)
			n += 1 + l + sovApiTokenService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApiTokenService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApiTokenService(x uint64) (n int) {
	return sovApiTokenService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenerateTokenRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiTokenService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenerateTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenerateTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiTokenService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiTokenService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiTokenService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiTokenService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = &types.Timestamp{}
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiTokenService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenerateTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiTokenService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenerateTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenerateTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiTokenService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiTokenService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &storage.TokenMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiTokenService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAPITokensRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiTokenService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAPITokensRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAPITokensRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revoked", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.RevokedOneof = &GetAPITokensRequest_Revoked{b}
		default:
			iNdEx = preIndex
			skippy, err := skipApiTokenService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAPITokensResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiTokenService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAPITokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAPITokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiTokenService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, &storage.TokenMetadata{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiTokenService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListAllowedTokenRolesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiTokenService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListAllowedTokenRolesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAllowedTokenRolesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiTokenService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleNames = append(m.RoleNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiTokenService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiTokenService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApiTokenService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApiTokenService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApiTokenService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApiTokenService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApiTokenService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApiTokenService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApiTokenService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApiTokenService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApiTokenService = fmt.Errorf("proto: unexpected end of group")
)
