package tracing

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go-opentracing"
	"github.com/oshankkumar/GatewayOmega/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func init() {
	var (
		tracer      opentracing.Tracer
		hostPort    = viper.GetString("http.port")
		hostIP      = utils.GetOutBoundIp()
		zipkinUrl   = "http://zipkin.iamplus.xyz/api/v1/spans"
		serviceName = fmt.Sprintf("%s-%s", "gateway", viper.GetString("env"))
	)
	logrus.WithFields(logrus.Fields{
		"hostPort" : hostPort,
		"hostIP"   : hostIP,
		"zipkinUrl" : zipkinUrl,
		"serviceName" : serviceName,
	}).Infof("initializing zipkin")
	collector, err := zipkintracer.NewHTTPCollector(
		zipkinUrl,
		zipkintracer.HTTPLogger(zipkintracer.LogWrapper(log.New(os.Stdout, "tracer:zipkin", log.LstdFlags))),
	)
	if err != nil {
		logrus.WithError(err).Errorf("error in creating http collector")
	}
	if tracer, err = zipkintracer.NewTracer(
		zipkintracer.NewRecorder(collector, false, fmt.Sprintf("%v:%v", hostIP, hostPort), serviceName),
		zipkintracer.ClientServerSameSpan(true),
		zipkintracer.TraceID128Bit(false),
	); err != nil {
		logrus.WithError(err).Errorf("error in creating new tracer")
	}
	opentracing.InitGlobalTracer(tracer)
}
