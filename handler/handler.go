package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	db "github.com/saltchang/magfile-server/db/sqlc"
	"github.com/saltchang/magfile-server/util"
)

// HTTPHandler handles the request from router
type HTTPHandler struct {
	sync.Mutex
	db *db.Queries
}

var queries *db.Queries

// GetUserByID get BlogUser from databae by the given user ID.
func (h *HTTPHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	idString := parts[2]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println("Wrong user id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.Lock()

	user, err := queries.GetBlogUser(context.Background(), id)
	h.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("BlogUser [id:%s] not found", idString)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

// CreateAnUser create a BlogUser by the request body raw.
func (h *HTTPHandler) CreateAnUser(w http.ResponseWriter, r *http.Request) {

	hashSalt := os.Getenv("HASH_SALT")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var user db.BlogUser
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	h.Lock()

	params := db.CreateBlogUserParams{
		Username:        user.Username,
		Email:           user.Email,
		FullName:        user.FullName,
		Gender:          user.Gender,
		CurrentLocation: user.CurrentLocation,
		PasswordHash:    util.GetPasswordHash(user.PasswordHash, hashSalt),
		LoginedAt:       time.Now().UTC(),
	}

	newUser, err := queries.CreateBlogUser(context.Background(), params)
	h.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("Creating BlogUser failed, error:%v", err)
		return
	}

	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
}

// NewHandler creates a new instance of httpHandler
func NewHandler(db *db.Queries) *HTTPHandler {
	handler := HTTPHandler{}
	queries = db

	return &handler
}
