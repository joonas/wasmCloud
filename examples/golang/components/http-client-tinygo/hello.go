//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world hello --out gen ./wit

package main

import (
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net/wasihttp"
)

const name = "github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo"

var (
	wasiTransport = &wasihttp.Transport{}
	httpClient    = &http.Client{Transport: wasiTransport}
	tracer = otel.Tracer(name)
)

func init() {
	setupOTelSDK(httpClient)
	wasihttp.HandleFunc(handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger := wasilog.ContextLogger("handler")
	_, span := tracer.Start(r.Context(), "hello-handler")
	defer span.End()

	url := "https://dog.ceo/api/breeds/image/random"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error("failed to create outbound request", "err", err)
		http.Error(w, fmt.Sprintf("handler: failed to create outbound request: %s", err), http.StatusInternalServerError)
		return
	}
	span.SetAttributes(attribute.String("url", url))

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error("failed to make outbound request", "err", err)
		http.Error(w, fmt.Sprintf("handler: failed to make outbound request: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

func main() {}
