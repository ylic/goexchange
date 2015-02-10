// Code generated by protoc-gen-go.
// source: commonOrder.proto
// DO NOT EDIT!

/*
Package commonOrder is a generated protocol buffer package.

It is generated from these files:
	commonOrder.proto

It has these top-level messages:
	CommonOrderPb
*/
package exchange

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CommonOrderPb struct {
	Oid              *int64   `protobuf:"varint,1,req" json:"Oid,omitempty"`
	Price            *float64 `protobuf:"fixed64,2,req" json:"Price,omitempty"`
	Code             *string  `protobuf:"bytes,3,req" json:"Code,omitempty"`
	BsFlag           *int32   `protobuf:"varint,4,req" json:"BsFlag,omitempty"`
	Count            *int32   `protobuf:"varint,5,req" json:"Count,omitempty"`
	OrderTime        *string  `protobuf:"bytes,6,req" json:"OrderTime,omitempty"`
	IsCancle         *bool    `protobuf:"varint,7,req" json:"IsCancle,omitempty"`
	Position         *int64   `protobuf:"varint,8,req" json:"Position,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CommonOrderPb) Reset()         { *m = CommonOrderPb{} }
func (m *CommonOrderPb) String() string { return proto.CompactTextString(m) }
func (*CommonOrderPb) ProtoMessage()    {}

func (m *CommonOrderPb) GetOid() int64 {
	if m != nil && m.Oid != nil {
		return *m.Oid
	}
	return 0
}

func (m *CommonOrderPb) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *CommonOrderPb) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *CommonOrderPb) GetBsFlag() int32 {
	if m != nil && m.BsFlag != nil {
		return *m.BsFlag
	}
	return 0
}

func (m *CommonOrderPb) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CommonOrderPb) GetOrderTime() string {
	if m != nil && m.OrderTime != nil {
		return *m.OrderTime
	}
	return ""
}

func (m *CommonOrderPb) GetIsCancle() bool {
	if m != nil && m.IsCancle != nil {
		return *m.IsCancle
	}
	return false
}

func (m *CommonOrderPb) GetPosition() int64 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

func init() {
}
