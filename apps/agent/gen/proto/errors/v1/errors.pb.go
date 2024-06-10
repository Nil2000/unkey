// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: proto/errors/v1/errors.proto

package errorsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Fault int32

const (
	Fault_FAULT_UNSPECIFIED Fault = 0
	Fault_FAULT_UNKNOWN     Fault = 1
	Fault_FAULT_PLANETSCALE Fault = 2
	Fault_FAULT_GITHUB      Fault = 3
)

// Enum value maps for Fault.
var (
	Fault_name = map[int32]string{
		0: "FAULT_UNSPECIFIED",
		1: "FAULT_UNKNOWN",
		2: "FAULT_PLANETSCALE",
		3: "FAULT_GITHUB",
	}
	Fault_value = map[string]int32{
		"FAULT_UNSPECIFIED": 0,
		"FAULT_UNKNOWN":     1,
		"FAULT_PLANETSCALE": 2,
		"FAULT_GITHUB":      3,
	}
)

func (x Fault) Enum() *Fault {
	p := new(Fault)
	*p = x
	return p
}

func (x Fault) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Fault) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_errors_v1_errors_proto_enumTypes[0].Descriptor()
}

func (Fault) Type() protoreflect.EnumType {
	return &file_proto_errors_v1_errors_proto_enumTypes[0]
}

func (x Fault) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Fault.Descriptor instead.
func (Fault) EnumDescriptor() ([]byte, []int) {
	return file_proto_errors_v1_errors_proto_rawDescGZIP(), []int{0}
}

type Service int32

const (
	Service_ServiceUnknown    Service = 0
	Service_ServiceAgent      Service = 1
	Service_ServiceAuth       Service = 2
	Service_ServiceCatalog    Service = 3
	Service_ServiceConfig     Service = 4
	Service_ServiceDNS        Service = 5
	Service_ServiceGateway    Service = 6
	Service_ServiceGitHub     Service = 7
	Service_ServiceKubernetes Service = 8
	Service_ServiceLog        Service = 9
	Service_ServiceMetrics    Service = 10
	Service_ServiceMonitor    Service = 11
	Service_ServiceNetwork    Service = 12
	Service_ServiceOperator   Service = 13
	Service_ServiceRegistry   Service = 14
	Service_ServiceSecret     Service = 15
	Service_ServiceStorage    Service = 16
	Service_ServiceSystem     Service = 17
	Service_ServiceTelemetry  Service = 18
	Service_ServiceToken      Service = 19
	Service_ServiceUser       Service = 20
	Service_ServiceVault      Service = 21
	Service_ServiceWebhook    Service = 22
)

// Enum value maps for Service.
var (
	Service_name = map[int32]string{
		0:  "ServiceUnknown",
		1:  "ServiceAgent",
		2:  "ServiceAuth",
		3:  "ServiceCatalog",
		4:  "ServiceConfig",
		5:  "ServiceDNS",
		6:  "ServiceGateway",
		7:  "ServiceGitHub",
		8:  "ServiceKubernetes",
		9:  "ServiceLog",
		10: "ServiceMetrics",
		11: "ServiceMonitor",
		12: "ServiceNetwork",
		13: "ServiceOperator",
		14: "ServiceRegistry",
		15: "ServiceSecret",
		16: "ServiceStorage",
		17: "ServiceSystem",
		18: "ServiceTelemetry",
		19: "ServiceToken",
		20: "ServiceUser",
		21: "ServiceVault",
		22: "ServiceWebhook",
	}
	Service_value = map[string]int32{
		"ServiceUnknown":    0,
		"ServiceAgent":      1,
		"ServiceAuth":       2,
		"ServiceCatalog":    3,
		"ServiceConfig":     4,
		"ServiceDNS":        5,
		"ServiceGateway":    6,
		"ServiceGitHub":     7,
		"ServiceKubernetes": 8,
		"ServiceLog":        9,
		"ServiceMetrics":    10,
		"ServiceMonitor":    11,
		"ServiceNetwork":    12,
		"ServiceOperator":   13,
		"ServiceRegistry":   14,
		"ServiceSecret":     15,
		"ServiceStorage":    16,
		"ServiceSystem":     17,
		"ServiceTelemetry":  18,
		"ServiceToken":      19,
		"ServiceUser":       20,
		"ServiceVault":      21,
		"ServiceWebhook":    22,
	}
)

func (x Service) Enum() *Service {
	p := new(Service)
	*p = x
	return p
}

func (x Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_errors_v1_errors_proto_enumTypes[1].Descriptor()
}

func (Service) Type() protoreflect.EnumType {
	return &file_proto_errors_v1_errors_proto_enumTypes[1]
}

func (x Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service.Descriptor instead.
func (Service) EnumDescriptor() ([]byte, []int) {
	return file_proto_errors_v1_errors_proto_rawDescGZIP(), []int{1}
}

type ErrorCode int32

const (
	ErrorCode_ErrorCodeUnspecified ErrorCode = 0
	ErrorCode_ErrorCodeInternal    ErrorCode = 1
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ErrorCodeUnspecified",
		1: "ErrorCodeInternal",
	}
	ErrorCode_value = map[string]int32{
		"ErrorCodeUnspecified": 0,
		"ErrorCodeInternal":    1,
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
	return file_proto_errors_v1_errors_proto_enumTypes[2].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_proto_errors_v1_errors_proto_enumTypes[2]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_errors_v1_errors_proto_rawDescGZIP(), []int{2}
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         *string `protobuf:"bytes,1,opt,name=url,proto3,oneof" json:"url,omitempty"`
	Label       string  `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_errors_v1_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_proto_errors_v1_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_proto_errors_v1_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *Action) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Action) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fault    Fault            `protobuf:"varint,1,opt,name=fault,proto3,enum=errors.v1.Fault" json:"fault,omitempty"`
	Group    string           `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Code     ErrorCode        `protobuf:"varint,3,opt,name=code,proto3,enum=errors.v1.ErrorCode" json:"code,omitempty"`
	Type     string           `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Metadata *structpb.Struct `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Suggested actions the user should take to resolve this error.
	// These actions are not guaranteed to resolve the error, but they are a good starting point.
	//
	// As a last resort, the user should contact support.
	//
	// The actions are ordered by importance, the first action should be presented first.
	Actions []*Action `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_errors_v1_errors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_errors_v1_errors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_errors_v1_errors_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetFault() Fault {
	if x != nil {
		return x.Fault
	}
	return Fault_FAULT_UNSPECIFIED
}

func (x *Error) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Error) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ErrorCodeUnspecified
}

func (x *Error) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Error) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Error) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_proto_errors_v1_errors_proto protoreflect.FileDescriptor

var file_proto_errors_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0xe5, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x05, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x28, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2a, 0x5a, 0x0a, 0x05, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x50, 0x4c, 0x41,
	0x4e, 0x45, 0x54, 0x53, 0x43, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x5f, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x10, 0x03, 0x2a, 0xc4, 0x03, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x10, 0x02, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x4e, 0x53, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x10, 0x07, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x6f, 0x67, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x10, 0x0c,
	0x12, 0x13, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x10, 0x0f, 0x12, 0x12, 0x0a,
	0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10,
	0x10, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x10, 0x12, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x13, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x10, 0x14, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x15, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x10, 0x16, 0x2a, 0x3c, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x14, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x55, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10,
	0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x75, 0x6e, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_errors_v1_errors_proto_rawDescOnce sync.Once
	file_proto_errors_v1_errors_proto_rawDescData = file_proto_errors_v1_errors_proto_rawDesc
)

func file_proto_errors_v1_errors_proto_rawDescGZIP() []byte {
	file_proto_errors_v1_errors_proto_rawDescOnce.Do(func() {
		file_proto_errors_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_errors_v1_errors_proto_rawDescData)
	})
	return file_proto_errors_v1_errors_proto_rawDescData
}

var file_proto_errors_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_errors_v1_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_errors_v1_errors_proto_goTypes = []interface{}{
	(Fault)(0),              // 0: errors.v1.Fault
	(Service)(0),            // 1: errors.v1.Service
	(ErrorCode)(0),          // 2: errors.v1.ErrorCode
	(*Action)(nil),          // 3: errors.v1.Action
	(*Error)(nil),           // 4: errors.v1.Error
	(*structpb.Struct)(nil), // 5: google.protobuf.Struct
}
var file_proto_errors_v1_errors_proto_depIdxs = []int32{
	0, // 0: errors.v1.Error.fault:type_name -> errors.v1.Fault
	2, // 1: errors.v1.Error.code:type_name -> errors.v1.ErrorCode
	5, // 2: errors.v1.Error.metadata:type_name -> google.protobuf.Struct
	3, // 3: errors.v1.Error.actions:type_name -> errors.v1.Action
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_errors_v1_errors_proto_init() }
func file_proto_errors_v1_errors_proto_init() {
	if File_proto_errors_v1_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_errors_v1_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_proto_errors_v1_errors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
	file_proto_errors_v1_errors_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_errors_v1_errors_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_errors_v1_errors_proto_goTypes,
		DependencyIndexes: file_proto_errors_v1_errors_proto_depIdxs,
		EnumInfos:         file_proto_errors_v1_errors_proto_enumTypes,
		MessageInfos:      file_proto_errors_v1_errors_proto_msgTypes,
	}.Build()
	File_proto_errors_v1_errors_proto = out.File
	file_proto_errors_v1_errors_proto_rawDesc = nil
	file_proto_errors_v1_errors_proto_goTypes = nil
	file_proto_errors_v1_errors_proto_depIdxs = nil
}