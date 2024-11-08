package wasitel

import (
	types "github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo/wasitel/types/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func convertSpans(spans []trace.ReadOnlySpan) []*types.ResourceSpans {
	if len(spans) == 0 {
		return nil
	}

	rsm := make(map[attribute.Distinct]*types.ResourceSpans)

	type key struct {
		r  attribute.Distinct
		is instrumentation.Scope
	}
	ssm := make(map[key]*types.ScopeSpans)

	var resources int
	for _, span := range spans {
		if span == nil {
			continue
		}

		rKey := span.Resource().Equivalent()
		k := key{
			r:  rKey,
			is: span.InstrumentationScope(),
		}
		scopeSpan, iOk := ssm[k]
		if !iOk {
			// Either the resource or instrumentation scope were unknown.
			scopeSpan = &types.ScopeSpans{
				Scope:     InstrumentationScope(span.InstrumentationScope()),
				Spans:     []*types.Span{},
				SchemaUrl: span.InstrumentationScope().SchemaURL,
			}
		}
		scopeSpan.Spans = append(scopeSpan.Spans, convertSpan(span))
		ssm[k] = scopeSpan

		rs, rOk := rsm[rKey]
		if !rOk {
			resources++
			// The resource was unknown.
			rs = &types.ResourceSpans{
				Resource:   Resource(span.Resource()),
				ScopeSpans: []*types.ScopeSpans{scopeSpan},
				SchemaUrl:  span.Resource().SchemaURL(),
			}
			rsm[rKey] = rs
			continue
		}

		// The resource has been seen before. Check if the instrumentation
		// library lookup was unknown because if so we need to add it to the
		// ResourceSpanss. Otherwise, the instrumentation library has already
		// been seen and the append we did above will be included it in the
		// ScopeSpans reference.
		if !iOk {
			rs.ScopeSpans = append(rs.ScopeSpans, scopeSpan)
		}
	}

	// Transform the categorized map into a slice
	rss := make([]*types.ResourceSpans, 0, resources)
	for _, rs := range rsm {
		rss = append(rss, rs)
	}
	return rss
}
