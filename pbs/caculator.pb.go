// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: caculator.proto

package pbs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_caculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_caculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SumRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num1          int32                  `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2          int32                  `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	mi := &file_caculator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{2}
}

func (x *SumRequest) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *SumRequest) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sum           int32                  `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	mi := &file_caculator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{3}
}

func (x *SumResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type ToPrimeNumberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToPrimeNumberRequest) Reset() {
	*x = ToPrimeNumberRequest{}
	mi := &file_caculator_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToPrimeNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToPrimeNumberRequest) ProtoMessage() {}

func (x *ToPrimeNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToPrimeNumberRequest.ProtoReflect.Descriptor instead.
func (*ToPrimeNumberRequest) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{4}
}

func (x *ToPrimeNumberRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ToPrimeNumberResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PrimeNumber   int32                  `protobuf:"varint,1,opt,name=prime_number,json=primeNumber,proto3" json:"prime_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToPrimeNumberResponse) Reset() {
	*x = ToPrimeNumberResponse{}
	mi := &file_caculator_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToPrimeNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToPrimeNumberResponse) ProtoMessage() {}

func (x *ToPrimeNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToPrimeNumberResponse.ProtoReflect.Descriptor instead.
func (*ToPrimeNumberResponse) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{5}
}

func (x *ToPrimeNumberResponse) GetPrimeNumber() int32 {
	if x != nil {
		return x.PrimeNumber
	}
	return 0
}

type AverageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num           int32                  `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	mi := &file_caculator_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{6}
}

func (x *AverageRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Average       float32                `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	mi := &file_caculator_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{7}
}

func (x *AverageResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

type FindMaxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num           int32                  `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindMaxRequest) Reset() {
	*x = FindMaxRequest{}
	mi := &file_caculator_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindMaxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxRequest) ProtoMessage() {}

func (x *FindMaxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxRequest.ProtoReflect.Descriptor instead.
func (*FindMaxRequest) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{8}
}

func (x *FindMaxRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type FindMaxResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Max           int32                  `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindMaxResponse) Reset() {
	*x = FindMaxResponse{}
	mi := &file_caculator_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindMaxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxResponse) ProtoMessage() {}

func (x *FindMaxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxResponse.ProtoReflect.Descriptor instead.
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{9}
}

func (x *FindMaxResponse) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type SquareRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num           int32                  `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SquareRequest) Reset() {
	*x = SquareRequest{}
	mi := &file_caculator_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SquareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRequest) ProtoMessage() {}

func (x *SquareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRequest.ProtoReflect.Descriptor instead.
func (*SquareRequest) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{10}
}

func (x *SquareRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type SquareResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Square        float64                `protobuf:"fixed64,1,opt,name=square,proto3" json:"square,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SquareResponse) Reset() {
	*x = SquareResponse{}
	mi := &file_caculator_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SquareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareResponse) ProtoMessage() {}

func (x *SquareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caculator_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareResponse.ProtoReflect.Descriptor instead.
func (*SquareResponse) Descriptor() ([]byte, []int) {
	return file_caculator_proto_rawDescGZIP(), []int{11}
}

func (x *SquareResponse) GetSquare() float64 {
	if x != nil {
		return x.Square
	}
	return 0
}

var File_caculator_proto protoreflect.FileDescriptor

var file_caculator_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x63, 0x61, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x34,
	0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x6e, 0x75, 0x6d, 0x32, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x2e, 0x0a, 0x14, 0x54, 0x6f, 0x50, 0x72, 0x69, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x15, 0x54, 0x6f, 0x50, 0x72, 0x69, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x22, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x23, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x21, 0x0a, 0x0d, 0x53,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x28,
	0x0a, 0x0e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x32, 0x8b, 0x03, 0x0a, 0x10, 0x43, 0x61, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x54, 0x6f, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x54, 0x6f, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x54, 0x6f, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x07, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x40, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x4d, 0x61, 0x78, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x53,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x66,
	0x75, 0x6e, 0x6e, 0x7a, 0x79, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x62, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_caculator_proto_rawDescOnce sync.Once
	file_caculator_proto_rawDescData []byte
)

func file_caculator_proto_rawDescGZIP() []byte {
	file_caculator_proto_rawDescOnce.Do(func() {
		file_caculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_caculator_proto_rawDesc), len(file_caculator_proto_rawDesc)))
	})
	return file_caculator_proto_rawDescData
}

var file_caculator_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_caculator_proto_goTypes = []any{
	(*HelloRequest)(nil),          // 0: protos.HelloRequest
	(*HelloResponse)(nil),         // 1: protos.HelloResponse
	(*SumRequest)(nil),            // 2: protos.SumRequest
	(*SumResponse)(nil),           // 3: protos.SumResponse
	(*ToPrimeNumberRequest)(nil),  // 4: protos.ToPrimeNumberRequest
	(*ToPrimeNumberResponse)(nil), // 5: protos.ToPrimeNumberResponse
	(*AverageRequest)(nil),        // 6: protos.AverageRequest
	(*AverageResponse)(nil),       // 7: protos.AverageResponse
	(*FindMaxRequest)(nil),        // 8: protos.FindMaxRequest
	(*FindMaxResponse)(nil),       // 9: protos.FindMaxResponse
	(*SquareRequest)(nil),         // 10: protos.SquareRequest
	(*SquareResponse)(nil),        // 11: protos.SquareResponse
}
var file_caculator_proto_depIdxs = []int32{
	0,  // 0: protos.CaculatorService.Hello:input_type -> protos.HelloRequest
	2,  // 1: protos.CaculatorService.Sum:input_type -> protos.SumRequest
	4,  // 2: protos.CaculatorService.ToPrimeNumber:input_type -> protos.ToPrimeNumberRequest
	6,  // 3: protos.CaculatorService.Average:input_type -> protos.AverageRequest
	8,  // 4: protos.CaculatorService.FindMax:input_type -> protos.FindMaxRequest
	10, // 5: protos.CaculatorService.Square:input_type -> protos.SquareRequest
	1,  // 6: protos.CaculatorService.Hello:output_type -> protos.HelloResponse
	3,  // 7: protos.CaculatorService.Sum:output_type -> protos.SumResponse
	5,  // 8: protos.CaculatorService.ToPrimeNumber:output_type -> protos.ToPrimeNumberResponse
	7,  // 9: protos.CaculatorService.Average:output_type -> protos.AverageResponse
	9,  // 10: protos.CaculatorService.FindMax:output_type -> protos.FindMaxResponse
	11, // 11: protos.CaculatorService.Square:output_type -> protos.SquareResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_caculator_proto_init() }
func file_caculator_proto_init() {
	if File_caculator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_caculator_proto_rawDesc), len(file_caculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_caculator_proto_goTypes,
		DependencyIndexes: file_caculator_proto_depIdxs,
		MessageInfos:      file_caculator_proto_msgTypes,
	}.Build()
	File_caculator_proto = out.File
	file_caculator_proto_goTypes = nil
	file_caculator_proto_depIdxs = nil
}
