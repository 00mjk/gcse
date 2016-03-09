// Code generated by protoc-gen-go.
// source: github.com/daviddengcn/gcse/proto/spider/spider.proto
// DO NOT EDIT!

/*
Package gcse_spider is a generated protocol buffer package.

It is generated from these files:
	github.com/daviddengcn/gcse/proto/spider/spider.proto

It has these top-level messages:
	GoFileInfo
	RepoInfo
	FolderInfo
	CrawlingInfo
	HistoryEvent
	HistoryInfo
*/
package gcse_spider

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type GoFileInfo_Status int32

const (
	GoFileInfo_Unknown      GoFileInfo_Status = 0
	GoFileInfo_ParseSuccess GoFileInfo_Status = 1
	GoFileInfo_ParseFailed  GoFileInfo_Status = 2
	GoFileInfo_ShouldIgnore GoFileInfo_Status = 3
)

var GoFileInfo_Status_name = map[int32]string{
	0: "Unknown",
	1: "ParseSuccess",
	2: "ParseFailed",
	3: "ShouldIgnore",
}
var GoFileInfo_Status_value = map[string]int32{
	"Unknown":      0,
	"ParseSuccess": 1,
	"ParseFailed":  2,
	"ShouldIgnore": 3,
}

func (x GoFileInfo_Status) String() string {
	return proto.EnumName(GoFileInfo_Status_name, int32(x))
}
func (GoFileInfo_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type HistoryEvent_Action_Enum int32

const (
	HistoryEvent_Action_None    HistoryEvent_Action_Enum = 0
	HistoryEvent_Action_Success HistoryEvent_Action_Enum = 1
	HistoryEvent_Action_Failed  HistoryEvent_Action_Enum = 2
	HistoryEvent_Action_Invalid HistoryEvent_Action_Enum = 3
)

var HistoryEvent_Action_Enum_name = map[int32]string{
	0: "None",
	1: "Success",
	2: "Failed",
	3: "Invalid",
}
var HistoryEvent_Action_Enum_value = map[string]int32{
	"None":    0,
	"Success": 1,
	"Failed":  2,
	"Invalid": 3,
}

func (x HistoryEvent_Action_Enum) String() string {
	return proto.EnumName(HistoryEvent_Action_Enum_name, int32(x))
}
func (HistoryEvent_Action_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0, 0}
}

type GoFileInfo struct {
	Status      GoFileInfo_Status `protobuf:"varint,1,opt,name=status,enum=gcse.spider.GoFileInfo_Status" json:"status,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	IsTest      bool              `protobuf:"varint,4,opt,name=is_test,json=isTest" json:"is_test,omitempty"`
	Imports     []string          `protobuf:"bytes,5,rep,name=imports" json:"imports,omitempty"`
}

func (m *GoFileInfo) Reset()                    { *m = GoFileInfo{} }
func (m *GoFileInfo) String() string            { return proto.CompactTextString(m) }
func (*GoFileInfo) ProtoMessage()               {}
func (*GoFileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RepoInfo struct {
	LastCrawled *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=last_crawled,json=lastCrawled" json:"last_crawled,omitempty"`
	Stars       int32                      `protobuf:"varint,2,opt,name=stars" json:"stars,omitempty"`
	Description string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Where this project was forked from, full path
	Source      string                     `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	LastUpdated *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *RepoInfo) Reset()                    { *m = RepoInfo{} }
func (m *RepoInfo) String() string            { return proto.CompactTextString(m) }
func (*RepoInfo) ProtoMessage()               {}
func (*RepoInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RepoInfo) GetLastCrawled() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastCrawled
	}
	return nil
}

func (m *RepoInfo) GetLastUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Information for non-repository folder.
type FolderInfo struct {
	// E.g. "sub"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// E.g. "spider/sub"
	Path    string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Sha     string `protobuf:"bytes,3,opt,name=sha" json:"sha,omitempty"`
	HtmlUrl string `protobuf:"bytes,4,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	// The timestamp this folder-info is crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
}

func (m *FolderInfo) Reset()                    { *m = FolderInfo{} }
func (m *FolderInfo) String() string            { return proto.CompactTextString(m) }
func (*FolderInfo) ProtoMessage()               {}
func (*FolderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FolderInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

type CrawlingInfo struct {
	// The timestamp the related entry was crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
	Etag         string                     `protobuf:"bytes,2,opt,name=etag" json:"etag,omitempty"`
}

func (m *CrawlingInfo) Reset()                    { *m = CrawlingInfo{} }
func (m *CrawlingInfo) String() string            { return proto.CompactTextString(m) }
func (*CrawlingInfo) ProtoMessage()               {}
func (*CrawlingInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CrawlingInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

type HistoryEvent struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Action    HistoryEvent_Action_Enum   `protobuf:"varint,2,opt,name=action,enum=gcse.spider.HistoryEvent_Action_Enum" json:"action,omitempty"`
}

func (m *HistoryEvent) Reset()                    { *m = HistoryEvent{} }
func (m *HistoryEvent) String() string            { return proto.CompactTextString(m) }
func (*HistoryEvent) ProtoMessage()               {}
func (*HistoryEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HistoryEvent) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type HistoryEvent_Action struct {
}

func (m *HistoryEvent_Action) Reset()                    { *m = HistoryEvent_Action{} }
func (m *HistoryEvent_Action) String() string            { return proto.CompactTextString(m) }
func (*HistoryEvent_Action) ProtoMessage()               {}
func (*HistoryEvent_Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type HistoryInfo struct {
	Events    []*HistoryEvent            `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	FoundTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=found_time,json=foundTime" json:"found_time,omitempty"`
	// Possible value:
	//   web                  added from web
	//   user:<user>          found from user crawling
	//   parent               found by crawling his parent
	//   imported:<pkg>       imported by a <pkg>
	//   testimported:<pkg>   test imported by a <pkg>
	//   package:<pkg>
	//   reference:<pkg>      referenced in the readme file of <pkg>
	FoundWay string `protobuf:"bytes,3,opt,name=found_way,json=foundWay" json:"found_way,omitempty"`
}

func (m *HistoryInfo) Reset()                    { *m = HistoryInfo{} }
func (m *HistoryInfo) String() string            { return proto.CompactTextString(m) }
func (*HistoryInfo) ProtoMessage()               {}
func (*HistoryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HistoryInfo) GetEvents() []*HistoryEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *HistoryInfo) GetFoundTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.FoundTime
	}
	return nil
}

func init() {
	proto.RegisterType((*GoFileInfo)(nil), "gcse.spider.GoFileInfo")
	proto.RegisterType((*RepoInfo)(nil), "gcse.spider.RepoInfo")
	proto.RegisterType((*FolderInfo)(nil), "gcse.spider.FolderInfo")
	proto.RegisterType((*CrawlingInfo)(nil), "gcse.spider.CrawlingInfo")
	proto.RegisterType((*HistoryEvent)(nil), "gcse.spider.HistoryEvent")
	proto.RegisterType((*HistoryEvent_Action)(nil), "gcse.spider.HistoryEvent.Action")
	proto.RegisterType((*HistoryInfo)(nil), "gcse.spider.HistoryInfo")
	proto.RegisterEnum("gcse.spider.GoFileInfo_Status", GoFileInfo_Status_name, GoFileInfo_Status_value)
	proto.RegisterEnum("gcse.spider.HistoryEvent_Action_Enum", HistoryEvent_Action_Enum_name, HistoryEvent_Action_Enum_value)
}

var fileDescriptor0 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x6b, 0xd4, 0x40,
	0x10, 0x37, 0x77, 0xd7, 0xf4, 0x3a, 0x39, 0xf5, 0x58, 0x44, 0xd3, 0x0a, 0x52, 0x02, 0x82, 0x4f,
	0x09, 0x56, 0x2c, 0x8a, 0x88, 0x4a, 0x69, 0xb5, 0x3e, 0x88, 0xa4, 0x2d, 0x3e, 0x1e, 0xdb, 0x64,
	0x9b, 0x5b, 0x4c, 0x76, 0xc3, 0xee, 0xa6, 0xe5, 0xbe, 0x89, 0xcf, 0x7e, 0x1b, 0xbf, 0x81, 0x5f,
	0xc5, 0x37, 0x67, 0x37, 0x49, 0xef, 0x44, 0xd4, 0xf3, 0x29, 0xf3, 0xe7, 0x37, 0x99, 0xdf, 0xfc,
	0x66, 0x16, 0x9e, 0x16, 0xdc, 0xcc, 0x9b, 0xf3, 0x38, 0x93, 0x55, 0x92, 0xd3, 0x4b, 0x9e, 0xe7,
	0x4c, 0x14, 0x99, 0x48, 0x8a, 0x4c, 0xb3, 0xa4, 0x56, 0xd2, 0xc8, 0x44, 0xd7, 0x3c, 0x67, 0xaa,
	0xfb, 0xc4, 0x2e, 0x46, 0x02, 0x9b, 0x8f, 0xdb, 0xd0, 0xce, 0x8b, 0x95, 0x7f, 0x14, 0xb2, 0xa4,
	0xa2, 0x68, 0x2b, 0xcf, 0x9b, 0x8b, 0xa4, 0x36, 0x8b, 0x9a, 0xe9, 0xc4, 0xf0, 0x8a, 0x69, 0x43,
	0xab, 0x7a, 0x69, 0xb5, 0x7f, 0x8a, 0x7e, 0x78, 0x00, 0x6f, 0xe5, 0x11, 0x2f, 0xd9, 0xb1, 0xb8,
	0x90, 0x64, 0x1f, 0x7c, 0xcc, 0x9a, 0x46, 0x87, 0xde, 0xae, 0xf7, 0xe8, 0xd6, 0xde, 0x83, 0x78,
	0xa5, 0x53, 0xbc, 0x04, 0xc6, 0x27, 0x0e, 0x95, 0x76, 0x68, 0x42, 0x60, 0x24, 0x68, 0xc5, 0xc2,
	0x01, 0x56, 0x6d, 0xa5, 0xce, 0x26, 0xbb, 0x10, 0xe4, 0x4c, 0x67, 0x8a, 0xd7, 0x86, 0x4b, 0x11,
	0x0e, 0x5d, 0x6a, 0x35, 0x44, 0xee, 0xc1, 0x26, 0xd7, 0x33, 0x83, 0x84, 0xc2, 0x11, 0x66, 0xc7,
	0xa9, 0xcf, 0xf5, 0x29, 0x7a, 0x24, 0xc4, 0x44, 0x55, 0x4b, 0x65, 0x74, 0xb8, 0xb1, 0x3b, 0xc4,
	0xb2, 0xde, 0x8d, 0xde, 0x83, 0xdf, 0xb6, 0x26, 0x01, 0x6c, 0x9e, 0x89, 0xcf, 0x42, 0x5e, 0x89,
	0xe9, 0x0d, 0x32, 0x85, 0xc9, 0x47, 0xaa, 0x34, 0x3b, 0x69, 0xb2, 0x8c, 0x69, 0x3d, 0xf5, 0xc8,
	0x6d, 0x08, 0x5c, 0xe4, 0x88, 0x22, 0xe5, 0x7c, 0x3a, 0xb0, 0x90, 0x93, 0xb9, 0x6c, 0xca, 0xfc,
	0xb8, 0x10, 0x52, 0xb1, 0xe9, 0x30, 0xfa, 0xee, 0xc1, 0x38, 0x65, 0xb5, 0x74, 0x93, 0xbf, 0x84,
	0x49, 0x49, 0xb5, 0x99, 0x65, 0x8a, 0x5e, 0x61, 0x81, 0x9b, 0x3f, 0xd8, 0xdb, 0x89, 0x0b, 0x29,
	0x8b, 0x92, 0xc5, 0xbd, 0xa2, 0xf1, 0x69, 0x2f, 0x60, 0x1a, 0x58, 0xfc, 0x41, 0x0b, 0x27, 0x77,
	0x60, 0x03, 0xa3, 0x4a, 0x3b, 0x05, 0x36, 0xd2, 0xd6, 0x59, 0x43, 0x82, 0xbb, 0x28, 0xb8, 0x6c,
	0x54, 0xc6, 0x70, 0x50, 0x9b, 0xec, 0xbc, 0x6b, 0x3a, 0x4d, 0x9d, 0x53, 0x83, 0x74, 0x46, 0xeb,
	0xd1, 0x39, 0x6b, 0xe1, 0xd1, 0x57, 0x5c, 0xeb, 0x91, 0x2c, 0x71, 0x69, 0x6e, 0xb8, 0x7e, 0x3d,
	0xde, 0xca, 0x7a, 0x30, 0x56, 0x53, 0x33, 0xef, 0x57, 0x66, 0x6d, 0xd4, 0x68, 0xa8, 0xe7, 0xb4,
	0xe3, 0x69, 0x4d, 0xb2, 0x0d, 0xe3, 0xb9, 0xa9, 0xca, 0x59, 0xa3, 0x4a, 0xc7, 0x01, 0x57, 0x61,
	0xfd, 0x33, 0x55, 0x92, 0x57, 0x70, 0xd3, 0x89, 0xc5, 0x45, 0x31, 0xb3, 0x67, 0xe5, 0x26, 0xf8,
	0x3b, 0xc7, 0x49, 0x5f, 0x60, 0x43, 0x51, 0x06, 0x93, 0x83, 0xce, 0x77, 0x2c, 0x7f, 0xfb, 0xa1,
	0xf7, 0x7f, 0x3f, 0xb4, 0x23, 0x31, 0x43, 0x8b, 0x7e, 0x24, 0x6b, 0x47, 0xdf, 0x3c, 0x98, 0xbc,
	0xe3, 0xda, 0x48, 0xb5, 0x38, 0xbc, 0x64, 0xc2, 0x90, 0x67, 0xb0, 0x75, 0xfd, 0x08, 0xd6, 0xe8,
	0xb0, 0x04, 0xe3, 0x4e, 0x7c, 0x9a, 0xb9, 0x45, 0x0e, 0xdc, 0xe3, 0x78, 0xf8, 0xcb, 0xe3, 0x58,
	0x6d, 0x12, 0xbf, 0x71, 0xb8, 0xf8, 0x50, 0x34, 0x55, 0xda, 0x15, 0xed, 0xbc, 0x06, 0xbf, 0x0d,
	0x47, 0xfb, 0x30, 0xb2, 0x19, 0x32, 0x86, 0xd1, 0x07, 0x29, 0x18, 0xde, 0x2f, 0x1e, 0xf3, 0xf2,
	0x74, 0x01, 0xfc, 0xeb, 0xab, 0xc5, 0xc4, 0xb1, 0xb8, 0xa4, 0x25, 0xcf, 0xf1, 0x60, 0xbf, 0x78,
	0x10, 0x74, 0x6d, 0x9c, 0x60, 0x8f, 0xc1, 0x67, 0xb6, 0x9d, 0x7d, 0xad, 0x43, 0x9c, 0x63, 0xfb,
	0x8f, 0x84, 0xd2, 0x0e, 0x48, 0x9e, 0x03, 0x5c, 0xc8, 0x46, 0xe4, 0xad, 0xc0, 0x83, 0x7f, 0x8f,
	0xef, 0xd0, 0x4e, 0xdd, 0xfb, 0xd0, 0x3a, 0xb3, 0x2b, 0xba, 0xe8, 0x4e, 0x64, 0xec, 0x02, 0x9f,
	0xe8, 0xe2, 0xdc, 0x77, 0xb5, 0x4f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xad, 0x58, 0x2a,
	0xd1, 0x04, 0x00, 0x00,
}
