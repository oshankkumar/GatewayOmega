package tracing

import (
	"github.com/opentracing/opentracing-go"
	ghttp "github.com/oshankkumar/GatewayOmega/http"
	"github.com/oshankkumar/GatewayOmega/middlewares"
	"github.com/oshankkumar/GatewayOmega/services"
	"net/http"
	"github.com/sirupsen/logrus"
)

func newTracingService(opts ...services.SerivceOptionFunc) services.Service {
	t := &Tracing{}
	opt := &services.ServiceOption{}
	for _, optFunc := range opts {
		optFunc(opt)
	}
	t.next = opt.Service
	return t
}


type Tracing struct {
	next services.Service
}

func (t *Tracing) Name() string {
	return t.next.Name()
}

func (t *Tracing) Send(r *ghttp.GatewayRequest) (*http.Response, error) {
	span,ctx := opentracing.StartSpanFromContext(r.Req.Context(),t.Name())
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx,span)
	r.Req = r.Req.WithContext(ctx)
	r.Req = middlewares.ToHTTPRequest(opentracing.GlobalTracer())(r.Req)
	return t.next.Send(r)
}

func init()  {
	logrus.WithField("service","zipkin").Infof("registering service")
	services.Register("zipkin",newTracingService)
}