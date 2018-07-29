package utils

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func GetAuthTokenFromHeader(header http.Header) string {
	if authHeaderVal := header.Get("Authorization"); authHeaderVal != "" {
		return strings.TrimPrefix(authHeaderVal, "Bearer ")
	}
	if authHeaderVal := header.Get("Auth-Token"); authHeaderVal != "" {
		if len(authHeaderVal) > 6 && strings.ToLower(authHeaderVal[:6]) == "bearer" {
			return authHeaderVal[7:]
		}
		return authHeaderVal
	}
	return ""
}

func PrepareRequest(serviceUrl string, r *http.Request) *http.Request {
	var r2Uri = new(url.URL)
	var r2Headers = make(http.Header)
	var body = &bytes.Buffer{}
	serviceUri, err := url.Parse(serviceUrl)
	if err != nil {
		return r
	}
	{
		r2Uri.Scheme = serviceUri.Scheme
		r2Uri.Host = serviceUri.Host
		r2Uri.Path = r.URL.Path
		r2Uri.RawQuery = r.URL.Query().Encode()
	}
	{
		for key, val := range r.Header {
			r2Headers[key] = val
		}
	}
	io.Copy(body, r.Body)
	r2, err := http.NewRequest(r.Method, r2Uri.String(), body)
	if err != nil {
		return r
	}
	r2.Header = r2Headers
	return r2
}

func GetOutBoundIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}

var noAuthPathMap = map[string]bool{
	"/liveness":  true,
	"/readiness": true,
}

func IsNoAuthPath(r *http.Request) bool {
	if _,ok := noAuthPathMap[r.URL.Path]; ok {
		return true
	}
	return false
}
