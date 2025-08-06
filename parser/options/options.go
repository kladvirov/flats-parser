package options

import (
	"io"
	"net/http"
)

type Option func(r *http.Request) error

func WithHeader(k, v string) Option {
	return func(r *http.Request) error { r.Header.Set(k, v); return nil }
}

func WithMethod(method string, body io.Reader) Option {
	return func(r *http.Request) error {
		*r = *r.Clone(r.Context())
		r.Method = method
		r.Body = io.NopCloser(body)
		return nil
	}
}
