// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/location/locations.proto

package location // import "google.golang.org/genproto/googleapis/cloud/location"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
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

// The request message for [Locations.ListLocations][google.cloud.location.Locations.ListLocations].
type ListLocationsRequest struct {
	// The resource that owns the locations collection, if applicable.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The standard list filter.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// The standard list page size.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The standard list page token.
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLocationsRequest) Reset()         { *m = ListLocationsRequest{} }
func (m *ListLocationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListLocationsRequest) ProtoMessage()    {}
func (*ListLocationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_locations_7a7af132c8d24683, []int{0}
}
func (m *ListLocationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLocationsRequest.Unmarshal(m, b)
}
func (m *ListLocationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLocationsRequest.Marshal(b, m, deterministic)
}
func (dst *ListLocationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLocationsRequest.Merge(dst, src)
}
func (m *ListLocationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListLocationsRequest.Size(m)
}
func (m *ListLocationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLocationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListLocationsRequest proto.InternalMessageInfo

func (m *ListLocationsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListLocationsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListLocationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListLocationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [Locations.ListLocations][google.cloud.location.Locations.ListLocations].
type ListLocationsResponse struct {
	// A list of locations that matches the specified filter in the request.
	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	// The standard List next-page token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLocationsResponse) Reset()         { *m = ListLocationsResponse{} }
func (m *ListLocationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListLocationsResponse) ProtoMessage()    {}
func (*ListLocationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_locations_7a7af132c8d24683, []int{1}
}
func (m *ListLocationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLocationsResponse.Unmarshal(m, b)
}
func (m *ListLocationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLocationsResponse.Marshal(b, m, deterministic)
}
func (dst *ListLocationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLocationsResponse.Merge(dst, src)
}
func (m *ListLocationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListLocationsResponse.Size(m)
}
func (m *ListLocationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLocationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLocationsResponse proto.InternalMessageInfo

func (m *ListLocationsResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *ListLocationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [Locations.GetLocation][google.cloud.location.Locations.GetLocation].
type GetLocationRequest struct {
	// Resource name for the location.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLocationRequest) Reset()         { *m = GetLocationRequest{} }
func (m *GetLocationRequest) String() string { return proto.CompactTextString(m) }
func (*GetLocationRequest) ProtoMessage()    {}
func (*GetLocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_locations_7a7af132c8d24683, []int{2}
}
func (m *GetLocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLocationRequest.Unmarshal(m, b)
}
func (m *GetLocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLocationRequest.Marshal(b, m, deterministic)
}
func (dst *GetLocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLocationRequest.Merge(dst, src)
}
func (m *GetLocationRequest) XXX_Size() int {
	return xxx_messageInfo_GetLocationRequest.Size(m)
}
func (m *GetLocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLocationRequest proto.InternalMessageInfo

func (m *GetLocationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A resource that represents Google Cloud Platform location.
type Location struct {
	// Resource name for the location, which may vary between implementations.
	// For example: `"projects/example-project/locations/us-east1"`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The canonical id for this location. For example: `"us-east1"`.
	LocationId string `protobuf:"bytes,4,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// The friendly name for this location, typically a nearby city name.
	// For example, "Tokyo".
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Cross-service attributes for the location. For example
	//
	//     {"cloud.googleapis.com/region": "us-east1"}
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Service-specific metadata. For example the available capacity at the given
	// location.
	Metadata             *any.Any `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_locations_7a7af132c8d24683, []int{3}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetLocationId() string {
	if m != nil {
		return m.LocationId
	}
	return ""
}

func (m *Location) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Location) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Location) GetMetadata() *any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*ListLocationsRequest)(nil), "google.cloud.location.ListLocationsRequest")
	proto.RegisterType((*ListLocationsResponse)(nil), "google.cloud.location.ListLocationsResponse")
	proto.RegisterType((*GetLocationRequest)(nil), "google.cloud.location.GetLocationRequest")
	proto.RegisterType((*Location)(nil), "google.cloud.location.Location")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.location.Location.LabelsEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LocationsClient is the client API for Locations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocationsClient interface {
	// Lists information about the supported locations for this service.
	ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error)
	// Gets information about a location.
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*Location, error)
}

type locationsClient struct {
	cc *grpc.ClientConn
}

func NewLocationsClient(cc *grpc.ClientConn) LocationsClient {
	return &locationsClient{cc}
}

func (c *locationsClient) ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error) {
	out := new(ListLocationsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.location.Locations/ListLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/google.cloud.location.Locations/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationsServer is the server API for Locations service.
type LocationsServer interface {
	// Lists information about the supported locations for this service.
	ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	// Gets information about a location.
	GetLocation(context.Context, *GetLocationRequest) (*Location, error)
}

func RegisterLocationsServer(s *grpc.Server, srv LocationsServer) {
	s.RegisterService(&_Locations_serviceDesc, srv)
}

func _Locations_ListLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).ListLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.location.Locations/ListLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).ListLocations(ctx, req.(*ListLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.location.Locations/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Locations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.location.Locations",
	HandlerType: (*LocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLocations",
			Handler:    _Locations_ListLocations_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _Locations_GetLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/location/locations.proto",
}

func init() {
	proto.RegisterFile("google/cloud/location/locations.proto", fileDescriptor_locations_7a7af132c8d24683)
}

var fileDescriptor_locations_7a7af132c8d24683 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0x3a, 0x4d, 0x94, 0x8c, 0x29, 0xa0, 0x55, 0x8a, 0xdc, 0x00, 0x4a, 0x62, 0x04, 0xa4,
	0x05, 0x79, 0x21, 0x5c, 0xf8, 0x51, 0x0e, 0x14, 0x21, 0x84, 0x14, 0xa1, 0xc8, 0x70, 0xe2, 0x12,
	0x6d, 0xe2, 0xad, 0x65, 0xea, 0xec, 0x1a, 0xef, 0xa6, 0xc2, 0x45, 0xed, 0x01, 0xf1, 0x06, 0xe5,
	0x21, 0x78, 0x1f, 0x5e, 0x81, 0x87, 0xe0, 0x88, 0xbc, 0xfe, 0x49, 0x28, 0x2e, 0xe5, 0xb6, 0x3b,
	0xf3, 0x7d, 0xf3, 0xcd, 0xb7, 0x33, 0x36, 0xdc, 0xf6, 0x85, 0xf0, 0x43, 0x46, 0xe6, 0xa1, 0x58,
	0x7a, 0x24, 0x14, 0x73, 0xaa, 0x02, 0xc1, 0xcb, 0x83, 0x74, 0xa2, 0x58, 0x28, 0x81, 0xb7, 0x32,
	0x98, 0xa3, 0x61, 0x4e, 0x91, 0xed, 0xdc, 0xc8, 0xd9, 0x34, 0x0a, 0x08, 0xe5, 0x5c, 0xa8, 0x75,
	0x52, 0x67, 0x3b, 0xcf, 0xea, 0xdb, 0x6c, 0xb9, 0x4f, 0x28, 0x4f, 0xb2, 0x94, 0x7d, 0x02, 0xed,
	0x71, 0x20, 0xd5, 0xb8, 0x90, 0x71, 0xd9, 0xc7, 0x25, 0x93, 0x0a, 0x63, 0xd8, 0xe0, 0x74, 0xc1,
	0x2c, 0xd4, 0x43, 0x83, 0x96, 0xab, 0xcf, 0xf8, 0x1a, 0x34, 0xf6, 0x83, 0x50, 0xb1, 0xd8, 0x32,
	0x74, 0x34, 0xbf, 0xe1, 0xeb, 0xd0, 0x8a, 0xa8, 0xcf, 0xa6, 0x32, 0x38, 0x62, 0x56, 0xad, 0x87,
	0x06, 0x75, 0xb7, 0x99, 0x06, 0xde, 0x06, 0x47, 0x0c, 0xdf, 0x04, 0xd0, 0x49, 0x25, 0x0e, 0x18,
	0xb7, 0x36, 0x34, 0x51, 0xc3, 0xdf, 0xa5, 0x01, 0xfb, 0x04, 0xb6, 0xce, 0xe8, 0xcb, 0x48, 0x70,
	0xc9, 0xf0, 0x08, 0x5a, 0xa5, 0x77, 0x0b, 0xf5, 0x6a, 0x03, 0x73, 0xd8, 0x75, 0x2a, 0xcd, 0x3b,
	0x05, 0xd9, 0x5d, 0x31, 0xf0, 0x1d, 0xb8, 0xc2, 0xd9, 0x27, 0x35, 0x5d, 0xd3, 0xce, 0x9a, 0xde,
	0x4c, 0xc3, 0x93, 0x52, 0x7f, 0x00, 0xf8, 0x15, 0x2b, 0xe5, 0xff, 0xe1, 0xde, 0xfe, 0x66, 0x40,
	0xb3, 0xc0, 0x55, 0x3e, 0x4f, 0x17, 0xcc, 0x42, 0x7f, 0x1a, 0x78, 0xb9, 0x55, 0x28, 0x42, 0xaf,
	0x3d, 0xdc, 0x87, 0x4b, 0x5e, 0x20, 0xa3, 0x90, 0x26, 0x53, 0x4d, 0xae, 0x6b, 0x84, 0x99, 0xc7,
	0xde, 0xa4, 0x35, 0x5e, 0x40, 0x23, 0xa4, 0x33, 0x16, 0x4a, 0xcb, 0xd0, 0x96, 0xef, 0x5d, 0x60,
	0xd9, 0x19, 0x6b, 0xf4, 0x4b, 0xae, 0xe2, 0xc4, 0xcd, 0xa9, 0xf8, 0x01, 0x34, 0x17, 0x4c, 0x51,
	0x8f, 0x2a, 0xaa, 0xc7, 0x61, 0x0e, 0xdb, 0x45, 0x99, 0x62, 0x03, 0x9c, 0xe7, 0x3c, 0x71, 0x4b,
	0x54, 0xe7, 0x09, 0x98, 0x6b, 0x85, 0xf0, 0x55, 0xa8, 0x1d, 0xb0, 0x24, 0x37, 0x97, 0x1e, 0x71,
	0x1b, 0xea, 0x87, 0x34, 0x5c, 0xb2, 0xfc, 0x11, 0xb3, 0xcb, 0x53, 0xe3, 0x31, 0x1a, 0x7e, 0x37,
	0xa0, 0x55, 0x4e, 0x0f, 0x9f, 0x22, 0xd8, 0xfc, 0x63, 0x9e, 0xf8, 0x5c, 0x07, 0x15, 0x5b, 0xd7,
	0xb9, 0xff, 0x7f, 0xe0, 0x6c, 0x45, 0xec, 0xbb, 0x5f, 0x7e, 0xfc, 0x3c, 0x35, 0xfa, 0xb8, 0x4b,
	0x0e, 0x1f, 0x92, 0xcf, 0xe9, 0x93, 0x8e, 0xa2, 0x58, 0x7c, 0x60, 0x73, 0x25, 0xc9, 0xee, 0xf1,
	0xea, 0xd3, 0xc1, 0x5f, 0x11, 0x98, 0x6b, 0x53, 0xc6, 0x3b, 0xe7, 0xc8, 0xfc, 0xbd, 0x09, 0x9d,
	0x8b, 0x76, 0xce, 0xde, 0xd1, 0x4d, 0xdc, 0xc2, 0xfd, 0xaa, 0x26, 0x56, 0x3d, 0x90, 0xdd, 0xe3,
	0x3d, 0x01, 0xdb, 0x73, 0xb1, 0xa8, 0x2e, 0xb8, 0x77, 0xb9, 0xf4, 0x37, 0x49, 0x67, 0x34, 0x41,
	0xef, 0x47, 0x39, 0xd0, 0x17, 0x21, 0xe5, 0xbe, 0x23, 0x62, 0x9f, 0xf8, 0x8c, 0xeb, 0x09, 0x92,
	0x2c, 0x45, 0xa3, 0x40, 0x9e, 0xf9, 0x61, 0x3c, 0x2b, 0x0e, 0xbf, 0x10, 0x9a, 0x35, 0x34, 0xf8,
	0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0xd7, 0x9e, 0x57, 0x5c, 0x04, 0x00, 0x00,
}
