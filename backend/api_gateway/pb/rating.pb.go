// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: proto/rating.proto

package pb

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

type HostRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HostId string `protobuf:"bytes,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rating int32  `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *HostRating) Reset() {
	*x = HostRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostRating) ProtoMessage() {}

func (x *HostRating) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostRating.ProtoReflect.Descriptor instead.
func (*HostRating) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{0}
}

func (x *HostRating) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HostRating) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *HostRating) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *HostRating) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type HostRatingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ratings []*HostRating `protobuf:"bytes,1,rep,name=ratings,proto3" json:"ratings,omitempty"`
}

func (x *HostRatingList) Reset() {
	*x = HostRatingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostRatingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostRatingList) ProtoMessage() {}

func (x *HostRatingList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostRatingList.ProtoReflect.Descriptor instead.
func (*HostRatingList) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{1}
}

func (x *HostRatingList) GetRatings() []*HostRating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type AccommodationRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccommodationId string `protobuf:"bytes,2,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	UserId          string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rating          int32  `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *AccommodationRating) Reset() {
	*x = AccommodationRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationRating) ProtoMessage() {}

func (x *AccommodationRating) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationRating.ProtoReflect.Descriptor instead.
func (*AccommodationRating) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{2}
}

func (x *AccommodationRating) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccommodationRating) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *AccommodationRating) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AccommodationRating) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type AccommodationRatingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ratings []*AccommodationRating `protobuf:"bytes,1,rep,name=ratings,proto3" json:"ratings,omitempty"`
}

func (x *AccommodationRatingList) Reset() {
	*x = AccommodationRatingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationRatingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationRatingList) ProtoMessage() {}

func (x *AccommodationRatingList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationRatingList.ProtoReflect.Descriptor instead.
func (*AccommodationRatingList) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{3}
}

func (x *AccommodationRatingList) GetRatings() []*AccommodationRating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // host_id or accommodation_id or user_id
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{4}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating float64 `protobuf:"fixed64,1,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RatingResponse) Reset() {
	*x = RatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingResponse) ProtoMessage() {}

func (x *RatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingResponse.ProtoReflect.Descriptor instead.
func (*RatingResponse) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{5}
}

func (x *RatingResponse) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type RateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // host_id or accommodation_id
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rating int32  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{6}
}

func (x *RateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RateRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated bool `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{7}
}

func (x *RateResponse) GetUpdated() bool {
	if x != nil {
		return x.Updated
	}
	return false
}

type RemoveRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // host_id or accommodation_id
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemoveRatingRequest) Reset() {
	*x = RemoveRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRatingRequest) ProtoMessage() {}

func (x *RemoveRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRatingRequest.ProtoReflect.Descriptor instead.
func (*RemoveRatingRequest) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveRatingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoveRatingRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RemoveRatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Removed bool `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
}

func (x *RemoveRatingResponse) Reset() {
	*x = RemoveRatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRatingResponse) ProtoMessage() {}

func (x *RemoveRatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRatingResponse.ProtoReflect.Descriptor instead.
func (*RemoveRatingResponse) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveRatingResponse) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

var File_proto_rating_proto protoreflect.FileDescriptor

var file_proto_rating_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x37, 0x0a, 0x0e,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a,
	0x10, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x49, 0x0a, 0x17, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4e, 0x0a, 0x0b, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x28, 0x0a, 0x0c, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x32, 0xd6, 0x03, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x48, 0x6f, 0x73,
	0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0a, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x08, 0x52, 0x61, 0x74, 0x65, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x0c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x48, 0x6f, 0x73,
	0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0a, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0a, 0x2e,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x11,
	0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0a, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rating_proto_rawDescOnce sync.Once
	file_proto_rating_proto_rawDescData = file_proto_rating_proto_rawDesc
)

func file_proto_rating_proto_rawDescGZIP() []byte {
	file_proto_rating_proto_rawDescOnce.Do(func() {
		file_proto_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rating_proto_rawDescData)
	})
	return file_proto_rating_proto_rawDescData
}

var file_proto_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_rating_proto_goTypes = []interface{}{
	(*HostRating)(nil),              // 0: HostRating
	(*HostRatingList)(nil),          // 1: HostRatingList
	(*AccommodationRating)(nil),     // 2: AccommodationRating
	(*AccommodationRatingList)(nil), // 3: AccommodationRatingList
	(*IdRequest)(nil),               // 4: IdRequest
	(*RatingResponse)(nil),          // 5: RatingResponse
	(*RateRequest)(nil),             // 6: RateRequest
	(*RateResponse)(nil),            // 7: RateResponse
	(*RemoveRatingRequest)(nil),     // 8: RemoveRatingRequest
	(*RemoveRatingResponse)(nil),    // 9: RemoveRatingResponse
}
var file_proto_rating_proto_depIdxs = []int32{
	0,  // 0: HostRatingList.ratings:type_name -> HostRating
	2,  // 1: AccommodationRatingList.ratings:type_name -> AccommodationRating
	4,  // 2: RatingService.HostRating:input_type -> IdRequest
	6,  // 3: RatingService.RateHost:input_type -> RateRequest
	8,  // 4: RatingService.RemoveHostRating:input_type -> RemoveRatingRequest
	4,  // 5: RatingService.GetMyHostRating:input_type -> IdRequest
	4,  // 6: RatingService.AccommodationRating:input_type -> IdRequest
	6,  // 7: RatingService.RateAccommodation:input_type -> RateRequest
	8,  // 8: RatingService.RemoveAccommodationRating:input_type -> RemoveRatingRequest
	4,  // 9: RatingService.GetMyAccommodationRating:input_type -> IdRequest
	5,  // 10: RatingService.HostRating:output_type -> RatingResponse
	7,  // 11: RatingService.RateHost:output_type -> RateResponse
	9,  // 12: RatingService.RemoveHostRating:output_type -> RemoveRatingResponse
	1,  // 13: RatingService.GetMyHostRating:output_type -> HostRatingList
	5,  // 14: RatingService.AccommodationRating:output_type -> RatingResponse
	7,  // 15: RatingService.RateAccommodation:output_type -> RateResponse
	9,  // 16: RatingService.RemoveAccommodationRating:output_type -> RemoveRatingResponse
	3,  // 17: RatingService.GetMyAccommodationRating:output_type -> AccommodationRatingList
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_rating_proto_init() }
func file_proto_rating_proto_init() {
	if File_proto_rating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostRating); i {
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
		file_proto_rating_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostRatingList); i {
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
		file_proto_rating_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationRating); i {
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
		file_proto_rating_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationRatingList); i {
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
		file_proto_rating_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_proto_rating_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingResponse); i {
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
		file_proto_rating_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRequest); i {
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
		file_proto_rating_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
		file_proto_rating_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRatingRequest); i {
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
		file_proto_rating_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRatingResponse); i {
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
			RawDescriptor: file_proto_rating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rating_proto_goTypes,
		DependencyIndexes: file_proto_rating_proto_depIdxs,
		MessageInfos:      file_proto_rating_proto_msgTypes,
	}.Build()
	File_proto_rating_proto = out.File
	file_proto_rating_proto_rawDesc = nil
	file_proto_rating_proto_goTypes = nil
	file_proto_rating_proto_depIdxs = nil
}
