package resource

import (
	"github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo/wasitel/types/common"
)

// Resource information.
type Resource struct {
	// Set of attributes that describe the resource.
	// Attribute keys MUST be unique (it is not allowed to have more than one
	// attribute with the same key).
	Attributes []*common.KeyValue `json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0, then
	// no attributes were dropped.
	DroppedAttributesCount uint32 `json:"dropped_attributes_count,omitempty"`
}
