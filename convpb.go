package convpb

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Nillable[T any] interface {
	EmptySetNil() T
	IsNil() bool
}

type Cloner[T any] interface {
	Clone() T
}

type Commoner[T any] interface {
	Cloner[T]
	Nillable[T]
}

// Unwrappers

func StringValue(value *wrapperspb.StringValue) StringValuer {
	return &stringValuer{value: value}
}
func BytesValue(value *wrapperspb.BytesValue) BytesValuer {
	return &bytesValuer{value: value}
}
func UInt32Value(value *wrapperspb.UInt32Value) NumberValuer {
	return &uint32Valuer{value: value}
}
func UInt64Value(value *wrapperspb.UInt64Value) NumberValuer {
	return &uint64Valuer{value: value}
}
func Int32Value(value *wrapperspb.Int32Value) NumberValuer {
	return &int32Valuer{value: value}
}
func Int64Value(value *wrapperspb.Int64Value) NumberValuer {
	return &int64Valuer{value: value}
}
func FloatValue(value *wrapperspb.FloatValue) NumberValuer {
	return &float32Value{value: value}
}
func DoubleValue(value *wrapperspb.DoubleValue) NumberValuer {
	return &float64Valuer{value: value}
}
func BoolValue(value *wrapperspb.BoolValue) BoolValuer {
	return &boolValuer{value: value}
}
func Timestamp(value *timestamppb.Timestamp) TimestampValuer {
	return &timestampValuer{value: value}
}
func Duration(value *durationpb.Duration) DurationValuer {
	return &durationValuer{value: value}
}

// Wrappers

func StringRef(value *string) StringWrapper {
	return &stringWrapper{value: value}
}
func String(value string) StringWrapper {
	return &stringWrapper{value: &value}
}
func BoolRef(value *bool) BoolWrapper {
	return &boolWrapper{value: value}
}
func Bool(value bool) BoolWrapper {
	return &boolWrapper{value: &value}
}
func Bytes(value []byte) BytesWrapper {
	return &bytesWrapper{value: value}
}
func NumberRef[T number](value *T) NumberWrapper {
	return &numberWrapper[T]{value: value}
}
func Number[T number](value T) NumberWrapper {
	return &numberWrapper[T]{value: &value}
}
func TimeRef(value *time.Time) TimeWrapper {
	return &timeWrapper{value: value}
}
func Time(value time.Time) TimeWrapper {
	return &timeWrapper{value: &value}
}
func TimeDurationRef(value *time.Duration) DurationWrapper {
	return &durationWrapper{value: value}
}
func TimeDuration(value time.Duration) DurationWrapper {
	return &durationWrapper{value: &value}
}
func SQLNullTime(value sql.NullTime) TimeWrapper {
	if !value.Valid {
		return TimeRef(nil)
	}
	return Time(value.Time)
}
func SQLNullBool(value sql.NullBool) BoolWrapper {
	if !value.Valid {
		return BoolRef(nil)
	}
	return Bool(value.Bool)
}
func SQLNullInt64(value sql.NullInt64) NumberWrapper {
	if !value.Valid {
		return NumberRef[int64](nil)
	}
	return Number(value.Int64)
}
func SQLNullInt32(value sql.NullInt32) NumberWrapper {
	if !value.Valid {
		return NumberRef[int32](nil)
	}
	return Number(value.Int32)
}
func SQLNullInt16(value sql.NullInt16) NumberWrapper {
	if !value.Valid {
		return NumberRef[int16](nil)
	}
	return Number(value.Int16)
}
func SQLNullFloat64(value sql.NullFloat64) NumberWrapper {
	if !value.Valid {
		return NumberRef[float64](nil)
	}
	return Number(value.Float64)
}
func SQLNullString(value sql.NullString) StringWrapper {
	if !value.Valid {
		return StringRef(nil)
	}
	return String(value.String)
}
func UUID(id uuid.UUID) UUIDWrapper {
	return &uuidWrapper{value: uuid.NullUUID{Valid: true, UUID: id}}
}
func NullUUID(id uuid.NullUUID) UUIDWrapper {
	return &uuidWrapper{value: id}
}
