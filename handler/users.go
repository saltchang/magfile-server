package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	db "github.com/saltchang/magfile-server/db/sqlc"
	"github.com/saltchang/magfile-server/util"
)

// UsersHandler handles all request to route "/users" or "/users/*"
func (h *HTTPHandler) UsersHandler(w http.ResponseWriter, r *Request) {
	eh := &errorHandler{w, r}
	log.Printf("%s was visited.", r.URL)

	switch r.Method {
	case http.MethodGet:
		// log.Println("GetUserByID")
		// h.GetUserByID(w, r)
		returnStringAsResponse(w, r, "Method: GET")
		return
	case http.MethodPost:
		// log.Println("CreateAnUser")
		// h.CreateAnUser(w, r)
		returnStringAsResponse(w, r, "Method: POST")
		return
	default:
		eh.httpMethodNotAllowed(nil)
		return
	}
}

// GetUserByID get BlogUser from databae by the given user ID.
func (h *HTTPHandler) GetUserByID(w http.ResponseWriter, r *Request) {
	eh := &errorHandler{w, r}

	idString, err := handleURLPattern(r.URL.String(), 3, 2)
	if err != nil {
		eh.httpMethodNotAllowed(err)
		return
	}

	log.Println(idString)

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
		log.Println("Wrong user id")
		eh.badRequest(err)
		return
	}

	h.Lock()

	user, err := h.db.GetBlogUser(context.Background(), id)
	h.Unlock()

	if err != nil {
		eh.notFound(err)
		log.Printf("BlogUser [id:%s] not found", idString)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		eh.internalServerError(err)
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// CreateAnUser create a BlogUser by the request body raw.
func (h *HTTPHandler) CreateAnUser(w http.ResponseWriter, r *Request) {
	log.Println("CreateAnUser")
	eh := &errorHandler{w, r}
	_, err := handleURLPattern(r.URL.String(), 2, 1)
	if err != nil {
		eh.httpMethodNotAllowed(err)
		return
	}

	hashSalt := os.Getenv("HASH_SALT")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		eh.internalServerError(err)
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		eh.unsupportedMediaType(err)
		return
	}

	var user db.BlogUser
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		eh.badRequest(err)
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

	newUser, err := h.db.CreateBlogUser(context.Background(), params)
	h.Unlock()

	if err != nil {
		eh.internalServerError(err)
		log.Printf("Creating BlogUser failed, error:%v", err)
		return
	}

	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		eh.internalServerError(err)
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
