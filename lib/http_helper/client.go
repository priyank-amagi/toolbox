package http_helper

import "net/http"

//go:generate mockery --name IHttpSvc --filename=client.go
type IHttpSvc interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpSvc struct {
	client *http.Client
}

var _ IHttpSvc = (*HttpSvc)(nil)

func NewHttpSvc() *HttpSvc {
	return &HttpSvc{client: &http.Client{}}
}

func (c *HttpSvc) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
