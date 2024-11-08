package wasitel

import (
	"github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo/wasitel/types/common"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

// Iterator transforms an attribute iterator into OTLP key-values.
func Iterator(iter attribute.Iterator) []*common.KeyValue {
	l := iter.Len()
	if l == 0 {
		return nil
	}

	out := make([]*common.KeyValue, 0, l)
	for iter.Next() {
		out = append(out, KeyValue(iter.Attribute()))
	}
	return out
}

// ResourceAttributes transforms a Resource OTLP key-values.
func ResourceAttributes(res *resource.Resource) []*common.KeyValue {
	return Iterator(res.Iter())
}

// KeyValues transforms a slice of attribute KeyValues into OTLP key-values.
func KeyValues(attrs []attribute.KeyValue) []*common.KeyValue {
	if len(attrs) == 0 {
		return nil
	}

	out := make([]*common.KeyValue, 0, len(attrs))
	for _, kv := range attrs {
		out = append(out, KeyValue(kv))
	}
	return out
}

// KeyValue transforms an attribute KeyValue into an OTLP key-value.
func KeyValue(kv attribute.KeyValue) *common.KeyValue {
	return &common.KeyValue{Key: string(kv.Key), Value: Value(kv.Value)}
}

// Value transforms an attribute Value into an OTLP AnyValue.
func Value(v attribute.Value) *common.AnyValue {
	av := new(common.AnyValue)
	switch v.Type() {
	case attribute.BOOL:
		av.BoolValue = v.AsBool()
	case attribute.BOOLSLICE:
		av.ArrayValue = &common.ArrayValue{
			Values: boolSliceValues(v.AsBoolSlice()),
		}
	case attribute.INT64:
		av.IntValue = v.AsInt64()
	case attribute.INT64SLICE:
		av.ArrayValue = &common.ArrayValue{
			Values: int64SliceValues(v.AsInt64Slice()),
		}
	case attribute.FLOAT64:
		av.DoubleValue = v.AsFloat64()
	case attribute.FLOAT64SLICE:
		av.ArrayValue = &common.ArrayValue{
			Values: float64SliceValues(v.AsFloat64Slice()),
		}
	case attribute.STRING:
		av.StringValue = v.AsString()
	case attribute.STRINGSLICE:
		av.ArrayValue = &common.ArrayValue{
			Values: stringSliceValues(v.AsStringSlice()),
		}
	default:
		av.StringValue = "INVALID"
	}
	return av
}

func boolSliceValues(vals []bool) []*common.AnyValue {
	converted := make([]*common.AnyValue, len(vals))
	for i, v := range vals {
		converted[i] = &common.AnyValue{
			BoolValue: v,
		}
	}
	return converted
}

func int64SliceValues(vals []int64) []*common.AnyValue {
	converted := make([]*common.AnyValue, len(vals))
	for i, v := range vals {
		converted[i] = &common.AnyValue{
			IntValue: v,
		}
	}
	return converted
}

func float64SliceValues(vals []float64) []*common.AnyValue {
	converted := make([]*common.AnyValue, len(vals))
	for i, v := range vals {
		converted[i] = &common.AnyValue{
			DoubleValue: v,
		}
	}
	return converted
}

func stringSliceValues(vals []string) []*common.AnyValue {
	converted := make([]*common.AnyValue, len(vals))
	for i, v := range vals {
		converted[i] = &common.AnyValue{
			StringValue: v,
		}
	}
	return converted
}
