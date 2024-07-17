package convpb

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestSQL(t *testing.T) {
	if SQLNullBool(sql.NullBool{}).ToBoolValue() != nil {
		t.Fatal("not nil")
	}
	if SQLNullString(sql.NullString{}).ToStringValue() != nil {
		t.Fatal("not nil")
	}
	if SQLNullInt16(sql.NullInt16{}).ToInt64Value() != nil {
		t.Fatal("not nil")
	}
	if SQLNullInt32(sql.NullInt32{}).ToInt64Value() != nil {
		t.Fatal("not nil")
	}
	if SQLNullInt64(sql.NullInt64{}).ToInt64Value() != nil {
		t.Fatal("not nil")
	}
	if SQLNullFloat64(sql.NullFloat64{}).ToInt64Value() != nil {
		t.Fatal("not nil")
	}
	if SQLNullTime(sql.NullTime{}).ToTimestamp() != nil {
		t.Fatal("not nil")
	}
}

func TestNumberWrapper_ToStringValue(t *testing.T) {
	type test struct {
		in  NumberWrapper
		out string
	}

	cases := []test{
		{in: Number(int8(31)), out: "31"},
		{in: Number(int16(32)), out: "32"},
		{in: Number(int32(33)), out: "33"},
		{in: Number(int64(34)), out: "34"},
		{in: Number(int(35)), out: "35"},
		{in: Number(uint8(36)), out: "36"},
		{in: Number(uint16(37)), out: "37"},
		{in: Number(uint32(38)), out: "38"},
		{in: Number(uint64(39)), out: "39"},
		{in: Number(uint(40)), out: "40"},
		{in: Number(uintptr(1)), out: "1"},
		{in: Number(float32(2.12304)), out: "2.12304"},
		{in: Number(float64(3.12304)), out: "3.12304"},
	}

	for _, cas := range cases {
		v := cas.in.ToStringValue().GetValue()
		if v != cas.out {
			t.Fatalf("v(%s) != cas.out(%s)", v, cas.out)
		}
	}
}

func TestUUIDWrapper(t *testing.T) {
	id := uuid.New()
	w := UUID(id)
	assert.False(t, w.IsNil())
	assert.Equal(t, w.ToStringValue(), wrapperspb.String(id.String()))
	assert.Equal(t, w.ToString(), id.String())
	assert.True(t, w.CondSetNil(id).IsNil())
}

func TestStringValuer_UUID(t *testing.T) {
	id := uuid.New()
	s := wrapperspb.String(id.String())
	v := StringValue(s)

	pid, err := v.ToUUID()
	require.NoError(t, err)
	assert.Equal(t, pid, id)
	assert.Equal(t, v.ToUUIDOrNil(), id)

	pnid, err := v.ToNullUUID()
	require.NoError(t, err)
	assert.Equal(t, pnid, uuid.NullUUID{UUID: id, Valid: true})
	assert.Equal(t, v.ToNullUUIDOrNil(), uuid.NullUUID{UUID: id, Valid: true})
}

func TestStringValuer_UUID_error(t *testing.T) {
	id := "asd"
	s := wrapperspb.String(id)
	v := StringValue(s)

	pid, err := v.ToUUID()
	require.Error(t, err)
	assert.Equal(t, pid, uuid.Nil)
	assert.Equal(t, v.ToUUIDOrNil(), uuid.Nil)

	pnid, err := v.ToNullUUID()
	require.Error(t, err)
	assert.Equal(t, pnid, uuid.NullUUID{})
	assert.Equal(t, v.ToNullUUIDOrNil(), uuid.NullUUID{})
}
