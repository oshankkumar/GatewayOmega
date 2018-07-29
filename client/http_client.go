package client

import (
	"bytes"
	"net/http"
	"net/url"
	"io"
)

type gatewayClient struct {
	method  string
	baseUrl string
	path    string
	query   string
	body    []byte
	header  http.Header
	client  *http.Client
}

var Default = New()

func New() GatewayClient {
	return &gatewayClient{
		client: http.DefaultClient,
		header: make(http.Header),
	}
}

func (c *gatewayClient) Reset() {
	client := c.client
	c = New().(*gatewayClient)
	if client != nil {
		c.client = client
	}
}

func (c *gatewayClient) Verb(m string) {
	c.method = m
}

func (c *gatewayClient) Path(p string) {
	c.path = p
}

func (c *gatewayClient) Base(base string) {
	c.baseUrl = base
}

func (c *gatewayClient) Query(q url.Values) {
	c.query = q.Encode()
}

func (c *gatewayClient) Header(h http.Header) {
	for key, vals := range h {
		for _, val := range vals {
			c.header.Add(key, val)
		}
	}
}

func (c *gatewayClient) Body(body []byte) {
	c.body = body
}

func (c *gatewayClient) Request() (*http.Request, error) {
	var body io.Reader
	url, err := url.Parse(c.baseUrl)
	if err != nil {
		return nil, err
	}
	url.Path = c.path
	url.RawQuery = c.query
	if len(c.body) != 0 {
		body = bytes.NewBuffer(c.body)
	}

	req,err := http.NewRequest(c.method, url.String(), body)
	if err != nil {
		return nil,err
	}
	req.Header = c.header
	return req,nil
}

func (c *gatewayClient) Client(httpC *http.Client) {
	c.client = httpC
}

func (c *gatewayClient) Do(r *http.Request, success interface{}) (*http.Response, error) {
	return c.client.Do(r)
}
