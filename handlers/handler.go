package handlers

import (
	ghttp "github.com/oshankkumar/GatewayOmega/http"
	"github.com/oshankkumar/GatewayOmega/services"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func NluHandlerFunc(w http.ResponseWriter, r *http.Request) {
	serviceFunc, err := services.Get("nlu")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	service := serviceFunc()
	resp, err := service.Send(&ghttp.GatewayRequest{Req: r})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logrus.WithFields(logrus.Fields{
		"service": "nlu",
		"status":  resp.Status,
	}).Infof("response received")
	for key, val := range resp.Header {
		w.Header().Set(key, val[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
