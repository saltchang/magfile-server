package router

import (
	"net/http"
	"regexp"

	"github.com/saltchang/magfile-server/handler"
)

var h *handler.HTTPHandler

// Router struct implements http.Handler, so that it may be used with the
// default http library.  It keeps a registry mapping regexes to functions for
// easier url parsing.
type Router struct {
	routes []*entry
}

type entry struct {
	route       *RouteRegex
	handlerFunc handler.Func
}

// NewRouter create a new instance of type Router
func NewRouter() *Router {
	return &Router{}
}

// ServeHTTP implements the http.Handler interface, so that we may use our router with
// the default http package.
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, fn := router.match(r)
	if req == nil {
		http.NotFound(w, r)
		return
	}
	fn(w, req)
}

// checks an incoming http request against our list of known routes.  If the
// request matches one of the routes, the request is transformed into a
// routes.Request, and its Args and Kwargs fields are filled in based on the
// url.  If no match is found, returns (nil, nil)
func (router *Router) match(req *http.Request) (*handler.Request, handler.Func) {
	for _, e := range router.routes {
		if match := e.route.Match(req.URL.Path); match != nil {
			return &handler.Request{Request: req, Args: match.Args, Kwargs: match.Kwargs}, e.handlerFunc
		}
	}
	return nil, nil
}

// AddRoute adds a regex-based route in the normal human fashion.
func (router *Router) AddRoute(pattern string, fn handler.Func) {
	router.routes = append(router.routes, &entry{route: NewRoute(pattern), handlerFunc: fn})
}

// RouteRegex is the struct of route
type RouteRegex struct {
	*regexp.Regexp
}

// RouteMatch is the struct of route matching
type RouteMatch struct {
	Args   []string
	Kwargs map[string]string
}

// NewRoute create a new instance of route
func NewRoute(pattern string) *RouteRegex {
	return &RouteRegex{regexp.MustCompile(pattern)}
}

// Match create a route matching
func (r *RouteRegex) Match(target string) *RouteMatch {
	submatches := r.FindStringSubmatch(target)
	if submatches == nil {
		return nil
	}

	if len(submatches) == 1 {
		return new(RouteMatch)
	}

	m := new(RouteMatch)
	submatches = submatches[1:]
	for i, name := range r.SubexpNames()[1:] {
		if name == "" {
			m.Args = append(m.Args, submatches[i])
		} else {
			if m.Kwargs == nil {
				m.Kwargs = make(map[string]string)
			}
			m.Kwargs[name] = submatches[i]
		}
	}
	return m
}

// UseRouter create a router for the server to handling all routes
func UseRouter(innerHandler *handler.HTTPHandler) *Router {
	router := NewRouter()

	h = innerHandler

	routes := map[string]interface{}{
		"^/$":             h.HomeHandler,
		"^/users(/)?$":    h.UsersHandler,
		"^/users/(\\d)+$": h.UsersHandler,
	}

	for pattern, fn := range routes {
		router.AddRoute(
			pattern, fn.(func(http.ResponseWriter, *handler.Request)),
		)
	}

	return router
}
