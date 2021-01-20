// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workspace.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PrepareForUserNSRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareForUserNSRequest) Reset()         { *m = PrepareForUserNSRequest{} }
func (m *PrepareForUserNSRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareForUserNSRequest) ProtoMessage()    {}
func (*PrepareForUserNSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{0}
}

func (m *PrepareForUserNSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareForUserNSRequest.Unmarshal(m, b)
}
func (m *PrepareForUserNSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareForUserNSRequest.Marshal(b, m, deterministic)
}
func (m *PrepareForUserNSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareForUserNSRequest.Merge(m, src)
}
func (m *PrepareForUserNSRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareForUserNSRequest.Size(m)
}
func (m *PrepareForUserNSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareForUserNSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareForUserNSRequest proto.InternalMessageInfo

type PrepareForUserNSResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareForUserNSResponse) Reset()         { *m = PrepareForUserNSResponse{} }
func (m *PrepareForUserNSResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareForUserNSResponse) ProtoMessage()    {}
func (*PrepareForUserNSResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{1}
}

func (m *PrepareForUserNSResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareForUserNSResponse.Unmarshal(m, b)
}
func (m *PrepareForUserNSResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareForUserNSResponse.Marshal(b, m, deterministic)
}
func (m *PrepareForUserNSResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareForUserNSResponse.Merge(m, src)
}
func (m *PrepareForUserNSResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareForUserNSResponse.Size(m)
}
func (m *PrepareForUserNSResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareForUserNSResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareForUserNSResponse proto.InternalMessageInfo

type WriteIDMappingResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ErrorCode            uint32   `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteIDMappingResponse) Reset()         { *m = WriteIDMappingResponse{} }
func (m *WriteIDMappingResponse) String() string { return proto.CompactTextString(m) }
func (*WriteIDMappingResponse) ProtoMessage()    {}
func (*WriteIDMappingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{2}
}

func (m *WriteIDMappingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteIDMappingResponse.Unmarshal(m, b)
}
func (m *WriteIDMappingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteIDMappingResponse.Marshal(b, m, deterministic)
}
func (m *WriteIDMappingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteIDMappingResponse.Merge(m, src)
}
func (m *WriteIDMappingResponse) XXX_Size() int {
	return xxx_messageInfo_WriteIDMappingResponse.Size(m)
}
func (m *WriteIDMappingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteIDMappingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteIDMappingResponse proto.InternalMessageInfo

func (m *WriteIDMappingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *WriteIDMappingResponse) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type WriteIDMappingRequest struct {
	Pid                  int64                            `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Gid                  bool                             `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	Mapping              []*WriteIDMappingRequest_Mapping `protobuf:"bytes,3,rep,name=mapping,proto3" json:"mapping,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WriteIDMappingRequest) Reset()         { *m = WriteIDMappingRequest{} }
func (m *WriteIDMappingRequest) String() string { return proto.CompactTextString(m) }
func (*WriteIDMappingRequest) ProtoMessage()    {}
func (*WriteIDMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{3}
}

func (m *WriteIDMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteIDMappingRequest.Unmarshal(m, b)
}
func (m *WriteIDMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteIDMappingRequest.Marshal(b, m, deterministic)
}
func (m *WriteIDMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteIDMappingRequest.Merge(m, src)
}
func (m *WriteIDMappingRequest) XXX_Size() int {
	return xxx_messageInfo_WriteIDMappingRequest.Size(m)
}
func (m *WriteIDMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteIDMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteIDMappingRequest proto.InternalMessageInfo

func (m *WriteIDMappingRequest) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *WriteIDMappingRequest) GetGid() bool {
	if m != nil {
		return m.Gid
	}
	return false
}

func (m *WriteIDMappingRequest) GetMapping() []*WriteIDMappingRequest_Mapping {
	if m != nil {
		return m.Mapping
	}
	return nil
}

type WriteIDMappingRequest_Mapping struct {
	ContainerId          uint32   `protobuf:"varint,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	HostId               uint32   `protobuf:"varint,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Size                 uint32   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteIDMappingRequest_Mapping) Reset()         { *m = WriteIDMappingRequest_Mapping{} }
func (m *WriteIDMappingRequest_Mapping) String() string { return proto.CompactTextString(m) }
func (*WriteIDMappingRequest_Mapping) ProtoMessage()    {}
func (*WriteIDMappingRequest_Mapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{3, 0}
}

func (m *WriteIDMappingRequest_Mapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteIDMappingRequest_Mapping.Unmarshal(m, b)
}
func (m *WriteIDMappingRequest_Mapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteIDMappingRequest_Mapping.Marshal(b, m, deterministic)
}
func (m *WriteIDMappingRequest_Mapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteIDMappingRequest_Mapping.Merge(m, src)
}
func (m *WriteIDMappingRequest_Mapping) XXX_Size() int {
	return xxx_messageInfo_WriteIDMappingRequest_Mapping.Size(m)
}
func (m *WriteIDMappingRequest_Mapping) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteIDMappingRequest_Mapping.DiscardUnknown(m)
}

var xxx_messageInfo_WriteIDMappingRequest_Mapping proto.InternalMessageInfo

func (m *WriteIDMappingRequest_Mapping) GetContainerId() uint32 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *WriteIDMappingRequest_Mapping) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *WriteIDMappingRequest_Mapping) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type MountProcRequest struct {
	Pid                  int64    `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MountProcRequest) Reset()         { *m = MountProcRequest{} }
func (m *MountProcRequest) String() string { return proto.CompactTextString(m) }
func (*MountProcRequest) ProtoMessage()    {}
func (*MountProcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{4}
}

func (m *MountProcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountProcRequest.Unmarshal(m, b)
}
func (m *MountProcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountProcRequest.Marshal(b, m, deterministic)
}
func (m *MountProcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountProcRequest.Merge(m, src)
}
func (m *MountProcRequest) XXX_Size() int {
	return xxx_messageInfo_MountProcRequest.Size(m)
}
func (m *MountProcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MountProcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MountProcRequest proto.InternalMessageInfo

func (m *MountProcRequest) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type MountProcResponse struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MountProcResponse) Reset()         { *m = MountProcResponse{} }
func (m *MountProcResponse) String() string { return proto.CompactTextString(m) }
func (*MountProcResponse) ProtoMessage()    {}
func (*MountProcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{5}
}

func (m *MountProcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountProcResponse.Unmarshal(m, b)
}
func (m *MountProcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountProcResponse.Marshal(b, m, deterministic)
}
func (m *MountProcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountProcResponse.Merge(m, src)
}
func (m *MountProcResponse) XXX_Size() int {
	return xxx_messageInfo_MountProcResponse.Size(m)
}
func (m *MountProcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MountProcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MountProcResponse proto.InternalMessageInfo

func (m *MountProcResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type TeardownRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeardownRequest) Reset()         { *m = TeardownRequest{} }
func (m *TeardownRequest) String() string { return proto.CompactTextString(m) }
func (*TeardownRequest) ProtoMessage()    {}
func (*TeardownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{6}
}

func (m *TeardownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeardownRequest.Unmarshal(m, b)
}
func (m *TeardownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeardownRequest.Marshal(b, m, deterministic)
}
func (m *TeardownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeardownRequest.Merge(m, src)
}
func (m *TeardownRequest) XXX_Size() int {
	return xxx_messageInfo_TeardownRequest.Size(m)
}
func (m *TeardownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeardownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeardownRequest proto.InternalMessageInfo

type TeardownResponse struct {
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeardownResponse) Reset()         { *m = TeardownResponse{} }
func (m *TeardownResponse) String() string { return proto.CompactTextString(m) }
func (*TeardownResponse) ProtoMessage()    {}
func (*TeardownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac718ecaafc2333, []int{7}
}

func (m *TeardownResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeardownResponse.Unmarshal(m, b)
}
func (m *TeardownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeardownResponse.Marshal(b, m, deterministic)
}
func (m *TeardownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeardownResponse.Merge(m, src)
}
func (m *TeardownResponse) XXX_Size() int {
	return xxx_messageInfo_TeardownResponse.Size(m)
}
func (m *TeardownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TeardownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TeardownResponse proto.InternalMessageInfo

func (m *TeardownResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*PrepareForUserNSRequest)(nil), "iws.PrepareForUserNSRequest")
	proto.RegisterType((*PrepareForUserNSResponse)(nil), "iws.PrepareForUserNSResponse")
	proto.RegisterType((*WriteIDMappingResponse)(nil), "iws.WriteIDMappingResponse")
	proto.RegisterType((*WriteIDMappingRequest)(nil), "iws.WriteIDMappingRequest")
	proto.RegisterType((*WriteIDMappingRequest_Mapping)(nil), "iws.WriteIDMappingRequest.Mapping")
	proto.RegisterType((*MountProcRequest)(nil), "iws.MountProcRequest")
	proto.RegisterType((*MountProcResponse)(nil), "iws.MountProcResponse")
	proto.RegisterType((*TeardownRequest)(nil), "iws.TeardownRequest")
	proto.RegisterType((*TeardownResponse)(nil), "iws.TeardownResponse")
}

func init() {
	proto.RegisterFile("workspace.proto", fileDescriptor_dac718ecaafc2333)
}

var fileDescriptor_dac718ecaafc2333 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0xa5, 0x99, 0xd6, 0xf6, 0xee, 0x57, 0x96, 0x59, 0x74, 0x0b, 0x81, 0x49, 0x25, 0x4f,
	0xe5, 0xcf, 0x52, 0x69, 0x3c, 0x21, 0xed, 0x09, 0x10, 0x52, 0x41, 0x43, 0x5b, 0x06, 0x9a, 0xe0,
	0xa5, 0xca, 0xe2, 0xab, 0xcc, 0x82, 0xfa, 0x1a, 0xdb, 0x25, 0x12, 0xdf, 0x89, 0x6f, 0xc3, 0x07,
	0x42, 0x35, 0x71, 0x06, 0x61, 0x7d, 0xbb, 0xe7, 0x5c, 0xdf, 0x13, 0xfb, 0x9c, 0x1b, 0xd8, 0xad,
	0x49, 0x7f, 0x36, 0xaa, 0x28, 0x31, 0x53, 0x9a, 0x2c, 0xb1, 0x50, 0xd4, 0x26, 0xbd, 0x07, 0x07,
	0x67, 0x1a, 0x55, 0xa1, 0xf1, 0x35, 0xe9, 0x0f, 0x06, 0xf5, 0xbb, 0x8b, 0x1c, 0xbf, 0xae, 0xd0,
	0xd8, 0x34, 0x81, 0xf8, 0xdf, 0x96, 0x51, 0x24, 0x0d, 0xa6, 0xe7, 0xb0, 0x7f, 0xa9, 0x85, 0xc5,
	0xf9, 0xab, 0xd3, 0x42, 0x29, 0x21, 0x2b, 0xdf, 0x61, 0x31, 0xf4, 0x97, 0x68, 0x4c, 0x51, 0x61,
	0x1c, 0x4c, 0x82, 0xe9, 0x30, 0xf7, 0x90, 0x1d, 0x02, 0xa0, 0xd6, 0xa4, 0x17, 0x25, 0x71, 0x8c,
	0x7b, 0x93, 0x60, 0x3a, 0xca, 0x87, 0x8e, 0x79, 0x49, 0x1c, 0xd3, 0x9f, 0x01, 0x8c, 0xbb, 0x9a,
	0xee, 0x22, 0x2c, 0x82, 0x50, 0x09, 0xee, 0xe4, 0xc2, 0x7c, 0x5d, 0xae, 0x99, 0x4a, 0x70, 0xa7,
	0x31, 0xc8, 0xd7, 0x25, 0x3b, 0x81, 0xfe, 0xf2, 0xf7, 0x54, 0x1c, 0x4e, 0xc2, 0xe9, 0xce, 0x71,
	0x9a, 0x89, 0xda, 0x64, 0xb7, 0x0a, 0x66, 0x1e, 0xfa, 0x91, 0xe4, 0x23, 0xf4, 0x1b, 0x8e, 0x3d,
	0x84, 0xff, 0x4b, 0x92, 0xb6, 0x10, 0x12, 0xf5, 0xa2, 0xf9, 0xea, 0x28, 0xdf, 0x69, 0xb9, 0x39,
	0x67, 0x07, 0xd0, 0xbf, 0x26, 0x63, 0x17, 0xcd, 0x0d, 0x46, 0xf9, 0xf6, 0x1a, 0xce, 0x39, 0x63,
	0xb0, 0x65, 0xc4, 0x77, 0x8c, 0x43, 0xc7, 0xba, 0x3a, 0x7d, 0x0c, 0xd1, 0x29, 0xad, 0xa4, 0x3d,
	0xd3, 0x54, 0x76, 0x1e, 0xd4, 0x6b, 0x1f, 0xf4, 0x66, 0x6b, 0x10, 0x44, 0xbd, 0x74, 0x06, 0x7b,
	0x7f, 0x9c, 0x6d, 0x0c, 0x4d, 0x60, 0xf0, 0x85, 0xca, 0xc2, 0x0a, 0x92, 0x8d, 0xa3, 0x2d, 0x4e,
	0xf7, 0x60, 0xf7, 0x3d, 0x16, 0x9a, 0x53, 0x2d, 0x7d, 0x6a, 0x4f, 0x21, 0xba, 0xa1, 0x6e, 0x32,
	0x31, 0xab, 0xb2, 0x44, 0x63, 0x1a, 0xcb, 0x3c, 0x3c, 0xfe, 0xd1, 0x03, 0x36, 0x97, 0x97, 0x7e,
	0x33, 0x2e, 0x50, 0x7f, 0x13, 0x25, 0xb2, 0x73, 0x88, 0xba, 0xd1, 0xb3, 0x07, 0xce, 0xd0, 0x0d,
	0xcb, 0x92, 0x1c, 0x6e, 0xe8, 0x36, 0xfb, 0xf2, 0x1f, 0x7b, 0x0b, 0x77, 0xfe, 0x0e, 0x83, 0x25,
	0x9b, 0x13, 0x4a, 0xee, 0xdf, 0xda, 0x6b, 0xc5, 0x4e, 0x60, 0xd8, 0x1a, 0xc5, 0xc6, 0xee, 0x6c,
	0xd7, 0xe4, 0x64, 0xbf, 0x4b, 0xb7, 0xd3, 0xcf, 0x61, 0xe0, 0x2d, 0x62, 0x77, 0xdd, 0xa9, 0x8e,
	0x89, 0xc9, 0xb8, 0xc3, 0xfa, 0xd1, 0x17, 0x4f, 0x3e, 0x3d, 0xaa, 0x84, 0xbd, 0x5e, 0x5d, 0x65,
	0x25, 0x2d, 0x67, 0x95, 0xb0, 0x8a, 0xf8, 0x91, 0xa0, 0xa6, 0x9a, 0xd5, 0xe6, 0x88, 0x17, 0xb8,
	0x24, 0x39, 0x2b, 0x94, 0xb8, 0xda, 0x76, 0xff, 0xd9, 0xb3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0x74, 0xb6, 0xeb, 0x7a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InWorkspaceServiceClient is the client API for InWorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InWorkspaceServiceClient interface {
	// PrepareForUserNS prepares a workspace container for wrapping it in a user namespace.
	// A container that called this function MUST call Teardown.
	//
	// This call will make the workspace container's rootfs shared, and mount the workspace
	// container's rootfs as a shiftfs mark under `/.workspace/mark` if the workspace has
	// the daemon hostPath mount. Can only be used once per workspace.
	PrepareForUserNS(ctx context.Context, in *PrepareForUserNSRequest, opts ...grpc.CallOption) (*PrepareForUserNSResponse, error)
	// WriteIDMapping writes a new user/group ID mapping to /proc/<pid>/uid_map (gid_map respectively). This is used
	// for user namespaces and is available four times every 10 seconds.
	WriteIDMapping(ctx context.Context, in *WriteIDMappingRequest, opts ...grpc.CallOption) (*WriteIDMappingResponse, error)
	// MountProc mounts a masked proc in the container's rootfs.
	// For now this can be used only once per workspace.
	MountProc(ctx context.Context, in *MountProcRequest, opts ...grpc.CallOption) (*MountProcResponse, error)
	// Teardown prepares workspace content backups and unmounts shiftfs mounts. The canary is supposed to be triggered
	// when the workspace is about to shut down, e.g. using the PreStop hook of a Kubernetes container.
	Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*TeardownResponse, error)
}

type inWorkspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInWorkspaceServiceClient(cc grpc.ClientConnInterface) InWorkspaceServiceClient {
	return &inWorkspaceServiceClient{cc}
}

func (c *inWorkspaceServiceClient) PrepareForUserNS(ctx context.Context, in *PrepareForUserNSRequest, opts ...grpc.CallOption) (*PrepareForUserNSResponse, error) {
	out := new(PrepareForUserNSResponse)
	err := c.cc.Invoke(ctx, "/iws.InWorkspaceService/PrepareForUserNS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inWorkspaceServiceClient) WriteIDMapping(ctx context.Context, in *WriteIDMappingRequest, opts ...grpc.CallOption) (*WriteIDMappingResponse, error) {
	out := new(WriteIDMappingResponse)
	err := c.cc.Invoke(ctx, "/iws.InWorkspaceService/WriteIDMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inWorkspaceServiceClient) MountProc(ctx context.Context, in *MountProcRequest, opts ...grpc.CallOption) (*MountProcResponse, error) {
	out := new(MountProcResponse)
	err := c.cc.Invoke(ctx, "/iws.InWorkspaceService/MountProc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inWorkspaceServiceClient) Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*TeardownResponse, error) {
	out := new(TeardownResponse)
	err := c.cc.Invoke(ctx, "/iws.InWorkspaceService/Teardown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InWorkspaceServiceServer is the server API for InWorkspaceService service.
type InWorkspaceServiceServer interface {
	// PrepareForUserNS prepares a workspace container for wrapping it in a user namespace.
	// A container that called this function MUST call Teardown.
	//
	// This call will make the workspace container's rootfs shared, and mount the workspace
	// container's rootfs as a shiftfs mark under `/.workspace/mark` if the workspace has
	// the daemon hostPath mount. Can only be used once per workspace.
	PrepareForUserNS(context.Context, *PrepareForUserNSRequest) (*PrepareForUserNSResponse, error)
	// WriteIDMapping writes a new user/group ID mapping to /proc/<pid>/uid_map (gid_map respectively). This is used
	// for user namespaces and is available four times every 10 seconds.
	WriteIDMapping(context.Context, *WriteIDMappingRequest) (*WriteIDMappingResponse, error)
	// MountProc mounts a masked proc in the container's rootfs.
	// For now this can be used only once per workspace.
	MountProc(context.Context, *MountProcRequest) (*MountProcResponse, error)
	// Teardown prepares workspace content backups and unmounts shiftfs mounts. The canary is supposed to be triggered
	// when the workspace is about to shut down, e.g. using the PreStop hook of a Kubernetes container.
	Teardown(context.Context, *TeardownRequest) (*TeardownResponse, error)
}

// UnimplementedInWorkspaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInWorkspaceServiceServer struct {
}

func (*UnimplementedInWorkspaceServiceServer) PrepareForUserNS(ctx context.Context, req *PrepareForUserNSRequest) (*PrepareForUserNSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareForUserNS not implemented")
}
func (*UnimplementedInWorkspaceServiceServer) WriteIDMapping(ctx context.Context, req *WriteIDMappingRequest) (*WriteIDMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteIDMapping not implemented")
}
func (*UnimplementedInWorkspaceServiceServer) MountProc(ctx context.Context, req *MountProcRequest) (*MountProcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MountProc not implemented")
}
func (*UnimplementedInWorkspaceServiceServer) Teardown(ctx context.Context, req *TeardownRequest) (*TeardownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Teardown not implemented")
}

func RegisterInWorkspaceServiceServer(s *grpc.Server, srv InWorkspaceServiceServer) {
	s.RegisterService(&_InWorkspaceService_serviceDesc, srv)
}

func _InWorkspaceService_PrepareForUserNS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareForUserNSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InWorkspaceServiceServer).PrepareForUserNS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iws.InWorkspaceService/PrepareForUserNS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InWorkspaceServiceServer).PrepareForUserNS(ctx, req.(*PrepareForUserNSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InWorkspaceService_WriteIDMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteIDMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InWorkspaceServiceServer).WriteIDMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iws.InWorkspaceService/WriteIDMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InWorkspaceServiceServer).WriteIDMapping(ctx, req.(*WriteIDMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InWorkspaceService_MountProc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountProcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InWorkspaceServiceServer).MountProc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iws.InWorkspaceService/MountProc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InWorkspaceServiceServer).MountProc(ctx, req.(*MountProcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InWorkspaceService_Teardown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeardownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InWorkspaceServiceServer).Teardown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iws.InWorkspaceService/Teardown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InWorkspaceServiceServer).Teardown(ctx, req.(*TeardownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InWorkspaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iws.InWorkspaceService",
	HandlerType: (*InWorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareForUserNS",
			Handler:    _InWorkspaceService_PrepareForUserNS_Handler,
		},
		{
			MethodName: "WriteIDMapping",
			Handler:    _InWorkspaceService_WriteIDMapping_Handler,
		},
		{
			MethodName: "MountProc",
			Handler:    _InWorkspaceService_MountProc_Handler,
		},
		{
			MethodName: "Teardown",
			Handler:    _InWorkspaceService_Teardown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace.proto",
}
