// Code generated by protoc-gen-go. DO NOT EDIT.
// source: funnel.proto

/*
Package funnel is a generated protocol buffer package.

It is generated from these files:
	funnel.proto

It has these top-level messages:
	Resources
	Worker
	TaskWrapper
	UpdateExecutorLogsRequest
	UpdateExecutorLogsResponse
	ListWorkersRequest
	ListWorkersResponse
	UpdateWorkerResponse
	GetWorkerRequest
	UpdateTaskLogsRequest
	UpdateTaskLogsResponse
*/
package funnel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tes "github.com/ohsu-comp-bio/funnel/proto/tes"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type WorkerState int32

const (
	WorkerState_UNINITIALIZED WorkerState = 0
	WorkerState_ALIVE         WorkerState = 1
	WorkerState_DEAD          WorkerState = 2
	WorkerState_GONE          WorkerState = 3
	WorkerState_INITIALIZING  WorkerState = 4
)

var WorkerState_name = map[int32]string{
	0: "UNINITIALIZED",
	1: "ALIVE",
	2: "DEAD",
	3: "GONE",
	4: "INITIALIZING",
}
var WorkerState_value = map[string]int32{
	"UNINITIALIZED": 0,
	"ALIVE":         1,
	"DEAD":          2,
	"GONE":          3,
	"INITIALIZING":  4,
}

func (x WorkerState) String() string {
	return proto.EnumName(WorkerState_name, int32(x))
}
func (WorkerState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Resources struct {
	Cpus uint32 `protobuf:"varint,1,opt,name=cpus" json:"cpus,omitempty"`
	// In GB
	RamGb float64 `protobuf:"fixed64,2,opt,name=ram_gb,json=ramGb" json:"ram_gb,omitempty"`
	// In GB
	DiskGb float64 `protobuf:"fixed64,3,opt,name=disk_gb,json=diskGb" json:"disk_gb,omitempty"`
}

func (m *Resources) Reset()                    { *m = Resources{} }
func (m *Resources) String() string            { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()               {}
func (*Resources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Resources) GetCpus() uint32 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

func (m *Resources) GetRamGb() float64 {
	if m != nil {
		return m.RamGb
	}
	return 0
}

func (m *Resources) GetDiskGb() float64 {
	if m != nil {
		return m.DiskGb
	}
	return 0
}

type Worker struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Task ID -> TaskWrapper
	Tasks       map[string]*TaskWrapper `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Resources   *Resources              `protobuf:"bytes,5,opt,name=resources" json:"resources,omitempty"`
	Available   *Resources              `protobuf:"bytes,6,opt,name=available" json:"available,omitempty"`
	LastPing    int64                   `protobuf:"varint,7,opt,name=last_ping,json=lastPing" json:"last_ping,omitempty"`
	State       WorkerState             `protobuf:"varint,8,opt,name=state,enum=funnel.WorkerState" json:"state,omitempty"`
	Preemptible bool                    `protobuf:"varint,9,opt,name=preemptible" json:"preemptible,omitempty"`
	// TODO where does this get updated?
	ActivePorts []int32 `protobuf:"varint,10,rep,packed,name=active_ports,json=activePorts" json:"active_ports,omitempty"`
	Zone        string  `protobuf:"bytes,11,opt,name=zone" json:"zone,omitempty"`
	// Hostname of the worker host.
	// TODO
	Hostname string `protobuf:"bytes,13,opt,name=hostname" json:"hostname,omitempty"`
	// Version of the record in the database. Used to prevent write conflicts.
	Version  int64             `protobuf:"varint,14,opt,name=version" json:"version,omitempty"`
	Metadata map[string]string `protobuf:"bytes,15,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Worker) Reset()                    { *m = Worker{} }
func (m *Worker) String() string            { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()               {}
func (*Worker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Worker) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Worker) GetTasks() map[string]*TaskWrapper {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Worker) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Worker) GetAvailable() *Resources {
	if m != nil {
		return m.Available
	}
	return nil
}

func (m *Worker) GetLastPing() int64 {
	if m != nil {
		return m.LastPing
	}
	return 0
}

func (m *Worker) GetState() WorkerState {
	if m != nil {
		return m.State
	}
	return WorkerState_UNINITIALIZED
}

func (m *Worker) GetPreemptible() bool {
	if m != nil {
		return m.Preemptible
	}
	return false
}

func (m *Worker) GetActivePorts() []int32 {
	if m != nil {
		return m.ActivePorts
	}
	return nil
}

func (m *Worker) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Worker) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Worker) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Worker) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// TODO is there a nice way to avoid this wrapper? Maybe protobuf extensions?
//      or use metadata field of Task?
type TaskWrapper struct {
	Task *tes.Task `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
	Auth string    `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
}

func (m *TaskWrapper) Reset()                    { *m = TaskWrapper{} }
func (m *TaskWrapper) String() string            { return proto.CompactTextString(m) }
func (*TaskWrapper) ProtoMessage()               {}
func (*TaskWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TaskWrapper) GetTask() *tes.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *TaskWrapper) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

type UpdateExecutorLogsRequest struct {
	Id       string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Step     int64            `protobuf:"varint,2,opt,name=step" json:"step,omitempty"`
	Log      *tes.ExecutorLog `protobuf:"bytes,4,opt,name=log" json:"log,omitempty"`
	WorkerId string           `protobuf:"bytes,5,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
}

func (m *UpdateExecutorLogsRequest) Reset()                    { *m = UpdateExecutorLogsRequest{} }
func (m *UpdateExecutorLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateExecutorLogsRequest) ProtoMessage()               {}
func (*UpdateExecutorLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateExecutorLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateExecutorLogsRequest) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *UpdateExecutorLogsRequest) GetLog() *tes.ExecutorLog {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *UpdateExecutorLogsRequest) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

type UpdateExecutorLogsResponse struct {
}

func (m *UpdateExecutorLogsResponse) Reset()                    { *m = UpdateExecutorLogsResponse{} }
func (m *UpdateExecutorLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateExecutorLogsResponse) ProtoMessage()               {}
func (*UpdateExecutorLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListWorkersRequest struct {
}

func (m *ListWorkersRequest) Reset()                    { *m = ListWorkersRequest{} }
func (m *ListWorkersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListWorkersRequest) ProtoMessage()               {}
func (*ListWorkersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListWorkersResponse struct {
	Workers []*Worker `protobuf:"bytes,1,rep,name=workers" json:"workers,omitempty"`
}

func (m *ListWorkersResponse) Reset()                    { *m = ListWorkersResponse{} }
func (m *ListWorkersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListWorkersResponse) ProtoMessage()               {}
func (*ListWorkersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListWorkersResponse) GetWorkers() []*Worker {
	if m != nil {
		return m.Workers
	}
	return nil
}

type UpdateWorkerResponse struct {
}

func (m *UpdateWorkerResponse) Reset()                    { *m = UpdateWorkerResponse{} }
func (m *UpdateWorkerResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateWorkerResponse) ProtoMessage()               {}
func (*UpdateWorkerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type GetWorkerRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetWorkerRequest) Reset()                    { *m = GetWorkerRequest{} }
func (m *GetWorkerRequest) String() string            { return proto.CompactTextString(m) }
func (*GetWorkerRequest) ProtoMessage()               {}
func (*GetWorkerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetWorkerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateTaskLogsRequest struct {
	Id      string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TaskLog *tes.TaskLog `protobuf:"bytes,2,opt,name=task_log,json=taskLog" json:"task_log,omitempty"`
}

func (m *UpdateTaskLogsRequest) Reset()                    { *m = UpdateTaskLogsRequest{} }
func (m *UpdateTaskLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateTaskLogsRequest) ProtoMessage()               {}
func (*UpdateTaskLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateTaskLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTaskLogsRequest) GetTaskLog() *tes.TaskLog {
	if m != nil {
		return m.TaskLog
	}
	return nil
}

type UpdateTaskLogsResponse struct {
}

func (m *UpdateTaskLogsResponse) Reset()                    { *m = UpdateTaskLogsResponse{} }
func (m *UpdateTaskLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateTaskLogsResponse) ProtoMessage()               {}
func (*UpdateTaskLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*Resources)(nil), "funnel.Resources")
	proto.RegisterType((*Worker)(nil), "funnel.Worker")
	proto.RegisterType((*TaskWrapper)(nil), "funnel.TaskWrapper")
	proto.RegisterType((*UpdateExecutorLogsRequest)(nil), "funnel.UpdateExecutorLogsRequest")
	proto.RegisterType((*UpdateExecutorLogsResponse)(nil), "funnel.UpdateExecutorLogsResponse")
	proto.RegisterType((*ListWorkersRequest)(nil), "funnel.ListWorkersRequest")
	proto.RegisterType((*ListWorkersResponse)(nil), "funnel.ListWorkersResponse")
	proto.RegisterType((*UpdateWorkerResponse)(nil), "funnel.UpdateWorkerResponse")
	proto.RegisterType((*GetWorkerRequest)(nil), "funnel.GetWorkerRequest")
	proto.RegisterType((*UpdateTaskLogsRequest)(nil), "funnel.UpdateTaskLogsRequest")
	proto.RegisterType((*UpdateTaskLogsResponse)(nil), "funnel.UpdateTaskLogsResponse")
	proto.RegisterEnum("funnel.WorkerState", WorkerState_name, WorkerState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SchedulerService service

type SchedulerServiceClient interface {
	UpdateExecutorLogs(ctx context.Context, in *UpdateExecutorLogsRequest, opts ...grpc.CallOption) (*UpdateExecutorLogsResponse, error)
	UpdateWorker(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*UpdateWorkerResponse, error)
	UpdateTaskLogs(ctx context.Context, in *UpdateTaskLogsRequest, opts ...grpc.CallOption) (*UpdateTaskLogsResponse, error)
	ListWorkers(ctx context.Context, in *ListWorkersRequest, opts ...grpc.CallOption) (*ListWorkersResponse, error)
	GetWorker(ctx context.Context, in *GetWorkerRequest, opts ...grpc.CallOption) (*Worker, error)
}

type schedulerServiceClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerServiceClient(cc *grpc.ClientConn) SchedulerServiceClient {
	return &schedulerServiceClient{cc}
}

func (c *schedulerServiceClient) UpdateExecutorLogs(ctx context.Context, in *UpdateExecutorLogsRequest, opts ...grpc.CallOption) (*UpdateExecutorLogsResponse, error) {
	out := new(UpdateExecutorLogsResponse)
	err := grpc.Invoke(ctx, "/funnel.SchedulerService/UpdateExecutorLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) UpdateWorker(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*UpdateWorkerResponse, error) {
	out := new(UpdateWorkerResponse)
	err := grpc.Invoke(ctx, "/funnel.SchedulerService/UpdateWorker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) UpdateTaskLogs(ctx context.Context, in *UpdateTaskLogsRequest, opts ...grpc.CallOption) (*UpdateTaskLogsResponse, error) {
	out := new(UpdateTaskLogsResponse)
	err := grpc.Invoke(ctx, "/funnel.SchedulerService/UpdateTaskLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) ListWorkers(ctx context.Context, in *ListWorkersRequest, opts ...grpc.CallOption) (*ListWorkersResponse, error) {
	out := new(ListWorkersResponse)
	err := grpc.Invoke(ctx, "/funnel.SchedulerService/ListWorkers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) GetWorker(ctx context.Context, in *GetWorkerRequest, opts ...grpc.CallOption) (*Worker, error) {
	out := new(Worker)
	err := grpc.Invoke(ctx, "/funnel.SchedulerService/GetWorker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SchedulerService service

type SchedulerServiceServer interface {
	UpdateExecutorLogs(context.Context, *UpdateExecutorLogsRequest) (*UpdateExecutorLogsResponse, error)
	UpdateWorker(context.Context, *Worker) (*UpdateWorkerResponse, error)
	UpdateTaskLogs(context.Context, *UpdateTaskLogsRequest) (*UpdateTaskLogsResponse, error)
	ListWorkers(context.Context, *ListWorkersRequest) (*ListWorkersResponse, error)
	GetWorker(context.Context, *GetWorkerRequest) (*Worker, error)
}

func RegisterSchedulerServiceServer(s *grpc.Server, srv SchedulerServiceServer) {
	s.RegisterService(&_SchedulerService_serviceDesc, srv)
}

func _SchedulerService_UpdateExecutorLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExecutorLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).UpdateExecutorLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funnel.SchedulerService/UpdateExecutorLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).UpdateExecutorLogs(ctx, req.(*UpdateExecutorLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_UpdateWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Worker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).UpdateWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funnel.SchedulerService/UpdateWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).UpdateWorker(ctx, req.(*Worker))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_UpdateTaskLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).UpdateTaskLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funnel.SchedulerService/UpdateTaskLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).UpdateTaskLogs(ctx, req.(*UpdateTaskLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_ListWorkers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ListWorkers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funnel.SchedulerService/ListWorkers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ListWorkers(ctx, req.(*ListWorkersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_GetWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).GetWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funnel.SchedulerService/GetWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).GetWorker(ctx, req.(*GetWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchedulerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "funnel.SchedulerService",
	HandlerType: (*SchedulerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateExecutorLogs",
			Handler:    _SchedulerService_UpdateExecutorLogs_Handler,
		},
		{
			MethodName: "UpdateWorker",
			Handler:    _SchedulerService_UpdateWorker_Handler,
		},
		{
			MethodName: "UpdateTaskLogs",
			Handler:    _SchedulerService_UpdateTaskLogs_Handler,
		},
		{
			MethodName: "ListWorkers",
			Handler:    _SchedulerService_ListWorkers_Handler,
		},
		{
			MethodName: "GetWorker",
			Handler:    _SchedulerService_GetWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "funnel.proto",
}

func init() { proto.RegisterFile("funnel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 835 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x45, 0xfd, 0x71, 0x28, 0xab, 0xcc, 0xc4, 0x71, 0x37, 0x8c, 0x13, 0x30, 0xbc, 0x94,
	0xed, 0xc1, 0x42, 0xd5, 0x4b, 0xd0, 0x02, 0x6d, 0x0d, 0x58, 0x10, 0x04, 0x38, 0x8e, 0xb3, 0x8e,
	0x1b, 0xa0, 0x40, 0x21, 0xac, 0xc4, 0xad, 0x4c, 0x58, 0x22, 0x59, 0xee, 0x4a, 0x6d, 0x5a, 0x14,
	0x28, 0xfa, 0x0a, 0x7d, 0xa0, 0x5e, 0x7b, 0xef, 0x2b, 0xf4, 0x41, 0x8a, 0xdd, 0x25, 0x65, 0x49,
	0xb6, 0x73, 0x9b, 0xfd, 0x66, 0xe6, 0x9b, 0xc1, 0x37, 0x1f, 0x41, 0xe8, 0xfc, 0xb8, 0x4c, 0x53,
	0x3e, 0x3f, 0xca, 0x8b, 0x4c, 0x66, 0xd8, 0x34, 0x2f, 0xdf, 0x91, 0x5c, 0x18, 0xc8, 0x3f, 0x9c,
	0x65, 0xd9, 0x6c, 0xce, 0x7b, 0x2c, 0x4f, 0x7a, 0x2c, 0x4d, 0x33, 0xc9, 0x64, 0x92, 0xa5, 0x65,
	0x36, 0x7c, 0x0d, 0x0e, 0xe5, 0x22, 0x5b, 0x16, 0x53, 0x2e, 0x10, 0xa1, 0x3e, 0xcd, 0x97, 0x82,
	0x58, 0x81, 0x15, 0xed, 0x51, 0x1d, 0xe3, 0x63, 0x68, 0x16, 0x6c, 0x31, 0x9e, 0x4d, 0x48, 0x2d,
	0xb0, 0x22, 0x8b, 0x36, 0x0a, 0xb6, 0x18, 0x4e, 0xf0, 0x63, 0x68, 0xc5, 0x89, 0xb8, 0x56, 0xb8,
	0xad, 0xf1, 0xa6, 0x7a, 0x0e, 0x27, 0xe1, 0x3f, 0x75, 0x68, 0xbe, 0xcb, 0x8a, 0x6b, 0x5e, 0x60,
	0x17, 0x6a, 0x49, 0xac, 0xc9, 0x1c, 0x5a, 0x4b, 0x62, 0xec, 0x41, 0x43, 0x32, 0x71, 0x2d, 0x48,
	0x2d, 0xb0, 0x23, 0xb7, 0xff, 0xe4, 0xa8, 0x5c, 0xdd, 0x94, 0x1f, 0xbd, 0x55, 0xb9, 0x41, 0x2a,
	0x8b, 0xf7, 0xd4, 0xd4, 0x61, 0x0f, 0x9c, 0xa2, 0x5a, 0x8e, 0x34, 0x02, 0x2b, 0x72, 0xfb, 0x0f,
	0xab, 0xa6, 0xf5, 0xd6, 0xf4, 0xa6, 0x46, 0x35, 0xb0, 0x15, 0x4b, 0xe6, 0x6c, 0x32, 0xe7, 0xa4,
	0x79, 0x6f, 0xc3, 0xba, 0x06, 0x9f, 0x82, 0x33, 0x67, 0x42, 0x8e, 0xf3, 0x24, 0x9d, 0x91, 0x56,
	0x60, 0x45, 0x36, 0x6d, 0x2b, 0xe0, 0x3c, 0x49, 0x67, 0xf8, 0x29, 0x34, 0x84, 0x64, 0x92, 0x93,
	0x76, 0x60, 0x45, 0xdd, 0xfe, 0xa3, 0xed, 0x7d, 0x2f, 0x54, 0x8a, 0x9a, 0x0a, 0x0c, 0xc0, 0xcd,
	0x0b, 0xce, 0x17, 0xb9, 0x4c, 0xd4, 0x68, 0x27, 0xb0, 0xa2, 0x36, 0xdd, 0x84, 0xf0, 0x05, 0x74,
	0xd8, 0x54, 0x26, 0x2b, 0x3e, 0xce, 0xb3, 0x42, 0x0a, 0x02, 0x81, 0x1d, 0x35, 0xa8, 0x6b, 0xb0,
	0x73, 0x05, 0x29, 0xf9, 0x7f, 0xcd, 0x52, 0x4e, 0x5c, 0xad, 0x98, 0x8e, 0xd1, 0x87, 0xf6, 0x55,
	0x26, 0x64, 0xca, 0x16, 0x9c, 0xec, 0x69, 0x7c, 0xfd, 0x46, 0x02, 0xad, 0x15, 0x2f, 0x44, 0x92,
	0xa5, 0xa4, 0xab, 0x57, 0xaf, 0x9e, 0xf8, 0x12, 0xda, 0x0b, 0x2e, 0x59, 0xcc, 0x24, 0x23, 0x1f,
	0x69, 0xb1, 0x0f, 0x77, 0xc4, 0x7e, 0x55, 0xa6, 0x8d, 0xde, 0xeb, 0x6a, 0xff, 0x15, 0xc0, 0xcd,
	0x1d, 0xd0, 0x03, 0xfb, 0x9a, 0xbf, 0x2f, 0x4f, 0xa8, 0x42, 0xa5, 0xc9, 0x8a, 0xcd, 0x97, 0x5c,
	0xbb, 0xc1, 0xbd, 0xd1, 0x44, 0x35, 0xbd, 0x2b, 0x58, 0x9e, 0xf3, 0x82, 0x9a, 0x8a, 0x2f, 0x6b,
	0x2f, 0x2d, 0xff, 0x2b, 0xd8, 0xdb, 0x9a, 0x74, 0x07, 0xe3, 0xfe, 0x26, 0xa3, 0xb3, 0xd1, 0x1c,
	0x7e, 0x0b, 0xee, 0x06, 0x2d, 0x3e, 0x83, 0xba, 0xb2, 0x85, 0xee, 0x75, 0xfb, 0xce, 0x91, 0xb2,
	0xb8, 0xca, 0x53, 0x0d, 0x2b, 0xf5, 0xd8, 0x52, 0x5e, 0x95, 0x34, 0x3a, 0x0e, 0xff, 0xb0, 0xe0,
	0xc9, 0x65, 0x1e, 0x33, 0xc9, 0x07, 0xbf, 0xf0, 0xe9, 0x52, 0x66, 0xc5, 0x69, 0x36, 0x13, 0x94,
	0xff, 0xb4, 0xe4, 0x42, 0xde, 0xf2, 0x27, 0x42, 0x5d, 0x48, 0x9e, 0x6b, 0x06, 0x9b, 0xea, 0x18,
	0x43, 0xb0, 0xe7, 0xd9, 0x8c, 0xd4, 0xf5, 0x4c, 0x4f, 0xcf, 0xdc, 0xa0, 0xa2, 0x2a, 0xa9, 0x4c,
	0xf4, 0xb3, 0x56, 0x75, 0x9c, 0xc4, 0xda, 0xa6, 0x0e, 0x6d, 0x1b, 0x60, 0x14, 0x87, 0x87, 0xe0,
	0xdf, 0xb5, 0x81, 0xc8, 0xb3, 0x54, 0xf0, 0x70, 0x1f, 0xf0, 0x34, 0x11, 0xd2, 0x1c, 0xa5, 0x5a,
	0x2c, 0xfc, 0x06, 0x1e, 0x6d, 0xa1, 0xa6, 0x18, 0x23, 0x68, 0x19, 0x5a, 0xf5, 0x85, 0xaa, 0xa3,
	0x76, 0xb7, 0x8f, 0x4a, 0xab, 0x74, 0x78, 0x00, 0xfb, 0x66, 0x68, 0x99, 0xa8, 0xc6, 0x85, 0xe0,
	0x0d, 0xb9, 0xac, 0xc0, 0x3b, 0x55, 0x08, 0xcf, 0xe1, 0xb1, 0xe9, 0x55, 0xda, 0x7e, 0x48, 0xae,
	0x4f, 0xa0, 0xad, 0x84, 0x1f, 0x2b, 0x7d, 0x8c, 0x1b, 0x3a, 0xeb, 0x9b, 0x28, 0x6d, 0x5a, 0xd2,
	0x04, 0x21, 0x81, 0x83, 0x5d, 0x46, 0xb3, 0xcf, 0x67, 0x6f, 0xc0, 0xdd, 0xf8, 0x98, 0xf0, 0x21,
	0xec, 0x5d, 0x9e, 0x8d, 0xce, 0x46, 0x6f, 0x47, 0xc7, 0xa7, 0xa3, 0xef, 0x07, 0x27, 0xde, 0x03,
	0x74, 0xa0, 0x71, 0x7c, 0x3a, 0xfa, 0x6e, 0xe0, 0x59, 0xd8, 0x86, 0xfa, 0xc9, 0xe0, 0xf8, 0xc4,
	0xab, 0xa9, 0x68, 0xf8, 0xfa, 0x6c, 0xe0, 0xd9, 0xe8, 0x41, 0x67, 0x5d, 0x3f, 0x3a, 0x1b, 0x7a,
	0xf5, 0xfe, 0xdf, 0x36, 0x78, 0x17, 0xd3, 0x2b, 0x1e, 0x2f, 0xe7, 0xbc, 0xb8, 0xe0, 0xc5, 0x2a,
	0x99, 0x72, 0xfc, 0x01, 0xf0, 0xf6, 0x11, 0xf0, 0x45, 0x25, 0xdf, 0xbd, 0x16, 0xf1, 0xc3, 0x0f,
	0x95, 0x94, 0xa2, 0x3e, 0xc0, 0xaf, 0xa1, 0xb3, 0x29, 0x37, 0xee, 0xdc, 0xc5, 0x3f, 0xdc, 0x66,
	0xd9, 0x39, 0xca, 0x03, 0x7c, 0x03, 0xdd, 0x6d, 0x81, 0xf0, 0xd9, 0x76, 0xc7, 0xce, 0x29, 0xfc,
	0xe7, 0xf7, 0xa5, 0xd7, 0x94, 0x13, 0x70, 0x37, 0x2c, 0x84, 0x7e, 0xd5, 0x70, 0xdb, 0x6d, 0xfe,
	0xd3, 0x3b, 0x73, 0x25, 0x93, 0xff, 0xe7, 0xbf, 0xff, 0xfd, 0x55, 0xdb, 0x47, 0xec, 0xad, 0x3e,
	0xef, 0x99, 0xba, 0x5e, 0xe9, 0x32, 0xbc, 0x04, 0x67, 0xed, 0x26, 0x24, 0x15, 0xcb, 0xae, 0xc1,
	0xfc, 0x1d, 0x35, 0xc2, 0xe7, 0x9a, 0x92, 0xe0, 0xc1, 0x2d, 0xca, 0xde, 0x6f, 0x49, 0xfc, 0xfb,
	0xa4, 0xa9, 0xff, 0x4c, 0x5f, 0xfc, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x79, 0xc5, 0x4e, 0x73, 0xda,
	0x06, 0x00, 0x00,
}