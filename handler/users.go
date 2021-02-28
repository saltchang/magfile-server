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
	"time"

	db "github.com/saltchang/magfile-server/db/sqlc"
	"github.com/saltchang/magfile-server/util"
)

// CreateAnUserRequestBody is a struct of the raw body content of the request of creating an user.
type CreateAnUserRequestBody struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	PasswordHash string `json:"password_hash"`
}

// UsersHandler handles all request to route "/users" or "/users/*"
func (h *HTTPHandler) UsersHandler(w http.ResponseWriter, r *Request) {
	eh := &errorHandler{w, r}

	path, _ := parseURL(*r.URL)

	patternCount := countURLPattern(path)

	if patternCount == 1 {
		switch r.Method {
		case http.MethodPost:
			log.Println("CreateAnUser")
			h.CreateAnUser(w, r)
			return
		default:
			eh.httpMethodNotAllowed(nil)
			return
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			log.Println("GetUserByID")
			// h.GetUserByID(w, r)
			returnStringAsResponse(w, r, fmt.Sprintf("Method: %s", r.Method))
			return
		default:
			eh.httpMethodNotAllowed(nil)
			return
		}
	}

}

// GetUserByID get BlogUser from databae by the given user ID.
func (h *HTTPHandler) GetUserByID(w http.ResponseWriter, r *Request) {
	eh := &errorHandler{w, r}

	idString := getURLPattern(r.URL.String(), 1)

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

	var requestedUser *CreateAnUserRequestBody
	err = json.Unmarshal(bodyBytes, &requestedUser)
	if err != nil {
		eh.badRequest(err)
		return
	}

	// Check username
	h.Lock()
	_, err = h.db.GetBlogUserByUsername(context.Background(), requestedUser.Username)
	h.Unlock()

	if err == nil {
		eh.duplicateUsername(nil)
		log.Printf("Creating BlogUser failed, duplicate username")
		return
	}

	// Check email
	h.Lock()
	_, err = h.db.GetBlogUserByEmail(context.Background(), requestedUser.Email)
	h.Unlock()

	if err == nil {
		eh.duplicateEmail(nil)
		log.Printf("Creating BlogUser failed, duplicate email")
		return
	}

	params := db.CreateBlogUserParams{
		Username:        requestedUser.Username,
		Email:           requestedUser.Email,
		FullName:        requestedUser.FullName,
		Gender:          "",
		CurrentLocation: "",
		PasswordHash:    util.GetPasswordHash(requestedUser.PasswordHash, hashSalt),
		LoginedAt:       time.Now().UTC(),
	}

	h.Lock()
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
