// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account_attachments.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// AccountAttachmentCreateRequest specifies what kind of AccountAttachments should be registered in
// the organizations fleet.
type AccountAttachmentCreateRequest struct {
	// Reserved for future use.
	Meta *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Parameters to define the new AccountAttachment.
	AccountAttachment    *AccountAttachment `protobuf:"bytes,2,opt,name=account_attachment,json=accountAttachment,proto3" json:"account_attachment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountAttachmentCreateRequest) Reset()         { *m = AccountAttachmentCreateRequest{} }
func (m *AccountAttachmentCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AccountAttachmentCreateRequest) ProtoMessage()    {}
func (*AccountAttachmentCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{0}
}

func (m *AccountAttachmentCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachmentCreateRequest.Unmarshal(m, b)
}
func (m *AccountAttachmentCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachmentCreateRequest.Marshal(b, m, deterministic)
}
func (m *AccountAttachmentCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachmentCreateRequest.Merge(m, src)
}
func (m *AccountAttachmentCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AccountAttachmentCreateRequest.Size(m)
}
func (m *AccountAttachmentCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachmentCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachmentCreateRequest proto.InternalMessageInfo

func (m *AccountAttachmentCreateRequest) GetMeta() *CreateRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountAttachmentCreateRequest) GetAccountAttachment() *AccountAttachment {
	if m != nil {
		return m.AccountAttachment
	}
	return nil
}

// AccountAttachmentCreateResponse reports how the AccountAttachments were created in the system.
type AccountAttachmentCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The created AccountAttachment.
	AccountAttachment *AccountAttachment `protobuf:"bytes,2,opt,name=account_attachment,json=accountAttachment,proto3" json:"account_attachment,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountAttachmentCreateResponse) Reset()         { *m = AccountAttachmentCreateResponse{} }
func (m *AccountAttachmentCreateResponse) String() string { return proto.CompactTextString(m) }
func (*AccountAttachmentCreateResponse) ProtoMessage()    {}
func (*AccountAttachmentCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{1}
}

func (m *AccountAttachmentCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachmentCreateResponse.Unmarshal(m, b)
}
func (m *AccountAttachmentCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachmentCreateResponse.Marshal(b, m, deterministic)
}
func (m *AccountAttachmentCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachmentCreateResponse.Merge(m, src)
}
func (m *AccountAttachmentCreateResponse) XXX_Size() int {
	return xxx_messageInfo_AccountAttachmentCreateResponse.Size(m)
}
func (m *AccountAttachmentCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachmentCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachmentCreateResponse proto.InternalMessageInfo

func (m *AccountAttachmentCreateResponse) GetMeta() *CreateResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountAttachmentCreateResponse) GetAccountAttachment() *AccountAttachment {
	if m != nil {
		return m.AccountAttachment
	}
	return nil
}

func (m *AccountAttachmentCreateResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountAttachmentGetRequest specifies which AccountAttachment to retrieve.
type AccountAttachmentGetRequest struct {
	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the AccountAttachment to retrieve.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountAttachmentGetRequest) Reset()         { *m = AccountAttachmentGetRequest{} }
func (m *AccountAttachmentGetRequest) String() string { return proto.CompactTextString(m) }
func (*AccountAttachmentGetRequest) ProtoMessage()    {}
func (*AccountAttachmentGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{2}
}

func (m *AccountAttachmentGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachmentGetRequest.Unmarshal(m, b)
}
func (m *AccountAttachmentGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachmentGetRequest.Marshal(b, m, deterministic)
}
func (m *AccountAttachmentGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachmentGetRequest.Merge(m, src)
}
func (m *AccountAttachmentGetRequest) XXX_Size() int {
	return xxx_messageInfo_AccountAttachmentGetRequest.Size(m)
}
func (m *AccountAttachmentGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachmentGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachmentGetRequest proto.InternalMessageInfo

func (m *AccountAttachmentGetRequest) GetMeta() *GetRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountAttachmentGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// AccountAttachmentGetResponse returns a requested AccountAttachment.
type AccountAttachmentGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested AccountAttachment.
	AccountAttachment *AccountAttachment `protobuf:"bytes,2,opt,name=account_attachment,json=accountAttachment,proto3" json:"account_attachment,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountAttachmentGetResponse) Reset()         { *m = AccountAttachmentGetResponse{} }
func (m *AccountAttachmentGetResponse) String() string { return proto.CompactTextString(m) }
func (*AccountAttachmentGetResponse) ProtoMessage()    {}
func (*AccountAttachmentGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{3}
}

func (m *AccountAttachmentGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachmentGetResponse.Unmarshal(m, b)
}
func (m *AccountAttachmentGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachmentGetResponse.Marshal(b, m, deterministic)
}
func (m *AccountAttachmentGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachmentGetResponse.Merge(m, src)
}
func (m *AccountAttachmentGetResponse) XXX_Size() int {
	return xxx_messageInfo_AccountAttachmentGetResponse.Size(m)
}
func (m *AccountAttachmentGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachmentGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachmentGetResponse proto.InternalMessageInfo

func (m *AccountAttachmentGetResponse) GetMeta() *GetResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountAttachmentGetResponse) GetAccountAttachment() *AccountAttachment {
	if m != nil {
		return m.AccountAttachment
	}
	return nil
}

func (m *AccountAttachmentGetResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountAttachmentDeleteRequest identifies a AccountAttachment by ID to delete.
type AccountAttachmentDeleteRequest struct {
	// Reserved for future use.
	Meta *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the AccountAttachment to delete.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountAttachmentDeleteRequest) Reset()         { *m = AccountAttachmentDeleteRequest{} }
func (m *AccountAttachmentDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AccountAttachmentDeleteRequest) ProtoMessage()    {}
func (*AccountAttachmentDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{4}
}

func (m *AccountAttachmentDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachmentDeleteRequest.Unmarshal(m, b)
}
func (m *AccountAttachmentDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachmentDeleteRequest.Marshal(b, m, deterministic)
}
func (m *AccountAttachmentDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachmentDeleteRequest.Merge(m, src)
}
func (m *AccountAttachmentDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_AccountAttachmentDeleteRequest.Size(m)
}
func (m *AccountAttachmentDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachmentDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachmentDeleteRequest proto.InternalMessageInfo

func (m *AccountAttachmentDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountAttachmentDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// AccountAttachmentDeleteResponse returns information about a AccountAttachment that was deleted.
type AccountAttachmentDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountAttachmentDeleteResponse) Reset()         { *m = AccountAttachmentDeleteResponse{} }
func (m *AccountAttachmentDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*AccountAttachmentDeleteResponse) ProtoMessage()    {}
func (*AccountAttachmentDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{5}
}

func (m *AccountAttachmentDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachmentDeleteResponse.Unmarshal(m, b)
}
func (m *AccountAttachmentDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachmentDeleteResponse.Marshal(b, m, deterministic)
}
func (m *AccountAttachmentDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachmentDeleteResponse.Merge(m, src)
}
func (m *AccountAttachmentDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_AccountAttachmentDeleteResponse.Size(m)
}
func (m *AccountAttachmentDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachmentDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachmentDeleteResponse proto.InternalMessageInfo

func (m *AccountAttachmentDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountAttachmentDeleteResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountAttachmentListRequest specifies criteria for retrieving a list of AccountAttachments.
type AccountAttachmentListRequest struct {
	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountAttachmentListRequest) Reset()         { *m = AccountAttachmentListRequest{} }
func (m *AccountAttachmentListRequest) String() string { return proto.CompactTextString(m) }
func (*AccountAttachmentListRequest) ProtoMessage()    {}
func (*AccountAttachmentListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{6}
}

func (m *AccountAttachmentListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachmentListRequest.Unmarshal(m, b)
}
func (m *AccountAttachmentListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachmentListRequest.Marshal(b, m, deterministic)
}
func (m *AccountAttachmentListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachmentListRequest.Merge(m, src)
}
func (m *AccountAttachmentListRequest) XXX_Size() int {
	return xxx_messageInfo_AccountAttachmentListRequest.Size(m)
}
func (m *AccountAttachmentListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachmentListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachmentListRequest proto.InternalMessageInfo

func (m *AccountAttachmentListRequest) GetMeta() *ListRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountAttachmentListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// AccountAttachmentListResponse returns a list of AccountAttachments that meet the criteria of a
// AccountAttachmentListRequest.
type AccountAttachmentListResponse struct {
	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	AccountAttachments []*AccountAttachment `protobuf:"bytes,2,rep,name=account_attachments,json=accountAttachments,proto3" json:"account_attachments,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountAttachmentListResponse) Reset()         { *m = AccountAttachmentListResponse{} }
func (m *AccountAttachmentListResponse) String() string { return proto.CompactTextString(m) }
func (*AccountAttachmentListResponse) ProtoMessage()    {}
func (*AccountAttachmentListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{7}
}

func (m *AccountAttachmentListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachmentListResponse.Unmarshal(m, b)
}
func (m *AccountAttachmentListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachmentListResponse.Marshal(b, m, deterministic)
}
func (m *AccountAttachmentListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachmentListResponse.Merge(m, src)
}
func (m *AccountAttachmentListResponse) XXX_Size() int {
	return xxx_messageInfo_AccountAttachmentListResponse.Size(m)
}
func (m *AccountAttachmentListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachmentListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachmentListResponse proto.InternalMessageInfo

func (m *AccountAttachmentListResponse) GetMeta() *ListResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountAttachmentListResponse) GetAccountAttachments() []*AccountAttachment {
	if m != nil {
		return m.AccountAttachments
	}
	return nil
}

func (m *AccountAttachmentListResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountAttachments assign an account to a role.
type AccountAttachment struct {
	// Unique identifier of the AccountAttachment.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the account of this AccountAttachment.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// The id of the attached role of this AccountAttachment.
	RoleId               string   `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountAttachment) Reset()         { *m = AccountAttachment{} }
func (m *AccountAttachment) String() string { return proto.CompactTextString(m) }
func (*AccountAttachment) ProtoMessage()    {}
func (*AccountAttachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bca51f1c3865f0, []int{8}
}

func (m *AccountAttachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAttachment.Unmarshal(m, b)
}
func (m *AccountAttachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAttachment.Marshal(b, m, deterministic)
}
func (m *AccountAttachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAttachment.Merge(m, src)
}
func (m *AccountAttachment) XXX_Size() int {
	return xxx_messageInfo_AccountAttachment.Size(m)
}
func (m *AccountAttachment) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAttachment.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAttachment proto.InternalMessageInfo

func (m *AccountAttachment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccountAttachment) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *AccountAttachment) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountAttachmentCreateRequest)(nil), "v1.AccountAttachmentCreateRequest")
	proto.RegisterType((*AccountAttachmentCreateResponse)(nil), "v1.AccountAttachmentCreateResponse")
	proto.RegisterType((*AccountAttachmentGetRequest)(nil), "v1.AccountAttachmentGetRequest")
	proto.RegisterType((*AccountAttachmentGetResponse)(nil), "v1.AccountAttachmentGetResponse")
	proto.RegisterType((*AccountAttachmentDeleteRequest)(nil), "v1.AccountAttachmentDeleteRequest")
	proto.RegisterType((*AccountAttachmentDeleteResponse)(nil), "v1.AccountAttachmentDeleteResponse")
	proto.RegisterType((*AccountAttachmentListRequest)(nil), "v1.AccountAttachmentListRequest")
	proto.RegisterType((*AccountAttachmentListResponse)(nil), "v1.AccountAttachmentListResponse")
	proto.RegisterType((*AccountAttachment)(nil), "v1.AccountAttachment")
}

func init() { proto.RegisterFile("account_attachments.proto", fileDescriptor_43bca51f1c3865f0) }

var fileDescriptor_43bca51f1c3865f0 = []byte{
	// 876 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0x96, 0x9d, 0x34, 0x66, 0x07, 0x38, 0xec, 0x20, 0xda, 0xc4, 0x0d, 0x59, 0x33, 0x0b, 0x12,
	0xda, 0x6e, 0x6c, 0x25, 0x44, 0x42, 0x32, 0x20, 0x48, 0x58, 0xa9, 0x5a, 0x69, 0x91, 0xaa, 0x48,
	0x9c, 0xa3, 0xa9, 0x3d, 0xcd, 0x5a, 0xb5, 0x3d, 0xc6, 0x33, 0x49, 0x2a, 0xaa, 0x5e, 0xb8, 0x71,
	0x2d, 0xfc, 0x05, 0x3d, 0x21, 0xb8, 0x20, 0x99, 0x03, 0xe2, 0x04, 0x3d, 0xf6, 0xca, 0x9f, 0x00,
	0x07, 0xce, 0xe4, 0x40, 0x17, 0x09, 0x09, 0xcd, 0xc4, 0xf9, 0xe1, 0xd8, 0xce, 0x0a, 0xad, 0x38,
	0x70, 0xda, 0xcd, 0xbc, 0xef, 0xbd, 0xef, 0xbd, 0xef, 0x7d, 0x1e, 0x1b, 0x34, 0xb0, 0xe3, 0xd0,
	0x49, 0xc8, 0x47, 0x98, 0x73, 0xec, 0x9c, 0x07, 0x24, 0xe4, 0xcc, 0x8c, 0x62, 0xca, 0x29, 0x54,
	0xa7, 0x1d, 0xbd, 0x39, 0xa6, 0x74, 0xec, 0x13, 0x0b, 0x47, 0x9e, 0x85, 0xc3, 0x90, 0x72, 0xcc,
	0x3d, 0x1a, 0xa6, 0x08, 0xfd, 0x58, 0xfe, 0x71, 0xda, 0x63, 0x12, 0xb6, 0xd9, 0x0c, 0x8f, 0xc7,
	0x24, 0xb6, 0x68, 0x24, 0x11, 0x05, 0xe8, 0x97, 0xd3, 0x50, 0xfa, 0x13, 0xb0, 0x88, 0x38, 0x8b,
	0xff, 0xd1, 0xf7, 0x0a, 0x68, 0xf5, 0x17, 0x8d, 0xf4, 0x57, 0x7d, 0x7c, 0x14, 0x13, 0xcc, 0xc9,
	0x90, 0x7c, 0x3a, 0x21, 0x8c, 0xc3, 0x36, 0xa8, 0x06, 0x84, 0xe3, 0xba, 0x62, 0x28, 0x6f, 0xbd,
	0xd8, 0x6d, 0x98, 0xd3, 0x8e, 0x99, 0x01, 0x7c, 0x4c, 0x38, 0x76, 0x31, 0xc7, 0x43, 0x09, 0x83,
	0x77, 0x00, 0xcc, 0x4f, 0x56, 0x57, 0x65, 0xf2, 0xab, 0x22, 0x39, 0x47, 0x37, 0x00, 0x7f, 0x3c,
	0x4f, 0xb4, 0x6b, 0xdf, 0xcd, 0x13, 0x4d, 0x19, 0xee, 0xe3, 0xed, 0xb0, 0xbd, 0xff, 0xd7, 0xf3,
	0x44, 0x7b, 0xe9, 0xc7, 0x79, 0xa2, 0x69, 0xe9, 0x20, 0xe8, 0x6f, 0x05, 0x1c, 0x94, 0xb6, 0xcd,
	0x22, 0x1a, 0x32, 0x02, 0xed, 0x4c, 0xdf, 0xfa, 0x66, 0xdf, 0x0b, 0xc4, 0xb2, 0xf1, 0x0c, 0xff,
	0x7f, 0x34, 0x04, 0xfc, 0x10, 0x80, 0x18, 0x73, 0x32, 0xf2, 0xbd, 0xc0, 0xe3, 0xf5, 0xca, 0xba,
	0xd2, 0x10, 0x73, 0x72, 0x26, 0x0e, 0x0b, 0xdb, 0xd9, 0x8b, 0x97, 0x61, 0x1b, 0x08, 0x19, 0xae,
	0x7d, 0x2d, 0x8e, 0x11, 0x01, 0x37, 0x73, 0x1d, 0xdc, 0x26, 0x7c, 0xb9, 0xb2, 0xa3, 0xcc, 0xe8,
	0xd7, 0x05, 0xcd, 0x3a, 0xba, 0xb5, 0x2f, 0x1d, 0xa8, 0x9e, 0x2b, 0x47, 0xdb, 0xcb, 0x30, 0xab,
	0x9e, 0x8b, 0xfe, 0x54, 0x40, 0xb3, 0x98, 0x27, 0xd5, 0xf8, 0x9d, 0x0c, 0xd1, 0x8d, 0x15, 0xd1,
	0xff, 0x5e, 0xe0, 0xfb, 0x05, 0x8f, 0xc5, 0x09, 0xf1, 0xc9, 0xce, 0xc7, 0x22, 0x03, 0xf8, 0x17,
	0x32, 0x7f, 0x5b, 0xe4, 0xe6, 0x65, 0xb1, 0x72, 0x37, 0x67, 0x11, 0x3b, 0xc4, 0xce, 0x4a, 0xa3,
	0x5e, 0x51, 0x1a, 0x5a, 0xe0, 0x89, 0x33, 0x8f, 0xad, 0xcc, 0x77, 0x2b, 0xef, 0x89, 0x8d, 0xf0,
	0x96, 0x2c, 0x08, 0xd4, 0xee, 0x79, 0x3e, 0x27, 0x71, 0x81, 0x34, 0x69, 0x04, 0xfd, 0xaa, 0x80,
	0xd7, 0x4a, 0x18, 0x53, 0x71, 0x8e, 0x33, 0x94, 0xf5, 0x35, 0x65, 0x56, 0x9a, 0x94, 0x73, 0x08,
	0x5e, 0x29, 0xb8, 0x7b, 0xeb, 0xaa, 0x51, 0xb9, 0xcc, 0x7c, 0x3f, 0xc8, 0xbe, 0x60, 0xce, 0x7c,
	0xec, 0xea, 0xee, 0x43, 0x5f, 0x54, 0xc0, 0x7e, 0x8e, 0x17, 0x9a, 0xd2, 0x36, 0x8a, 0xd4, 0xa6,
	0x25, 0x12, 0x1b, 0x4f, 0xe6, 0x89, 0xa6, 0x9e, 0x9e, 0xc8, 0xfc, 0x67, 0xf3, 0x44, 0x7b, 0xe1,
	0x13, 0x46, 0xe2, 0x21, 0xf5, 0x89, 0xb0, 0x12, 0x1c, 0x00, 0xb0, 0x9c, 0x6d, 0x65, 0xb7, 0x43,
	0x91, 0xd7, 0x12, 0x79, 0x35, 0x81, 0x4d, 0x73, 0x7f, 0x5a, 0x16, 0xa8, 0xca, 0x02, 0x7b, 0x69,
	0xda, 0xa9, 0x0b, 0xdf, 0x03, 0x5a, 0x4c, 0x7d, 0x22, 0x0a, 0x54, 0xb6, 0x0b, 0x08, 0xa2, 0x5c,
	0x01, 0xc9, 0x5e, 0x13, 0x39, 0xa7, 0xae, 0xfd, 0xb3, 0xf2, 0xb8, 0xff, 0x41, 0xf7, 0x7d, 0xf8,
	0xee, 0x43, 0x03, 0x79, 0x2e, 0xb2, 0x0d, 0x84, 0x71, 0xbb, 0xdb, 0xeb, 0xa1, 0x63, 0x03, 0xa5,
	0x65, 0xc5, 0x61, 0xdc, 0xee, 0x74, 0x3a, 0xe2, 0x6c, 0xdd, 0x2e, 0xb2, 0x11, 0x6e, 0xf7, 0x7a,
	0x3d, 0xf4, 0x48, 0xb8, 0xed, 0xb3, 0x27, 0x9b, 0xc3, 0x49, 0xeb, 0x3d, 0x9d, 0x27, 0x9a, 0x23,
	0x8e, 0x6f, 0xf1, 0x7b, 0x23, 0xf2, 0x00, 0x07, 0x91, 0x4f, 0x98, 0x95, 0xdf, 0xe5, 0x28, 0x26,
	0x8c, 0x4e, 0x62, 0x87, 0x98, 0xfc, 0x01, 0xff, 0x66, 0x9e, 0x68, 0xe6, 0x25, 0x78, 0xb1, 0x97,
	0xd1, 0x3a, 0xa5, 0x3b, 0xaf, 0x02, 0xd8, 0xcf, 0x2f, 0xf9, 0x77, 0x05, 0xd4, 0x16, 0xaf, 0x10,
	0x88, 0x0a, 0x6d, 0x92, 0x79, 0x2f, 0xea, 0x87, 0x3b, 0x31, 0x0b, 0x6b, 0xa2, 0xaf, 0x94, 0xc7,
	0x7d, 0x8a, 0x02, 0xf0, 0xe6, 0x19, 0xc1, 0x71, 0x68, 0x9c, 0xd3, 0x99, 0xc1, 0xa9, 0x11, 0xe0,
	0xfb, 0xc4, 0xc0, 0x46, 0xde, 0x0f, 0x27, 0xe7, 0x9c, 0x47, 0xcc, 0xb6, 0xac, 0xd9, 0x6c, 0x66,
	0x32, 0x1e, 0xd3, 0x70, 0xec, 0x06, 0xa6, 0x43, 0x03, 0xcb, 0xa5, 0x0e, 0x93, 0x5f, 0x0b, 0x8c,
	0xc4, 0x53, 0xcf, 0x21, 0xcc, 0xca, 0x4f, 0x70, 0xb8, 0x20, 0xff, 0xfc, 0x97, 0xdf, 0xbe, 0x54,
	0x9b, 0xe8, 0x86, 0x35, 0xed, 0x14, 0xa8, 0xc1, 0x6c, 0xe5, 0x08, 0x86, 0xa0, 0x72, 0x9b, 0x70,
	0x78, 0x50, 0x38, 0xc2, 0xfa, 0x5d, 0xa2, 0x1b, 0xe5, 0x80, 0x74, 0xc0, 0x37, 0x24, 0x5b, 0x0b,
	0x36, 0x4b, 0xd8, 0xac, 0x87, 0x9e, 0xfb, 0x08, 0xce, 0x40, 0x6d, 0x71, 0x9d, 0x95, 0x28, 0x9b,
	0xb9, 0x5a, 0x4b, 0x94, 0xcd, 0xde, 0x87, 0x4b, 0xe2, 0xa3, 0xdd, 0xc4, 0x3e, 0xa8, 0x8a, 0xab,
	0x02, 0x16, 0x0f, 0xb2, 0x71, 0x71, 0xe9, 0xaf, 0xef, 0x40, 0xa4, 0x94, 0x07, 0x92, 0xb2, 0x01,
	0xcb, 0x94, 0xd5, 0x6f, 0x3e, 0xbb, 0x48, 0xb4, 0xeb, 0x4f, 0x2f, 0x12, 0x2d, 0xff, 0xac, 0x0f,
	0x6c, 0xd0, 0x74, 0x68, 0xb0, 0xde, 0x2a, 0x8e, 0x3c, 0x41, 0x19, 0xf9, 0x93, 0xe0, 0xae, 0x17,
	0x8e, 0x07, 0x7a, 0x7e, 0xa1, 0x77, 0xd2, 0xd8, 0xdd, 0x9a, 0xfc, 0x9c, 0x7b, 0xfb, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x08, 0x2d, 0xdc, 0x43, 0x56, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountAttachmentsClient is the client API for AccountAttachments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountAttachmentsClient interface {
	// Create registers a new AccountAttachment.
	Create(ctx context.Context, in *AccountAttachmentCreateRequest, opts ...grpc.CallOption) (*AccountAttachmentCreateResponse, error)
	// Get reads one AccountAttachment by ID.
	Get(ctx context.Context, in *AccountAttachmentGetRequest, opts ...grpc.CallOption) (*AccountAttachmentGetResponse, error)
	// Delete removes a AccountAttachment by ID.
	Delete(ctx context.Context, in *AccountAttachmentDeleteRequest, opts ...grpc.CallOption) (*AccountAttachmentDeleteResponse, error)
	// List gets a list of AccountAttachments matching a given set of criteria.
	List(ctx context.Context, in *AccountAttachmentListRequest, opts ...grpc.CallOption) (*AccountAttachmentListResponse, error)
}

type accountAttachmentsClient struct {
	cc *grpc.ClientConn
}

func NewAccountAttachmentsClient(cc *grpc.ClientConn) AccountAttachmentsClient {
	return &accountAttachmentsClient{cc}
}

func (c *accountAttachmentsClient) Create(ctx context.Context, in *AccountAttachmentCreateRequest, opts ...grpc.CallOption) (*AccountAttachmentCreateResponse, error) {
	out := new(AccountAttachmentCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountAttachments/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAttachmentsClient) Get(ctx context.Context, in *AccountAttachmentGetRequest, opts ...grpc.CallOption) (*AccountAttachmentGetResponse, error) {
	out := new(AccountAttachmentGetResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountAttachments/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAttachmentsClient) Delete(ctx context.Context, in *AccountAttachmentDeleteRequest, opts ...grpc.CallOption) (*AccountAttachmentDeleteResponse, error) {
	out := new(AccountAttachmentDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountAttachments/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAttachmentsClient) List(ctx context.Context, in *AccountAttachmentListRequest, opts ...grpc.CallOption) (*AccountAttachmentListResponse, error) {
	out := new(AccountAttachmentListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountAttachments/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountAttachmentsServer is the server API for AccountAttachments service.
type AccountAttachmentsServer interface {
	// Create registers a new AccountAttachment.
	Create(context.Context, *AccountAttachmentCreateRequest) (*AccountAttachmentCreateResponse, error)
	// Get reads one AccountAttachment by ID.
	Get(context.Context, *AccountAttachmentGetRequest) (*AccountAttachmentGetResponse, error)
	// Delete removes a AccountAttachment by ID.
	Delete(context.Context, *AccountAttachmentDeleteRequest) (*AccountAttachmentDeleteResponse, error)
	// List gets a list of AccountAttachments matching a given set of criteria.
	List(context.Context, *AccountAttachmentListRequest) (*AccountAttachmentListResponse, error)
}

// UnimplementedAccountAttachmentsServer can be embedded to have forward compatible implementations.
type UnimplementedAccountAttachmentsServer struct {
}

func (*UnimplementedAccountAttachmentsServer) Create(ctx context.Context, req *AccountAttachmentCreateRequest) (*AccountAttachmentCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccountAttachmentsServer) Get(ctx context.Context, req *AccountAttachmentGetRequest) (*AccountAttachmentGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccountAttachmentsServer) Delete(ctx context.Context, req *AccountAttachmentDeleteRequest) (*AccountAttachmentDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedAccountAttachmentsServer) List(ctx context.Context, req *AccountAttachmentListRequest) (*AccountAttachmentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterAccountAttachmentsServer(s *grpc.Server, srv AccountAttachmentsServer) {
	s.RegisterService(&_AccountAttachments_serviceDesc, srv)
}

func _AccountAttachments_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAttachmentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAttachmentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountAttachments/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAttachmentsServer).Create(ctx, req.(*AccountAttachmentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAttachments_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAttachmentGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAttachmentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountAttachments/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAttachmentsServer).Get(ctx, req.(*AccountAttachmentGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAttachments_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAttachmentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAttachmentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountAttachments/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAttachmentsServer).Delete(ctx, req.(*AccountAttachmentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAttachments_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAttachmentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAttachmentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountAttachments/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAttachmentsServer).List(ctx, req.(*AccountAttachmentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountAttachments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccountAttachments",
	HandlerType: (*AccountAttachmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AccountAttachments_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccountAttachments_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccountAttachments_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AccountAttachments_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_attachments.proto",
}