//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world hello --out gen ./wit

package main

import (
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/otel"

	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net/wasihttp"
)

const name = "github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo"

var (
	wasiTransport = &wasihttp.Transport{}
	httpClient    = &http.Client{Transport: wasiTransport}
	tracer        = otel.Tracer(name)
)

func init() {
	wasihttp.HandleFunc(handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	setupOTelSDK(httpClient)
	logger := wasilog.ContextLogger("handler")
	logger.Info("handler")
	_, span := tracer.Start(r.Context(), "hello-handler")
	logger.Info("tracer.Start")
	span.End()
	logger.Info("span.End")

	url := "https://dog.ceo/api/breeds/image/random"
	logger.Info("url")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	logger.Info("http.NewRequest")
	if err != nil {
		logger.Error("failed to create outbound request", "err", err)
		http.Error(w, fmt.Sprintf("handler: failed to create outbound request: %s", err), http.StatusInternalServerError)
		return
	}
	// span.SetAttributes(attribute.String("url", url))
	logger.Info("span.SetAttributes")

	resp, err := httpClient.Do(req)
	logger.Info("httpClient.Do")
	if err != nil {
		logger.Error("failed to make outbound request", "err", err)
		http.Error(w, fmt.Sprintf("handler: failed to make outbound request: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	logger.Info("w.WriteHeader")

	_, _ = io.Copy(w, resp.Body)
}

func main() {}
