package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data any `json:"data"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any

	if payload != nil {
		response = &ResponseData{
			Data: payload,
		}
	} else {
		response = &ResponseError{
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}