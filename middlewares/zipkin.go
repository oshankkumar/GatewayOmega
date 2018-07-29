package middlewares

import (
	"net/http"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/opentracing/opentracing-go/ext"
	"net"
	"strconv"
	"github.com/openzipkin/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"github.com/oshankkumar/GatewayOmega/utils"
)

func ZipkinTracing(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if utils.IsNoAuthPath(r){
			next.ServeHTTP(w,r)
			return
		}
		wireCtx,err := opentracing.GlobalTracer().Extract(
			opentracing.TextMap,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
		if err != nil {
			logrus.WithError(err).Errorf("encountterd an error while extracting span from req")
		}
		span := opentracing.GlobalTracer().StartSpan(
			fmt.Sprintf("%s %s",r.Method,r.URL.Path),
			ext.RPCServerOption(wireCtx),
		)
		defer span.Finish()
		ctx := opentracing.ContextWithSpan(r.Context(),span)
		req := r.WithContext(ctx)
		next.ServeHTTP(w,req)

	})
}

type RequestFunc func(req *http.Request) *http.Request

func ToHTTPRequest(tracer opentracing.Tracer) RequestFunc {
	return func(req *http.Request) *http.Request {
		// Retrieve the Span from context.
		if span := opentracing.SpanFromContext(req.Context()); span != nil {

			// We are going to use this span in a client request, so mark as such.
			ext.SpanKindRPCClient.Set(span)

			// Add some standard OpenTracing tags, useful in an HTTP request.
			ext.HTTPMethod.Set(span, req.Method)
			span.SetTag(zipkincore.HTTP_HOST, req.URL.Host)
			span.SetTag(zipkincore.HTTP_PATH, req.URL.Path)
			ext.HTTPUrl.Set(
				span,
				fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path),
			)

			// Add information on the peer service we're about to contact.
			if host, portString, err := net.SplitHostPort(req.URL.Host); err == nil {
				ext.PeerHostname.Set(span, host)
				if port, err := strconv.Atoi(portString); err != nil {
					ext.PeerPort.Set(span, uint16(port))
				}
			} else {
				ext.PeerHostname.Set(span, req.URL.Host)
			}

			// Inject the Span context into the outgoing HTTP Request.
			if err := tracer.Inject(
				span.Context(),
				opentracing.TextMap,
				opentracing.HTTPHeadersCarrier(req.Header),
			); err != nil {
				fmt.Printf("error encountered while trying to inject span: %+v\n", err)
			}
		}
		return req
	}
}
