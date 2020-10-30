// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: darkbase/darkbase.proto

package darkbase

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ErrorCode int32

const (
	ErrorCode_Success          ErrorCode = 0
	ErrorCode_NotFound         ErrorCode = 5404
	ErrorCode_UnknownError     ErrorCode = 5300
	ErrorCode_ResponseRequired ErrorCode = 5301
	ErrorCode_HugePipeline     ErrorCode = 5302
	ErrorCode_ServerTimeout    ErrorCode = 5303
	// table is nil or table not in cluster
	ErrorCode_TableNilOrNotMatch ErrorCode = 5304
	ErrorCode_ClientClosed       ErrorCode = 5305
	// parameters error
	ErrorCode_InvalidArgument ErrorCode = 5308
	// probably parse int failed
	ErrorCode_InvalidResponseType ErrorCode = 5309
	ErrorCode_ServerBroken        ErrorCode = 5310
	ErrorCode_GdprAuthFail        ErrorCode = 5311
	ErrorCode_QueueChaos          ErrorCode = 5312
	// parameters error
	ErrorCode_BadRequest ErrorCode = 5400
	// big value
	ErrorCode_HugeFrame        ErrorCode = 5401
	ErrorCode_MethodNotAllowed ErrorCode = 5405
	// generation mismatch. in xget/xset
	ErrorCode_InvalidGeneration   ErrorCode = 5497
	ErrorCode_InternalServerError ErrorCode = 5500
	// method not implemented in server
	ErrorCode_NotImplemented     ErrorCode = 5501
	ErrorCode_ServiceUnavailable ErrorCode = 5503
	// for scan
	ErrorCode_ScanTimeout ErrorCode = 5510
	// task queue is full
	ErrorCode_TaskQueueFull ErrorCode = 5511
	// slave is stale
	ErrorCode_OutOfSync ErrorCode = 5599
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:    "Success",
		5404: "NotFound",
		5300: "UnknownError",
		5301: "ResponseRequired",
		5302: "HugePipeline",
		5303: "ServerTimeout",
		5304: "TableNilOrNotMatch",
		5305: "ClientClosed",
		5308: "InvalidArgument",
		5309: "InvalidResponseType",
		5310: "ServerBroken",
		5311: "GdprAuthFail",
		5312: "QueueChaos",
		5400: "BadRequest",
		5401: "HugeFrame",
		5405: "MethodNotAllowed",
		5497: "InvalidGeneration",
		5500: "InternalServerError",
		5501: "NotImplemented",
		5503: "ServiceUnavailable",
		5510: "ScanTimeout",
		5511: "TaskQueueFull",
		5599: "OutOfSync",
	}
	ErrorCode_value = map[string]int32{
		"Success":             0,
		"NotFound":            5404,
		"UnknownError":        5300,
		"ResponseRequired":    5301,
		"HugePipeline":        5302,
		"ServerTimeout":       5303,
		"TableNilOrNotMatch":  5304,
		"ClientClosed":        5305,
		"InvalidArgument":     5308,
		"InvalidResponseType": 5309,
		"ServerBroken":        5310,
		"GdprAuthFail":        5311,
		"QueueChaos":          5312,
		"BadRequest":          5400,
		"HugeFrame":           5401,
		"MethodNotAllowed":    5405,
		"InvalidGeneration":   5497,
		"InternalServerError": 5500,
		"NotImplemented":      5501,
		"ServiceUnavailable":  5503,
		"ScanTimeout":         5510,
		"TaskQueueFull":       5511,
		"OutOfSync":           5599,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_darkbase_darkbase_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_darkbase_darkbase_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_darkbase_darkbase_proto_rawDescGZIP(), []int{0}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	TableName string `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkbase_darkbase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_darkbase_darkbase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_darkbase_darkbase_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode ErrorCode `protobuf:"varint,1,opt,name=errCode,proto3,enum=ErrorCode" json:"errCode,omitempty"`
	Value   []byte    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkbase_darkbase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_darkbase_darkbase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_darkbase_darkbase_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetErrCode() ErrorCode {
	if x != nil {
		return x.ErrCode
	}
	return ErrorCode_Success
}

func (x *GetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type BatchGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string        `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Reqs      []*GetRequest `protobuf:"bytes,2,rep,name=reqs,proto3" json:"reqs,omitempty"`
}

func (x *BatchGetRequest) Reset() {
	*x = BatchGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkbase_darkbase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetRequest) ProtoMessage() {}

func (x *BatchGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_darkbase_darkbase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetRequest.ProtoReflect.Descriptor instead.
func (*BatchGetRequest) Descriptor() ([]byte, []int) {
	return file_darkbase_darkbase_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGetRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *BatchGetRequest) GetReqs() []*GetRequest {
	if x != nil {
		return x.Reqs
	}
	return nil
}

type BatchGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp []*GetResponse `protobuf:"bytes,1,rep,name=resp,proto3" json:"resp,omitempty"`
}

func (x *BatchGetResponse) Reset() {
	*x = BatchGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkbase_darkbase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetResponse) ProtoMessage() {}

func (x *BatchGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_darkbase_darkbase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetResponse.ProtoReflect.Descriptor instead.
func (*BatchGetResponse) Descriptor() ([]byte, []int) {
	return file_darkbase_darkbase_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetResponse) GetResp() []*GetResponse {
	if x != nil {
		return x.Resp
	}
	return nil
}

type BatchPutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string        `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Req       []*PutRequest `protobuf:"bytes,2,rep,name=req,proto3" json:"req,omitempty"`
}

func (x *BatchPutRequest) Reset() {
	*x = BatchPutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkbase_darkbase_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchPutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchPutRequest) ProtoMessage() {}

func (x *BatchPutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_darkbase_darkbase_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchPutRequest.ProtoReflect.Descriptor instead.
func (*BatchPutRequest) Descriptor() ([]byte, []int) {
	return file_darkbase_darkbase_proto_rawDescGZIP(), []int{4}
}

func (x *BatchPutRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *BatchPutRequest) GetReq() []*PutRequest {
	if x != nil {
		return x.Req
	}
	return nil
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	TableName string `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Value     []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkbase_darkbase_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_darkbase_darkbase_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_darkbase_darkbase_proto_rawDescGZIP(), []int{5}
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *PutRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode ErrorCode `protobuf:"varint,1,opt,name=errCode,proto3,enum=ErrorCode" json:"errCode,omitempty"`
}

func (x *PutResponse) Reset() {
	*x = PutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkbase_darkbase_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponse) ProtoMessage() {}

func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_darkbase_darkbase_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponse.ProtoReflect.Descriptor instead.
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_darkbase_darkbase_proto_rawDescGZIP(), []int{6}
}

func (x *PutResponse) GetErrCode() ErrorCode {
	if x != nil {
		return x.ErrCode
	}
	return ErrorCode_Success
}

var File_darkbase_darkbase_proto protoreflect.FileDescriptor

var file_darkbase_darkbase_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x61, 0x72, 0x6b, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x72, 0x6b, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x51, 0x0a, 0x0f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x65, 0x71, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x72, 0x65, 0x71, 0x73, 0x22, 0x34, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x72, 0x65,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x22, 0x4f, 0x0a, 0x0f,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x53, 0x0a,
	0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x33, 0x0a, 0x0b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0xd9, 0x03, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x9c,
	0x2a, 0x12, 0x11, 0x0a, 0x0c, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0xb4, 0x29, 0x12, 0x15, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x10, 0xb5, 0x29, 0x12, 0x11, 0x0a, 0x0c, 0x48,
	0x75, 0x67, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0xb6, 0x29, 0x12, 0x12,
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10,
	0xb7, 0x29, 0x12, 0x17, 0x0a, 0x12, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x69, 0x6c, 0x4f, 0x72,
	0x4e, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0xb8, 0x29, 0x12, 0x11, 0x0a, 0x0c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x10, 0xb9, 0x29, 0x12, 0x14,
	0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x10, 0xbc, 0x29, 0x12, 0x18, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0xbd, 0x29, 0x12, 0x11,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0xbe,
	0x29, 0x12, 0x11, 0x0a, 0x0c, 0x47, 0x64, 0x70, 0x72, 0x41, 0x75, 0x74, 0x68, 0x46, 0x61, 0x69,
	0x6c, 0x10, 0xbf, 0x29, 0x12, 0x0f, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x68, 0x61,
	0x6f, 0x73, 0x10, 0xc0, 0x29, 0x12, 0x0f, 0x0a, 0x0a, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x10, 0x98, 0x2a, 0x12, 0x0e, 0x0a, 0x09, 0x48, 0x75, 0x67, 0x65, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x10, 0x99, 0x2a, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4e, 0x6f, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x10, 0x9d, 0x2a, 0x12, 0x16, 0x0a,
	0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0xf9, 0x2a, 0x12, 0x18, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xfc, 0x2a, 0x12,
	0x13, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65,
	0x64, 0x10, 0xfd, 0x2a, 0x12, 0x17, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55,
	0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0xff, 0x2a, 0x12, 0x10, 0x0a,
	0x0b, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x86, 0x2b, 0x12,
	0x12, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x46, 0x75, 0x6c, 0x6c,
	0x10, 0x87, 0x2b, 0x12, 0x0e, 0x0a, 0x09, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x53, 0x79, 0x6e, 0x63,
	0x10, 0xdf, 0x2b, 0x32, 0xb6, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x72, 0x6b, 0x62, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x50, 0x75, 0x74,
	0x12, 0x0b, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_darkbase_darkbase_proto_rawDescOnce sync.Once
	file_darkbase_darkbase_proto_rawDescData = file_darkbase_darkbase_proto_rawDesc
)

func file_darkbase_darkbase_proto_rawDescGZIP() []byte {
	file_darkbase_darkbase_proto_rawDescOnce.Do(func() {
		file_darkbase_darkbase_proto_rawDescData = protoimpl.X.CompressGZIP(file_darkbase_darkbase_proto_rawDescData)
	})
	return file_darkbase_darkbase_proto_rawDescData
}

var file_darkbase_darkbase_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_darkbase_darkbase_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_darkbase_darkbase_proto_goTypes = []interface{}{
	(ErrorCode)(0),           // 0: ErrorCode
	(*GetRequest)(nil),       // 1: GetRequest
	(*GetResponse)(nil),      // 2: GetResponse
	(*BatchGetRequest)(nil),  // 3: BatchGetRequest
	(*BatchGetResponse)(nil), // 4: BatchGetResponse
	(*BatchPutRequest)(nil),  // 5: BatchPutRequest
	(*PutRequest)(nil),       // 6: PutRequest
	(*PutResponse)(nil),      // 7: PutResponse
}
var file_darkbase_darkbase_proto_depIdxs = []int32{
	0, // 0: GetResponse.errCode:type_name -> ErrorCode
	1, // 1: BatchGetRequest.reqs:type_name -> GetRequest
	2, // 2: BatchGetResponse.resp:type_name -> GetResponse
	6, // 3: BatchPutRequest.req:type_name -> PutRequest
	0, // 4: PutResponse.errCode:type_name -> ErrorCode
	1, // 5: DarkbaseService.Get:input_type -> GetRequest
	6, // 6: DarkbaseService.Put:input_type -> PutRequest
	3, // 7: DarkbaseService.BatchGet:input_type -> BatchGetRequest
	5, // 8: DarkbaseService.BatchPut:input_type -> BatchPutRequest
	2, // 9: DarkbaseService.Get:output_type -> GetResponse
	7, // 10: DarkbaseService.Put:output_type -> PutResponse
	4, // 11: DarkbaseService.BatchGet:output_type -> BatchGetResponse
	5, // 12: DarkbaseService.BatchPut:output_type -> BatchPutRequest
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_darkbase_darkbase_proto_init() }
func file_darkbase_darkbase_proto_init() {
	if File_darkbase_darkbase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_darkbase_darkbase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_darkbase_darkbase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_darkbase_darkbase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetRequest); i {
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
		file_darkbase_darkbase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetResponse); i {
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
		file_darkbase_darkbase_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchPutRequest); i {
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
		file_darkbase_darkbase_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_darkbase_darkbase_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_darkbase_darkbase_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_darkbase_darkbase_proto_goTypes,
		DependencyIndexes: file_darkbase_darkbase_proto_depIdxs,
		EnumInfos:         file_darkbase_darkbase_proto_enumTypes,
		MessageInfos:      file_darkbase_darkbase_proto_msgTypes,
	}.Build()
	File_darkbase_darkbase_proto = out.File
	file_darkbase_darkbase_proto_rawDesc = nil
	file_darkbase_darkbase_proto_goTypes = nil
	file_darkbase_darkbase_proto_depIdxs = nil
}