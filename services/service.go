package services

import (
	"fmt"
	ghttp "github.com/oshankkumar/GatewayOmega/http"
	"net/http"
)

type Service interface {
	Send(r *ghttp.GatewayRequest) (*http.Response, error)
	Name() string
}

type NewServiceFunc func(...SerivceOptionFunc) Service

var ServiceFactory = make(map[string]NewServiceFunc)

type ServiceOption struct {
	Service Service
}

type SerivceOptionFunc func(*ServiceOption)

func WithService(srv Service)SerivceOptionFunc{
	return func(option *ServiceOption) {
		option.Service = srv
	}
}

func Register(name string, s NewServiceFunc) {
	ServiceFactory[name] = s
}

func Get(name string) (NewServiceFunc, error) {
	if srv, ok := ServiceFactory[name]; ok {
		return srv, nil
	}
	return nil, fmt.Errorf("service not exist")
}
