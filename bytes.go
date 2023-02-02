package convwpb

import (
	"bytes"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type BytesWrapper interface {
	Commoner[BytesWrapper]
	ToBytesValue() *wrapperspb.BytesValue
	ToStringValue() *wrapperspb.StringValue
}

type bytesWrapper struct {
	value []byte
}

func (b *bytesWrapper) IsNil() bool {
	return b == nil || b.value == nil
}
func (b *bytesWrapper) Clone() BytesWrapper {
	if b.IsNil() {
		return &bytesWrapper{}
	}
	cop := make([]byte, len(b.value))
	copy(b.value, cop)
	return &bytesWrapper{value: cop}
}

func (b *bytesWrapper) EmptySetNil() BytesWrapper {
	if b.IsNil() {
		return b
	}
	if len(b.value) == 0 {
		b.value = nil
	}
	return b
}

func (b *bytesWrapper) ToBytesValue() *wrapperspb.BytesValue {
	if b.IsNil() {
		return nil
	}
	return wrapperspb.Bytes(b.value)
}

func (b *bytesWrapper) ToStringValue() *wrapperspb.StringValue {
	if b.IsNil() {
		return nil
	}
	return wrapperspb.String(string(b.value))
}

type BytesValuer interface {
	Commoner[BytesValuer]
	ToBuffer() *bytes.Buffer
	ToBytes() []byte
	ToStringRef() *string
	ToString() string
}

type bytesValuer struct {
	value *wrapperspb.BytesValue
}

func (b *bytesValuer) IsNil() bool {
	return b == nil || b.value == nil
}
func (b *bytesValuer) Clone() BytesValuer {
	if b.IsNil() {
		return &bytesValuer{}
	}

	cop := make([]byte, len(b.value.Value))
	copy(b.value.Value, cop)
	return &bytesValuer{value: wrapperspb.Bytes(cop)}
}
func (b *bytesValuer) EmptySetNil() BytesValuer {
	if b.IsNil() {
		return b
	}
	if len(b.value.Value) == 0 {
		b.value = nil
	}
	return b
}
func (b *bytesValuer) ToBuffer() *bytes.Buffer {
	if b.IsNil() {
		return nil
	}
	return bytes.NewBuffer(b.value.Value)
}

func (b *bytesValuer) ToBytes() []byte {
	if b.IsNil() {
		return nil
	}
	return b.value.Value
}
func (b *bytesValuer) ToStringRef() *string {
	if b.IsNil() {
		return nil
	}
	v := string(b.value.Value)
	return &v
}
func (b *bytesValuer) ToString() string {
	if b.IsNil() {
		return ""
	}
	return string(b.value.Value)
}
