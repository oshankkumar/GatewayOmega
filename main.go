package main

import _ "github.com/oshankkumar/GatewayOmega/config"
import _ "github.com/oshankkumar/GatewayOmega/log"
import _ "github.com/oshankkumar/GatewayOmega/tracing"
import _ "github.com/oshankkumar/GatewayOmega/services/nlu"
import _ "github.com/oshankkumar/GatewayOmega/services/tracing"

import (
	"github.com/oshankkumar/GatewayOmega/handlers"
	"github.com/oshankkumar/GatewayOmega/server"
	"github.com/oshankkumar/GatewayOmega/server/route"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

var routes = route.Routes{
	{
		Name:        "nlu",
		Description: "Subdomain for NLU",
		Pattern:     "/api/v1/ai/",
		IsSubdomain: true,
		Handler:     http.HandlerFunc(handlers.NluHandlerFunc),
	},
}

func main() {
	gatewayServer := server.NewGatewayServer(":" + viper.GetString("http.port"))
	gatewayServer.RegisterRoutes(routes)
	logrus.Info("started listening on ", viper.GetString("http.port"))
	if err := gatewayServer.Start(); err != nil {
		logrus.Errorf("got error in starting server")
	}
}
