// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: proto/reservation.proto

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

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	HostId          string  `protobuf:"bytes,3,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	AccommodationId string  `protobuf:"bytes,4,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	StartDate       string  `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate         string  `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Price           float64 `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	NumberOfGuests  int32   `protobuf:"varint,8,opt,name=number_of_guests,json=numberOfGuests,proto3" json:"number_of_guests,omitempty"`
	Status          string  `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reservation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Reservation) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *Reservation) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *Reservation) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Reservation) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Reservation) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Reservation) GetNumberOfGuests() int32 {
	if x != nil {
		return x.NumberOfGuests
	}
	return 0
}

func (x *Reservation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationId string `protobuf:"bytes,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
}

func (x *GetReservationRequest) Reset() {
	*x = GetReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationRequest) ProtoMessage() {}

func (x *GetReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationRequest.ProtoReflect.Descriptor instead.
func (*GetReservationRequest) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *GetReservationRequest) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

type CreateReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	HostId          string  `protobuf:"bytes,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	AccommodationId string  `protobuf:"bytes,3,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	StartDate       string  `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate         string  `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Price           float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	NumberOfGuests  int32   `protobuf:"varint,7,opt,name=number_of_guests,json=numberOfGuests,proto3" json:"number_of_guests,omitempty"`
}

func (x *CreateReservationRequest) Reset() {
	*x = CreateReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationRequest) ProtoMessage() {}

func (x *CreateReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationRequest.ProtoReflect.Descriptor instead.
func (*CreateReservationRequest) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *CreateReservationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateReservationRequest) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *CreateReservationRequest) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *CreateReservationRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CreateReservationRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *CreateReservationRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateReservationRequest) GetNumberOfGuests() int32 {
	if x != nil {
		return x.NumberOfGuests
	}
	return 0
}

type ReservationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReservationStatus) Reset() {
	*x = ReservationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationStatus) ProtoMessage() {}

func (x *ReservationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationStatus.ProtoReflect.Descriptor instead.
func (*ReservationStatus) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{3}
}

func (x *ReservationStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ReservationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *ReservationList) Reset() {
	*x = ReservationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationList) ProtoMessage() {}

func (x *ReservationList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationList.ProtoReflect.Descriptor instead.
func (*ReservationList) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{4}
}

func (x *ReservationList) GetReservations() []*Reservation {
	if x != nil {
		return x.Reservations
	}
	return nil
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[5]
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
	return file_proto_reservation_proto_rawDescGZIP(), []int{5}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IdList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdList) Reset() {
	*x = IdList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdList) ProtoMessage() {}

func (x *IdList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdList.ProtoReflect.Descriptor instead.
func (*IdList) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{6}
}

func (x *IdList) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type FilterTakenAccommodationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationIds []string `protobuf:"bytes,1,rep,name=accommodation_ids,json=accommodationIds,proto3" json:"accommodation_ids,omitempty"`
	StartDate        string   `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate          string   `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *FilterTakenAccommodationsRequest) Reset() {
	*x = FilterTakenAccommodationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterTakenAccommodationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterTakenAccommodationsRequest) ProtoMessage() {}

func (x *FilterTakenAccommodationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterTakenAccommodationsRequest.ProtoReflect.Descriptor instead.
func (*FilterTakenAccommodationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{7}
}

func (x *FilterTakenAccommodationsRequest) GetAccommodationIds() []string {
	if x != nil {
		return x.AccommodationIds
	}
	return nil
}

func (x *FilterTakenAccommodationsRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *FilterTakenAccommodationsRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type IntervalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId string `protobuf:"bytes,2,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	StartDate       string `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate         string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *IntervalRequest) Reset() {
	*x = IntervalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntervalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntervalRequest) ProtoMessage() {}

func (x *IntervalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntervalRequest.ProtoReflect.Descriptor instead.
func (*IntervalRequest) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{8}
}

func (x *IntervalRequest) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *IntervalRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *IntervalRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type BoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BoolResponse) Reset() {
	*x = BoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reservation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResponse) ProtoMessage() {}

func (x *BoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResponse.ProtoReflect.Descriptor instead.
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{9}
}

func (x *BoolResponse) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

var File_proto_reservation_proto protoreflect.FileDescriptor

var file_proto_reservation_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02, 0x0a, 0x0b, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x6f, 0x66, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xf1, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x2b, 0x0a, 0x11,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0c,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1b,
	0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x06, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x20, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x76, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0xaf, 0x06, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40,
	0x0a, 0x12, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3f, 0x0a, 0x11, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3f, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x37, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0a, 0x2e,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a,
	0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x1c,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x07, 0x2e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x1e, 0x48, 0x61, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x10, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x21, 0x48,
	0x61, 0x73, 0x47, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x0a, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x20, 0x48,
	0x61, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x07, 0x2e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x46, 0x6f, 0x72, 0x53, 0x75, 0x70, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x75, 0x70, 0x65, 0x72, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x07, 0x2e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_reservation_proto_rawDescOnce sync.Once
	file_proto_reservation_proto_rawDescData = file_proto_reservation_proto_rawDesc
)

func file_proto_reservation_proto_rawDescGZIP() []byte {
	file_proto_reservation_proto_rawDescOnce.Do(func() {
		file_proto_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reservation_proto_rawDescData)
	})
	return file_proto_reservation_proto_rawDescData
}

var file_proto_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_reservation_proto_goTypes = []interface{}{
	(*Reservation)(nil),                      // 0: Reservation
	(*GetReservationRequest)(nil),            // 1: GetReservationRequest
	(*CreateReservationRequest)(nil),         // 2: CreateReservationRequest
	(*ReservationStatus)(nil),                // 3: ReservationStatus
	(*ReservationList)(nil),                  // 4: ReservationList
	(*IdRequest)(nil),                        // 5: IdRequest
	(*IdList)(nil),                           // 6: IdList
	(*FilterTakenAccommodationsRequest)(nil), // 7: FilterTakenAccommodationsRequest
	(*IntervalRequest)(nil),                  // 8: IntervalRequest
	(*BoolResponse)(nil),                     // 9: BoolResponse
}
var file_proto_reservation_proto_depIdxs = []int32{
	0,  // 0: ReservationList.reservations:type_name -> Reservation
	1,  // 1: ReservationService.GetReservation:input_type -> GetReservationRequest
	2,  // 2: ReservationService.CreateReservation:input_type -> CreateReservationRequest
	1,  // 3: ReservationService.ApproveReservation:input_type -> GetReservationRequest
	1,  // 4: ReservationService.RejectReservation:input_type -> GetReservationRequest
	1,  // 5: ReservationService.CancelReservation:input_type -> GetReservationRequest
	5,  // 6: ReservationService.GetReservationsForGuest:input_type -> IdRequest
	5,  // 7: ReservationService.GetReservationsForAccommodation:input_type -> IdRequest
	7,  // 8: ReservationService.FilterOutTakenAccommodations:input_type -> FilterTakenAccommodationsRequest
	8,  // 9: ReservationService.HasActiveReservationInInterval:input_type -> IntervalRequest
	5,  // 10: ReservationService.HasGuestActiveReservationInFuture:input_type -> IdRequest
	6,  // 11: ReservationService.HasHostActiveReservationInFuture:input_type -> IdList
	5,  // 12: ReservationService.CheckForSuperHost:input_type -> IdRequest
	6,  // 13: ReservationService.GetHostIdsForSuperHost:input_type -> IdList
	0,  // 14: ReservationService.GetReservation:output_type -> Reservation
	0,  // 15: ReservationService.CreateReservation:output_type -> Reservation
	3,  // 16: ReservationService.ApproveReservation:output_type -> ReservationStatus
	3,  // 17: ReservationService.RejectReservation:output_type -> ReservationStatus
	3,  // 18: ReservationService.CancelReservation:output_type -> ReservationStatus
	4,  // 19: ReservationService.GetReservationsForGuest:output_type -> ReservationList
	4,  // 20: ReservationService.GetReservationsForAccommodation:output_type -> ReservationList
	6,  // 21: ReservationService.FilterOutTakenAccommodations:output_type -> IdList
	9,  // 22: ReservationService.HasActiveReservationInInterval:output_type -> BoolResponse
	9,  // 23: ReservationService.HasGuestActiveReservationInFuture:output_type -> BoolResponse
	9,  // 24: ReservationService.HasHostActiveReservationInFuture:output_type -> BoolResponse
	9,  // 25: ReservationService.CheckForSuperHost:output_type -> BoolResponse
	6,  // 26: ReservationService.GetHostIdsForSuperHost:output_type -> IdList
	14, // [14:27] is the sub-list for method output_type
	1,  // [1:14] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_reservation_proto_init() }
func file_proto_reservation_proto_init() {
	if File_proto_reservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_reservation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
		file_proto_reservation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReservationRequest); i {
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
		file_proto_reservation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReservationRequest); i {
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
		file_proto_reservation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationStatus); i {
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
		file_proto_reservation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationList); i {
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
		file_proto_reservation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_reservation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdList); i {
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
		file_proto_reservation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterTakenAccommodationsRequest); i {
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
		file_proto_reservation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntervalRequest); i {
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
		file_proto_reservation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResponse); i {
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
			RawDescriptor: file_proto_reservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_reservation_proto_goTypes,
		DependencyIndexes: file_proto_reservation_proto_depIdxs,
		MessageInfos:      file_proto_reservation_proto_msgTypes,
	}.Build()
	File_proto_reservation_proto = out.File
	file_proto_reservation_proto_rawDesc = nil
	file_proto_reservation_proto_goTypes = nil
	file_proto_reservation_proto_depIdxs = nil
}
