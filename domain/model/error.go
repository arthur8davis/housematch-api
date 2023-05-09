package model

const InternalServerErrorCode = "INTERNAL_SERVER_ERROR"
const InternalServerErrorDescription = "An internal server error has occurred"

type ErrorApi struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseError struct {
	Error string `json:"error"`
}
