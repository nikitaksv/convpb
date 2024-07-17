package convpb

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UUIDWrapper interface {
	Commoner[UUIDWrapper]
	CondSetNil(cond uuid.UUID) UUIDWrapper
	ToStringValue() *wrapperspb.StringValue
	ToString() string
}

type uuidWrapper struct {
	value uuid.NullUUID
}

func (s *uuidWrapper) IsNil() bool {
	return s == nil || !s.value.Valid
}
func (s *uuidWrapper) Clone() UUIDWrapper {
	if s.IsNil() {
		return &uuidWrapper{}
	}
	var c uuid.UUID
	copy(c[:], s.value.UUID[:])
	return &uuidWrapper{value: uuid.NullUUID{UUID: c, Valid: true}}
}

func (s *uuidWrapper) EmptySetNil() UUIDWrapper {
	if s.IsNil() {
		return s
	}
	if s.value.UUID == uuid.Nil {
		s.value = uuid.NullUUID{}
	}
	return s
}
func (s *uuidWrapper) CondSetNil(cond uuid.UUID) UUIDWrapper {
	if s.IsNil() {
		return s
	}
	if s.value.UUID == cond {
		s.value = uuid.NullUUID{}
	}
	return s
}
func (s *uuidWrapper) ToStringValue() *wrapperspb.StringValue {
	if s.IsNil() {
		return nil
	}
	return wrapperspb.String(s.value.UUID.String())
}
func (s *uuidWrapper) ToString() string {
	if s.IsNil() {
		return ""
	}
	return s.value.UUID.String()
}
