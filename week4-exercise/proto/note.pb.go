// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/note.proto

/*
Package note is a generated protocol buffer package.

It is generated from these files:
	proto/note.proto

It has these top-level messages:
	Note
	NoteReq
	NoteFindReq
	NoteDelReq
	NoteDelRes
*/
package note

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Note struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Completed bool   `protobuf:"varint,3,opt,name=completed" json:"completed,omitempty"`
	// bool complete = 3; // Only change the name
	// int32 completed = 3; // Change data type
	CreatedAt   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Description string                     `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
}

func (m *Note) Reset()                    { *m = Note{} }
func (m *Note) String() string            { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()               {}
func (*Note) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Note) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Note) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Note) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Note) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type NoteReq struct {
	Title     string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Completed bool   `protobuf:"varint,2,opt,name=completed" json:"completed,omitempty"`
}

func (m *NoteReq) Reset()                    { *m = NoteReq{} }
func (m *NoteReq) String() string            { return proto.CompactTextString(m) }
func (*NoteReq) ProtoMessage()               {}
func (*NoteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NoteReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NoteReq) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type NoteFindReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *NoteFindReq) Reset()                    { *m = NoteFindReq{} }
func (m *NoteFindReq) String() string            { return proto.CompactTextString(m) }
func (*NoteFindReq) ProtoMessage()               {}
func (*NoteFindReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NoteFindReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NoteDelReq struct {
	Id int32 `protobuf:"varint,5,opt,name=id" json:"id,omitempty"`
}

func (m *NoteDelReq) Reset()                    { *m = NoteDelReq{} }
func (m *NoteDelReq) String() string            { return proto.CompactTextString(m) }
func (*NoteDelReq) ProtoMessage()               {}
func (*NoteDelReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NoteDelReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NoteDelRes struct {
	Success      bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
}

func (m *NoteDelRes) Reset()                    { *m = NoteDelRes{} }
func (m *NoteDelRes) String() string            { return proto.CompactTextString(m) }
func (*NoteDelRes) ProtoMessage()               {}
func (*NoteDelRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NoteDelRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *NoteDelRes) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Note)(nil), "note.Note")
	proto.RegisterType((*NoteReq)(nil), "note.NoteReq")
	proto.RegisterType((*NoteFindReq)(nil), "note.NoteFindReq")
	proto.RegisterType((*NoteDelReq)(nil), "note.NoteDelReq")
	proto.RegisterType((*NoteDelRes)(nil), "note.NoteDelRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NoteService service

type NoteServiceClient interface {
	Create(ctx context.Context, in *NoteReq, opts ...grpc.CallOption) (*Note, error)
	Find(ctx context.Context, in *NoteFindReq, opts ...grpc.CallOption) (*Note, error)
	// rpc Find(NoteFindReq) returns (Note) {
	//   option (google.api.http) = {
	//     get: "/v1/note/{id}"
	//   };
	// }
	Delete(ctx context.Context, in *NoteDelReq, opts ...grpc.CallOption) (*NoteDelRes, error)
}

type noteServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoteServiceClient(cc *grpc.ClientConn) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) Create(ctx context.Context, in *NoteReq, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := grpc.Invoke(ctx, "/note.NoteService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) Find(ctx context.Context, in *NoteFindReq, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := grpc.Invoke(ctx, "/note.NoteService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) Delete(ctx context.Context, in *NoteDelReq, opts ...grpc.CallOption) (*NoteDelRes, error) {
	out := new(NoteDelRes)
	err := grpc.Invoke(ctx, "/note.NoteService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NoteService service

type NoteServiceServer interface {
	Create(context.Context, *NoteReq) (*Note, error)
	Find(context.Context, *NoteFindReq) (*Note, error)
	// rpc Find(NoteFindReq) returns (Note) {
	//   option (google.api.http) = {
	//     get: "/v1/note/{id}"
	//   };
	// }
	Delete(context.Context, *NoteDelReq) (*NoteDelRes, error)
}

func RegisterNoteServiceServer(s *grpc.Server, srv NoteServiceServer) {
	s.RegisterService(&_NoteService_serviceDesc, srv)
}

func _NoteService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Create(ctx, req.(*NoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteFindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Find(ctx, req.(*NoteFindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Delete(ctx, req.(*NoteDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "note.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NoteService_Create_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _NoteService_Find_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NoteService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/note.proto",
}

func init() { proto.RegisterFile("proto/note.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0x86, 0x2b, 0xd7, 0x76, 0x92, 0x49, 0x52, 0x52, 0xd1, 0x83, 0x31, 0x29, 0x35, 0x2e, 0xa5,
	0xa6, 0x07, 0x1b, 0xd2, 0x53, 0x0f, 0x85, 0x86, 0x86, 0x5e, 0x4a, 0x7b, 0x70, 0x7b, 0x0f, 0x8e,
	0x35, 0x0d, 0x02, 0xdb, 0x72, 0x2d, 0xa5, 0x97, 0xb2, 0x97, 0x7d, 0x85, 0x7d, 0x80, 0x7d, 0xa8,
	0x7d, 0x83, 0x65, 0x1f, 0x64, 0x91, 0x6c, 0xaf, 0x83, 0x61, 0xd9, 0x9b, 0xe6, 0x9f, 0x7f, 0xe4,
	0xff, 0x1b, 0x0b, 0x56, 0x75, 0x23, 0x94, 0x48, 0x2a, 0xa1, 0x30, 0x36, 0x47, 0x6a, 0xeb, 0xb3,
	0xff, 0xe6, 0x28, 0xc4, 0xb1, 0xc0, 0xc4, 0x68, 0x87, 0xd3, 0x9f, 0x44, 0xf1, 0x12, 0xa5, 0xca,
	0xca, 0xba, 0xb5, 0xf9, 0xeb, 0xce, 0x90, 0xd5, 0x3c, 0xc9, 0xaa, 0x4a, 0xa8, 0x4c, 0x71, 0x51,
	0xc9, 0xb6, 0x1b, 0xde, 0x12, 0xb0, 0x7f, 0x0a, 0x85, 0xf4, 0x05, 0x58, 0x9c, 0x79, 0x24, 0x20,
	0x91, 0x93, 0x5a, 0x9c, 0xd1, 0x57, 0xe0, 0x28, 0xae, 0x0a, 0xf4, 0xac, 0x80, 0x44, 0xb3, 0xb4,
	0x2d, 0xe8, 0x1a, 0x66, 0xb9, 0x28, 0xeb, 0x02, 0x15, 0x32, 0xef, 0x79, 0x40, 0xa2, 0x69, 0x3a,
	0x08, 0xf4, 0x13, 0x40, 0xde, 0x60, 0xa6, 0x90, 0xed, 0x33, 0xe5, 0xd9, 0x01, 0x89, 0xe6, 0x1b,
	0x3f, 0x6e, 0xbf, 0x1f, 0xf7, 0x01, 0xe3, 0xdf, 0x7d, 0xc0, 0x74, 0xd6, 0xb9, 0xb7, 0x4a, 0x8f,
	0x9e, 0x6a, 0xd6, 0x8f, 0x3a, 0x4f, 0x8f, 0x76, 0xee, 0xad, 0xa2, 0x01, 0xcc, 0x19, 0xca, 0xbc,
	0xe1, 0xb5, 0x06, 0xf3, 0x5c, 0x93, 0xf7, 0x5c, 0x0a, 0x3f, 0xc3, 0x44, 0x33, 0xa6, 0xf8, 0x77,
	0xc0, 0x22, 0x8f, 0x62, 0x59, 0x23, 0xac, 0xf0, 0x35, 0xcc, 0xf5, 0xf8, 0x37, 0x5e, 0x31, 0x7d,
	0xc5, 0x68, 0x53, 0xe1, 0x1a, 0x40, 0xb7, 0x77, 0x58, 0x0c, 0x5d, 0xe7, 0xa1, 0xfb, 0xfd, 0xac,
	0x2b, 0xa9, 0x07, 0x13, 0x79, 0xca, 0x73, 0x94, 0xd2, 0x5c, 0x30, 0x4d, 0xfb, 0x92, 0xbe, 0x85,
	0x25, 0x36, 0x8d, 0x68, 0xf6, 0x25, 0x4a, 0x99, 0x1d, 0xfb, 0xbd, 0x2f, 0x8c, 0xf8, 0xa3, 0xd5,
	0x36, 0xd7, 0xa4, 0x8d, 0xf2, 0x0b, 0x9b, 0x7f, 0x3c, 0x47, 0xfa, 0x0e, 0xdc, 0xaf, 0x66, 0x85,
	0x74, 0x19, 0x9b, 0x97, 0xd1, 0x61, 0xfa, 0x30, 0x94, 0xe1, 0x33, 0xfa, 0x1e, 0x6c, 0x1d, 0x9e,
	0xbe, 0x1c, 0xd4, 0x0e, 0x66, 0x64, 0xfc, 0x02, 0xee, 0x0e, 0x35, 0x34, 0x5d, 0x0d, 0x7a, 0x0b,
	0xe6, 0x8f, 0x15, 0x19, 0xd2, 0xcb, 0x9b, 0xbb, 0x2b, 0x6b, 0xf1, 0x01, 0xcc, 0xab, 0x4c, 0xfe,
	0x73, 0x76, 0x71, 0x70, 0xcd, 0xbf, 0xfa, 0x78, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x12, 0x23, 0xa4,
	0xdf, 0xaf, 0x02, 0x00, 0x00,
}
