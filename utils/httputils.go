package utils

import (
	"io"
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
	var r2Uri *url.URL
	serviceUri, err := url.Parse(serviceUrl)
	if err != nil {
		return r
	}
	{
		r2Uri = r.URL
		r2Uri.Scheme = serviceUri.Scheme
		r2Uri.Host = serviceUri.Host
		r2Uri.RawQuery = r.URL.Query().Encode()
	}

	pr, pw := io.Pipe()
	go func() {
		io.Copy(pw, r.Body)
		pw.Close()
	}()
	r2, err := http.NewRequest(r.Method, r2Uri.String(), pr)
	if err != nil {
		return r
	}
	r2.Header = r.Header
	return r2
}
