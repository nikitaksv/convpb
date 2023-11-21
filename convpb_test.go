package convpb

import (
	"database/sql"
	"testing"
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
