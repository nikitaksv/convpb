package convpb

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type BoolWrapper interface {
	Commoner[BoolWrapper]
	ToBoolValue() *wrapperspb.BoolValue
}

type boolWrapper struct {
	value *bool
}

func (b *boolWrapper) IsNil() bool {
	return b == nil || b.value == nil
}

func (b *boolWrapper) Clone() BoolWrapper {
	if b.IsNil() {
		return &boolWrapper{}
	}
	cop := *b.value
	return &boolWrapper{value: &cop}
}

func (b *boolWrapper) EmptySetNil() BoolWrapper {
	if b.IsNil() {
		return b
	}
	if !*b.value {
		b.value = nil
	}
	return b
}

func (b *boolWrapper) ToBoolValue() *wrapperspb.BoolValue {
	if b.IsNil() {
		return nil
	}
	return wrapperspb.Bool(*b.value)
}

type BoolValuer interface {
	Commoner[BoolValuer]
	ToBoolRef() *bool
	ToBool() bool
	ToSQLNullBool() sql.NullBool
}

type boolValuer struct {
	value *wrapperspb.BoolValue
}

func (b *boolValuer) IsNil() bool {
	return b == nil || b.value == nil
}
func (b *boolValuer) Clone() BoolValuer {
	if b.IsNil() {
		return &boolValuer{}
	}
	return &boolValuer{value: wrapperspb.Bool(b.value.Value)}
}
func (b *boolValuer) EmptySetNil() BoolValuer {
	if b.IsNil() {
		return b
	}
	if !b.value.Value {
		b.value = nil
	}
	return b
}
func (b *boolValuer) ToBoolRef() *bool {
	if b.IsNil() {
		return nil
	}
	v := b.value.Value
	return &v
}
func (b *boolValuer) ToBool() bool {
	if b.IsNil() {
		return false
	}
	return b.value.Value
}

func (b *boolValuer) ToSQLNullBool() sql.NullBool {
	return sql.NullBool{
		Bool:  b.ToBool(),
		Valid: !b.IsNil(),
	}
}
