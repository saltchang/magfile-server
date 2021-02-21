package handler

import "net/http"

// ResponseCode is a type of response code, not http status code
type ResponseCode int

// One of the collections of ResponseCode
const (
	// Normal
	RcSuccess ResponseCode = 1200

	// Error

	RcBadRequest      ResponseCode = 2400
	RcWrongRequestURL ResponseCode = 2000
	RcWrongParams     ResponseCode = 2100

	// HTTP Status: 405 Method Not Allowed
	RcMethodNotAllowed ResponseCode = http.StatusMethodNotAllowed

	// Undefined Error
	RcUndefined ResponseCode = 9999
)

// ErrorCodeList is a list that collects all types of error code
var ErrorCodeList = []ResponseCode{
	RcWrongRequestURL,
	RcUndefined,
}

// ResponseMessage is a type of response message
type ResponseMessage string

// One of the collections of ResponseMessage
const (
	RmWrongRequestURL ResponseMessage = "Wrong request URL."
	RmWrongParams     ResponseMessage = "Bad request: Wrong parameters."

	// HTTP Status: 405 Method Not Allowed
	RmMethodNotAllowed ResponseMessage = "Method Not Allowed"
)

// Response is the response struct, contains basic response and data.
type Response struct {
	ResponseBasic
	Data interface{} `json:"data"`
}

// ResponseBasic is the basic struct of response, contains code, message, and remark.
type ResponseBasic struct {
	Code    ResponseCode    `json:"code"`
	Message ResponseMessage `json:"message"`
	Remark  string          `json:"remark"`
}
