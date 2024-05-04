package handlers

import "net/http"

func (app *Handler) About(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data:           responseMessage{Message: "Welcome to about page"},
			statusCode:     http.StatusOK,
		})

		return

	}

	JsonResponseWriter(JsonResponse{
		ResponseWriter: w,
		Data:           responseError{Error: http.StatusText(http.StatusMethodNotAllowed)},
		statusCode:     http.StatusMethodNotAllowed,
	})
}
