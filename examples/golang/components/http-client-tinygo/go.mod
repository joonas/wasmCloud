module github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo

go 1.23.0

toolchain go1.23.2

require (
	github.com/bytecodealliance/wasm-tools-go v0.3.0
	go.opentelemetry.io/otel v1.31.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.31.0
	go.opentelemetry.io/otel/sdk v1.31.0
	go.opentelemetry.io/otel/trace v1.31.0
	go.wasmcloud.dev/component v0.0.3
)

replace go.wasmcloud.dev/component => github.com/wasmcloud/component-sdk-go v0.0.4-0.20241028180114-9bb991602394

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/regclient/regclient v0.7.1 // indirect
	github.com/samber/lo v1.47.0 // indirect
	github.com/samber/slog-common v0.17.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/ulikunitz/xz v0.5.12 // indirect
	github.com/urfave/cli/v3 v3.0.0-alpha9 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	go.opentelemetry.io/otel/metric v1.31.0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)
