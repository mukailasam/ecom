package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ftsog/ecom/customerrors"
	"github.com/go-chi/chi/v5"
)

func (app *Handler) ReadProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		productID := chi.URLParam(r, "id")
		pID, err := strconv.ParseInt(productID, 10, 64)
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusNotFound,
					Status:  http.StatusText(http.StatusNotFound),
					Message: "404 Not Found",
					Detail:  "the resource you are looking for not availale kindly check your url",
				},
				statusCode: http.StatusNotFound,
			})

			return
		}

		pd, err := app.Db.ReadProduct(pID)
		if err == sql.ErrNoRows {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusNotFound,
					Status:  http.StatusText(http.StatusNotFound),
					Message: "404 Not Found",
					Detail:  "the resource you are looking for not availale kindly check your url",
				},
				statusCode: http.StatusNotFound,
			})

			return

		} else if err != nil {
			app.Logger.ErrorLog.Println(err.Error())
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusInternalServerError),
					Status:  http.StatusInternalServerError,
					Message: customerrors.InternalErrorMessage,
					Detail:  customerrors.InternalErrorDetail,
				},
				statusCode: http.StatusInternalServerError,
			})

			return
		}

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data:           pd,
			statusCode:     http.StatusOK,
		})

		return
	}

	JsonResponseWriter(JsonResponse{
		ResponseWriter: w,
		Data: responseError{
			Error:    http.StatusText(http.StatusMethodNotAllowed),
			Status:   http.StatusMethodNotAllowed,
			Message:  "Wrong HTTP VERB",
			Detail:   "Make use of the appropriate HTTP verb/method",
			Path:     "/register",
			Redirect: "/register",
		},
		statusCode: http.StatusMethodNotAllowed,
	})
}
