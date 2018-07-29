package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GatewayRequest struct {
	Req  *http.Request
	body []byte
}

func (g *GatewayRequest) Body() ([]byte, error) {
	if g.body != nil {
		return g.body, nil
	}
	data, err := ioutil.ReadAll(g.Req.Body)
	if err == nil {
		g.body = data
		return data, nil
	}
	return nil, err
}

func (g *GatewayRequest) UnmarshalJson(i interface{}) error {
	data, err := g.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, i)
}
