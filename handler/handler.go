package handler

import (
	"net/http"
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

func handleURLPattern(URL string, expected, target int) (string, error) {
	parts := strings.Split(URL, "/")
	if len(parts) != expected {
		return "", nil
	}

	return parts[target], nil
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
