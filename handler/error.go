package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponse ResponseBasic

func (e *errorResponse) json() string {

	er := &errorResponse{e.Code, e.Message, e.Remark}
	errStr, err := json.Marshal(er)
	if err != nil {
		errStr = []byte("{\"code\":0,\"message\":\"INTERNAL SERVER ERROR: Error occurred during error converting.\",\"remark\":\"\"}")
	}

	return string(errStr)
}

func (e *errorResponse) New(code ResponseCode, message ResponseMessage, remark string) *errorResponse {
	var eR errorResponse
	eR.Code = code
	eR.Message = message
	eR.Remark = remark
	return &eR
}

type errorHandler struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *errorHandler) writer(status int, err error, code ResponseCode, message ResponseMessage) {
	var eR *errorResponse
	var remarkStr string
	if err == nil {
		remarkStr = ""
	} else {
		remarkStr = err.Error()
	}
	errorResponse := eR.New(code, message, remarkStr)

	h.w.Header().Set("Content-type", "application/json")
	h.w.WriteHeader(status)
	h.w.Write([]byte(errorResponse.json()))
}

func (h *errorHandler) httpMethodNotAllowed(err error) {
	h.writer(http.StatusMethodNotAllowed, err, RcMethodNotAllowed, RmMethodNotAllowed)
}

func (h *errorHandler) notFound(err error) {
	h.writer(http.StatusNotFound, err, RcWrongRequestURL, RmWrongRequestURL)
}

func (h *errorHandler) badRequest(err error) {
	h.writer(http.StatusBadRequest, err, RcWrongRequestURL, RmWrongRequestURL)
}

func (h *errorHandler) internalServerError(err error) {
	h.writer(http.StatusInternalServerError, err, RcWrongRequestURL, RmWrongRequestURL)
}

func (h *errorHandler) unsupportedMediaType(err error) {
	h.writer(http.StatusUnsupportedMediaType, err, RcWrongRequestURL, RmWrongRequestURL)
}
