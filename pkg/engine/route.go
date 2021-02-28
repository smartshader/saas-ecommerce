package engine

import "net/http"

// Route defines route and has a Handler to build
// the response pipeline and chain middlewares
type Route struct {
	Logger  bool
	Handler http.Handler
}
