package nlu

import (
	"github.com/oshankkumar/GatewayOmega/client"
	ghttp "github.com/oshankkumar/GatewayOmega/http"
	"github.com/oshankkumar/GatewayOmega/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

type Nlu struct{}

func newNLUService(opts ...services.SerivceOptionFunc) services.Service {
	nlu := &Nlu{}
	opt := &services.ServiceOption{}
	for _, optFunc := range opts {
		optFunc(opt)
	}
	return nlu
}

func (n *Nlu) Name() string {
	return "nlu"
}

func (n *Nlu) Send(r *ghttp.GatewayRequest) (*http.Response, error) {
	client.Default.Reset()
	client.Default.Verb(r.Req.Method)
	client.Default.Base(viper.GetString("services.nlu.addr"))
	client.Default.Path(r.Req.URL.Path)
	client.Default.Query(r.Req.URL.Query())
	client.Default.Header(r.Req.Header)
	if body, err := r.Body(); err == nil {
		client.Default.Body(body)
	}
	req, err := client.Default.Request()
	logrus.WithFields(logrus.Fields{
		"service" : "nlu",
		"url"     : req.URL.String(),
		"headers" : req.Header,
	}).Infof("sending request")
	if err != nil {
		return nil, err
	}
	return client.Default.Do(req, nil)
}

func init() {
	logrus.WithField("service", "nlu").Infof("registering service")
	services.Register("nlu", newNLUService)
}
