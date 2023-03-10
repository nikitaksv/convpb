package convpb

import (
	"database/sql"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type TimeWrapper interface {
	Commoner[TimeWrapper]
	ToTimestamp() *timestamppb.Timestamp
	ToStringValue(format string) *wrapperspb.StringValue
}

type timeWrapper struct {
	value *time.Time
}

func (t *timeWrapper) IsNil() bool {
	return t == nil || t.value == nil
}
func (t *timeWrapper) Clone() TimeWrapper {
	if t.IsNil() {
		return &timeWrapper{}
	}

	cop := *t.value
	return &timeWrapper{value: &cop}
}

func (t *timeWrapper) EmptySetNil() TimeWrapper {
	if t.IsNil() {
		return t
	}
	if t.value.IsZero() {
		t.value = nil
	}
	return t
}

func (t *timeWrapper) ToTimestamp() *timestamppb.Timestamp {
	if t.IsNil() {
		return nil
	}
	return timestamppb.New(*t.value)
}

func (t *timeWrapper) ToStringValue(layout string) *wrapperspb.StringValue {
	if t.IsNil() {
		return nil
	}
	return wrapperspb.String(t.value.Format(layout))
}

type TimestampValuer interface {
	Commoner[TimestampValuer]
	ToTimeRef() *time.Time
	ToTime() time.Time
	ToSQLNullTime() sql.NullTime
}

type timestampValuer struct {
	value *timestamppb.Timestamp
}

func (t *timestampValuer) IsNil() bool {
	return t == nil || t.value == nil
}
func (t *timestampValuer) Clone() TimestampValuer {
	if t.IsNil() {
		return &timestampValuer{}
	}
	return &timestampValuer{value: timestamppb.New(t.value.AsTime())}
}

func (t *timestampValuer) EmptySetNil() TimestampValuer {
	if t.IsNil() {
		return t
	}
	if !t.value.IsValid() || t.value.AsTime().IsZero() {
		t.value = nil
	}
	return t
}

func (t *timestampValuer) ToTimeRef() *time.Time {
	if t.IsNil() {
		return nil
	}
	v := t.value.AsTime()
	return &v
}

func (t *timestampValuer) ToTime() time.Time {
	if t.IsNil() {
		return time.Time{}
	}
	return t.value.AsTime()
}

func (t *timestampValuer) ToSQLNullTime() sql.NullTime {
	return sql.NullTime{
		Time:  t.ToTime(),
		Valid: !t.IsNil(),
	}
}
