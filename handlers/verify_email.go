package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/ftsog/ecom/customerrors"
	"github.com/go-chi/chi/v5"
)

func (app *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		username := chi.URLParam(r, "username")
		pToken := chi.URLParam(r, "token")
		username = strings.TrimSpace(username)
		pToken = strings.TrimSpace(pToken)

		if username == "" || pToken == "" {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data:           responseError{},
				statusCode:     http.StatusBadRequest,
			})
		}

		expired, err := app.Db.TokenExpired(username)
		if err != nil {
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

		if *expired == true {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseMessage{
					Status:  http.StatusNotFound,
					Message: "Expired email verification token",
				},
				statusCode: http.StatusNotFound,
			})
			return
		}

		token, err := app.Db.GetToken(username)
		if err == sql.ErrNoRows {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "Invalid link",
					Detail:  "wrong verification link",
				},
				statusCode: http.StatusBadRequest,
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

		if pToken == token {
			err := app.Db.VerifyEmail(username)
			if err != nil {
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

			err = app.Db.ExpiredToken(username)
			if err != nil {
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
		}

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data: responseMessage{
				Status:  http.StatusOK,
				Message: "Successfully verified your email",
			},
			statusCode: http.StatusOK,
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
