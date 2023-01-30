package convwpb

import (
	"database/sql"
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type NumberValuer interface {
	// Clone cloning value
	Clone() NumberValuer
	// EmptySetNil empty number (0) set is nil
	EmptySetNil() NumberValuer
	IsNil() bool

	uintRefValuer
	uintValuer
	intRefValuer
	intValuer
	floatRefValuer
	floatValuer
	sqlValuer
}

type uintRefValuer interface {
	ToUintRef() *uint
	ToUint8Ref() *uint8
	ToUint16Ref() *uint16
	ToUint32Ref() *uint32
	ToUint64Ref() *uint64
}
type uintValuer interface {
	ToUint() uint
	ToUint8() uint8
	ToUint16() uint16
	ToUint32() uint32
	ToUint64() uint64
}
type intRefValuer interface {
	ToIntRef() *int
	ToInt8Ref() *int8
	ToInt16Ref() *int16
	ToInt32Ref() *int32
	ToInt64Ref() *int64
}
type intValuer interface {
	ToInt() int
	ToInt8() int8
	ToInt16() int16
	ToInt32() int32
	ToInt64() int64
}
type floatRefValuer interface {
	ToFloat32Ref() *float32
	ToFloat64Ref() *float64
}
type floatValuer interface {
	ToFloat32() float32
	ToFloat64() float64
}
type sqlValuer interface {
	ToSQLNullInt16() sql.NullInt16
	ToSQLNullInt32() sql.NullInt32
	ToSQLNullInt64() sql.NullInt64
}

type NumberWrapper interface {
	Clone() NumberWrapper
	EmptySetNil() NumberWrapper
	IsNil() bool

	ToUInt32Value() *wrapperspb.UInt32Value
	ToUInt64Value() *wrapperspb.UInt64Value
	ToInt32Value() *wrapperspb.Int32Value
	ToInt64Value() *wrapperspb.Int64Value
	ToFloatValue() *wrapperspb.FloatValue
	ToDoubleValue() *wrapperspb.DoubleValue
	ToStringValue() *wrapperspb.StringValue
}

type numberWrapper[T number] struct {
	value *T
}

func (w *numberWrapper[T]) IsNil() bool {
	return w == nil || w.value == nil
}
func (w *numberWrapper[T]) Clone() NumberWrapper {
	if w.IsNil() {
		return &numberWrapper[T]{}
	}
	v := *w.value
	return &numberWrapper[T]{value: &v}
}
func (w *numberWrapper[T]) EmptySetNil() NumberWrapper {
	if w.IsNil() {
		return w
	}
	if *w.value == 0 {
		w.value = nil
	}
	return w
}
func (w *numberWrapper[T]) ToUInt32Value() *wrapperspb.UInt32Value {
	if w.IsNil() {
		return nil
	}
	return wrapperspb.UInt32(uint32(*w.value))
}
func (w *numberWrapper[T]) ToUInt64Value() *wrapperspb.UInt64Value {
	if w.IsNil() {
		return nil
	}
	return wrapperspb.UInt64(uint64(*w.value))
}
func (w *numberWrapper[T]) ToInt32Value() *wrapperspb.Int32Value {
	if w.IsNil() {
		return nil
	}
	return wrapperspb.Int32(int32(*w.value))
}
func (w *numberWrapper[T]) ToInt64Value() *wrapperspb.Int64Value {
	if w.IsNil() {
		return nil
	}
	return wrapperspb.Int64(int64(*w.value))
}
func (w *numberWrapper[T]) ToFloatValue() *wrapperspb.FloatValue {
	if w.IsNil() {
		return nil
	}
	return wrapperspb.Float(float32(*w.value))
}
func (w *numberWrapper[T]) ToDoubleValue() *wrapperspb.DoubleValue {
	if w.IsNil() {
		return nil
	}
	return wrapperspb.Double(float64(*w.value))
}
func (w *numberWrapper[T]) ToStringValue() *wrapperspb.StringValue {
	if w.IsNil() {
		return nil
	}
	return wrapperspb.String(fmt.Sprintf("%d", *w.value))
}