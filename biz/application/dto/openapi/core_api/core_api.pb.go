// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.27.3
// source: core_api.proto

package core_api

import (
	_ "github.com/xh-polaris/openapi-core-api/biz/application/dto/http"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_core_api_proto protoreflect.FileDescriptor

var file_core_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x1a, 0x1d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x5e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x71, 0x1a, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f,
	0x75, 0x70, 0x32, 0xa8, 0x04, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5e, 0x0a, 0x0b, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x6b, 0x65,
	0x79, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x0c, 0xca, 0xc1, 0x18, 0x08, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x67, 0x65, 0x74,
	0x12, 0x58, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0xd2, 0xc1, 0x18, 0x0b, 0x2f,
	0x6b, 0x65, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xca, 0xc1, 0x18, 0x0c, 0x2f, 0x6b, 0x65, 0x79, 0x2f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x5a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x0e, 0xd2, 0xc1, 0x18, 0x0a, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x58, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0xca, 0xc1,
	0x18, 0x0b, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x32, 0xa6, 0x0a,
	0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x6d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x28, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x6d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x6d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x10, 0xca, 0xc1, 0x18, 0x0c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x73, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0xca, 0xc1, 0x18,
	0x09, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x6d, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x28, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x66, 0x75,
	0x6c, 0x6c, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x6d, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x28, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x66, 0x75, 0x6c,
	0x6c, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x66, 0x75,
	0x6c, 0x6c, 0x2f, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x6d, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x28, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xca, 0xc1, 0x18, 0x0c, 0x2f, 0x66, 0x75, 0x6c,
	0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x73, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x26, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d,
	0xca, 0xc1, 0x18, 0x09, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x67, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f,
	0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x65, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x11, 0xca, 0xc1, 0x18, 0x0d, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x32, 0x6c, 0x0a, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x64,
	0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x22, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0a, 0xd2, 0xc1, 0x18, 0x06, 0x2f, 0x63,
	0x61, 0x6c, 0x6c, 0x2f, 0x42, 0x83, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70,
	0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x61, 0x70, 0x69, 0x42, 0x0c, 0x43,
	0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c,
	0x61, 0x72, 0x69, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_core_api_proto_goTypes = []interface{}{
	(*SignUpReq)(nil),              // 0: openapi.core_api.SignUpReq
	(*GenerateKeyReq)(nil),         // 1: openapi.core_api.GenerateKeyReq
	(*GetKeysReq)(nil),             // 2: openapi.core_api.GetKeysReq
	(*UpdateKeyReq)(nil),           // 3: openapi.core_api.UpdateKeyReq
	(*RefreshKeyReq)(nil),          // 4: openapi.core_api.RefreshKeyReq
	(*UpdateHostReq)(nil),          // 5: openapi.core_api.UpdateHostReq
	(*DeleteKeyReq)(nil),           // 6: openapi.core_api.DeleteKeyReq
	(*CreateBaseInterfaceReq)(nil), // 7: openapi.core_api.CreateBaseInterfaceReq
	(*UpdateBaseInterfaceReq)(nil), // 8: openapi.core_api.UpdateBaseInterfaceReq
	(*DeleteBaseInterfaceReq)(nil), // 9: openapi.core_api.DeleteBaseInterfaceReq
	(*GetBaseInterfacesReq)(nil),   // 10: openapi.core_api.GetBaseInterfacesReq
	(*CreateFullInterfaceReq)(nil), // 11: openapi.core_api.CreateFullInterfaceReq
	(*UpdateFullInterfaceReq)(nil), // 12: openapi.core_api.UpdateFullInterfaceReq
	(*UpdateMarginReq)(nil),        // 13: openapi.core_api.UpdateMarginReq
	(*DeleteFullInterfaceReq)(nil), // 14: openapi.core_api.DeleteFullInterfaceReq
	(*GetFullInterfacesReq)(nil),   // 15: openapi.core_api.GetFullInterfacesReq
	(*CreateGradientReq)(nil),      // 16: openapi.core_api.CreateGradientReq
	(*UpdateGradientReq)(nil),      // 17: openapi.core_api.UpdateGradientReq
	(*GetGradientReq)(nil),         // 18: openapi.core_api.GetGradientReq
	(*CallInterfaceReq)(nil),       // 19: openapi.core_api.CallInterfaceReq
	(*SignUpResp)(nil),             // 20: openapi.core_api.SignUpResp
	(*Response)(nil),               // 21: openapi.core_api.Response
	(*GetKeysResp)(nil),            // 22: openapi.core_api.GetKeysResp
	(*GetBaseInterfacesResp)(nil),  // 23: openapi.core_api.GetBaseInterfacesResp
	(*GetFullInterfacesResp)(nil),  // 24: openapi.core_api.GetFullInterfacesResp
	(*GetGradientResp)(nil),        // 25: openapi.core_api.GetGradientResp
	(*CallInterfaceResp)(nil),      // 26: openapi.core_api.CallInterfaceResp
}
var file_core_api_proto_depIdxs = []int32{
	0,  // 0: openapi.core_api.user.SignUp:input_type -> openapi.core_api.SignUpReq
	1,  // 1: openapi.core_api.key.GenerateKey:input_type -> openapi.core_api.GenerateKeyReq
	2,  // 2: openapi.core_api.key.GetKeys:input_type -> openapi.core_api.GetKeysReq
	3,  // 3: openapi.core_api.key.UpdateKey:input_type -> openapi.core_api.UpdateKeyReq
	4,  // 4: openapi.core_api.key.RefreshKey:input_type -> openapi.core_api.RefreshKeyReq
	5,  // 5: openapi.core_api.key.UpdateHosts:input_type -> openapi.core_api.UpdateHostReq
	6,  // 6: openapi.core_api.key.DeleteKey:input_type -> openapi.core_api.DeleteKeyReq
	7,  // 7: openapi.core_api.charge.CreateBaseInterface:input_type -> openapi.core_api.CreateBaseInterfaceReq
	8,  // 8: openapi.core_api.charge.UpdateBaseInterface:input_type -> openapi.core_api.UpdateBaseInterfaceReq
	9,  // 9: openapi.core_api.charge.DeleteBaseInterface:input_type -> openapi.core_api.DeleteBaseInterfaceReq
	10, // 10: openapi.core_api.charge.GetBaseInterfaces:input_type -> openapi.core_api.GetBaseInterfacesReq
	11, // 11: openapi.core_api.charge.CreateFullInterface:input_type -> openapi.core_api.CreateFullInterfaceReq
	12, // 12: openapi.core_api.charge.UpdateFullInterface:input_type -> openapi.core_api.UpdateFullInterfaceReq
	13, // 13: openapi.core_api.charge.UpdateMargin:input_type -> openapi.core_api.UpdateMarginReq
	14, // 14: openapi.core_api.charge.DeleteFullInterface:input_type -> openapi.core_api.DeleteFullInterfaceReq
	15, // 15: openapi.core_api.charge.GetFullInterfaces:input_type -> openapi.core_api.GetFullInterfacesReq
	16, // 16: openapi.core_api.charge.CreateGradient:input_type -> openapi.core_api.CreateGradientReq
	17, // 17: openapi.core_api.charge.UpdateGradient:input_type -> openapi.core_api.UpdateGradientReq
	18, // 18: openapi.core_api.charge.GetGradient:input_type -> openapi.core_api.GetGradientReq
	19, // 19: openapi.core_api.call.CallInterface:input_type -> openapi.core_api.CallInterfaceReq
	20, // 20: openapi.core_api.user.SignUp:output_type -> openapi.core_api.SignUpResp
	21, // 21: openapi.core_api.key.GenerateKey:output_type -> openapi.core_api.Response
	22, // 22: openapi.core_api.key.GetKeys:output_type -> openapi.core_api.GetKeysResp
	21, // 23: openapi.core_api.key.UpdateKey:output_type -> openapi.core_api.Response
	21, // 24: openapi.core_api.key.RefreshKey:output_type -> openapi.core_api.Response
	21, // 25: openapi.core_api.key.UpdateHosts:output_type -> openapi.core_api.Response
	21, // 26: openapi.core_api.key.DeleteKey:output_type -> openapi.core_api.Response
	21, // 27: openapi.core_api.charge.CreateBaseInterface:output_type -> openapi.core_api.Response
	21, // 28: openapi.core_api.charge.UpdateBaseInterface:output_type -> openapi.core_api.Response
	21, // 29: openapi.core_api.charge.DeleteBaseInterface:output_type -> openapi.core_api.Response
	23, // 30: openapi.core_api.charge.GetBaseInterfaces:output_type -> openapi.core_api.GetBaseInterfacesResp
	21, // 31: openapi.core_api.charge.CreateFullInterface:output_type -> openapi.core_api.Response
	21, // 32: openapi.core_api.charge.UpdateFullInterface:output_type -> openapi.core_api.Response
	21, // 33: openapi.core_api.charge.UpdateMargin:output_type -> openapi.core_api.Response
	21, // 34: openapi.core_api.charge.DeleteFullInterface:output_type -> openapi.core_api.Response
	24, // 35: openapi.core_api.charge.GetFullInterfaces:output_type -> openapi.core_api.GetFullInterfacesResp
	21, // 36: openapi.core_api.charge.CreateGradient:output_type -> openapi.core_api.Response
	21, // 37: openapi.core_api.charge.UpdateGradient:output_type -> openapi.core_api.Response
	25, // 38: openapi.core_api.charge.GetGradient:output_type -> openapi.core_api.GetGradientResp
	26, // 39: openapi.core_api.call.CallInterface:output_type -> openapi.core_api.CallInterfaceResp
	20, // [20:40] is the sub-list for method output_type
	0,  // [0:20] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_core_api_proto_init() }
func file_core_api_proto_init() {
	if File_core_api_proto != nil {
		return
	}
	file_openapi_core_api_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_core_api_proto_goTypes,
		DependencyIndexes: file_core_api_proto_depIdxs,
	}.Build()
	File_core_api_proto = out.File
	file_core_api_proto_rawDesc = nil
	file_core_api_proto_goTypes = nil
	file_core_api_proto_depIdxs = nil
}
