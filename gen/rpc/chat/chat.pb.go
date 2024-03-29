// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: rpc/chat/chat.proto

package chat

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

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChannelType   int32    `protobuf:"varint,2,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	Type          int32    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Content       string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	From          string   `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To            string   `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	SendTime      int64    `protobuf:"varint,7,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	Uuid          string   `protobuf:"bytes,8,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AtUserList    []string `protobuf:"bytes,9,rep,name=at_user_list,json=atUserList,proto3" json:"at_user_list,omitempty"`
	IsTransparent bool     `protobuf:"varint,10,opt,name=is_transparent,json=isTransparent,proto3" json:"is_transparent,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Msg) GetChannelType() int32 {
	if x != nil {
		return x.ChannelType
	}
	return 0
}

func (x *Msg) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Msg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Msg) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Msg) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Msg) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *Msg) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Msg) GetAtUserList() []string {
	if x != nil {
		return x.AtUserList
	}
	return nil
}

func (x *Msg) GetIsTransparent() bool {
	if x != nil {
		return x.IsTransparent
	}
	return false
}

type SendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelType   int32    `protobuf:"varint,1,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	MsgType       int32    `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	From          string   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string   `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Content       string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Uuid          string   `protobuf:"bytes,6,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AtUserList    []string `protobuf:"bytes,7,rep,name=at_user_list,json=atUserList,proto3" json:"at_user_list,omitempty"`
	IsTransparent bool     `protobuf:"varint,8,opt,name=is_transparent,json=isTransparent,proto3" json:"is_transparent,omitempty"`
}

func (x *SendReq) Reset() {
	*x = SendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReq) ProtoMessage() {}

func (x *SendReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReq.ProtoReflect.Descriptor instead.
func (*SendReq) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SendReq) GetChannelType() int32 {
	if x != nil {
		return x.ChannelType
	}
	return 0
}

func (x *SendReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *SendReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendReq) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SendReq) GetAtUserList() []string {
	if x != nil {
		return x.AtUserList
	}
	return nil
}

func (x *SendReq) GetIsTransparent() bool {
	if x != nil {
		return x.IsTransparent
	}
	return false
}

type SendRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Id         int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	SendTime   int64  `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	ClientUuid string `protobuf:"bytes,5,opt,name=client_uuid,json=clientUuid,proto3" json:"client_uuid,omitempty"`
}

func (x *SendRsp) Reset() {
	*x = SendRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRsp) ProtoMessage() {}

func (x *SendRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRsp.ProtoReflect.Descriptor instead.
func (*SendRsp) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *SendRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SendRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendRsp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SendRsp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *SendRsp) GetClientUuid() string {
	if x != nil {
		return x.ClientUuid
	}
	return ""
}

type SyncMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin    string `protobuf:"bytes,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SyncMsgReq) Reset() {
	*x = SyncMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMsgReq) ProtoMessage() {}

func (x *SyncMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMsgReq.ProtoReflect.Descriptor instead.
func (*SyncMsgReq) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *SyncMsgReq) GetUin() string {
	if x != nil {
		return x.Uin
	}
	return ""
}

func (x *SyncMsgReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SyncMsgReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SyncMsgRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Msg `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *SyncMsgRsp) Reset() {
	*x = SyncMsgRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMsgRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMsgRsp) ProtoMessage() {}

func (x *SyncMsgRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMsgRsp.ProtoReflect.Descriptor instead.
func (*SyncMsgRsp) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *SyncMsgRsp) GetList() []*Msg {
	if x != nil {
		return x.List
	}
	return nil
}

type MsgAckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin string `protobuf:"bytes,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MsgAckReq) Reset() {
	*x = MsgAckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAckReq) ProtoMessage() {}

func (x *MsgAckReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAckReq.ProtoReflect.Descriptor instead.
func (*MsgAckReq) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{5}
}

func (x *MsgAckReq) GetUin() string {
	if x != nil {
		return x.Uin
	}
	return ""
}

func (x *MsgAckReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MsgAckRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgAckRsp) Reset() {
	*x = MsgAckRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAckRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAckRsp) ProtoMessage() {}

func (x *MsgAckRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAckRsp.ProtoReflect.Descriptor instead.
func (*MsgAckRsp) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{6}
}

type RecallReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin string `protobuf:"bytes,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RecallReq) Reset() {
	*x = RecallReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecallReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecallReq) ProtoMessage() {}

func (x *RecallReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecallReq.ProtoReflect.Descriptor instead.
func (*RecallReq) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{7}
}

func (x *RecallReq) GetUin() string {
	if x != nil {
		return x.Uin
	}
	return ""
}

func (x *RecallReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RecallRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecallRsp) Reset() {
	*x = RecallRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecallRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecallRsp) ProtoMessage() {}

func (x *RecallRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecallRsp.ProtoReflect.Descriptor instead.
func (*RecallRsp) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{8}
}

type DeleteMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin string  `protobuf:"bytes,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Ids []int64 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteMsgReq) Reset() {
	*x = DeleteMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMsgReq) ProtoMessage() {}

func (x *DeleteMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMsgReq.ProtoReflect.Descriptor instead.
func (*DeleteMsgReq) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMsgReq) GetUin() string {
	if x != nil {
		return x.Uin
	}
	return ""
}

func (x *DeleteMsgReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteMsgRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMsgRsp) Reset() {
	*x = DeleteMsgRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMsgRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMsgRsp) ProtoMessage() {}

func (x *DeleteMsgRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMsgRsp.ProtoReflect.Descriptor instead.
func (*DeleteMsgRsp) Descriptor() ([]byte, []int) {
	return file_rpc_chat_chat_proto_rawDescGZIP(), []int{10}
}

var File_rpc_chat_chat_proto protoreflect.FileDescriptor

var file_rpc_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x22,
	0x84, 0x02, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0c, 0x61, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0xe2, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x07,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55,
	0x75, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x2f, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x52, 0x73, 0x70, 0x12,
	0x21, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x2d, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x0b, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x41, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x2d,
	0x0a, 0x09, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0b, 0x0a,
	0x09, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x73, 0x70, 0x22, 0x32, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x0e,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x73, 0x70, 0x32, 0x9d,
	0x02, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x73, 0x67, 0x12, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x53, 0x79,
	0x6e, 0x63, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x52, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x41, 0x63, 0x6b, 0x12, 0x13, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x73,
	0x67, 0x41, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x52, 0x65, 0x63,
	0x61, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52,
	0x65, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x81,
	0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x42,
	0x09, 0x43, 0x68, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x62, 0x72, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x2f, 0x7a, 0x69, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0xa2, 0x02, 0x03, 0x52, 0x43, 0x58, 0xaa, 0x02, 0x08, 0x52, 0x70, 0x63, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0xca, 0x02, 0x08, 0x52, 0x70, 0x63, 0x5c, 0x43, 0x68, 0x61, 0x74, 0xe2,
	0x02, 0x14, 0x52, 0x70, 0x63, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x43, 0x68,
	0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_chat_chat_proto_rawDescOnce sync.Once
	file_rpc_chat_chat_proto_rawDescData = file_rpc_chat_chat_proto_rawDesc
)

func file_rpc_chat_chat_proto_rawDescGZIP() []byte {
	file_rpc_chat_chat_proto_rawDescOnce.Do(func() {
		file_rpc_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_chat_chat_proto_rawDescData)
	})
	return file_rpc_chat_chat_proto_rawDescData
}

var file_rpc_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_rpc_chat_chat_proto_goTypes = []interface{}{
	(*Msg)(nil),          // 0: rpc.chat.Msg
	(*SendReq)(nil),      // 1: rpc.chat.SendReq
	(*SendRsp)(nil),      // 2: rpc.chat.SendRsp
	(*SyncMsgReq)(nil),   // 3: rpc.chat.SyncMsgReq
	(*SyncMsgRsp)(nil),   // 4: rpc.chat.SyncMsgRsp
	(*MsgAckReq)(nil),    // 5: rpc.chat.MsgAckReq
	(*MsgAckRsp)(nil),    // 6: rpc.chat.MsgAckRsp
	(*RecallReq)(nil),    // 7: rpc.chat.RecallReq
	(*RecallRsp)(nil),    // 8: rpc.chat.RecallRsp
	(*DeleteMsgReq)(nil), // 9: rpc.chat.DeleteMsgReq
	(*DeleteMsgRsp)(nil), // 10: rpc.chat.DeleteMsgRsp
}
var file_rpc_chat_chat_proto_depIdxs = []int32{
	0,  // 0: rpc.chat.SyncMsgRsp.list:type_name -> rpc.chat.Msg
	1,  // 1: rpc.chat.Chat.SendMsg:input_type -> rpc.chat.SendReq
	3,  // 2: rpc.chat.Chat.SyncMsg:input_type -> rpc.chat.SyncMsgReq
	5,  // 3: rpc.chat.Chat.MsgAck:input_type -> rpc.chat.MsgAckReq
	7,  // 4: rpc.chat.Chat.Recall:input_type -> rpc.chat.RecallReq
	9,  // 5: rpc.chat.Chat.DeleteMsg:input_type -> rpc.chat.DeleteMsgReq
	2,  // 6: rpc.chat.Chat.SendMsg:output_type -> rpc.chat.SendRsp
	4,  // 7: rpc.chat.Chat.SyncMsg:output_type -> rpc.chat.SyncMsgRsp
	6,  // 8: rpc.chat.Chat.MsgAck:output_type -> rpc.chat.MsgAckRsp
	8,  // 9: rpc.chat.Chat.Recall:output_type -> rpc.chat.RecallRsp
	10, // 10: rpc.chat.Chat.DeleteMsg:output_type -> rpc.chat.DeleteMsgRsp
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_chat_chat_proto_init() }
func file_rpc_chat_chat_proto_init() {
	if File_rpc_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_rpc_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReq); i {
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
		file_rpc_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRsp); i {
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
		file_rpc_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMsgReq); i {
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
		file_rpc_chat_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMsgRsp); i {
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
		file_rpc_chat_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAckReq); i {
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
		file_rpc_chat_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAckRsp); i {
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
		file_rpc_chat_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecallReq); i {
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
		file_rpc_chat_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecallRsp); i {
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
		file_rpc_chat_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMsgReq); i {
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
		file_rpc_chat_chat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMsgRsp); i {
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
			RawDescriptor: file_rpc_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_chat_chat_proto_goTypes,
		DependencyIndexes: file_rpc_chat_chat_proto_depIdxs,
		MessageInfos:      file_rpc_chat_chat_proto_msgTypes,
	}.Build()
	File_rpc_chat_chat_proto = out.File
	file_rpc_chat_chat_proto_rawDesc = nil
	file_rpc_chat_chat_proto_goTypes = nil
	file_rpc_chat_chat_proto_depIdxs = nil
}
