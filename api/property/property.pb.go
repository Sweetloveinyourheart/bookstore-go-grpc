// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: property/property.proto

package proto

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

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{0}
}

func (x *ResultResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetType string `protobuf:"bytes,1,opt,name=assetType,proto3" json:"assetType,omitempty"`
}

func (x *GetAssetsRequest) Reset() {
	*x = GetAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetsRequest) ProtoMessage() {}

func (x *GetAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetsRequest.ProtoReflect.Descriptor instead.
func (*GetAssetsRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{1}
}

func (x *GetAssetsRequest) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

type Assets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	IconURL string `protobuf:"bytes,2,opt,name=iconURL,proto3" json:"iconURL,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Assets) Reset() {
	*x = Assets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assets) ProtoMessage() {}

func (x *Assets) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assets.ProtoReflect.Descriptor instead.
func (*Assets) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{2}
}

func (x *Assets) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Assets) GetIconURL() string {
	if x != nil {
		return x.IconURL
	}
	return ""
}

func (x *Assets) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assets []*Assets `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (x *GetAssetsResponse) Reset() {
	*x = GetAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetsResponse) ProtoMessage() {}

func (x *GetAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetsResponse.ProtoReflect.Descriptor instead.
func (*GetAssetsResponse) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssetsResponse) GetAssets() []*Assets {
	if x != nil {
		return x.Assets
	}
	return nil
}

type AddAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetType string `protobuf:"bytes,1,opt,name=assetType,proto3" json:"assetType,omitempty"`
	IconURL   string `protobuf:"bytes,2,opt,name=iconURL,proto3" json:"iconURL,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddAssetRequest) Reset() {
	*x = AddAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAssetRequest) ProtoMessage() {}

func (x *AddAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAssetRequest.ProtoReflect.Descriptor instead.
func (*AddAssetRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{4}
}

func (x *AddAssetRequest) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *AddAssetRequest) GetIconURL() string {
	if x != nil {
		return x.IconURL
	}
	return ""
}

func (x *AddAssetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ModifyAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetType string `protobuf:"bytes,1,opt,name=assetType,proto3" json:"assetType,omitempty"`
	Id        uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	IconURL   string `protobuf:"bytes,3,opt,name=iconURL,proto3" json:"iconURL,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ModifyAssetRequest) Reset() {
	*x = ModifyAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyAssetRequest) ProtoMessage() {}

func (x *ModifyAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyAssetRequest.ProtoReflect.Descriptor instead.
func (*ModifyAssetRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{5}
}

func (x *ModifyAssetRequest) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *ModifyAssetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ModifyAssetRequest) GetIconURL() string {
	if x != nil {
		return x.IconURL
	}
	return ""
}

func (x *ModifyAssetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DisableAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetType string `protobuf:"bytes,1,opt,name=assetType,proto3" json:"assetType,omitempty"`
	Id        uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DisableAssetRequest) Reset() {
	*x = DisableAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableAssetRequest) ProtoMessage() {}

func (x *DisableAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableAssetRequest.ProtoReflect.Descriptor instead.
func (*DisableAssetRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{6}
}

func (x *DisableAssetRequest) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *DisableAssetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NewDestinationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country   string  `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	City      string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Latitude  float32 `protobuf:"fixed32,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *NewDestinationRequest) Reset() {
	*x = NewDestinationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewDestinationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewDestinationRequest) ProtoMessage() {}

func (x *NewDestinationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewDestinationRequest.ProtoReflect.Descriptor instead.
func (*NewDestinationRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{7}
}

func (x *NewDestinationRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *NewDestinationRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *NewDestinationRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *NewDestinationRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type NewPropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId      uint32                 `protobuf:"varint,1,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	CategoryId  uint32                 `protobuf:"varint,5,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	AmenityId   uint32                 `protobuf:"varint,6,opt,name=amenity_id,json=amenityId,proto3" json:"amenity_id,omitempty"`
	Destination *NewDestinationRequest `protobuf:"bytes,7,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *NewPropertyRequest) Reset() {
	*x = NewPropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPropertyRequest) ProtoMessage() {}

func (x *NewPropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPropertyRequest.ProtoReflect.Descriptor instead.
func (*NewPropertyRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{8}
}

func (x *NewPropertyRequest) GetHostId() uint32 {
	if x != nil {
		return x.HostId
	}
	return 0
}

func (x *NewPropertyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewPropertyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewPropertyRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *NewPropertyRequest) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *NewPropertyRequest) GetAmenityId() uint32 {
	if x != nil {
		return x.AmenityId
	}
	return 0
}

func (x *NewPropertyRequest) GetDestination() *NewDestinationRequest {
	if x != nil {
		return x.Destination
	}
	return nil
}

type EditDestinationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country   string  `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	City      string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Latitude  float32 `protobuf:"fixed32,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *EditDestinationRequest) Reset() {
	*x = EditDestinationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditDestinationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditDestinationRequest) ProtoMessage() {}

func (x *EditDestinationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditDestinationRequest.ProtoReflect.Descriptor instead.
func (*EditDestinationRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{9}
}

func (x *EditDestinationRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *EditDestinationRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *EditDestinationRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *EditDestinationRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type EditPropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string                  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32                 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	CategoryId  uint32                  `protobuf:"varint,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	AmenityId   uint32                  `protobuf:"varint,5,opt,name=amenity_id,json=amenityId,proto3" json:"amenity_id,omitempty"`
	Destination *EditDestinationRequest `protobuf:"bytes,6,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *EditPropertyRequest) Reset() {
	*x = EditPropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditPropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditPropertyRequest) ProtoMessage() {}

func (x *EditPropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditPropertyRequest.ProtoReflect.Descriptor instead.
func (*EditPropertyRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{10}
}

func (x *EditPropertyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EditPropertyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EditPropertyRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *EditPropertyRequest) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *EditPropertyRequest) GetAmenityId() uint32 {
	if x != nil {
		return x.AmenityId
	}
	return 0
}

func (x *EditPropertyRequest) GetDestination() *EditDestinationRequest {
	if x != nil {
		return x.Destination
	}
	return nil
}

type SyncPropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncPropertiesRequest) Reset() {
	*x = SyncPropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPropertiesRequest) ProtoMessage() {}

func (x *SyncPropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPropertiesRequest.ProtoReflect.Descriptor instead.
func (*SyncPropertiesRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{11}
}

var File_property_property_proto protoreflect.FileDescriptor

var file_property_property_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x06, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x06, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x22, 0x5d, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x52,
	0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x52, 0x4c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e,
	0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55,
	0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7f, 0x0a, 0x15, 0x4e,
	0x65, 0x77, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xf5, 0x01, 0x0a,
	0x12, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x6d, 0x65, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x61, 0x6d, 0x65, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x4e, 0x65, 0x77, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x6d, 0x65, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x61, 0x6d, 0x65, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x79, 0x6e, 0x63,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0x80, 0x03, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x32,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x10,
	0x2e, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x12, 0x13, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x13, 0x2e, 0x4e,
	0x65, 0x77, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x12, 0x14, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x53, 0x79, 0x6e,
	0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x77, 0x65, 0x65, 0x74, 0x6c, 0x6f, 0x76, 0x65, 0x69, 0x6e, 0x79, 0x6f,
	0x75, 0x72, 0x68, 0x65, 0x61, 0x72, 0x74, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_property_property_proto_rawDescOnce sync.Once
	file_property_property_proto_rawDescData = file_property_property_proto_rawDesc
)

func file_property_property_proto_rawDescGZIP() []byte {
	file_property_property_proto_rawDescOnce.Do(func() {
		file_property_property_proto_rawDescData = protoimpl.X.CompressGZIP(file_property_property_proto_rawDescData)
	})
	return file_property_property_proto_rawDescData
}

var file_property_property_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_property_property_proto_goTypes = []interface{}{
	(*ResultResponse)(nil),         // 0: ResultResponse
	(*GetAssetsRequest)(nil),       // 1: GetAssetsRequest
	(*Assets)(nil),                 // 2: Assets
	(*GetAssetsResponse)(nil),      // 3: GetAssetsResponse
	(*AddAssetRequest)(nil),        // 4: AddAssetRequest
	(*ModifyAssetRequest)(nil),     // 5: ModifyAssetRequest
	(*DisableAssetRequest)(nil),    // 6: DisableAssetRequest
	(*NewDestinationRequest)(nil),  // 7: NewDestinationRequest
	(*NewPropertyRequest)(nil),     // 8: NewPropertyRequest
	(*EditDestinationRequest)(nil), // 9: EditDestinationRequest
	(*EditPropertyRequest)(nil),    // 10: EditPropertyRequest
	(*SyncPropertiesRequest)(nil),  // 11: SyncPropertiesRequest
}
var file_property_property_proto_depIdxs = []int32{
	2,  // 0: GetAssetsResponse.assets:type_name -> Assets
	7,  // 1: NewPropertyRequest.destination:type_name -> NewDestinationRequest
	9,  // 2: EditPropertyRequest.destination:type_name -> EditDestinationRequest
	1,  // 3: Property.GetAssets:input_type -> GetAssetsRequest
	4,  // 4: Property.AddAsset:input_type -> AddAssetRequest
	5,  // 5: Property.ModifyAsset:input_type -> ModifyAssetRequest
	6,  // 6: Property.DisableAsset:input_type -> DisableAssetRequest
	8,  // 7: Property.AddProperty:input_type -> NewPropertyRequest
	10, // 8: Property.EditProperty:input_type -> EditPropertyRequest
	11, // 9: Property.SyncProperties:input_type -> SyncPropertiesRequest
	3,  // 10: Property.GetAssets:output_type -> GetAssetsResponse
	0,  // 11: Property.AddAsset:output_type -> ResultResponse
	0,  // 12: Property.ModifyAsset:output_type -> ResultResponse
	0,  // 13: Property.DisableAsset:output_type -> ResultResponse
	0,  // 14: Property.AddProperty:output_type -> ResultResponse
	0,  // 15: Property.EditProperty:output_type -> ResultResponse
	0,  // 16: Property.SyncProperties:output_type -> ResultResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_property_property_proto_init() }
func file_property_property_proto_init() {
	if File_property_property_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_property_property_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResponse); i {
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
		file_property_property_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetsRequest); i {
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
		file_property_property_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Assets); i {
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
		file_property_property_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetsResponse); i {
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
		file_property_property_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAssetRequest); i {
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
		file_property_property_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyAssetRequest); i {
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
		file_property_property_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableAssetRequest); i {
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
		file_property_property_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewDestinationRequest); i {
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
		file_property_property_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPropertyRequest); i {
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
		file_property_property_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditDestinationRequest); i {
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
		file_property_property_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditPropertyRequest); i {
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
		file_property_property_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPropertiesRequest); i {
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
			RawDescriptor: file_property_property_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_property_property_proto_goTypes,
		DependencyIndexes: file_property_property_proto_depIdxs,
		MessageInfos:      file_property_property_proto_msgTypes,
	}.Build()
	File_property_property_proto = out.File
	file_property_property_proto_rawDesc = nil
	file_property_property_proto_goTypes = nil
	file_property_property_proto_depIdxs = nil
}
