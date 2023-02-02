package convwpb

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type DurationWrapper interface {
	Commoner[DurationWrapper]
	ToDuration() *durationpb.Duration
	ToTimestamp() *timestamppb.Timestamp
	ToStringValue() *wrapperspb.StringValue
	ToFormatStringValue(format string) *wrapperspb.StringValue
}

type durationWrapper struct {
	value *time.Duration
}

func (t *durationWrapper) IsNil() bool {
	return t == nil || t.value == nil
}
func (t *durationWrapper) Clone() DurationWrapper {
	if t.IsNil() {
		return &durationWrapper{}
	}

	cop := *t.value
	return &durationWrapper{value: &cop}
}

func (t *durationWrapper) EmptySetNil() DurationWrapper {
	if t.IsNil() {
		return t
	}
	if *t.value == 0 {
		t.value = nil
	}
	return t
}

func (t *durationWrapper) ToDuration() *durationpb.Duration {
	if t.IsNil() {
		return nil
	}
	return durationpb.New(*t.value)
}

func (t *durationWrapper) ToTimestamp() *timestamppb.Timestamp {
	if t.IsNil() {
		return nil
	}
	tm := time.Time{}
	time.Now()
	tm.Add(*t.value)
	return timestamppb.New(tm)
}

func (t *durationWrapper) ToStringValue() *wrapperspb.StringValue {
	if t.IsNil() {
		return nil
	}
	return wrapperspb.String(t.value.String())
}

func (t *durationWrapper) ToFormatStringValue(layout string) *wrapperspb.StringValue {
	if t.IsNil() {
		return nil
	}
	tm := time.Time{}
	tm.Add(*t.value)
	return wrapperspb.String(tm.Format(layout))
}

type DurationValuer interface {
	Commoner[DurationValuer]
	ToDurationRef() *time.Duration
	ToDuration() time.Duration
}

type durationValuer struct {
	value *durationpb.Duration
}

func (t *durationValuer) IsNil() bool {
	return t == nil || t.value == nil
}
func (t *durationValuer) Clone() DurationValuer {
	if t.IsNil() {
		return &durationValuer{}
	}
	return &durationValuer{value: durationpb.New(t.value.AsDuration())}
}

func (t *durationValuer) EmptySetNil() DurationValuer {
	if t.IsNil() {
		return t
	}
	if !t.value.IsValid() {
		t.value = nil
	}
	return t
}

func (t *durationValuer) ToDurationRef() *time.Duration {
	if t.IsNil() {
		return nil
	}
	v := t.value.AsDuration()
	return &v
}

func (t *durationValuer) ToDuration() time.Duration {
	if t.IsNil() {
		var d time.Duration
		return d
	}
	return t.value.AsDuration()
}
