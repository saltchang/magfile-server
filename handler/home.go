package handler

import (
	"log"
	"net/http"
)

// HomeHandler handles all request to route "/users" or "/users/*"
func (h *HTTPHandler) HomeHandler(w http.ResponseWriter, r *Request) {
	eh := &errorHandler{w, r}
	log.Printf("%s was visited.", r.URL)

	switch r.Method {
	case http.MethodGet:
		// log.Println("GetUserByID")
		// h.GetUserByID(w, r)
		returnStringAsResponse(w, r, "Welcome to home, Method: GET")
		return
	default:
		eh.httpMethodNotAllowed(nil)
		return
	}
}
