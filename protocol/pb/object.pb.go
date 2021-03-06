// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: object.proto

package pb

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

type PutObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName  string            `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	ObjectName  string            `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	UserDefined map[string]string `protobuf:"bytes,4,rep,name=user_defined,json=userDefined,proto3" json:"user_defined,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserTags    string            `protobuf:"bytes,5,opt,name=user_tags,json=userTags,proto3" json:"user_tags,omitempty"`
	ObjectsSize int64             `protobuf:"varint,6,opt,name=objects_size,json=objectsSize,proto3" json:"objects_size,omitempty"`
	ContentType string            `protobuf:"bytes,7,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Message     string            `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PutObjectRequest) Reset() {
	*x = PutObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutObjectRequest) ProtoMessage() {}

func (x *PutObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutObjectRequest.ProtoReflect.Descriptor instead.
func (*PutObjectRequest) Descriptor() ([]byte, []int) {
	return file_object_proto_rawDescGZIP(), []int{0}
}

func (x *PutObjectRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *PutObjectRequest) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *PutObjectRequest) GetUserDefined() map[string]string {
	if x != nil {
		return x.UserDefined
	}
	return nil
}

func (x *PutObjectRequest) GetUserTags() string {
	if x != nil {
		return x.UserTags
	}
	return ""
}

func (x *PutObjectRequest) GetObjectsSize() int64 {
	if x != nil {
		return x.ObjectsSize
	}
	return 0
}

func (x *PutObjectRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *PutObjectRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PutObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName  string            `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	ObjectName  string            `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	ObjectId    int64             `protobuf:"varint,3,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	UserDefined map[string]string `protobuf:"bytes,4,rep,name=user_defined,json=userDefined,proto3" json:"user_defined,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserTags    string            `protobuf:"bytes,5,opt,name=user_tags,json=userTags,proto3" json:"user_tags,omitempty"`
	BlockId     string            `protobuf:"bytes,6,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	StartPos    int64             `protobuf:"varint,7,opt,name=start_pos,json=startPos,proto3" json:"start_pos,omitempty"`
	ObjectsSize int64             `protobuf:"varint,8,opt,name=objects_size,json=objectsSize,proto3" json:"objects_size,omitempty"`
	ContentType string            `protobuf:"bytes,9,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Message     string            `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PutObjectResponse) Reset() {
	*x = PutObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutObjectResponse) ProtoMessage() {}

func (x *PutObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutObjectResponse.ProtoReflect.Descriptor instead.
func (*PutObjectResponse) Descriptor() ([]byte, []int) {
	return file_object_proto_rawDescGZIP(), []int{1}
}

func (x *PutObjectResponse) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *PutObjectResponse) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *PutObjectResponse) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *PutObjectResponse) GetUserDefined() map[string]string {
	if x != nil {
		return x.UserDefined
	}
	return nil
}

func (x *PutObjectResponse) GetUserTags() string {
	if x != nil {
		return x.UserTags
	}
	return ""
}

func (x *PutObjectResponse) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *PutObjectResponse) GetStartPos() int64 {
	if x != nil {
		return x.StartPos
	}
	return 0
}

func (x *PutObjectResponse) GetObjectsSize() int64 {
	if x != nil {
		return x.ObjectsSize
	}
	return 0
}

func (x *PutObjectResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *PutObjectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//bucket name
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	//format: {bucket-name}.id
	ObjectId int64  `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetObjectRequest) Reset() {
	*x = GetObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectRequest) ProtoMessage() {}

func (x *GetObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_object_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectRequest.ProtoReflect.Descriptor instead.
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return file_object_proto_rawDescGZIP(), []int{2}
}

func (x *GetObjectRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *GetObjectRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *GetObjectRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetObjectReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//bucket name
	BucketName  string            `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	ObjectId    int64             `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	UserDefined map[string]string `protobuf:"bytes,3,rep,name=user_defined,json=userDefined,proto3" json:"user_defined,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserTags    string            `protobuf:"bytes,4,opt,name=user_tags,json=userTags,proto3" json:"user_tags,omitempty"`
	BlockId     string            `protobuf:"bytes,5,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	StartPos    int64             `protobuf:"varint,6,opt,name=start_pos,json=startPos,proto3" json:"start_pos,omitempty"`
	ObjectsSize int64             `protobuf:"varint,7,opt,name=objects_size,json=objectsSize,proto3" json:"objects_size,omitempty"`
	ContentType string            `protobuf:"bytes,8,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Message     string            `protobuf:"bytes,9,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetObjectReponse) Reset() {
	*x = GetObjectReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectReponse) ProtoMessage() {}

func (x *GetObjectReponse) ProtoReflect() protoreflect.Message {
	mi := &file_object_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectReponse.ProtoReflect.Descriptor instead.
func (*GetObjectReponse) Descriptor() ([]byte, []int) {
	return file_object_proto_rawDescGZIP(), []int{3}
}

func (x *GetObjectReponse) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *GetObjectReponse) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *GetObjectReponse) GetUserDefined() map[string]string {
	if x != nil {
		return x.UserDefined
	}
	return nil
}

func (x *GetObjectReponse) GetUserTags() string {
	if x != nil {
		return x.UserTags
	}
	return ""
}

func (x *GetObjectReponse) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *GetObjectReponse) GetStartPos() int64 {
	if x != nil {
		return x.StartPos
	}
	return 0
}

func (x *GetObjectReponse) GetObjectsSize() int64 {
	if x != nil {
		return x.ObjectsSize
	}
	return 0
}

func (x *GetObjectReponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *GetObjectReponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	ObjectId   int64  `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteObjectRequest) Reset() {
	*x = DeleteObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectRequest) ProtoMessage() {}

func (x *DeleteObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_object_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteObjectRequest) Descriptor() ([]byte, []int) {
	return file_object_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteObjectRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *DeleteObjectRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *DeleteObjectRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	ObjectId   int64  `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteObjectResponse) Reset() {
	*x = DeleteObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectResponse) ProtoMessage() {}

func (x *DeleteObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_object_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectResponse.ProtoReflect.Descriptor instead.
func (*DeleteObjectResponse) Descriptor() ([]byte, []int) {
	return file_object_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteObjectResponse) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *DeleteObjectResponse) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *DeleteObjectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_object_proto protoreflect.FileDescriptor

var file_object_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xdb, 0x02, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x3e, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xb2, 0x03, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x50, 0x6f, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x8f, 0x03, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x70, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x6f, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x6d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x6e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_object_proto_rawDescOnce sync.Once
	file_object_proto_rawDescData = file_object_proto_rawDesc
)

func file_object_proto_rawDescGZIP() []byte {
	file_object_proto_rawDescOnce.Do(func() {
		file_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_object_proto_rawDescData)
	})
	return file_object_proto_rawDescData
}

var file_object_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_object_proto_goTypes = []interface{}{
	(*PutObjectRequest)(nil),     // 0: pb.PutObjectRequest
	(*PutObjectResponse)(nil),    // 1: pb.PutObjectResponse
	(*GetObjectRequest)(nil),     // 2: pb.GetObjectRequest
	(*GetObjectReponse)(nil),     // 3: pb.GetObjectReponse
	(*DeleteObjectRequest)(nil),  // 4: pb.DeleteObjectRequest
	(*DeleteObjectResponse)(nil), // 5: pb.DeleteObjectResponse
	nil,                          // 6: pb.PutObjectRequest.UserDefinedEntry
	nil,                          // 7: pb.PutObjectResponse.UserDefinedEntry
	nil,                          // 8: pb.GetObjectReponse.UserDefinedEntry
}
var file_object_proto_depIdxs = []int32{
	6, // 0: pb.PutObjectRequest.user_defined:type_name -> pb.PutObjectRequest.UserDefinedEntry
	7, // 1: pb.PutObjectResponse.user_defined:type_name -> pb.PutObjectResponse.UserDefinedEntry
	8, // 2: pb.GetObjectReponse.user_defined:type_name -> pb.GetObjectReponse.UserDefinedEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_object_proto_init() }
func file_object_proto_init() {
	if File_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutObjectRequest); i {
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
		file_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutObjectResponse); i {
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
		file_object_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectRequest); i {
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
		file_object_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectReponse); i {
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
		file_object_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectRequest); i {
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
		file_object_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectResponse); i {
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
			RawDescriptor: file_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_object_proto_goTypes,
		DependencyIndexes: file_object_proto_depIdxs,
		MessageInfos:      file_object_proto_msgTypes,
	}.Build()
	File_object_proto = out.File
	file_object_proto_rawDesc = nil
	file_object_proto_goTypes = nil
	file_object_proto_depIdxs = nil
}
