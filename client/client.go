package client

import (
	"net/http"
	"net/url"
)

type GatewayClient interface {
	Do(r *http.Request, success interface{}) (*http.Response, error)
	Verb(method string)
	Base(url string)
	Path(path string)
	Query(q url.Values)
	Header(header http.Header)
	Body(data []byte)
	Request() (*http.Request, error)
	Client(client *http.Client)
	Reset()
}
