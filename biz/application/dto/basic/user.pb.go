// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.27.3
// source: basic/user.proto

package basic

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

type UserMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string          `protobuf:"bytes,1,opt,name=userId,proto3" form:"userId" json:"userId" query:"userId"`
	AppId           APP             `protobuf:"varint,2,opt,name=appId,proto3,enum=basic.APP" form:"appId" json:"appId" query:"appId"`
	DeviceId        string          `protobuf:"bytes,3,opt,name=deviceId,proto3" form:"deviceId" json:"deviceId" query:"deviceId"`
	SessionUserId   string          `protobuf:"bytes,4,opt,name=sessionUserId,proto3" form:"sessionUserId" json:"sessionUserId" query:"sessionUserId"`
	SessionAppId    APP             `protobuf:"varint,5,opt,name=sessionAppId,proto3,enum=basic.APP" form:"sessionAppId" json:"sessionAppId" query:"sessionAppId"`
	SessionDeviceId string          `protobuf:"bytes,6,opt,name=sessionDeviceId,proto3" form:"sessionDeviceId" json:"sessionDeviceId" query:"sessionDeviceId"`
	IsLogin         bool            `protobuf:"varint,7,opt,name=isLogin,proto3" form:"isLogin" json:"isLogin" query:"isLogin"`
	WechatUserMeta  *WechatUserMeta `protobuf:"bytes,8,opt,name=wechatUserMeta,proto3,oneof" form:"wechatUserMeta" json:"wechatUserMeta" query:"wechatUserMeta"`
}

func (x *UserMeta) Reset() {
	*x = UserMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMeta) ProtoMessage() {}

func (x *UserMeta) ProtoReflect() protoreflect.Message {
	mi := &file_basic_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMeta.ProtoReflect.Descriptor instead.
func (*UserMeta) Descriptor() ([]byte, []int) {
	return file_basic_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserMeta) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserMeta) GetAppId() APP {
	if x != nil {
		return x.AppId
	}
	return APP_Unknown
}

func (x *UserMeta) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *UserMeta) GetSessionUserId() string {
	if x != nil {
		return x.SessionUserId
	}
	return ""
}

func (x *UserMeta) GetSessionAppId() APP {
	if x != nil {
		return x.SessionAppId
	}
	return APP_Unknown
}

func (x *UserMeta) GetSessionDeviceId() string {
	if x != nil {
		return x.SessionDeviceId
	}
	return ""
}

func (x *UserMeta) GetIsLogin() bool {
	if x != nil {
		return x.IsLogin
	}
	return false
}

func (x *UserMeta) GetWechatUserMeta() *WechatUserMeta {
	if x != nil {
		return x.WechatUserMeta
	}
	return nil
}

// 客户端透传
type Extra struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIP       string `protobuf:"bytes,1,opt,name=clientIP,proto3" form:"clientIP" json:"clientIP" query:"clientIP"`
	Language       string `protobuf:"bytes,2,opt,name=language,proto3" form:"language" json:"language" query:"language"`
	Resolution     string `protobuf:"bytes,3,opt,name=resolution,proto3" form:"resolution" json:"resolution" query:"resolution"`                 // 像素比
	DevicePlatform string `protobuf:"bytes,4,opt,name=devicePlatform,proto3" form:"devicePlatform" json:"devicePlatform" query:"devicePlatform"` // 设备平台
	DeviceBrand    string `protobuf:"bytes,5,opt,name=deviceBrand,proto3" form:"deviceBrand" json:"deviceBrand" query:"deviceBrand"`             // 设备品牌
	DeviceId       string `protobuf:"bytes,6,opt,name=deviceId,proto3" form:"deviceId" json:"deviceId" query:"deviceId"`
	OperateSystem  string `protobuf:"bytes,7,opt,name=operateSystem,proto3" form:"operateSystem" json:"operateSystem" query:"operateSystem"`
}

func (x *Extra) Reset() {
	*x = Extra{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extra) ProtoMessage() {}

func (x *Extra) ProtoReflect() protoreflect.Message {
	mi := &file_basic_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extra.ProtoReflect.Descriptor instead.
func (*Extra) Descriptor() ([]byte, []int) {
	return file_basic_user_proto_rawDescGZIP(), []int{1}
}

func (x *Extra) GetClientIP() string {
	if x != nil {
		return x.ClientIP
	}
	return ""
}

func (x *Extra) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Extra) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *Extra) GetDevicePlatform() string {
	if x != nil {
		return x.DevicePlatform
	}
	return ""
}

func (x *Extra) GetDeviceBrand() string {
	if x != nil {
		return x.DeviceBrand
	}
	return ""
}

func (x *Extra) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Extra) GetOperateSystem() string {
	if x != nil {
		return x.OperateSystem
	}
	return ""
}

type WechatUserMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      string `protobuf:"bytes,1,opt,name=appId,proto3" form:"appId" json:"appId" query:"appId"`
	OpenId     string `protobuf:"bytes,2,opt,name=openId,proto3" form:"openId" json:"openId" query:"openId"`
	PlatformId string `protobuf:"bytes,3,opt,name=platformId,proto3" form:"platformId" json:"platformId" query:"platformId"`
	UnionId    string `protobuf:"bytes,4,opt,name=unionId,proto3" form:"unionId" json:"unionId" query:"unionId"`
}

func (x *WechatUserMeta) Reset() {
	*x = WechatUserMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WechatUserMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WechatUserMeta) ProtoMessage() {}

func (x *WechatUserMeta) ProtoReflect() protoreflect.Message {
	mi := &file_basic_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WechatUserMeta.ProtoReflect.Descriptor instead.
func (*WechatUserMeta) Descriptor() ([]byte, []int) {
	return file_basic_user_proto_rawDescGZIP(), []int{2}
}

func (x *WechatUserMeta) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *WechatUserMeta) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *WechatUserMeta) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

func (x *WechatUserMeta) GetUnionId() string {
	if x != nil {
		return x.UnionId
	}
	return ""
}

var File_basic_user_proto protoreflect.FileDescriptor

var file_basic_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x1a, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x50, 0x50, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x70,
	0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x41, 0x50, 0x50, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x42, 0x0a, 0x0e, 0x77, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x22, 0xeb,
	0x01, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x50, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x78, 0x0a, 0x0e,
	0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x6b, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68,
	0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68,
	0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basic_user_proto_rawDescOnce sync.Once
	file_basic_user_proto_rawDescData = file_basic_user_proto_rawDesc
)

func file_basic_user_proto_rawDescGZIP() []byte {
	file_basic_user_proto_rawDescOnce.Do(func() {
		file_basic_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_user_proto_rawDescData)
	})
	return file_basic_user_proto_rawDescData
}

var file_basic_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_basic_user_proto_goTypes = []interface{}{
	(*UserMeta)(nil),       // 0: basic.UserMeta
	(*Extra)(nil),          // 1: basic.Extra
	(*WechatUserMeta)(nil), // 2: basic.WechatUserMeta
	(APP)(0),               // 3: basic.APP
}
var file_basic_user_proto_depIdxs = []int32{
	3, // 0: basic.UserMeta.appId:type_name -> basic.APP
	3, // 1: basic.UserMeta.sessionAppId:type_name -> basic.APP
	2, // 2: basic.UserMeta.wechatUserMeta:type_name -> basic.WechatUserMeta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_basic_user_proto_init() }
func file_basic_user_proto_init() {
	if File_basic_user_proto != nil {
		return
	}
	file_basic_app_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_basic_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMeta); i {
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
		file_basic_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extra); i {
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
		file_basic_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WechatUserMeta); i {
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
	file_basic_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basic_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basic_user_proto_goTypes,
		DependencyIndexes: file_basic_user_proto_depIdxs,
		MessageInfos:      file_basic_user_proto_msgTypes,
	}.Build()
	File_basic_user_proto = out.File
	file_basic_user_proto_rawDesc = nil
	file_basic_user_proto_goTypes = nil
	file_basic_user_proto_depIdxs = nil
}
