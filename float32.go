package convwpb

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type float32Value struct {
	value *wrapperspb.FloatValue
}

func (u *float32Value) IsNil() bool {
	return u == nil || u.value == nil
}

func (u *float32Value) Clone() NumberValuer {
	if u.IsNil() {
		return &float32Value{}
	}

	return &float32Value{value: wrapperspb.Float(u.value.Value)}
}

func (u *float32Value) EmptySetNil() NumberValuer {
	if u.IsNil() {
		return u
	}
	if u.value.Value == 0 {
		u.value = nil
	}
	return u
}

func (u *float32Value) ToUintRef() *uint {
	if u.IsNil() {
		return nil
	}
	i := uint(u.value.Value)
	return &i
}

func (u *float32Value) ToUint8Ref() *uint8 {
	if u.IsNil() {
		return nil
	}
	i := uint8(u.value.Value)
	return &i
}

func (u *float32Value) ToUint16Ref() *uint16 {
	if u.IsNil() {
		return nil
	}
	i := uint16(u.value.Value)
	return &i
}

func (u *float32Value) ToUint32Ref() *uint32 {
	if u.IsNil() {
		return nil
	}
	i := uint32(u.value.Value)
	return &i
}

func (u *float32Value) ToUint64Ref() *uint64 {
	if u.IsNil() {
		return nil
	}
	i := uint64(u.value.Value)
	return &i
}

func (u *float32Value) ToUint() uint {
	if u.IsNil() {
		return 0
	}
	return uint(u.value.Value)
}

func (u *float32Value) ToUint8() uint8 {
	if u.IsNil() {
		return 0
	}
	return uint8(u.value.Value)
}

func (u *float32Value) ToUint16() uint16 {
	if u.IsNil() {
		return 0
	}
	return uint16(u.value.Value)
}

func (u *float32Value) ToUint32() uint32 {
	if u.IsNil() {
		return 0
	}
	return uint32(u.value.Value)
}

func (u *float32Value) ToUint64() uint64 {
	if u.IsNil() {
		return 0
	}
	return uint64(u.value.Value)
}

func (u *float32Value) ToIntRef() *int {
	if u.IsNil() {
		return nil
	}
	i := int(u.value.Value)
	return &i
}

func (u *float32Value) ToInt8Ref() *int8 {
	if u.IsNil() {
		return nil
	}
	i := int8(u.value.Value)
	return &i
}

func (u *float32Value) ToInt16Ref() *int16 {
	if u.IsNil() {
		return nil
	}
	i := int16(u.value.Value)
	return &i
}

func (u *float32Value) ToInt32Ref() *int32 {
	if u.IsNil() {
		return nil
	}
	i := int32(u.value.Value)
	return &i
}

func (u *float32Value) ToInt64Ref() *int64 {
	if u.IsNil() {
		return nil
	}
	i := int64(u.value.Value)
	return &i
}

func (u *float32Value) ToInt() int {
	if u.IsNil() {
		return 0
	}
	return int(u.value.Value)
}

func (u *float32Value) ToInt8() int8 {
	if u.IsNil() {
		return 0
	}
	return int8(u.value.Value)
}

func (u *float32Value) ToInt16() int16 {
	if u.IsNil() {
		return 0
	}
	return int16(u.value.Value)
}

func (u *float32Value) ToInt32() int32 {
	if u.IsNil() {
		return 0
	}
	return int32(u.value.Value)
}

func (u *float32Value) ToInt64() int64 {
	if u.IsNil() {
		return 0
	}
	return int64(u.value.Value)
}

func (u *float32Value) ToFloat32Ref() *float32 {
	if u.IsNil() {
		return nil
	}
	f := u.value.Value
	return &f
}

func (u *float32Value) ToFloat64Ref() *float64 {
	if u.IsNil() {
		return nil
	}
	f := float64(u.value.Value)
	return &f
}

func (u *float32Value) ToFloat32() float32 {
	if u.IsNil() {
		return 0
	}
	return u.value.Value
}

func (u *float32Value) ToFloat64() float64 {
	if u.IsNil() {
		return 0
	}
	return float64(u.value.Value)
}

func (u *float32Value) ToSQLNullInt16() sql.NullInt16 {
	return sql.NullInt16{
		Int16: u.ToInt16(),
		Valid: !u.IsNil(),
	}
}

func (u *float32Value) ToSQLNullInt32() sql.NullInt32 {
	return sql.NullInt32{
		Int32: u.ToInt32(),
		Valid: !u.IsNil(),
	}
}

func (u *float32Value) ToSQLNullInt64() sql.NullInt64 {
	return sql.NullInt64{
		Int64: u.ToInt64(),
		Valid: !u.IsNil(),
	}
}
