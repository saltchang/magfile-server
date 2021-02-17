package router

import (
	"encoding/json"
	"net/http"

	db "github.com/saltchang/magfile-server/db/sqlc"
)

// UserHandler handle the request of route: "/users"
type UserHandler struct {
	store map[string]db.BlogUser
}

func (h *UserHandler) get(w http.ResponseWriter, r *http.Request) {
	users := make([]db.BlogUser, len(h.store))

	i := 0
	for _, user := range h.store {
		users[i] = user
		i++
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		// TODO
	}
	w.Write(jsonBytes)
}

// NewUserHandler returns a new userHandler instance
func NewUserHandler(store map[string]db.BlogUser) *UserHandler {
	return &UserHandler{store: store}
}
