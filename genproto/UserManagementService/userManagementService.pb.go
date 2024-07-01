// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: UserManagementService/userManagementService.proto

package UserManagementService

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

type IdUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *IdUserRequest) Reset() {
	*x = IdUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserManagementService_userManagementService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdUserRequest) ProtoMessage() {}

func (x *IdUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserManagementService_userManagementService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdUserRequest.ProtoReflect.Descriptor instead.
func (*IdUserRequest) Descriptor() ([]byte, []int) {
	return file_UserManagementService_userManagementService_proto_rawDescGZIP(), []int{0}
}

func (x *IdUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserManagementService_userManagementService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserManagementService_userManagementService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_UserManagementService_userManagementService_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserManagementService_userManagementService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserManagementService_userManagementService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_UserManagementService_userManagementService_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserManagementService_userManagementService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserManagementService_userManagementService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_UserManagementService_userManagementService_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName  string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Bio       string `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	Expertise string `protobuf:"bytes,4,opt,name=expertise,proto3" json:"expertise,omitempty"`
	Location  string `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	AvatarUrl string `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserManagementService_userManagementService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_UserManagementService_userManagementService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_UserManagementService_userManagementService_proto_rawDescGZIP(), []int{4}
}

func (x *UserProfile) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserProfile) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *UserProfile) GetExpertise() string {
	if x != nil {
		return x.Expertise
	}
	return ""
}

func (x *UserProfile) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UserProfile) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type UserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserProfile *UserProfile `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
}

func (x *UserProfileRequest) Reset() {
	*x = UserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserManagementService_userManagementService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileRequest) ProtoMessage() {}

func (x *UserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserManagementService_userManagementService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileRequest.ProtoReflect.Descriptor instead.
func (*UserProfileRequest) Descriptor() ([]byte, []int) {
	return file_UserManagementService_userManagementService_proto_rawDescGZIP(), []int{5}
}

func (x *UserProfileRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserProfileRequest) GetUserProfile() *UserProfile {
	if x != nil {
		return x.UserProfile
	}
	return nil
}

type UserProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserProfile *UserProfile `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
}

func (x *UserProfileResponse) Reset() {
	*x = UserProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserManagementService_userManagementService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileResponse) ProtoMessage() {}

func (x *UserProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserManagementService_userManagementService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileResponse.ProtoReflect.Descriptor instead.
func (*UserProfileResponse) Descriptor() ([]byte, []int) {
	return file_UserManagementService_userManagementService_proto_rawDescGZIP(), []int{6}
}

func (x *UserProfileResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserProfileResponse) GetUserProfile() *UserProfile {
	if x != nil {
		return x.UserProfile
	}
	return nil
}

var File_UserManagementService_userManagementService_proto protoreflect.FileDescriptor

var file_UserManagementService_userManagementService_proto_rawDesc = []byte{
	0x0a, 0x31, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x0d, 0x49, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x97, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72,
	0x6c, 0x22, 0x63, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x64, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xfd, 0x03, 0x0a,
	0x15, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x17, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x2e, 0x49, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UserManagementService_userManagementService_proto_rawDescOnce sync.Once
	file_UserManagementService_userManagementService_proto_rawDescData = file_UserManagementService_userManagementService_proto_rawDesc
)

func file_UserManagementService_userManagementService_proto_rawDescGZIP() []byte {
	file_UserManagementService_userManagementService_proto_rawDescOnce.Do(func() {
		file_UserManagementService_userManagementService_proto_rawDescData = protoimpl.X.CompressGZIP(file_UserManagementService_userManagementService_proto_rawDescData)
	})
	return file_UserManagementService_userManagementService_proto_rawDescData
}

var file_UserManagementService_userManagementService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_UserManagementService_userManagementService_proto_goTypes = []any{
	(*IdUserRequest)(nil),       // 0: User.IdUserRequest
	(*UpdateUserRequest)(nil),   // 1: User.UpdateUserRequest
	(*DeleteUserResponse)(nil),  // 2: User.DeleteUserResponse
	(*UserResponse)(nil),        // 3: User.UserResponse
	(*UserProfile)(nil),         // 4: User.UserProfile
	(*UserProfileRequest)(nil),  // 5: User.UserProfileRequest
	(*UserProfileResponse)(nil), // 6: User.UserProfileResponse
}
var file_UserManagementService_userManagementService_proto_depIdxs = []int32{
	4, // 0: User.UserProfileRequest.user_profile:type_name -> User.UserProfile
	4, // 1: User.UserProfileResponse.user_profile:type_name -> User.UserProfile
	0, // 2: User.UserManagementService.GetUserById:input_type -> User.IdUserRequest
	1, // 3: User.UserManagementService.UpdateUserById:input_type -> User.UpdateUserRequest
	0, // 4: User.UserManagementService.DeleteUserById:input_type -> User.IdUserRequest
	0, // 5: User.UserManagementService.GetUserProfile:input_type -> User.IdUserRequest
	5, // 6: User.UserManagementService.UpdateUserProfile:input_type -> User.UserProfileRequest
	0, // 7: User.UserManagementService.GetUserProfileById:input_type -> User.IdUserRequest
	5, // 8: User.UserManagementService.UpdateUserProfileById:input_type -> User.UserProfileRequest
	3, // 9: User.UserManagementService.GetUserById:output_type -> User.UserResponse
	3, // 10: User.UserManagementService.UpdateUserById:output_type -> User.UserResponse
	2, // 11: User.UserManagementService.DeleteUserById:output_type -> User.DeleteUserResponse
	6, // 12: User.UserManagementService.GetUserProfile:output_type -> User.UserProfileResponse
	6, // 13: User.UserManagementService.UpdateUserProfile:output_type -> User.UserProfileResponse
	6, // 14: User.UserManagementService.GetUserProfileById:output_type -> User.UserProfileResponse
	6, // 15: User.UserManagementService.UpdateUserProfileById:output_type -> User.UserProfileResponse
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_UserManagementService_userManagementService_proto_init() }
func file_UserManagementService_userManagementService_proto_init() {
	if File_UserManagementService_userManagementService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UserManagementService_userManagementService_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IdUserRequest); i {
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
		file_UserManagementService_userManagementService_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserRequest); i {
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
		file_UserManagementService_userManagementService_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserResponse); i {
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
		file_UserManagementService_userManagementService_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserResponse); i {
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
		file_UserManagementService_userManagementService_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserProfile); i {
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
		file_UserManagementService_userManagementService_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UserProfileRequest); i {
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
		file_UserManagementService_userManagementService_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserProfileResponse); i {
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
			RawDescriptor: file_UserManagementService_userManagementService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_UserManagementService_userManagementService_proto_goTypes,
		DependencyIndexes: file_UserManagementService_userManagementService_proto_depIdxs,
		MessageInfos:      file_UserManagementService_userManagementService_proto_msgTypes,
	}.Build()
	File_UserManagementService_userManagementService_proto = out.File
	file_UserManagementService_userManagementService_proto_rawDesc = nil
	file_UserManagementService_userManagementService_proto_goTypes = nil
	file_UserManagementService_userManagementService_proto_depIdxs = nil
}