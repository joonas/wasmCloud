package wasitel

import (
	"github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo/wasitel/types/common"
	"go.opentelemetry.io/otel/sdk/instrumentation"
)

func InstrumentationScope(il instrumentation.Scope) *common.InstrumentationScope {
	if il == (instrumentation.Scope{}) {
		return nil
	}
	return &common.InstrumentationScope{
		Name:    il.Name,
		Version: il.Version,
	}
}
