package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/oshankkumar/GatewayOmega/middlewares"
	"github.com/oshankkumar/GatewayOmega/server/route"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type GatewayServer struct {
	Router *mux.Router
	server *http.Server
}

func NewGatewayServer(addr string) *GatewayServer {
	return &GatewayServer{
		server: &http.Server{Addr: addr},
		Router: mux.NewRouter(),
	}
}

func (g *GatewayServer) Start() error {
	g.server.Handler = alice.New(
		middlewares.Authentication,
	).Then(g.Router)
	return g.server.ListenAndServe()
}

func (g *GatewayServer) Stop() {
	g.server.Shutdown(context.Background())
}

func (g *GatewayServer) RegisterRoutes(routes route.Routes) {
	for _, route := range routes {
		log.WithFields(log.Fields{
			"name": route.Name,
			"path": route.Pattern,
		}).Info("registering new route")
		if route.IsSubdomain {
			g.Router.
				PathPrefix(route.Pattern).
				Name(route.Name).
				Handler(http.StripPrefix(route.Pattern, route.Handler))
		} else {
			g.Router.
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.Handler)
		}

	}
}
