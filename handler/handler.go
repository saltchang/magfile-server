package handler

import (
	"net/http"
	"net/url"
	"strings"
	"sync"

	db "github.com/saltchang/magfile-server/db/sqlc"
)

// Request struct wraps the http.Request struct, providing a slice of
// strings representing the positional arguments found in a url pattern, and a
// map[string]string called kwargs representing the named parameters captured
// in url parsing.
type Request struct {
	*http.Request
	Args   []string
	Kwargs map[string]string
}

// Func is nearly the same as http.HandlerFunc, it simply takes a
// routes.Request object instead of an http.Request object.
type Func func(http.ResponseWriter, *Request)

// HTTPHandler handles the request from router
type HTTPHandler struct {
	sync.Mutex
	db *db.Queries
}

// countURLPattern counts the amount of patterns in the given URL.
func countURLPattern(URL string) int {
	path := strings.TrimPrefix(strings.TrimSuffix(URL, "/"), "/")
	parts := strings.Split(path, "/")

	return len(parts)
}

// getURLPattern parses then given URL string by expected index of target pattern, return the target string.
func getURLPattern(URL string, target int) string {
	path := strings.TrimPrefix(strings.TrimSuffix(URL, "/"), "/")
	parts := strings.Split(path, "/")

	return parts[target]
}

func parseURL(URL url.URL) (string, string) {
	query := URL.RawQuery
	URL.RawQuery = ""
	path := URL.Path

	return path, query
}

// NewHandler creates a new instance of httpHandler
func NewHandler(db *db.Queries) *HTTPHandler {
	handler := HTTPHandler{}
	handler.db = db

	return &handler
}

func returnStringAsResponse(w http.ResponseWriter, r *Request, s string) {
	w.Write([]byte(s))
}

// Router handles all request to server by the request path
// func (h *HTTPHandler) Router(w http.ResponseWriter, r *Request) {
// 	eh := &errorHandler{w, r}

// 	path := r.URL.Path

// 	log.Printf("route: \"%s\" was visited.", path)

// 	switch path {
// 	case "/users/":
// 		h.GetUserByID(w, r)
// 		return
// 	case "/users":
// 		h.CreateAnUser(w, r)
// 		return
// 	default:
// 		eh.httpMethodNotAllowed(nil)
// 		return
// 	}
// }
