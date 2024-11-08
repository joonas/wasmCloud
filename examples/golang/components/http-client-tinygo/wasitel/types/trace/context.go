package trace

import (
	"encoding/json"

	"go.wasmcloud.dev/component/log/wasilog"
)

func NewTraceID(str string) *TraceID {
	logger := wasilog.ContextLogger("NewTraceID")
	logger.Info("NewTraceID", "trace_id", str)
	return &TraceID{
		value: []byte(str),
	}
}

type TraceID struct {
	value []byte
}

func (tid *TraceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(tid.value))
}

func NewSpanID(str string) *SpanID {
	logger := wasilog.ContextLogger("NewSpanID")
	logger.Info("NewSpanID", "span_id", str)
	return &SpanID{
		value: []byte(str),
	}
}

type SpanID struct {
	value []byte
}

func (sid *SpanID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(sid.value))
}
