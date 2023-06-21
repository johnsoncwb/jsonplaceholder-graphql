package utils

import (
	"context"
	"crypto/tls"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
	"time"
)

type HttpClient struct {
	http.Client
	ctx context.Context
}

func NewHttpClient(ctx context.Context, timeout time.Duration) *HttpClient {
	client := http.Client{
		Timeout: timeout,
	}

	return &HttpClient{
		Client: client,
		ctx:    ctx,
	}
}

func (h *HttpClient) WithSkipVerify() *HttpClient {
	transport := http.DefaultTransport
	tlsClientConfig := &tls.Config{InsecureSkipVerify: true}
	transport.(*http.Transport).TLSClientConfig = tlsClientConfig
	transport.(*http.Transport).DisableKeepAlives = true

	h.Transport = transport
	return h
}

func (h *HttpClient) Do(req *http.Request) (*http.Response, error) {
	txn := newrelic.FromContext(h.ctx)

	var exSeg *newrelic.ExternalSegment
	if txn != nil {
		exSeg = newrelic.StartExternalSegment(txn, req)
		defer exSeg.End()
	}

	res, err := h.Client.Do(req)

	if exSeg != nil {
		exSeg.Response = res
	}

	return res, err
}
