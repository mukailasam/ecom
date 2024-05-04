package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ftsog/ecom/utils"
)

type responseMessage struct {
	Status   interface{}
	Message  interface{}
	Detail   interface{}
	Path     interface{}
	Redirect interface{}
}

type responseError struct {
	Error    interface{}
	Status   interface{}
	Message  interface{}
	Detail   interface{}
	Path     interface{}
	Redirect interface{}
}

type userDetails struct {
	Username    string
	Email       string
	FirstName   string
	LastName    string
	Password    string
	PhoneNumber string
}

type JsonResponse struct {
	ResponseWriter http.ResponseWriter
	Data           any
	statusCode     int
}

func JsonResponseWriter(response JsonResponse) {
	response.ResponseWriter.Header().Set("Content-Type", "application/json")
	if response.Data != nil {
		resp, err := json.Marshal(response.Data)
		if err != nil {
			utils.InternalServerError(response.ResponseWriter, nil)
			response.ResponseWriter.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		utils.GeneralStatus(response.ResponseWriter, nil, response.statusCode)
		response.ResponseWriter.Write(resp)
		return
	}

	utils.InternalServerError(response.ResponseWriter, nil)
	response.ResponseWriter.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
