package wasitel

import (
	types "github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo/wasitel/types/resource"
	"go.opentelemetry.io/otel/sdk/resource"
)

// Resource transforms a Resource into an OTLP Resource.
func Resource(r *resource.Resource) *types.Resource {
	if r == nil {
		return nil
	}
	return &types.Resource{Attributes: ResourceAttributes(r)}
}
