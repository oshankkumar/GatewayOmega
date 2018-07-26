package route

import "net/http"

type Route struct {
	Name        string
	Description string
	Pattern     string
	IsSubdomain bool
	Handler     http.Handler
}

// Routes is a table of many Route's
type Routes []Route
