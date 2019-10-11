// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/type/dayofweek.proto

package google_type

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a day of week.
type DayOfWeek int32

const (
	// The unspecified day-of-week.
	DAY_OF_WEEK_UNSPECIFIED DayOfWeek = 0
	// The day-of-week of Monday.
	MONDAY DayOfWeek = 1
	// The day-of-week of Tuesday.
	TUESDAY DayOfWeek = 2
	// The day-of-week of Wednesday.
	WEDNESDAY DayOfWeek = 3
	// The day-of-week of Thursday.
	THURSDAY DayOfWeek = 4
	// The day-of-week of Friday.
	FRIDAY DayOfWeek = 5
	// The day-of-week of Saturday.
	SATURDAY DayOfWeek = 6
	// The day-of-week of Sunday.
	SUNDAY DayOfWeek = 7
)

var DayOfWeek_name = map[int32]string{
	0: "DAY_OF_WEEK_UNSPECIFIED",
	1: "MONDAY",
	2: "TUESDAY",
	3: "WEDNESDAY",
	4: "THURSDAY",
	5: "FRIDAY",
	6: "SATURDAY",
	7: "SUNDAY",
}
var DayOfWeek_value = map[string]int32{
	"DAY_OF_WEEK_UNSPECIFIED": 0,
	"MONDAY":                  1,
	"TUESDAY":                 2,
	"WEDNESDAY":               3,
	"THURSDAY":                4,
	"FRIDAY":                  5,
	"SATURDAY":                6,
	"SUNDAY":                  7,
}

func (DayOfWeek) EnumDescriptor() ([]byte, []int) { return fileDescriptorDayofweek, []int{0} }

func init() {
	proto.RegisterEnum("google.type.DayOfWeek", DayOfWeek_name, DayOfWeek_value)
}
func (x DayOfWeek) String() string {
	s, ok := DayOfWeek_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("google/type/dayofweek.proto", fileDescriptorDayofweek) }

var fileDescriptorDayofweek = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0x49, 0xac, 0xcc, 0x4f, 0x2b, 0x4f, 0x4d,
	0xcd, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x48, 0xea, 0x81, 0x24, 0xb5, 0x5a,
	0x18, 0xb9, 0x38, 0x5d, 0x12, 0x2b, 0xfd, 0xd3, 0xc2, 0x53, 0x53, 0xb3, 0x85, 0xa4, 0xb9, 0xc4,
	0x5d, 0x1c, 0x23, 0xe3, 0xfd, 0xdd, 0xe2, 0xc3, 0x5d, 0x5d, 0xbd, 0xe3, 0x43, 0xfd, 0x82, 0x03,
	0x5c, 0x9d, 0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x84, 0xb8, 0xb8, 0xd8, 0x7c, 0xfd, 0xfd,
	0x5c, 0x1c, 0x23, 0x05, 0x18, 0x85, 0xb8, 0xb9, 0xd8, 0x43, 0x42, 0x5d, 0x83, 0x41, 0x1c, 0x26,
	0x21, 0x5e, 0x2e, 0xce, 0x70, 0x57, 0x17, 0x3f, 0x08, 0x97, 0x59, 0x88, 0x87, 0x8b, 0x23, 0xc4,
	0x23, 0x34, 0x08, 0xcc, 0x63, 0x01, 0xe9, 0x72, 0x0b, 0xf2, 0x04, 0xb1, 0x59, 0x41, 0x32, 0xc1,
	0x8e, 0x21, 0xa1, 0x41, 0x20, 0x1e, 0x1b, 0x48, 0x26, 0x38, 0x14, 0x6c, 0x1e, 0xbb, 0x53, 0xe8,
	0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e,
	0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xe0, 0xe2,
	0x4f, 0xce, 0xcf, 0xd5, 0x43, 0xf2, 0x85, 0x13, 0x1f, 0xdc, 0x0b, 0x01, 0x20, 0x2f, 0x06, 0x30,
	0x2e, 0x62, 0x62, 0x76, 0x0f, 0x09, 0x48, 0x62, 0x03, 0xfb, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x79, 0xc5, 0xab, 0x16, 0x10, 0x01, 0x00, 0x00,
}