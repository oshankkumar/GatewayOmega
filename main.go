package main

import (
	"github.com/oshankkumar/GatewayOmega/server"
	"github.com/oshankkumar/GatewayOmega/server/route"
	"github.com/oshankkumar/GatewayOmega/services/nlu"
	"github.com/sirupsen/logrus"
)

var routes = route.Routes{
	{
		Name:        "nlu",
		Description: "Subdomain for NLU",
		Pattern:     "/api/v1/ai/",
		IsSubdomain: true,
		Handler:     &nlu.NluHandler{},
	},
}

func main() {
	gatewayServer := server.NewGatewayServer(":8080")
	gatewayServer.RegisterRoutes(routes)
	logrus.Info("started listening on 8080")
	if err := gatewayServer.Start(); err != nil {
		logrus.Errorf("got error in starting server")
	}
}
