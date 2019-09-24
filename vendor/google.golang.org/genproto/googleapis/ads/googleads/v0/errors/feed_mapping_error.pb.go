// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/feed_mapping_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible feed item errors.
type FeedMappingErrorEnum_FeedMappingError int32

const (
	// Enum unspecified.
	FeedMappingErrorEnum_UNSPECIFIED FeedMappingErrorEnum_FeedMappingError = 0
	// The received error code is not known in this version.
	FeedMappingErrorEnum_UNKNOWN FeedMappingErrorEnum_FeedMappingError = 1
	// The given placeholder field does not exist.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_FIELD FeedMappingErrorEnum_FeedMappingError = 2
	// The given criterion field does not exist.
	FeedMappingErrorEnum_INVALID_CRITERION_FIELD FeedMappingErrorEnum_FeedMappingError = 3
	// The given placeholder type does not exist.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE FeedMappingErrorEnum_FeedMappingError = 4
	// The given criterion type does not exist.
	FeedMappingErrorEnum_INVALID_CRITERION_TYPE FeedMappingErrorEnum_FeedMappingError = 5
	// A feed mapping must contain at least one attribute field mapping.
	FeedMappingErrorEnum_NO_ATTRIBUTE_FIELD_MAPPINGS FeedMappingErrorEnum_FeedMappingError = 7
	// The type of the feed attribute referenced in the attribute field mapping
	// must match the type of the placeholder field.
	FeedMappingErrorEnum_FEED_ATTRIBUTE_TYPE_MISMATCH FeedMappingErrorEnum_FeedMappingError = 8
	// A feed mapping for a system generated feed cannot be operated on.
	FeedMappingErrorEnum_CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED FeedMappingErrorEnum_FeedMappingError = 9
	// Only one feed mapping for a placeholder type is allowed per feed or
	// customer (depending on the placeholder type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE FeedMappingErrorEnum_FeedMappingError = 10
	// Only one feed mapping for a criterion type is allowed per customer.
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE FeedMappingErrorEnum_FeedMappingError = 11
	// Only one feed attribute mapping for a placeholder field is allowed
	// (depending on the placeholder type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD FeedMappingErrorEnum_FeedMappingError = 12
	// Only one feed attribute mapping for a criterion field is allowed
	// (depending on the criterion type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD FeedMappingErrorEnum_FeedMappingError = 13
	// This feed mapping may not contain any explicit attribute field mappings.
	FeedMappingErrorEnum_UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS FeedMappingErrorEnum_FeedMappingError = 14
	// Location placeholder feed mappings can only be created for Places feeds.
	FeedMappingErrorEnum_LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS FeedMappingErrorEnum_FeedMappingError = 15
	// Mappings for typed feeds cannot be modified.
	FeedMappingErrorEnum_CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED FeedMappingErrorEnum_FeedMappingError = 16
	// The given placeholder type can only be mapped to system generated feeds.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED FeedMappingErrorEnum_FeedMappingError = 17
	// The given placeholder type cannot be mapped to a system generated feed
	// with the given type.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE FeedMappingErrorEnum_FeedMappingError = 18
)

var FeedMappingErrorEnum_FeedMappingError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INVALID_PLACEHOLDER_FIELD",
	3:  "INVALID_CRITERION_FIELD",
	4:  "INVALID_PLACEHOLDER_TYPE",
	5:  "INVALID_CRITERION_TYPE",
	7:  "NO_ATTRIBUTE_FIELD_MAPPINGS",
	8:  "FEED_ATTRIBUTE_TYPE_MISMATCH",
	9:  "CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED",
	10: "MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE",
	11: "MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE",
	12: "MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD",
	13: "MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD",
	14: "UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS",
	15: "LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS",
	16: "CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED",
	17: "INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED",
	18: "INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE",
}
var FeedMappingErrorEnum_FeedMappingError_value = map[string]int32{
	"UNSPECIFIED":                  0,
	"UNKNOWN":                      1,
	"INVALID_PLACEHOLDER_FIELD":    2,
	"INVALID_CRITERION_FIELD":      3,
	"INVALID_PLACEHOLDER_TYPE":     4,
	"INVALID_CRITERION_TYPE":       5,
	"NO_ATTRIBUTE_FIELD_MAPPINGS":  7,
	"FEED_ATTRIBUTE_TYPE_MISMATCH": 8,
	"CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED":    9,
	"MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE":                  10,
	"MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE":                    11,
	"MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD":                 12,
	"MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD":                   13,
	"UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS":                     14,
	"LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS":              15,
	"CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED":                   16,
	"INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED":  17,
	"INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE": 18,
}

func (x FeedMappingErrorEnum_FeedMappingError) String() string {
	return proto.EnumName(FeedMappingErrorEnum_FeedMappingError_name, int32(x))
}
func (FeedMappingErrorEnum_FeedMappingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_error_05e4b3c477cdcd72, []int{0, 0}
}

// Container for enum describing possible feed item errors.
type FeedMappingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedMappingErrorEnum) Reset()         { *m = FeedMappingErrorEnum{} }
func (m *FeedMappingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedMappingErrorEnum) ProtoMessage()    {}
func (*FeedMappingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_error_05e4b3c477cdcd72, []int{0}
}
func (m *FeedMappingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingErrorEnum.Unmarshal(m, b)
}
func (m *FeedMappingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingErrorEnum.Marshal(b, m, deterministic)
}
func (dst *FeedMappingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingErrorEnum.Merge(dst, src)
}
func (m *FeedMappingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedMappingErrorEnum.Size(m)
}
func (m *FeedMappingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FeedMappingErrorEnum)(nil), "google.ads.googleads.v0.errors.FeedMappingErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.FeedMappingErrorEnum_FeedMappingError", FeedMappingErrorEnum_FeedMappingError_name, FeedMappingErrorEnum_FeedMappingError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/feed_mapping_error.proto", fileDescriptor_feed_mapping_error_05e4b3c477cdcd72)
}

var fileDescriptor_feed_mapping_error_05e4b3c477cdcd72 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x69, 0x47, 0x19, 0x9c, 0x02, 0x0b, 0x16, 0xff, 0x37, 0x06, 0x2a, 0x7f, 0x06, 0x43,
	0x4a, 0x2b, 0x81, 0x18, 0xca, 0xae, 0xdc, 0xc4, 0xe9, 0x2c, 0x12, 0xdb, 0x4a, 0xdc, 0x42, 0x51,
	0x25, 0xab, 0x90, 0x10, 0x4d, 0x5a, 0x9b, 0xaa, 0x81, 0x3d, 0x10, 0x97, 0x3c, 0x0a, 0x17, 0xbc,
	0x06, 0x12, 0x57, 0x3c, 0x02, 0x4a, 0xdc, 0x96, 0xb5, 0x6b, 0xb7, 0xab, 0x1e, 0xf9, 0x7c, 0xbf,
	0xcf, 0xc7, 0x5f, 0x73, 0x60, 0x2f, 0x49, 0xd3, 0xe4, 0x28, 0xae, 0xf7, 0xa3, 0xac, 0xae, 0xcb,
	0xbc, 0x3a, 0x6e, 0xd4, 0xe3, 0xf1, 0x38, 0x1d, 0x67, 0xf5, 0x2f, 0x71, 0x1c, 0xa9, 0x41, 0x7f,
	0x34, 0x3a, 0x1c, 0x26, 0xaa, 0x38, 0x33, 0x47, 0xe3, 0xf4, 0x6b, 0x8a, 0xb6, 0xb5, 0xda, 0xec,
	0x47, 0x99, 0x39, 0x03, 0xcd, 0xe3, 0x86, 0xa9, 0xc1, 0xda, 0xef, 0x0a, 0xdc, 0x74, 0xe3, 0x38,
	0xf2, 0x35, 0x4b, 0xf2, 0x53, 0x32, 0xfc, 0x36, 0xa8, 0xfd, 0xaa, 0x80, 0xb1, 0xd8, 0x40, 0x1b,
	0x50, 0x6d, 0xb3, 0x50, 0x10, 0x9b, 0xba, 0x94, 0x38, 0xc6, 0x05, 0x54, 0x85, 0xf5, 0x36, 0x7b,
	0xc7, 0xf8, 0x7b, 0x66, 0x94, 0xd0, 0x03, 0xb8, 0x47, 0x59, 0x07, 0x7b, 0xd4, 0x51, 0xc2, 0xc3,
	0x36, 0x39, 0xe0, 0x9e, 0x43, 0x02, 0xe5, 0x52, 0xe2, 0x39, 0x46, 0x19, 0x6d, 0xc2, 0x9d, 0x69,
	0xdb, 0x0e, 0xa8, 0x24, 0x01, 0xe5, 0x6c, 0xd2, 0x5c, 0x43, 0x5b, 0x70, 0x77, 0x19, 0x2b, 0xbb,
	0x82, 0x18, 0x17, 0xd1, 0x7d, 0xb8, 0x7d, 0x1a, 0x2d, 0x7a, 0x15, 0xf4, 0x10, 0x36, 0x19, 0x57,
	0x58, 0xca, 0x80, 0x36, 0xdb, 0x92, 0x68, 0x47, 0xe5, 0x63, 0x21, 0x28, 0x6b, 0x85, 0xc6, 0x3a,
	0x7a, 0x04, 0x5b, 0x2e, 0x21, 0xce, 0x09, 0x49, 0x4e, 0x2a, 0x9f, 0x86, 0x3e, 0x96, 0xf6, 0x81,
	0x71, 0x19, 0xbd, 0x85, 0xd7, 0x36, 0x66, 0x8c, 0x4b, 0xc5, 0x05, 0x09, 0xb0, 0x24, 0x8a, 0xb3,
	0x99, 0x83, 0x72, 0x79, 0xa0, 0xc2, 0x6e, 0x28, 0x89, 0xaf, 0x5a, 0x84, 0x15, 0x7d, 0x47, 0xe5,
	0x8e, 0xc6, 0x15, 0xb4, 0x0b, 0xcf, 0xfc, 0xb6, 0x27, 0xa9, 0xf0, 0xc8, 0x3c, 0x70, 0xea, 0x11,
	0x80, 0x9e, 0xc3, 0x93, 0xe5, 0xda, 0x85, 0x27, 0x55, 0xd1, 0x4b, 0xd8, 0x39, 0xdf, 0x55, 0x27,
	0x77, 0x15, 0xbd, 0x80, 0xa7, 0xe7, 0xd9, 0x6a, 0xe9, 0x35, 0xb4, 0x03, 0x8f, 0xdb, 0x8c, 0x7c,
	0x10, 0xc4, 0x96, 0x73, 0x79, 0x2c, 0x44, 0x76, 0x1d, 0x99, 0xb0, 0xeb, 0x71, 0x1b, 0xcb, 0x1c,
	0x3e, 0x79, 0x27, 0x67, 0x5e, 0xf7, 0xff, 0x20, 0x61, 0x91, 0x42, 0x68, 0x6c, 0xe4, 0x33, 0x4c,
	0x02, 0xf4, 0xb9, 0x43, 0xdd, 0xee, 0xfc, 0x20, 0xf9, 0xab, 0x26, 0x89, 0x19, 0xc8, 0x82, 0x37,
	0xab, 0xfe, 0xe8, 0x42, 0xcc, 0x38, 0x5b, 0x91, 0xf6, 0x0d, 0xb4, 0x0f, 0x7b, 0x67, 0xb2, 0x4b,
	0x39, 0x1d, 0x2a, 0x6a, 0xfe, 0x2d, 0x41, 0xed, 0x73, 0x3a, 0x30, 0xcf, 0x5e, 0x88, 0xe6, 0xad,
	0xc5, 0x8f, 0x5e, 0xe4, 0x7b, 0x24, 0x4a, 0x1f, 0x9d, 0x09, 0x98, 0xa4, 0x47, 0xfd, 0x61, 0x62,
	0xa6, 0xe3, 0xa4, 0x9e, 0xc4, 0xc3, 0x62, 0xcb, 0xa6, 0x2b, 0x39, 0x3a, 0xcc, 0x56, 0x6d, 0xe8,
	0xbe, 0xfe, 0xf9, 0x5e, 0x5e, 0x6b, 0x61, 0xfc, 0xa3, 0xbc, 0xdd, 0xd2, 0x66, 0x38, 0xca, 0x4c,
	0x5d, 0xe6, 0x55, 0xa7, 0x61, 0x16, 0x57, 0x66, 0x3f, 0xa7, 0x82, 0x1e, 0x8e, 0xb2, 0xde, 0x4c,
	0xd0, 0xeb, 0x34, 0x7a, 0x5a, 0xf0, 0xa7, 0x5c, 0xd3, 0xa7, 0x96, 0x85, 0xa3, 0xcc, 0xb2, 0x66,
	0x12, 0xcb, 0xea, 0x34, 0x2c, 0x4b, 0x8b, 0x3e, 0x5d, 0x2a, 0xa6, 0x7b, 0xf5, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x2a, 0x93, 0xb4, 0xdd, 0x3e, 0x04, 0x00, 0x00,
}
