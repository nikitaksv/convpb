# convpb

---

Converter for protobuf types

## Installation

```shell
go get github.com/nikitaksv/convpb
```

## Usage

### Convert types from proto to golang

```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nikitaksv/convpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ProtoMsg struct {
	ID        *wrapperspb.StringValue
	String    *wrapperspb.StringValue
	Double    *wrapperspb.DoubleValue
	Timestamp *timestamppb.Timestamp
	Duration  *durationpb.Duration
}

type Model struct {
	ID       string
	String   *string
	Float32  float32
	Time     sql.NullTime
	Duration *time.Duration
}

func main() {
	msg := &ProtoMsg{
		ID:        wrapperspb.String("0000-0000-0000-0000"),
		String:    wrapperspb.String("string"),
		Double:    wrapperspb.Double(11.99),
		Timestamp: nil,
		Duration:  durationpb.New(time.Second * 1000),
	}

	m := &Model{
		ID:       convpb.StringValue(msg.ID).CondSetNil("0000-0000-0000-0000").ToStringRef(),
		String:   convpb.StringValue(msg.String).Clone().ToStringRef(),
		Float32:  convpb.DoubleValue(msg.Double).ToFloat32(),
		Time:     convpb.Timestamp(msg.Timestamp).ToSQLNullTime(),
		Duration: convpb.Duration(msg.Duration).ToDurationRef(),
	}

	bs, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

```

### Convert types from golang to proto

```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ProtoMsg struct {
	ID        *wrapperspb.StringValue
	String    *wrapperspb.StringValue
	Double    *wrapperspb.DoubleValue
	Timestamp *timestamppb.Timestamp
	Duration  *durationpb.Duration
}

type Model struct {
	ID       string
	String   *string
	Float32  float32
	Time     sql.NullTime
	Duration *time.Duration
}

func main() {
	str := "string"
	dur := time.Second * 100

	m := &Model{
		ID:       "0000-0000-0000-0000",
		String:   &str,
		Float32:  123.123,
		Time:     sql.NullTime{Time: time.Now(), Valie: true},
		Duration: &dur,
	}

	msg := &ProtoMsg{
		ID:        convpb.String(m.ID).CondSetNil("0000-0000-0000-0000").ToStringValue(),
		String:    convpb.String(m.ID).ToStringValue(),
		Double:    convpb.Number(m.Float32).ToDoubleValue(),
		Timestamp: convpb.SQLNullTime(m.Time).ToTimestamp(),
		Duration:  convpb.TimeDurationRef(m.dur).ToDuration(),
	}

	bs, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

```

