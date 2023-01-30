package convwpb

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type int32Valuer struct {
	value *wrapperspb.Int32Value
}

func (u *int32Valuer) IsNil() bool {
	return u == nil || u.value == nil
}

func (u *int32Valuer) Clone() NumberValuer {
	if u.IsNil() {
		return &int32Valuer{}
	}

	return &int32Valuer{value: wrapperspb.Int32(u.value.Value)}
}

func (u *int32Valuer) EmptySetNil() NumberValuer {
	if u.IsNil() {
		return u
	}
	if u.value.Value == 0 {
		u.value = nil
	}
	return u
}

func (u *int32Valuer) ToUintRef() *uint {
	if u.IsNil() {
		return nil
	}
	i := uint(u.value.Value)
	return &i
}

func (u *int32Valuer) ToUint8Ref() *uint8 {
	if u.IsNil() {
		return nil
	}
	i := uint8(u.value.Value)
	return &i
}

func (u *int32Valuer) ToUint16Ref() *uint16 {
	if u.IsNil() {
		return nil
	}
	i := uint16(u.value.Value)
	return &i
}

func (u *int32Valuer) ToUint32Ref() *uint32 {
	if u.IsNil() {
		return nil
	}
	i := uint32(u.value.Value)
	return &i
}

func (u *int32Valuer) ToUint64Ref() *uint64 {
	if u.IsNil() {
		return nil
	}
	i := uint64(u.value.Value)
	return &i
}

func (u *int32Valuer) ToUint() uint {
	if u.IsNil() {
		return 0
	}
	return uint(u.value.Value)
}

func (u *int32Valuer) ToUint8() uint8 {
	if u.IsNil() {
		return 0
	}
	return uint8(u.value.Value)
}

func (u *int32Valuer) ToUint16() uint16 {
	if u.IsNil() {
		return 0
	}
	return uint16(u.value.Value)
}

func (u *int32Valuer) ToUint32() uint32 {
	if u.IsNil() {
		return 0
	}
	return uint32(u.value.Value)
}

func (u *int32Valuer) ToUint64() uint64 {
	if u.IsNil() {
		return 0
	}
	return uint64(u.value.Value)
}

func (u *int32Valuer) ToIntRef() *int {
	if u.IsNil() {
		return nil
	}
	i := int(u.value.Value)
	return &i
}

func (u *int32Valuer) ToInt8Ref() *int8 {
	if u.IsNil() {
		return nil
	}
	i := int8(u.value.Value)
	return &i
}

func (u *int32Valuer) ToInt16Ref() *int16 {
	if u.IsNil() {
		return nil
	}
	i := int16(u.value.Value)
	return &i
}

func (u *int32Valuer) ToInt32Ref() *int32 {
	if u.IsNil() {
		return nil
	}
	i := u.value.Value
	return &i
}

func (u *int32Valuer) ToInt64Ref() *int64 {
	if u.IsNil() {
		return nil
	}
	i := int64(u.value.Value)
	return &i
}

func (u *int32Valuer) ToInt() int {
	if u.IsNil() {
		return 0
	}
	return int(u.value.Value)
}

func (u *int32Valuer) ToInt8() int8 {
	if u.IsNil() {
		return 0
	}
	return int8(u.value.Value)
}

func (u *int32Valuer) ToInt16() int16 {
	if u.IsNil() {
		return 0
	}
	return int16(u.value.Value)
}

func (u *int32Valuer) ToInt32() int32 {
	if u.IsNil() {
		return 0
	}
	return u.value.Value
}

func (u *int32Valuer) ToInt64() int64 {
	if u.IsNil() {
		return 0
	}
	return int64(u.value.Value)
}

func (u *int32Valuer) ToFloat32Ref() *float32 {
	if u.IsNil() {
		return nil
	}
	f := float32(u.value.Value)
	return &f
}

func (u *int32Valuer) ToFloat64Ref() *float64 {
	if u.IsNil() {
		return nil
	}
	f := float64(u.value.Value)
	return &f
}

func (u *int32Valuer) ToFloat32() float32 {
	if u.IsNil() {
		return 0
	}
	return float32(u.value.Value)
}

func (u *int32Valuer) ToFloat64() float64 {
	if u.IsNil() {
		return 0
	}
	return float64(u.value.Value)
}

func (u *int32Valuer) ToSQLNullInt16() sql.NullInt16 {
	return sql.NullInt16{
		Int16: u.ToInt16(),
		Valid: !u.IsNil(),
	}
}

func (u *int32Valuer) ToSQLNullInt32() sql.NullInt32 {
	return sql.NullInt32{
		Int32: u.ToInt32(),
		Valid: !u.IsNil(),
	}
}

func (u *int32Valuer) ToSQLNullInt64() sql.NullInt64 {
	return sql.NullInt64{
		Int64: u.ToInt64(),
		Valid: !u.IsNil(),
	}
}
