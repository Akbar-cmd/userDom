// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: note.proto

package gRPC

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

type BikeInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BikeId        uint64                 `protobuf:"varint,1,opt,name=bike_id,json=bikeId,proto3" json:"bike_id,omitempty"`
	Model         string                 `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Color         string                 `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	IsWork        bool                   `protobuf:"varint,4,opt,name=is_work,json=isWork,proto3" json:"is_work,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BikeInfo) Reset() {
	*x = BikeInfo{}
	mi := &file_note_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BikeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BikeInfo) ProtoMessage() {}

func (x *BikeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BikeInfo.ProtoReflect.Descriptor instead.
func (*BikeInfo) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{0}
}

func (x *BikeInfo) GetBikeId() uint64 {
	if x != nil {
		return x.BikeId
	}
	return 0
}

func (x *BikeInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *BikeInfo) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *BikeInfo) GetIsWork() bool {
	if x != nil {
		return x.IsWork
	}
	return false
}

type GetBikeByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBikeByUserRequest) Reset() {
	*x = GetBikeByUserRequest{}
	mi := &file_note_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBikeByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBikeByUserRequest) ProtoMessage() {}

func (x *GetBikeByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBikeByUserRequest.ProtoReflect.Descriptor instead.
func (*GetBikeByUserRequest) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{1}
}

func (x *GetBikeByUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetBikeByUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bikes         []*BikeInfo            `protobuf:"bytes,1,rep,name=bikes,proto3" json:"bikes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBikeByUserResponse) Reset() {
	*x = GetBikeByUserResponse{}
	mi := &file_note_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBikeByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBikeByUserResponse) ProtoMessage() {}

func (x *GetBikeByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBikeByUserResponse.ProtoReflect.Descriptor instead.
func (*GetBikeByUserResponse) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{2}
}

func (x *GetBikeByUserResponse) GetBikes() []*BikeInfo {
	if x != nil {
		return x.Bikes
	}
	return nil
}

type CreateBikeByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Model         string                 `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Color         string                 `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	IsWork        bool                   `protobuf:"varint,4,opt,name=is_work,json=isWork,proto3" json:"is_work,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBikeByUserRequest) Reset() {
	*x = CreateBikeByUserRequest{}
	mi := &file_note_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBikeByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBikeByUserRequest) ProtoMessage() {}

func (x *CreateBikeByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBikeByUserRequest.ProtoReflect.Descriptor instead.
func (*CreateBikeByUserRequest) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBikeByUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateBikeByUserRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CreateBikeByUserRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *CreateBikeByUserRequest) GetIsWork() bool {
	if x != nil {
		return x.IsWork
	}
	return false
}

type CreateBikeByUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BikeId        uint64                 `protobuf:"varint,1,opt,name=bike_id,json=bikeId,proto3" json:"bike_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBikeByUserResponse) Reset() {
	*x = CreateBikeByUserResponse{}
	mi := &file_note_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBikeByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBikeByUserResponse) ProtoMessage() {}

func (x *CreateBikeByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBikeByUserResponse.ProtoReflect.Descriptor instead.
func (*CreateBikeByUserResponse) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBikeByUserResponse) GetBikeId() uint64 {
	if x != nil {
		return x.BikeId
	}
	return 0
}

type UpdateBikeByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BikeId        uint64                 `protobuf:"varint,2,opt,name=bike_id,json=bikeId,proto3" json:"bike_id,omitempty"`
	Model         string                 `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Color         string                 `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	IsWork        bool                   `protobuf:"varint,5,opt,name=is_work,json=isWork,proto3" json:"is_work,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBikeByUserRequest) Reset() {
	*x = UpdateBikeByUserRequest{}
	mi := &file_note_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBikeByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBikeByUserRequest) ProtoMessage() {}

func (x *UpdateBikeByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBikeByUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateBikeByUserRequest) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBikeByUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateBikeByUserRequest) GetBikeId() uint64 {
	if x != nil {
		return x.BikeId
	}
	return 0
}

func (x *UpdateBikeByUserRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *UpdateBikeByUserRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *UpdateBikeByUserRequest) GetIsWork() bool {
	if x != nil {
		return x.IsWork
	}
	return false
}

type UpdateBikeByUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBikeByUserResponse) Reset() {
	*x = UpdateBikeByUserResponse{}
	mi := &file_note_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBikeByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBikeByUserResponse) ProtoMessage() {}

func (x *UpdateBikeByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBikeByUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateBikeByUserResponse) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBikeByUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteBikeByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BikeId        uint64                 `protobuf:"varint,2,opt,name=bike_id,json=bikeId,proto3" json:"bike_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBikeByUserRequest) Reset() {
	*x = DeleteBikeByUserRequest{}
	mi := &file_note_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBikeByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBikeByUserRequest) ProtoMessage() {}

func (x *DeleteBikeByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBikeByUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteBikeByUserRequest) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBikeByUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteBikeByUserRequest) GetBikeId() uint64 {
	if x != nil {
		return x.BikeId
	}
	return 0
}

type DeleteBikeByUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBikeByUserResponse) Reset() {
	*x = DeleteBikeByUserResponse{}
	mi := &file_note_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBikeByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBikeByUserResponse) ProtoMessage() {}

func (x *DeleteBikeByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBikeByUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteBikeByUserResponse) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBikeByUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_note_proto protoreflect.FileDescriptor

const file_note_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"note.proto\x12\x04gRPC\"h\n" +
	"\bBikeInfo\x12\x17\n" +
	"\abike_id\x18\x01 \x01(\x04R\x06bikeId\x12\x14\n" +
	"\x05model\x18\x02 \x01(\tR\x05model\x12\x14\n" +
	"\x05color\x18\x03 \x01(\tR\x05color\x12\x17\n" +
	"\ais_work\x18\x04 \x01(\bR\x06isWork\"/\n" +
	"\x14GetBikeByUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\"=\n" +
	"\x15GetBikeByUserResponse\x12$\n" +
	"\x05bikes\x18\x01 \x03(\v2\x0e.gRPC.BikeInfoR\x05bikes\"w\n" +
	"\x17CreateBikeByUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x14\n" +
	"\x05model\x18\x02 \x01(\tR\x05model\x12\x14\n" +
	"\x05color\x18\x03 \x01(\tR\x05color\x12\x17\n" +
	"\ais_work\x18\x04 \x01(\bR\x06isWork\"3\n" +
	"\x18CreateBikeByUserResponse\x12\x17\n" +
	"\abike_id\x18\x01 \x01(\x04R\x06bikeId\"\x90\x01\n" +
	"\x17UpdateBikeByUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x17\n" +
	"\abike_id\x18\x02 \x01(\x04R\x06bikeId\x12\x14\n" +
	"\x05model\x18\x03 \x01(\tR\x05model\x12\x14\n" +
	"\x05color\x18\x04 \x01(\tR\x05color\x12\x17\n" +
	"\ais_work\x18\x05 \x01(\bR\x06isWork\"4\n" +
	"\x18UpdateBikeByUserResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"K\n" +
	"\x17DeleteBikeByUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x17\n" +
	"\abike_id\x18\x02 \x01(\x04R\x06bikeId\"4\n" +
	"\x18DeleteBikeByUserResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xc9\x02\n" +
	"\x04Bike\x12H\n" +
	"\rGetBikeByUser\x12\x1a.gRPC.GetBikeByUserRequest\x1a\x1b.gRPC.GetBikeByUserResponse\x12Q\n" +
	"\x10CreateBikeByUser\x12\x1d.gRPC.CreateBikeByUserRequest\x1a\x1e.gRPC.CreateBikeByUserResponse\x12Q\n" +
	"\x10UpdateBikeByUser\x12\x1d.gRPC.UpdateBikeByUserRequest\x1a\x1e.gRPC.UpdateBikeByUserResponse\x12Q\n" +
	"\x10DeleteBikeByUser\x12\x1d.gRPC.DeleteBikeByUserRequest\x1a\x1e.gRPC.DeleteBikeByUserResponseB/Z-weBike/webike_User-domain_akbar/pkg/gRPC;gRPCb\x06proto3"

var (
	file_note_proto_rawDescOnce sync.Once
	file_note_proto_rawDescData []byte
)

func file_note_proto_rawDescGZIP() []byte {
	file_note_proto_rawDescOnce.Do(func() {
		file_note_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_note_proto_rawDesc), len(file_note_proto_rawDesc)))
	})
	return file_note_proto_rawDescData
}

var file_note_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_note_proto_goTypes = []any{
	(*BikeInfo)(nil),                 // 0: gRPC.BikeInfo
	(*GetBikeByUserRequest)(nil),     // 1: gRPC.GetBikeByUserRequest
	(*GetBikeByUserResponse)(nil),    // 2: gRPC.GetBikeByUserResponse
	(*CreateBikeByUserRequest)(nil),  // 3: gRPC.CreateBikeByUserRequest
	(*CreateBikeByUserResponse)(nil), // 4: gRPC.CreateBikeByUserResponse
	(*UpdateBikeByUserRequest)(nil),  // 5: gRPC.UpdateBikeByUserRequest
	(*UpdateBikeByUserResponse)(nil), // 6: gRPC.UpdateBikeByUserResponse
	(*DeleteBikeByUserRequest)(nil),  // 7: gRPC.DeleteBikeByUserRequest
	(*DeleteBikeByUserResponse)(nil), // 8: gRPC.DeleteBikeByUserResponse
}
var file_note_proto_depIdxs = []int32{
	0, // 0: gRPC.GetBikeByUserResponse.bikes:type_name -> gRPC.BikeInfo
	1, // 1: gRPC.Bike.GetBikeByUser:input_type -> gRPC.GetBikeByUserRequest
	3, // 2: gRPC.Bike.CreateBikeByUser:input_type -> gRPC.CreateBikeByUserRequest
	5, // 3: gRPC.Bike.UpdateBikeByUser:input_type -> gRPC.UpdateBikeByUserRequest
	7, // 4: gRPC.Bike.DeleteBikeByUser:input_type -> gRPC.DeleteBikeByUserRequest
	2, // 5: gRPC.Bike.GetBikeByUser:output_type -> gRPC.GetBikeByUserResponse
	4, // 6: gRPC.Bike.CreateBikeByUser:output_type -> gRPC.CreateBikeByUserResponse
	6, // 7: gRPC.Bike.UpdateBikeByUser:output_type -> gRPC.UpdateBikeByUserResponse
	8, // 8: gRPC.Bike.DeleteBikeByUser:output_type -> gRPC.DeleteBikeByUserResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_note_proto_init() }
func file_note_proto_init() {
	if File_note_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_note_proto_rawDesc), len(file_note_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_note_proto_goTypes,
		DependencyIndexes: file_note_proto_depIdxs,
		MessageInfos:      file_note_proto_msgTypes,
	}.Build()
	File_note_proto = out.File
	file_note_proto_goTypes = nil
	file_note_proto_depIdxs = nil
}
