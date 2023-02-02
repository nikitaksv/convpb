package convwpb

import (
	"bytes"
	"database/sql"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type StringWrapper interface {
	Commoner[StringWrapper]
	CondSetNil(cond string) StringWrapper
	ToStringValue() *wrapperspb.StringValue
	ToTimestamp(layout string) (*timestamppb.Timestamp, error)
}

type stringWrapper struct {
	value *string
}

func (s *stringWrapper) IsNil() bool {
	return s == nil || s.value == nil
}
func (s *stringWrapper) Clone() StringWrapper {
	if s.IsNil() {
		return &stringWrapper{}
	}
	cop := strings.Clone(*s.value)
	return &stringWrapper{value: &cop}
}
func (s *stringWrapper) EmptySetNil() StringWrapper {
	if s.IsNil() {
		return s
	}
	if *s.value == "" {
		s.value = nil
	}
	return s
}
func (s *stringWrapper) CondSetNil(cond string) StringWrapper {
	if s.IsNil() {
		return s
	}
	if *s.value == cond {
		s.value = nil
	}
	return s
}
func (s *stringWrapper) ToStringValue() *wrapperspb.StringValue {
	if s.IsNil() {
		return nil
	}
	return wrapperspb.String(*s.value)
}
func (s *stringWrapper) ToTimestamp(layout string) (*timestamppb.Timestamp, error) {
	if s.IsNil() {
		return nil, nil
	}
	t, err := time.Parse(layout, *s.value)
	if err != nil {
		return nil, err
	}
	return timestamppb.New(t), nil
}

type StringValuer interface {
	Commoner[StringValuer]
	CondSetNil(cond string) StringValuer
	// ToStringRef returned string reference
	ToStringRef() *string
	// ToString returned string. If string empty or nil return empty string
	ToString() string
	// ToSQLNullString returned sql.NullString. If string nil return valid=false
	ToSQLNullString() sql.NullString
	ToTimeRef(layout string) (*time.Time, error)
	ToTime(layout string) (time.Time, error)
	// ToBuffer returned bytes buffer
	ToBuffer() *bytes.Buffer
	// ToBytes returned slice of bytes
	ToBytes() []byte
}

type stringValuer struct {
	value *wrapperspb.StringValue
}

func (s *stringValuer) IsNil() bool {
	return s == nil || s.value == nil
}
func (s *stringValuer) Clone() StringValuer {
	if s.IsNil() {
		return &stringValuer{}
	}

	cop := strings.Clone(s.value.GetValue())
	return &stringValuer{value: wrapperspb.String(cop)}
}
func (s *stringValuer) EmptySetNil() StringValuer {
	if s.IsNil() {
		return s
	}
	if s.value.GetValue() == "" {
		s.value = nil
	}
	return s
}
func (s *stringValuer) CondSetNil(cond string) StringValuer {
	if s.IsNil() {
		return s
	}
	if s.value.GetValue() == cond {
		s.value = nil
	}
	return s
}
func (s *stringValuer) ToStringRef() *string {
	if s.IsNil() {
		return nil
	}
	return &s.value.Value
}
func (s *stringValuer) ToString() string {
	if s == nil {
		return ""
	}
	return s.value.GetValue()
}
func (s *stringValuer) ToTimeRef(layout string) (*time.Time, error) {
	t, err := s.ToTime(layout)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
func (s *stringValuer) ToTime(layout string) (time.Time, error) {
	if s == nil {
		return time.Time{}, nil
	}

	return time.Parse(layout, s.value.GetValue())
}
func (s *stringValuer) ToSQLNullString() sql.NullString {
	return sql.NullString{
		String: s.ToString(),
		Valid:  s != nil && s.value != nil,
	}
}
func (s *stringValuer) ToBuffer() *bytes.Buffer {
	if s.IsNil() {
		return nil
	}

	return bytes.NewBufferString(s.ToString())
}
func (s *stringValuer) ToBytes() []byte {
	if s.IsNil() {
		return nil
	}

	return []byte(s.ToString())
}
