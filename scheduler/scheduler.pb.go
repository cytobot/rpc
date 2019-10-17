// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler.proto

package com_cyto_scheduler

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ScheduledTask_IntervalType int32

const (
	ScheduledTask_NONE    ScheduledTask_IntervalType = 0
	ScheduledTask_SECONDS ScheduledTask_IntervalType = 1
	ScheduledTask_MINUTES ScheduledTask_IntervalType = 2
	ScheduledTask_HOURS   ScheduledTask_IntervalType = 3
	ScheduledTask_DAYS    ScheduledTask_IntervalType = 4
	ScheduledTask_WEEKS   ScheduledTask_IntervalType = 5
	ScheduledTask_MONTHS  ScheduledTask_IntervalType = 6
	ScheduledTask_YEARS   ScheduledTask_IntervalType = 7
)

var ScheduledTask_IntervalType_name = map[int32]string{
	0: "NONE",
	1: "SECONDS",
	2: "MINUTES",
	3: "HOURS",
	4: "DAYS",
	5: "WEEKS",
	6: "MONTHS",
	7: "YEARS",
}

var ScheduledTask_IntervalType_value = map[string]int32{
	"NONE":    0,
	"SECONDS": 1,
	"MINUTES": 2,
	"HOURS":   3,
	"DAYS":    4,
	"WEEKS":   5,
	"MONTHS":  6,
	"YEARS":   7,
}

func (x ScheduledTask_IntervalType) String() string {
	return proto.EnumName(ScheduledTask_IntervalType_name, int32(x))
}

func (ScheduledTask_IntervalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{4, 0}
}

type ScheduledTask_EventType int32

const (
	ScheduledTask_CHANNEL_REMINDER ScheduledTask_EventType = 0
	ScheduledTask_DIRECT_REMINDER  ScheduledTask_EventType = 1
)

var ScheduledTask_EventType_name = map[int32]string{
	0: "CHANNEL_REMINDER",
	1: "DIRECT_REMINDER",
}

var ScheduledTask_EventType_value = map[string]int32{
	"CHANNEL_REMINDER": 0,
	"DIRECT_REMINDER":  1,
}

func (x ScheduledTask_EventType) String() string {
	return proto.EnumName(ScheduledTask_EventType_name, int32(x))
}

func (ScheduledTask_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{4, 1}
}

type GuildQuery struct {
	GuildID              string   `protobuf:"bytes,1,opt,name=guildID,proto3" json:"guildID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GuildQuery) Reset()         { *m = GuildQuery{} }
func (m *GuildQuery) String() string { return proto.CompactTextString(m) }
func (*GuildQuery) ProtoMessage()    {}
func (*GuildQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{0}
}

func (m *GuildQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuildQuery.Unmarshal(m, b)
}
func (m *GuildQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuildQuery.Marshal(b, m, deterministic)
}
func (m *GuildQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuildQuery.Merge(m, src)
}
func (m *GuildQuery) XXX_Size() int {
	return xxx_messageInfo_GuildQuery.Size(m)
}
func (m *GuildQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GuildQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GuildQuery proto.InternalMessageInfo

func (m *GuildQuery) GetGuildID() string {
	if m != nil {
		return m.GuildID
	}
	return ""
}

type UserQuery struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserQuery) Reset()         { *m = UserQuery{} }
func (m *UserQuery) String() string { return proto.CompactTextString(m) }
func (*UserQuery) ProtoMessage()    {}
func (*UserQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{1}
}

func (m *UserQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserQuery.Unmarshal(m, b)
}
func (m *UserQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserQuery.Marshal(b, m, deterministic)
}
func (m *UserQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserQuery.Merge(m, src)
}
func (m *UserQuery) XXX_Size() int {
	return xxx_messageInfo_UserQuery.Size(m)
}
func (m *UserQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_UserQuery.DiscardUnknown(m)
}

var xxx_messageInfo_UserQuery proto.InternalMessageInfo

func (m *UserQuery) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type ScheduledTaskQuery struct {
	ScheduledTaskID      string   `protobuf:"bytes,1,opt,name=scheduledTaskID,proto3" json:"scheduledTaskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduledTaskQuery) Reset()         { *m = ScheduledTaskQuery{} }
func (m *ScheduledTaskQuery) String() string { return proto.CompactTextString(m) }
func (*ScheduledTaskQuery) ProtoMessage()    {}
func (*ScheduledTaskQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{2}
}

func (m *ScheduledTaskQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduledTaskQuery.Unmarshal(m, b)
}
func (m *ScheduledTaskQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduledTaskQuery.Marshal(b, m, deterministic)
}
func (m *ScheduledTaskQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledTaskQuery.Merge(m, src)
}
func (m *ScheduledTaskQuery) XXX_Size() int {
	return xxx_messageInfo_ScheduledTaskQuery.Size(m)
}
func (m *ScheduledTaskQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledTaskQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledTaskQuery proto.InternalMessageInfo

func (m *ScheduledTaskQuery) GetScheduledTaskID() string {
	if m != nil {
		return m.ScheduledTaskID
	}
	return ""
}

type ScheduledTaskList struct {
	Tasks                []*ScheduledTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ScheduledTaskList) Reset()         { *m = ScheduledTaskList{} }
func (m *ScheduledTaskList) String() string { return proto.CompactTextString(m) }
func (*ScheduledTaskList) ProtoMessage()    {}
func (*ScheduledTaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{3}
}

func (m *ScheduledTaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduledTaskList.Unmarshal(m, b)
}
func (m *ScheduledTaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduledTaskList.Marshal(b, m, deterministic)
}
func (m *ScheduledTaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledTaskList.Merge(m, src)
}
func (m *ScheduledTaskList) XXX_Size() int {
	return xxx_messageInfo_ScheduledTaskList.Size(m)
}
func (m *ScheduledTaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledTaskList.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledTaskList proto.InternalMessageInfo

func (m *ScheduledTaskList) GetTasks() []*ScheduledTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type ScheduledTask struct {
	ScheduledTaskID      string                     `protobuf:"bytes,1,opt,name=scheduledTaskID,proto3" json:"scheduledTaskID,omitempty"`
	CreatedUserID        string                     `protobuf:"bytes,2,opt,name=createdUserID,proto3" json:"createdUserID,omitempty"`
	CreatedDateUtc       *timestamp.Timestamp       `protobuf:"bytes,3,opt,name=createdDateUtc,proto3" json:"createdDateUtc,omitempty"`
	GuildID              string                     `protobuf:"bytes,4,opt,name=guildID,proto3" json:"guildID,omitempty"`
	ChannelID            string                     `protobuf:"bytes,5,opt,name=channelID,proto3" json:"channelID,omitempty"`
	IntervalType         ScheduledTask_IntervalType `protobuf:"varint,6,opt,name=intervalType,proto3,enum=com.cyto.scheduler.ScheduledTask_IntervalType" json:"intervalType,omitempty"`
	Interval             int32                      `protobuf:"varint,7,opt,name=interval,proto3" json:"interval,omitempty"`
	EventType            ScheduledTask_EventType    `protobuf:"varint,8,opt,name=eventType,proto3,enum=com.cyto.scheduler.ScheduledTask_EventType" json:"eventType,omitempty"`
	AdditionalPayload    string                     `protobuf:"bytes,9,opt,name=additionalPayload,proto3" json:"additionalPayload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ScheduledTask) Reset()         { *m = ScheduledTask{} }
func (m *ScheduledTask) String() string { return proto.CompactTextString(m) }
func (*ScheduledTask) ProtoMessage()    {}
func (*ScheduledTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{4}
}

func (m *ScheduledTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduledTask.Unmarshal(m, b)
}
func (m *ScheduledTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduledTask.Marshal(b, m, deterministic)
}
func (m *ScheduledTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledTask.Merge(m, src)
}
func (m *ScheduledTask) XXX_Size() int {
	return xxx_messageInfo_ScheduledTask.Size(m)
}
func (m *ScheduledTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledTask.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledTask proto.InternalMessageInfo

func (m *ScheduledTask) GetScheduledTaskID() string {
	if m != nil {
		return m.ScheduledTaskID
	}
	return ""
}

func (m *ScheduledTask) GetCreatedUserID() string {
	if m != nil {
		return m.CreatedUserID
	}
	return ""
}

func (m *ScheduledTask) GetCreatedDateUtc() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedDateUtc
	}
	return nil
}

func (m *ScheduledTask) GetGuildID() string {
	if m != nil {
		return m.GuildID
	}
	return ""
}

func (m *ScheduledTask) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *ScheduledTask) GetIntervalType() ScheduledTask_IntervalType {
	if m != nil {
		return m.IntervalType
	}
	return ScheduledTask_NONE
}

func (m *ScheduledTask) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *ScheduledTask) GetEventType() ScheduledTask_EventType {
	if m != nil {
		return m.EventType
	}
	return ScheduledTask_CHANNEL_REMINDER
}

func (m *ScheduledTask) GetAdditionalPayload() string {
	if m != nil {
		return m.AdditionalPayload
	}
	return ""
}

func init() {
	proto.RegisterEnum("com.cyto.scheduler.ScheduledTask_IntervalType", ScheduledTask_IntervalType_name, ScheduledTask_IntervalType_value)
	proto.RegisterEnum("com.cyto.scheduler.ScheduledTask_EventType", ScheduledTask_EventType_name, ScheduledTask_EventType_value)
	proto.RegisterType((*GuildQuery)(nil), "com.cyto.scheduler.GuildQuery")
	proto.RegisterType((*UserQuery)(nil), "com.cyto.scheduler.UserQuery")
	proto.RegisterType((*ScheduledTaskQuery)(nil), "com.cyto.scheduler.ScheduledTaskQuery")
	proto.RegisterType((*ScheduledTaskList)(nil), "com.cyto.scheduler.ScheduledTaskList")
	proto.RegisterType((*ScheduledTask)(nil), "com.cyto.scheduler.ScheduledTask")
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor_2b3fc28395a6d9c5) }

var fileDescriptor_2b3fc28395a6d9c5 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x7f, 0x6f, 0x93, 0x50,
	0x14, 0x2d, 0x5b, 0x69, 0xc7, 0xdd, 0xaf, 0xb7, 0x37, 0x63, 0xb0, 0xd1, 0xad, 0xa2, 0x2e, 0x24,
	0x1a, 0x96, 0xd4, 0x44, 0xff, 0x33, 0xa9, 0x03, 0x37, 0xe2, 0x46, 0xf5, 0x41, 0x63, 0xe6, 0x3f,
	0x0b, 0x85, 0x67, 0x8b, 0xa3, 0xa5, 0x81, 0xd7, 0x26, 0xfd, 0x8e, 0x7e, 0x0f, 0xbf, 0x86, 0x01,
	0x0a, 0x2d, 0x5d, 0x0d, 0x33, 0x71, 0x7f, 0xde, 0xf3, 0xce, 0x79, 0xe7, 0xde, 0x93, 0xf7, 0x2e,
	0xec, 0x47, 0xce, 0x80, 0xba, 0x13, 0x9f, 0x86, 0xca, 0x38, 0x0c, 0x58, 0x80, 0xb1, 0x13, 0x0c,
	0x15, 0x67, 0xc6, 0x02, 0x25, 0x3f, 0x69, 0x1c, 0xf7, 0x83, 0xa0, 0xef, 0xd3, 0xd3, 0x84, 0xd1,
	0x9b, 0xfc, 0x38, 0x65, 0xde, 0x90, 0x46, 0xcc, 0x1e, 0x8e, 0x53, 0x91, 0x74, 0x02, 0x70, 0x3e,
	0xf1, 0x7c, 0xf7, 0xeb, 0x84, 0x86, 0x33, 0x2c, 0x42, 0xbd, 0x1f, 0x57, 0xba, 0x2a, 0x72, 0x4d,
	0x4e, 0x16, 0x48, 0x56, 0x4a, 0x2f, 0x40, 0xe8, 0x46, 0x34, 0x4c, 0x69, 0x8f, 0xa1, 0x36, 0x89,
	0x68, 0x98, 0xb3, 0xe6, 0x95, 0xf4, 0x01, 0xb0, 0x39, 0xb7, 0x76, 0x2d, 0x3b, 0xba, 0x4d, 0xd9,
	0xf2, 0xa2, 0xd5, 0x04, 0xcd, 0x65, 0xab, 0xb0, 0x74, 0x09, 0x07, 0x05, 0xfd, 0xa5, 0x17, 0x31,
	0xfc, 0x1e, 0x78, 0x66, 0x47, 0xb7, 0x91, 0xc8, 0x35, 0x37, 0xe5, 0xed, 0xd6, 0x73, 0xe5, 0xee,
	0x98, 0x4a, 0x41, 0x45, 0x52, 0xbe, 0xf4, 0xbb, 0x0a, 0xbb, 0x85, 0x83, 0xfb, 0x77, 0x82, 0x5f,
	0xc2, 0xae, 0x13, 0x52, 0x9b, 0x51, 0xb7, 0x9b, 0x0e, 0xba, 0x91, 0xf0, 0x8a, 0x20, 0xfe, 0x08,
	0x7b, 0x73, 0x40, 0xb5, 0x19, 0xed, 0x32, 0x47, 0xdc, 0x6c, 0x72, 0xf2, 0x76, 0xab, 0xa1, 0xa4,
	0xb1, 0x2b, 0x59, 0xec, 0x8a, 0x95, 0xc5, 0x4e, 0x56, 0x14, 0xcb, 0x91, 0x57, 0x0b, 0x91, 0xe3,
	0xa7, 0x20, 0x38, 0x03, 0x7b, 0x34, 0xa2, 0xbe, 0xae, 0x8a, 0x7c, 0x72, 0xb6, 0x00, 0x30, 0x81,
	0x1d, 0x6f, 0xc4, 0x68, 0x38, 0xb5, 0x7d, 0x6b, 0x36, 0xa6, 0x62, 0xad, 0xc9, 0xc9, 0x7b, 0x2d,
	0xa5, 0x34, 0x1d, 0x45, 0x5f, 0x52, 0x91, 0xc2, 0x1d, 0xb8, 0x01, 0x5b, 0x59, 0x2d, 0xd6, 0x9b,
	0x9c, 0xcc, 0x93, 0xbc, 0xc6, 0x3a, 0x08, 0x74, 0x4a, 0x47, 0x2c, 0x31, 0xdb, 0x4a, 0xcc, 0x5e,
	0x97, 0x9b, 0x69, 0x99, 0x84, 0x2c, 0xd4, 0xf8, 0x0d, 0x1c, 0xd8, 0xae, 0xeb, 0x31, 0x2f, 0x18,
	0xd9, 0xfe, 0x17, 0x7b, 0xe6, 0x07, 0xb6, 0x2b, 0x0a, 0xc9, 0x80, 0x77, 0x0f, 0x24, 0x0f, 0x76,
	0x96, 0x5b, 0xc6, 0x5b, 0x50, 0x35, 0x3a, 0x86, 0x86, 0x2a, 0x78, 0x1b, 0xea, 0xa6, 0x76, 0xd6,
	0x31, 0x54, 0x13, 0x71, 0x71, 0x71, 0xa5, 0x1b, 0x5d, 0x4b, 0x33, 0xd1, 0x06, 0x16, 0x80, 0xbf,
	0xe8, 0x74, 0x89, 0x89, 0x36, 0x63, 0xba, 0xda, 0xbe, 0x36, 0x51, 0x35, 0x06, 0xbf, 0x69, 0xda,
	0x67, 0x13, 0xf1, 0x18, 0xa0, 0x76, 0xd5, 0x31, 0xac, 0x0b, 0x13, 0xd5, 0x62, 0xf8, 0x5a, 0x6b,
	0x13, 0x13, 0xd5, 0xa5, 0x77, 0x20, 0xe4, 0x0d, 0xe3, 0x47, 0x80, 0xce, 0x2e, 0xda, 0x86, 0xa1,
	0x5d, 0xde, 0x10, 0xed, 0x4a, 0x37, 0x54, 0x8d, 0xa0, 0x0a, 0x3e, 0x84, 0x7d, 0x55, 0x27, 0xda,
	0x99, 0xb5, 0x00, 0xb9, 0xd6, 0x2f, 0x1e, 0x84, 0x6c, 0xee, 0x10, 0x7f, 0x07, 0x64, 0x52, 0x56,
	0x7c, 0x79, 0xe5, 0xaf, 0xb6, 0x51, 0x4e, 0x91, 0x2a, 0xf8, 0x06, 0xd0, 0xf9, 0xea, 0xdd, 0x27,
	0xa5, 0xc2, 0xe4, 0x1f, 0xde, 0xcf, 0xc0, 0x85, 0x27, 0xab, 0x06, 0xd1, 0xa7, 0x20, 0x4c, 0x76,
	0x04, 0x3e, 0x5a, 0x77, 0xc3, 0x62, 0x7d, 0x34, 0x5e, 0x95, 0x3a, 0xc4, 0x3f, 0x5a, 0xaa, 0xe0,
	0x1e, 0x88, 0xeb, 0x5c, 0xe2, 0x6f, 0x85, 0x9f, 0xad, 0xbb, 0x24, 0xdf, 0x3d, 0xff, 0xe2, 0x71,
	0x48, 0xe8, 0x30, 0x98, 0xd2, 0x07, 0x4c, 0xeb, 0x27, 0x1c, 0xa7, 0x1e, 0x6d, 0xdf, 0x7f, 0xe8,
	0xcc, 0x06, 0x70, 0xf4, 0x77, 0xaf, 0xff, 0x99, 0x5c, 0xaf, 0x96, 0xac, 0xad, 0xb7, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x3e, 0x9a, 0x6e, 0x50, 0x62, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulerClient interface {
	SetScheduledTask(ctx context.Context, in *ScheduledTask, opts ...grpc.CallOption) (*ScheduledTask, error)
	GetScheduledTask(ctx context.Context, in *ScheduledTaskQuery, opts ...grpc.CallOption) (*ScheduledTask, error)
	GetScheduledTasksForGuild(ctx context.Context, in *GuildQuery, opts ...grpc.CallOption) (*ScheduledTaskList, error)
	GetScheduledTasksForUser(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*ScheduledTaskList, error)
	RemoveScheduledTask(ctx context.Context, in *ScheduledTaskQuery, opts ...grpc.CallOption) (*ScheduledTask, error)
	RemoveAllScheduledTasksForGuild(ctx context.Context, in *GuildQuery, opts ...grpc.CallOption) (*ScheduledTaskList, error)
	RemoveAllScheduledTasksForUser(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*ScheduledTaskList, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) SetScheduledTask(ctx context.Context, in *ScheduledTask, opts ...grpc.CallOption) (*ScheduledTask, error) {
	out := new(ScheduledTask)
	err := c.cc.Invoke(ctx, "/com.cyto.scheduler.Scheduler/SetScheduledTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetScheduledTask(ctx context.Context, in *ScheduledTaskQuery, opts ...grpc.CallOption) (*ScheduledTask, error) {
	out := new(ScheduledTask)
	err := c.cc.Invoke(ctx, "/com.cyto.scheduler.Scheduler/GetScheduledTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetScheduledTasksForGuild(ctx context.Context, in *GuildQuery, opts ...grpc.CallOption) (*ScheduledTaskList, error) {
	out := new(ScheduledTaskList)
	err := c.cc.Invoke(ctx, "/com.cyto.scheduler.Scheduler/GetScheduledTasksForGuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetScheduledTasksForUser(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*ScheduledTaskList, error) {
	out := new(ScheduledTaskList)
	err := c.cc.Invoke(ctx, "/com.cyto.scheduler.Scheduler/GetScheduledTasksForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) RemoveScheduledTask(ctx context.Context, in *ScheduledTaskQuery, opts ...grpc.CallOption) (*ScheduledTask, error) {
	out := new(ScheduledTask)
	err := c.cc.Invoke(ctx, "/com.cyto.scheduler.Scheduler/RemoveScheduledTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) RemoveAllScheduledTasksForGuild(ctx context.Context, in *GuildQuery, opts ...grpc.CallOption) (*ScheduledTaskList, error) {
	out := new(ScheduledTaskList)
	err := c.cc.Invoke(ctx, "/com.cyto.scheduler.Scheduler/RemoveAllScheduledTasksForGuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) RemoveAllScheduledTasksForUser(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*ScheduledTaskList, error) {
	out := new(ScheduledTaskList)
	err := c.cc.Invoke(ctx, "/com.cyto.scheduler.Scheduler/RemoveAllScheduledTasksForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
type SchedulerServer interface {
	SetScheduledTask(context.Context, *ScheduledTask) (*ScheduledTask, error)
	GetScheduledTask(context.Context, *ScheduledTaskQuery) (*ScheduledTask, error)
	GetScheduledTasksForGuild(context.Context, *GuildQuery) (*ScheduledTaskList, error)
	GetScheduledTasksForUser(context.Context, *UserQuery) (*ScheduledTaskList, error)
	RemoveScheduledTask(context.Context, *ScheduledTaskQuery) (*ScheduledTask, error)
	RemoveAllScheduledTasksForGuild(context.Context, *GuildQuery) (*ScheduledTaskList, error)
	RemoveAllScheduledTasksForUser(context.Context, *UserQuery) (*ScheduledTaskList, error)
}

// UnimplementedSchedulerServer can be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (*UnimplementedSchedulerServer) SetScheduledTask(ctx context.Context, req *ScheduledTask) (*ScheduledTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetScheduledTask not implemented")
}
func (*UnimplementedSchedulerServer) GetScheduledTask(ctx context.Context, req *ScheduledTaskQuery) (*ScheduledTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledTask not implemented")
}
func (*UnimplementedSchedulerServer) GetScheduledTasksForGuild(ctx context.Context, req *GuildQuery) (*ScheduledTaskList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledTasksForGuild not implemented")
}
func (*UnimplementedSchedulerServer) GetScheduledTasksForUser(ctx context.Context, req *UserQuery) (*ScheduledTaskList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledTasksForUser not implemented")
}
func (*UnimplementedSchedulerServer) RemoveScheduledTask(ctx context.Context, req *ScheduledTaskQuery) (*ScheduledTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveScheduledTask not implemented")
}
func (*UnimplementedSchedulerServer) RemoveAllScheduledTasksForGuild(ctx context.Context, req *GuildQuery) (*ScheduledTaskList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAllScheduledTasksForGuild not implemented")
}
func (*UnimplementedSchedulerServer) RemoveAllScheduledTasksForUser(ctx context.Context, req *UserQuery) (*ScheduledTaskList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAllScheduledTasksForUser not implemented")
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_SetScheduledTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduledTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SetScheduledTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.cyto.scheduler.Scheduler/SetScheduledTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SetScheduledTask(ctx, req.(*ScheduledTask))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetScheduledTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduledTaskQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetScheduledTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.cyto.scheduler.Scheduler/GetScheduledTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetScheduledTask(ctx, req.(*ScheduledTaskQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetScheduledTasksForGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuildQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetScheduledTasksForGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.cyto.scheduler.Scheduler/GetScheduledTasksForGuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetScheduledTasksForGuild(ctx, req.(*GuildQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetScheduledTasksForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetScheduledTasksForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.cyto.scheduler.Scheduler/GetScheduledTasksForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetScheduledTasksForUser(ctx, req.(*UserQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_RemoveScheduledTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduledTaskQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).RemoveScheduledTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.cyto.scheduler.Scheduler/RemoveScheduledTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).RemoveScheduledTask(ctx, req.(*ScheduledTaskQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_RemoveAllScheduledTasksForGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuildQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).RemoveAllScheduledTasksForGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.cyto.scheduler.Scheduler/RemoveAllScheduledTasksForGuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).RemoveAllScheduledTasksForGuild(ctx, req.(*GuildQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_RemoveAllScheduledTasksForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).RemoveAllScheduledTasksForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.cyto.scheduler.Scheduler/RemoveAllScheduledTasksForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).RemoveAllScheduledTasksForUser(ctx, req.(*UserQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.cyto.scheduler.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetScheduledTask",
			Handler:    _Scheduler_SetScheduledTask_Handler,
		},
		{
			MethodName: "GetScheduledTask",
			Handler:    _Scheduler_GetScheduledTask_Handler,
		},
		{
			MethodName: "GetScheduledTasksForGuild",
			Handler:    _Scheduler_GetScheduledTasksForGuild_Handler,
		},
		{
			MethodName: "GetScheduledTasksForUser",
			Handler:    _Scheduler_GetScheduledTasksForUser_Handler,
		},
		{
			MethodName: "RemoveScheduledTask",
			Handler:    _Scheduler_RemoveScheduledTask_Handler,
		},
		{
			MethodName: "RemoveAllScheduledTasksForGuild",
			Handler:    _Scheduler_RemoveAllScheduledTasksForGuild_Handler,
		},
		{
			MethodName: "RemoveAllScheduledTasksForUser",
			Handler:    _Scheduler_RemoveAllScheduledTasksForUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}
