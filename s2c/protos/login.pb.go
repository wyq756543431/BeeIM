// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

protoc --go_out=. login.proto

It is generated from these files:
	login.proto

It has these top-level messages:
	LoginReq
	LoginRes
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LoginReq struct {
	MessageFrom      *string `protobuf:"bytes,1,req" json:"MessageFrom,omitempty"`
	MessageTo        *string `protobuf:"bytes,2,req" json:"MessageTo,omitempty"`
	CryptoKey        *string `protobuf:"bytes,3,opt" json:"CryptoKey,omitempty"`
	UserPass         *string `protobuf:"bytes,4,opt" json:"UserPass,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}

func (m *LoginReq) GetMessageFrom() string {
	if m != nil && m.MessageFrom != nil {
		return *m.MessageFrom
	}
	return ""
}

func (m *LoginReq) GetMessageTo() string {
	if m != nil && m.MessageTo != nil {
		return *m.MessageTo
	}
	return ""
}

func (m *LoginReq) GetCryptoKey() string {
	if m != nil && m.CryptoKey != nil {
		return *m.CryptoKey
	}
	return ""
}

func (m *LoginReq) GetUserPass() string {
	if m != nil && m.UserPass != nil {
		return *m.UserPass
	}
	return ""
}

type LoginRes struct {
	MessageFrom      *string `protobuf:"bytes,1,req" json:"MessageFrom,omitempty"`
	MessageTo        *string `protobuf:"bytes,2,req" json:"MessageTo,omitempty"`
	Result           *string `protobuf:"bytes,3,opt" json:"Result,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}

func (m *LoginRes) GetMessageFrom() string {
	if m != nil && m.MessageFrom != nil {
		return *m.MessageFrom
	}
	return ""
}

func (m *LoginRes) GetMessageTo() string {
	if m != nil && m.MessageTo != nil {
		return *m.MessageTo
	}
	return ""
}

func (m *LoginRes) GetResult() string {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return ""
}
