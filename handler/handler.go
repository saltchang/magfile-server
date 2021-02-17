package handler

import (
	"net/http"

	db "github.com/saltchang/magfile-server/db/sqlc"
)

var dbInstance db.Database

// func NewHandler(db db.Database) http.Handler {
// 	dbInstance = db
// }
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
}
