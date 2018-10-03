// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2beta2/task.proto

package tasks

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The view specifies a subset of [Task][google.cloud.tasks.v2beta2.Task] data.
//
// When a task is returned in a response, not all
// information is retrieved by default because some data, such as
// payloads, might be desirable to return only when needed because
// of its large size or because of the sensitivity of data that it
// contains.
type Task_View int32

const (
	// Unspecified. Defaults to BASIC.
	Task_VIEW_UNSPECIFIED Task_View = 0
	// The basic view omits fields which can be large or can contain
	// sensitive data.
	//
	// This view does not include the
	// ([payload in AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest] and
	// [payload in PullMessage][google.cloud.tasks.v2beta2.PullMessage.payload]). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	Task_BASIC Task_View = 1
	// All information is returned.
	//
	// Authorization for [FULL][google.cloud.tasks.v2beta2.Task.View.FULL] requires
	// `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/)
	// permission on the [Queue][google.cloud.tasks.v2beta2.Queue] resource.
	Task_FULL Task_View = 2
)

var Task_View_name = map[int32]string{
	0: "VIEW_UNSPECIFIED",
	1: "BASIC",
	2: "FULL",
}

var Task_View_value = map[string]int32{
	"VIEW_UNSPECIFIED": 0,
	"BASIC":            1,
	"FULL":             2,
}

func (x Task_View) String() string {
	return proto.EnumName(Task_View_name, int32(x))
}

func (Task_View) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3fffa1a9946502fd, []int{0, 0}
}

// A unit of scheduled work.
type Task struct {
	// Optionally caller-specified in [CreateTask][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	//
	// The task name.
	//
	// The task name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the task's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	// * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//   hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	//
	// The task's payload is used by the task's target to process the task.
	// A payload is valid only if it is compatible with the queue's target.
	//
	// Types that are valid to be assigned to PayloadType:
	//	*Task_AppEngineHttpRequest
	//	*Task_PullMessage
	PayloadType isTask_PayloadType `protobuf_oneof:"payload_type"`
	// The time when the task is scheduled to be attempted.
	//
	// For App Engine queues, this is when the task will be attempted or retried.
	//
	// For pull queues, this is the time when the task is available to
	// be leased; if a task is currently leased, this is the time when
	// the current lease expires, that is, the time that the task was
	// leased plus the [lease_duration][google.cloud.tasks.v2beta2.LeaseTasksRequest.lease_duration].
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that the task was created.
	//
	// `create_time` will be truncated to the nearest second.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The task status.
	Status *TaskStatus `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	// Output only. The view specifies which subset of the [Task][google.cloud.tasks.v2beta2.Task] has
	// been returned.
	View                 Task_View `protobuf:"varint,8,opt,name=view,proto3,enum=google.cloud.tasks.v2beta2.Task_View" json:"view,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fffa1a9946502fd, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTask_PayloadType interface {
	isTask_PayloadType()
}

type Task_AppEngineHttpRequest struct {
	AppEngineHttpRequest *AppEngineHttpRequest `protobuf:"bytes,3,opt,name=app_engine_http_request,json=appEngineHttpRequest,proto3,oneof"`
}

type Task_PullMessage struct {
	PullMessage *PullMessage `protobuf:"bytes,4,opt,name=pull_message,json=pullMessage,proto3,oneof"`
}

func (*Task_AppEngineHttpRequest) isTask_PayloadType() {}

func (*Task_PullMessage) isTask_PayloadType() {}

func (m *Task) GetPayloadType() isTask_PayloadType {
	if m != nil {
		return m.PayloadType
	}
	return nil
}

func (m *Task) GetAppEngineHttpRequest() *AppEngineHttpRequest {
	if x, ok := m.GetPayloadType().(*Task_AppEngineHttpRequest); ok {
		return x.AppEngineHttpRequest
	}
	return nil
}

func (m *Task) GetPullMessage() *PullMessage {
	if x, ok := m.GetPayloadType().(*Task_PullMessage); ok {
		return x.PullMessage
	}
	return nil
}

func (m *Task) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *Task) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatus() *TaskStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetView() Task_View {
	if m != nil {
		return m.View
	}
	return Task_VIEW_UNSPECIFIED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Task) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Task_OneofMarshaler, _Task_OneofUnmarshaler, _Task_OneofSizer, []interface{}{
		(*Task_AppEngineHttpRequest)(nil),
		(*Task_PullMessage)(nil),
	}
}

func _Task_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Task)
	// payload_type
	switch x := m.PayloadType.(type) {
	case *Task_AppEngineHttpRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppEngineHttpRequest); err != nil {
			return err
		}
	case *Task_PullMessage:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PullMessage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Task.PayloadType has unexpected type %T", x)
	}
	return nil
}

func _Task_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Task)
	switch tag {
	case 3: // payload_type.app_engine_http_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AppEngineHttpRequest)
		err := b.DecodeMessage(msg)
		m.PayloadType = &Task_AppEngineHttpRequest{msg}
		return true, err
	case 4: // payload_type.pull_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PullMessage)
		err := b.DecodeMessage(msg)
		m.PayloadType = &Task_PullMessage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Task_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Task)
	// payload_type
	switch x := m.PayloadType.(type) {
	case *Task_AppEngineHttpRequest:
		s := proto.Size(x.AppEngineHttpRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Task_PullMessage:
		s := proto.Size(x.PullMessage)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Status of the task.
type TaskStatus struct {
	// Output only. The number of attempts dispatched.
	//
	// This count includes tasks which have been dispatched but haven't
	// received a response.
	AttemptDispatchCount int32 `protobuf:"varint,1,opt,name=attempt_dispatch_count,json=attemptDispatchCount,proto3" json:"attempt_dispatch_count,omitempty"`
	// Output only. The number of attempts which have received a response.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	AttemptResponseCount int32 `protobuf:"varint,2,opt,name=attempt_response_count,json=attemptResponseCount,proto3" json:"attempt_response_count,omitempty"`
	// Output only. The status of the task's first attempt.
	//
	// Only [dispatch_time][google.cloud.tasks.v2beta2.AttemptStatus.dispatch_time] will be set.
	// The other [AttemptStatus][google.cloud.tasks.v2beta2.AttemptStatus] information is not retained by Cloud Tasks.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	FirstAttemptStatus *AttemptStatus `protobuf:"bytes,3,opt,name=first_attempt_status,json=firstAttemptStatus,proto3" json:"first_attempt_status,omitempty"`
	// Output only. The status of the task's last attempt.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	LastAttemptStatus    *AttemptStatus `protobuf:"bytes,4,opt,name=last_attempt_status,json=lastAttemptStatus,proto3" json:"last_attempt_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskStatus) Reset()         { *m = TaskStatus{} }
func (m *TaskStatus) String() string { return proto.CompactTextString(m) }
func (*TaskStatus) ProtoMessage()    {}
func (*TaskStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fffa1a9946502fd, []int{1}
}

func (m *TaskStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskStatus.Unmarshal(m, b)
}
func (m *TaskStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskStatus.Marshal(b, m, deterministic)
}
func (m *TaskStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskStatus.Merge(m, src)
}
func (m *TaskStatus) XXX_Size() int {
	return xxx_messageInfo_TaskStatus.Size(m)
}
func (m *TaskStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TaskStatus proto.InternalMessageInfo

func (m *TaskStatus) GetAttemptDispatchCount() int32 {
	if m != nil {
		return m.AttemptDispatchCount
	}
	return 0
}

func (m *TaskStatus) GetAttemptResponseCount() int32 {
	if m != nil {
		return m.AttemptResponseCount
	}
	return 0
}

func (m *TaskStatus) GetFirstAttemptStatus() *AttemptStatus {
	if m != nil {
		return m.FirstAttemptStatus
	}
	return nil
}

func (m *TaskStatus) GetLastAttemptStatus() *AttemptStatus {
	if m != nil {
		return m.LastAttemptStatus
	}
	return nil
}

// The status of a task attempt.
type AttemptStatus struct {
	// Output only. The time that this attempt was scheduled.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that this attempt was dispatched.
	//
	// `dispatch_time` will be truncated to the nearest microsecond.
	DispatchTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=dispatch_time,json=dispatchTime,proto3" json:"dispatch_time,omitempty"`
	// Output only. The time that this attempt response was received.
	//
	// `response_time` will be truncated to the nearest microsecond.
	ResponseTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	// Output only. The response from the target for this attempt.
	//
	// If the task has not been attempted or the task is currently running
	// then the response status is unset.
	ResponseStatus       *status.Status `protobuf:"bytes,4,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AttemptStatus) Reset()         { *m = AttemptStatus{} }
func (m *AttemptStatus) String() string { return proto.CompactTextString(m) }
func (*AttemptStatus) ProtoMessage()    {}
func (*AttemptStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fffa1a9946502fd, []int{2}
}

func (m *AttemptStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttemptStatus.Unmarshal(m, b)
}
func (m *AttemptStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttemptStatus.Marshal(b, m, deterministic)
}
func (m *AttemptStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttemptStatus.Merge(m, src)
}
func (m *AttemptStatus) XXX_Size() int {
	return xxx_messageInfo_AttemptStatus.Size(m)
}
func (m *AttemptStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AttemptStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AttemptStatus proto.InternalMessageInfo

func (m *AttemptStatus) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *AttemptStatus) GetDispatchTime() *timestamp.Timestamp {
	if m != nil {
		return m.DispatchTime
	}
	return nil
}

func (m *AttemptStatus) GetResponseTime() *timestamp.Timestamp {
	if m != nil {
		return m.ResponseTime
	}
	return nil
}

func (m *AttemptStatus) GetResponseStatus() *status.Status {
	if m != nil {
		return m.ResponseStatus
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.tasks.v2beta2.Task_View", Task_View_name, Task_View_value)
	proto.RegisterType((*Task)(nil), "google.cloud.tasks.v2beta2.Task")
	proto.RegisterType((*TaskStatus)(nil), "google.cloud.tasks.v2beta2.TaskStatus")
	proto.RegisterType((*AttemptStatus)(nil), "google.cloud.tasks.v2beta2.AttemptStatus")
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2beta2/task.proto", fileDescriptor_3fffa1a9946502fd)
}

var fileDescriptor_3fffa1a9946502fd = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x6f, 0xda, 0x3c,
	0x14, 0xc6, 0x1b, 0x9a, 0xf6, 0x2d, 0xe6, 0xcf, 0xcb, 0x3c, 0xb4, 0x46, 0x68, 0xda, 0x10, 0x52,
	0x57, 0x76, 0x93, 0x6c, 0x6c, 0x37, 0x53, 0xa5, 0xa1, 0x42, 0xa9, 0x40, 0x62, 0x13, 0x0a, 0x6d,
	0xa7, 0x6d, 0x17, 0x91, 0x09, 0x6e, 0x88, 0x9a, 0xc4, 0x5e, 0x7c, 0xd2, 0xaa, 0x9f, 0xa2, 0x5f,
	0x65, 0x1f, 0x71, 0x8a, 0xe3, 0xd0, 0xa1, 0x76, 0xa0, 0xde, 0x71, 0xce, 0x79, 0x7e, 0x8f, 0xed,
	0x47, 0x87, 0xa0, 0x03, 0x8f, 0x31, 0x2f, 0xa0, 0x96, 0x1b, 0xb0, 0x64, 0x6e, 0x01, 0x11, 0x57,
	0xc2, 0xba, 0xee, 0xcc, 0x28, 0x90, 0x8e, 0xac, 0x4c, 0x1e, 0x33, 0x60, 0xb8, 0x91, 0xc9, 0x4c,
	0x29, 0x33, 0xa5, 0xcc, 0x54, 0xb2, 0xc6, 0x4b, 0x65, 0x41, 0xb8, 0x6f, 0x91, 0x28, 0x62, 0x40,
	0xc0, 0x67, 0x91, 0xc8, 0xc8, 0xc6, 0xe1, 0xda, 0x03, 0x62, 0x8f, 0x82, 0x12, 0xbe, 0x56, 0x42,
	0x59, 0xcd, 0x92, 0x4b, 0x0b, 0xfc, 0x90, 0x0a, 0x20, 0x21, 0x57, 0x82, 0x7d, 0x25, 0x88, 0xb9,
	0x6b, 0x09, 0x20, 0x90, 0xa8, 0x23, 0x5a, 0x77, 0x3a, 0xd2, 0xcf, 0x88, 0xb8, 0xc2, 0x18, 0xe9,
	0x11, 0x09, 0xa9, 0xa1, 0x35, 0xb5, 0x76, 0xd1, 0x96, 0xbf, 0xb1, 0x8f, 0xf6, 0x09, 0xe7, 0x0e,
	0x8d, 0x3c, 0x3f, 0xa2, 0xce, 0x02, 0x80, 0x3b, 0x31, 0xfd, 0x95, 0x50, 0x01, 0xc6, 0x76, 0x53,
	0x6b, 0x97, 0x3a, 0xef, 0xcc, 0x7f, 0xbf, 0xcd, 0x3c, 0xe6, 0x7c, 0x20, 0xc9, 0x21, 0x00, 0xb7,
	0x33, 0x6e, 0xb8, 0x65, 0xd7, 0xc9, 0x23, 0x7d, 0x3c, 0x46, 0x65, 0x9e, 0x04, 0x81, 0x13, 0x52,
	0x21, 0x88, 0x47, 0x0d, 0x5d, 0xfa, 0x1f, 0xae, 0xf3, 0x9f, 0x24, 0x41, 0xf0, 0x25, 0x93, 0x0f,
	0xb7, 0xec, 0x12, 0xbf, 0x2f, 0x71, 0x17, 0x55, 0x84, 0xbb, 0xa0, 0xf3, 0x24, 0xa0, 0x4e, 0x1a,
	0x85, 0xb1, 0x23, 0xed, 0x1a, 0xb9, 0x5d, 0x9e, 0x93, 0x79, 0x96, 0xe7, 0x64, 0x97, 0x73, 0x20,
	0x6d, 0xe1, 0x23, 0x54, 0x72, 0x63, 0x4a, 0x40, 0xe1, 0xbb, 0x1b, 0x71, 0x94, 0xc9, 0x25, 0xfc,
	0x19, 0xed, 0x66, 0x19, 0x1b, 0xff, 0x49, 0xee, 0xcd, 0xba, 0x57, 0xa4, 0xe1, 0x4f, 0xa5, 0xda,
	0x56, 0x14, 0xfe, 0x84, 0xf4, 0x6b, 0x9f, 0xde, 0x18, 0x7b, 0x4d, 0xad, 0x5d, 0xed, 0x1c, 0x6c,
	0xa2, 0xcd, 0x0b, 0x9f, 0xde, 0xd8, 0x12, 0x69, 0xbd, 0x47, 0x7a, 0x5a, 0xe1, 0x3a, 0xaa, 0x5d,
	0x8c, 0x06, 0xdf, 0x9c, 0xf3, 0xaf, 0xd3, 0xc9, 0xa0, 0x3f, 0x3a, 0x1d, 0x0d, 0x4e, 0x6a, 0x5b,
	0xb8, 0x88, 0x76, 0x7a, 0xc7, 0xd3, 0x51, 0xbf, 0xa6, 0xe1, 0x3d, 0xa4, 0x9f, 0x9e, 0x8f, 0xc7,
	0xb5, 0x42, 0xaf, 0x8a, 0xca, 0x9c, 0xdc, 0x06, 0x8c, 0xcc, 0x1d, 0xb8, 0xe5, 0xb4, 0xf5, 0xbb,
	0x80, 0xd0, 0xfd, 0xa5, 0xf0, 0x47, 0xf4, 0x82, 0x00, 0xd0, 0x90, 0x83, 0x33, 0xf7, 0x05, 0x27,
	0xe0, 0x2e, 0x1c, 0x97, 0x25, 0x11, 0xc8, 0x4d, 0xd9, 0xb1, 0xeb, 0x6a, 0x7a, 0xa2, 0x86, 0xfd,
	0x74, 0xf6, 0x37, 0x15, 0x53, 0xc1, 0x59, 0x24, 0xa8, 0xa2, 0x0a, 0x2b, 0x94, 0xad, 0x86, 0x19,
	0xf5, 0x13, 0xd5, 0x2f, 0xfd, 0x58, 0x80, 0x93, 0xb3, 0x2a, 0xc6, 0x6c, 0xd9, 0xde, 0xae, 0x5d,
	0xb6, 0x8c, 0x50, 0x49, 0x62, 0x69, 0xb3, 0xd2, 0xc3, 0xdf, 0xd1, 0xf3, 0x80, 0x3c, 0xf4, 0xd6,
	0x9f, 0xea, 0xfd, 0x2c, 0x75, 0x59, 0x69, 0xb5, 0xee, 0x0a, 0xa8, 0xb2, 0x7a, 0xd8, 0x83, 0x05,
	0xd4, 0x9e, 0xb8, 0x80, 0x5d, 0x54, 0x59, 0xc6, 0x2d, 0x0d, 0x0a, 0x9b, 0x0d, 0x72, 0x20, 0x37,
	0x58, 0x26, 0x2f, 0x0d, 0xb6, 0x37, 0x1b, 0xe4, 0x80, 0xfa, 0x0b, 0xfc, 0xbf, 0x34, 0x58, 0xc9,
	0x0a, 0xe7, 0x16, 0x31, 0x77, 0x4d, 0x15, 0x4a, 0x35, 0x97, 0x66, 0x75, 0x2f, 0x42, 0xaf, 0x5c,
	0x16, 0xae, 0x09, 0xb5, 0x57, 0x4c, 0x77, 0x6c, 0x92, 0x5e, 0x62, 0xa2, 0xfd, 0xe8, 0x2a, 0xa1,
	0xc7, 0x02, 0x12, 0x79, 0x26, 0x8b, 0x3d, 0xcb, 0xa3, 0x91, 0xbc, 0xa2, 0x95, 0x8d, 0x08, 0xf7,
	0xc5, 0x63, 0xdf, 0xc1, 0x23, 0x59, 0xcd, 0x76, 0xa5, 0xf6, 0xc3, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x67, 0x07, 0xb1, 0x59, 0x93, 0x05, 0x00, 0x00,
}
