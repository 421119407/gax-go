// Copyright 2021 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: apierror/internal/proto/error.proto

package jsonerror

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The error format v2 for Google JSON REST APIs.
// Copied from https://cloud.google.com/apis/design/errors#http_mapping.
//
// NOTE: This schema is not used for other wire protocols.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actual error payload. The nested message structure is for backward
	// compatibility with Google API client libraries. It also makes the error
	// more readable to developers.
	Error *Error_Status `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apierror_internal_proto_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_apierror_internal_proto_error_proto_msgTypes[0]
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
	return file_apierror_internal_proto_error_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetError() *Error_Status {
	if x != nil {
		return x.Error
	}
	return nil
}

// This message has the same semantics as `google.rpc.Status`. It uses HTTP
// status code instead of gRPC status code. It has an extra field `status`
// for backward compatibility with Google API Client Libraries.
type Error_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP status code that corresponds to `google.rpc.Status.code`.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// This corresponds to `google.rpc.Status.message`.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// This is the enum version for `google.rpc.Status.code`.
	Status code.Code `protobuf:"varint,4,opt,name=status,proto3,enum=google.rpc.Code" json:"status,omitempty"`
	// This corresponds to `google.rpc.Status.details`.
	Details []*anypb.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Error_Status) Reset() {
	*x = Error_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apierror_internal_proto_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error_Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error_Status) ProtoMessage() {}

func (x *Error_Status) ProtoReflect() protoreflect.Message {
	mi := &file_apierror_internal_proto_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error_Status.ProtoReflect.Descriptor instead.
func (*Error_Status) Descriptor() ([]byte, []int) {
	return file_apierror_internal_proto_error_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Error_Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error_Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error_Status) GetStatus() code.Code {
	if x != nil {
		return x.Status
	}
	return code.Code(0)
}

func (x *Error_Status) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_apierror_internal_proto_error_proto protoreflect.FileDescriptor

var file_apierror_internal_proto_error_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5,
	0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x1a, 0x90, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x61, 0x78, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x6a, 0x73, 0x6f, 0x6e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apierror_internal_proto_error_proto_rawDescOnce sync.Once
	file_apierror_internal_proto_error_proto_rawDescData = file_apierror_internal_proto_error_proto_rawDesc
)

func file_apierror_internal_proto_error_proto_rawDescGZIP() []byte {
	file_apierror_internal_proto_error_proto_rawDescOnce.Do(func() {
		file_apierror_internal_proto_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_apierror_internal_proto_error_proto_rawDescData)
	})
	return file_apierror_internal_proto_error_proto_rawDescData
}

var file_apierror_internal_proto_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apierror_internal_proto_error_proto_goTypes = []interface{}{
	(*Error)(nil),        // 0: error.Error
	(*Error_Status)(nil), // 1: error.Error.Status
	(code.Code)(0),       // 2: google.rpc.Code
	(*anypb.Any)(nil),    // 3: google.protobuf.Any
}
var file_apierror_internal_proto_error_proto_depIdxs = []int32{
	1, // 0: error.Error.error:type_name -> error.Error.Status
	2, // 1: error.Error.Status.status:type_name -> google.rpc.Code
	3, // 2: error.Error.Status.details:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apierror_internal_proto_error_proto_init() }
func file_apierror_internal_proto_error_proto_init() {
	if File_apierror_internal_proto_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apierror_internal_proto_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_apierror_internal_proto_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error_Status); i {
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
			RawDescriptor: file_apierror_internal_proto_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apierror_internal_proto_error_proto_goTypes,
		DependencyIndexes: file_apierror_internal_proto_error_proto_depIdxs,
		MessageInfos:      file_apierror_internal_proto_error_proto_msgTypes,
	}.Build()
	File_apierror_internal_proto_error_proto = out.File
	file_apierror_internal_proto_error_proto_rawDesc = nil
	file_apierror_internal_proto_error_proto_goTypes = nil
	file_apierror_internal_proto_error_proto_depIdxs = nil
}
