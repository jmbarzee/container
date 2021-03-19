// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.6
// source: bitbox.proto

package grpc

import (
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

type StatusReply_Status int32

const (
	// Running indicates that the process is running.
	StatusReply_Running StatusReply_Status = 0
	// Exited indicates that the process returned a non-zero exit code.
	StatusReply_Exited StatusReply_Status = 1
	// Stopped indicates that the process returned no exit code.
	StatusReply_Stopped StatusReply_Status = 2
)

// Enum value maps for StatusReply_Status.
var (
	StatusReply_Status_name = map[int32]string{
		0: "Running",
		1: "Exited",
		2: "Stopped",
	}
	StatusReply_Status_value = map[string]int32{
		"Running": 0,
		"Exited":  1,
		"Stopped": 2,
	}
)

func (x StatusReply_Status) Enum() *StatusReply_Status {
	p := new(StatusReply_Status)
	*p = x
	return p
}

func (x StatusReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_bitbox_proto_enumTypes[0].Descriptor()
}

func (StatusReply_Status) Type() protoreflect.EnumType {
	return &file_bitbox_proto_enumTypes[0]
}

func (x StatusReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusReply_Status.Descriptor instead.
func (StatusReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{5, 0}
}

// StartRequest holds the command and nessecary parameters to run it.
type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Command is an arbitrary linux command.
	// Command should be a fullpath or found in a directory in PATH.
	Command string `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	// Parameters are handed to the command as ARGV.
	// Glob patterns, pipelines, and redirections are not handled
	// and will instead be passed as literal strings to the command.
	Parameters []string `protobuf:"bytes,2,rep,name=Parameters,proto3" json:"Parameters,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *StartRequest) GetParameters() []string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

// StartReply holds the ID of the started process.
type StartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of the process.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *StartReply) Reset() {
	*x = StartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartReply) ProtoMessage() {}

func (x *StartReply) ProtoReflect() protoreflect.Message {
	mi := &file_bitbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartReply.ProtoReflect.Descriptor instead.
func (*StartReply) Descriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{1}
}

func (x *StartReply) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// StopRequest holds the ID of the process that should be stopped.
type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of the process.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"` // TODO: Consider making the type of signal sent to the process configurable, not just kill.
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{2}
}

func (x *StopRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// StopReply is empty.
type StopReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopReply) Reset() {
	*x = StopReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopReply) ProtoMessage() {}

func (x *StopReply) ProtoReflect() protoreflect.Message {
	mi := &file_bitbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopReply.ProtoReflect.Descriptor instead.
func (*StopReply) Descriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{3}
}

// StatusRequest holds the ID of the process to check the status of
type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of the process.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitbox_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitbox_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{4}
}

func (x *StatusRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// StatusReply holds the status of the reply
type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitbox_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_bitbox_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{5}
}

// QueryRequest holds the ID of the process whose output is being queried.
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of the process.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitbox_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitbox_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{6}
}

func (x *QueryRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// QueryReply holds output of the process queried.
type QueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output is one of the forms of output that a process may utilize.
	//
	// Types that are assignable to Output:
	//	*QueryReply_Stdout
	//	*QueryReply_Stderr
	//	*QueryReply_ExitCode
	Output isQueryReply_Output `protobuf_oneof:"Output"`
}

func (x *QueryReply) Reset() {
	*x = QueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitbox_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReply) ProtoMessage() {}

func (x *QueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_bitbox_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReply.ProtoReflect.Descriptor instead.
func (*QueryReply) Descriptor() ([]byte, []int) {
	return file_bitbox_proto_rawDescGZIP(), []int{7}
}

func (m *QueryReply) GetOutput() isQueryReply_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (x *QueryReply) GetStdout() string {
	if x, ok := x.GetOutput().(*QueryReply_Stdout); ok {
		return x.Stdout
	}
	return ""
}

func (x *QueryReply) GetStderr() string {
	if x, ok := x.GetOutput().(*QueryReply_Stderr); ok {
		return x.Stderr
	}
	return ""
}

func (x *QueryReply) GetExitCode() uint32 {
	if x, ok := x.GetOutput().(*QueryReply_ExitCode); ok {
		return x.ExitCode
	}
	return 0
}

type isQueryReply_Output interface {
	isQueryReply_Output()
}

type QueryReply_Stdout struct {
	// Stdout is any output from the process which was written to Stdout.
	Stdout string `protobuf:"bytes,1,opt,name=Stdout,proto3,oneof"`
}

type QueryReply_Stderr struct {
	// Stderr is any output from the process which was written to Stderr.
	Stderr string `protobuf:"bytes,2,opt,name=Stderr,proto3,oneof"`
}

type QueryReply_ExitCode struct {
	// ExitCode is the code the process exited with.
	ExitCode uint32 `protobuf:"varint,3,opt,name=ExitCode,proto3,oneof"`
}

func (*QueryReply_Stdout) isQueryReply_Output() {}

func (*QueryReply_Stderr) isQueryReply_Output() {}

func (*QueryReply_ExitCode) isQueryReply_Output() {}

var File_bitbox_proto protoreflect.FileDescriptor

var file_bitbox_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x69, 0x74, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x48, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x1c,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1d, 0x0a, 0x0b,
	0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x0b, 0x0a, 0x09, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3d, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x10, 0x02, 0x22, 0x1e, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x68, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x06, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x12, 0x18, 0x0a, 0x06, 0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x08, 0x45, 0x78,
	0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08,
	0x45, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x32, 0xce, 0x01, 0x0a, 0x06, 0x42, 0x69, 0x74, 0x42, 0x6f, 0x78, 0x12, 0x2f, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2c,
	0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6a, 0x6d, 0x62, 0x61, 0x72, 0x7a, 0x65, 0x65, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x6f,
	0x78, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bitbox_proto_rawDescOnce sync.Once
	file_bitbox_proto_rawDescData = file_bitbox_proto_rawDesc
)

func file_bitbox_proto_rawDescGZIP() []byte {
	file_bitbox_proto_rawDescOnce.Do(func() {
		file_bitbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_bitbox_proto_rawDescData)
	})
	return file_bitbox_proto_rawDescData
}

var file_bitbox_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bitbox_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bitbox_proto_goTypes = []interface{}{
	(StatusReply_Status)(0), // 0: grpc.StatusReply.Status
	(*StartRequest)(nil),    // 1: grpc.StartRequest
	(*StartReply)(nil),      // 2: grpc.StartReply
	(*StopRequest)(nil),     // 3: grpc.StopRequest
	(*StopReply)(nil),       // 4: grpc.StopReply
	(*StatusRequest)(nil),   // 5: grpc.StatusRequest
	(*StatusReply)(nil),     // 6: grpc.StatusReply
	(*QueryRequest)(nil),    // 7: grpc.QueryRequest
	(*QueryReply)(nil),      // 8: grpc.QueryReply
}
var file_bitbox_proto_depIdxs = []int32{
	1, // 0: grpc.BitBox.Start:input_type -> grpc.StartRequest
	3, // 1: grpc.BitBox.Stop:input_type -> grpc.StopRequest
	5, // 2: grpc.BitBox.Status:input_type -> grpc.StatusRequest
	7, // 3: grpc.BitBox.Query:input_type -> grpc.QueryRequest
	2, // 4: grpc.BitBox.Start:output_type -> grpc.StartReply
	4, // 5: grpc.BitBox.Stop:output_type -> grpc.StopReply
	6, // 6: grpc.BitBox.Status:output_type -> grpc.StatusReply
	8, // 7: grpc.BitBox.Query:output_type -> grpc.QueryReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bitbox_proto_init() }
func file_bitbox_proto_init() {
	if File_bitbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bitbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_bitbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartReply); i {
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
		file_bitbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_bitbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopReply); i {
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
		file_bitbox_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_bitbox_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
		file_bitbox_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_bitbox_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReply); i {
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
	file_bitbox_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*QueryReply_Stdout)(nil),
		(*QueryReply_Stderr)(nil),
		(*QueryReply_ExitCode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bitbox_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bitbox_proto_goTypes,
		DependencyIndexes: file_bitbox_proto_depIdxs,
		EnumInfos:         file_bitbox_proto_enumTypes,
		MessageInfos:      file_bitbox_proto_msgTypes,
	}.Build()
	File_bitbox_proto = out.File
	file_bitbox_proto_rawDesc = nil
	file_bitbox_proto_goTypes = nil
	file_bitbox_proto_depIdxs = nil
}