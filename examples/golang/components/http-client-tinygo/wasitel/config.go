package wasitel

import "net/http"

// TODO: Implement configurable fields and way to change them
type config struct {
	HttpClient *http.Client
}

func newConfig(opts ...Option) config {
	cfg := config{}
	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}
	return cfg
}

type Option interface {
	apply(config) config
}

func WithHTTPClient(client *http.Client) Option {
	return httpClientOption{client}
}

type httpClientOption struct {
	HttpClient *http.Client
}

func (h httpClientOption) apply(config config) config {
	config.HttpClient = h.HttpClient
	return config
}
