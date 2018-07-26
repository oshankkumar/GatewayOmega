package nlu

import (
	"fmt"
	"github.com/oshankkumar/GatewayOmega/utils"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type NluHandler struct {
}

var client = &http.Client{}

func (nlu *NluHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r2 := utils.PrepareRequest("http://nlu-dev.aneeda.ai", r)
	fmt.Println(r2)
	logrus.WithField("service", "nlu").Infof("forwarding request")
	resp, err := client.Do(r2)
	if err != nil {
		logrus.WithError(err).Errorf("error in forwarding to nlu service")
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}
	logrus.WithFields(logrus.Fields{
		"status": resp.Status,
		"header": resp.Header,
	}).Infof("response recieved")
	for key, val := range resp.Header {
		w.Header().Set(key, val[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
