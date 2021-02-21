package handler

import (
	"log"
	"net/http"
	"strings"
	"sync"

	db "github.com/saltchang/magfile-server/db/sqlc"
)

// HTTPHandler handles the request from router
type HTTPHandler struct {
	sync.Mutex
	db *db.Queries
}

var queries *db.Queries

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
	queries = db

	return &handler
}

func returnString(w http.ResponseWriter, r *http.Request, s string) {
	w.Write([]byte(s))
}

// Router handles all request to server by the request path
func (h *HTTPHandler) Router(w http.ResponseWriter, r *http.Request) {
	// eh := &errorHandler{w, r}

	path := r.URL.Path

	log.Printf("route: \"%s\" was visited.", path)
	returnString(w, r, path)

	return

	// switch path {
	// case "/user":
	// 	// log.Println("GetUserByID")
	// 	// h.GetUserByID(w, r)
	// 	returnString(w, r, path)
	// 	return
	// case http.MethodPost:
	// 	// log.Println("CreateAnUser")
	// 	// h.CreateAnUser(w, r)
	// 	returnString(w, r, "Method: POST")
	// 	return
	// default:
	// 	eh.httpMethodNotAllowed(nil)
	// 	return
	// }
}
