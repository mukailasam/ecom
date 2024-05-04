package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ftsog/ecom/customerrors"
	"github.com/go-chi/chi/v5"
)

func (app *Handler) ReadProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		username := chi.URLParam(r, "username")
		profile, err := app.Db.ReadProfile(username)
		if err == sql.ErrNoRows {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseMessage{
					Status:  http.StatusNotFound,
					Message: "Profile Does Not Exist",
				},

				statusCode: http.StatusNotFound,
			})

			return
		}

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

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data:           profile,
			statusCode:     http.StatusOK,
		})

		return
	}

	JsonResponseWriter(JsonResponse{
		ResponseWriter: w,
		Data: responseError{
			Error:   http.StatusText(http.StatusMethodNotAllowed),
			Status:  http.StatusMethodNotAllowed,
			Message: "Wrong HTTP VERB",
			Detail:  "Make use of the appropriate HTTP verb/method",
		},
		statusCode: http.StatusMethodNotAllowed,
	})
}

func (app *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		checker, _ := app.RDsession.CheckSession(w, r)
		if !(*checker) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusUnauthorized),
					Status:  http.StatusUnauthorized,
					Message: "Unauthorize",
					Detail:  "you are not allowed to this resource, account requred",
				},
				statusCode: http.StatusUnauthorized,
			})

			return
		}

		usr, err := app.RDsession.GetUserFromSession(w, r)
		if err != nil {
			app.Logger.ErrorLog.Println(err)
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

		jr, err := JsonRequestDecoder(r)
		if err != nil {
			app.Logger.ErrorLog.Println(err.Error())
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data:           responseError{Error: http.StatusText(http.StatusBadRequest)},
				statusCode:     http.StatusBadRequest,
			})

			return
		}

		firsName, ok2 := GetValue(jr, "FirstName")
		lastName, ok3 := GetValue(jr, "LastName")
		phone, ok4 := GetValue(jr, "PhoneNumber")
		address, ok5 := GetValue(jr, "Address")

		if ok2 != true && ok3 != true && ok4 != true && ok5 != true {
			app.Logger.ErrorLog.Println("Bad Request")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "check the json data you are sending",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		if len(jr.Data) <= 0 || len(jr.Data) < 4 || len(jr.Data) > 4 {
			app.Logger.ErrorLog.Println("Bad Request")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "check the json data you are sending",
					Detail:  "make sure the json data contains two field and not less or more",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		if firsName == "" || lastName == "" || phone == "" || address == "" {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data:           responseError{Error: "field can't be empty"},
				statusCode:     http.StatusBadRequest,
			})

			return
		}

		pass, err := app.Db.UpdateProfile(firsName, lastName, phone, address, *usr)
		if err == sql.ErrNoRows {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusForbidden,
					Status:  http.StatusText(http.StatusForbidden),
					Message: "Access denied",
				},
				statusCode: http.StatusForbidden,
			})

			return
		}

		if err != nil {
			app.Logger.ErrorLog.Println(err)
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

		if *pass == "pass" {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseMessage{
					Status:  http.StatusOK,
					Message: "Profile successfully updated",
				},
				statusCode: http.StatusOK,
			})

			return
		}

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data: responseError{
				Error:   http.StatusForbidden,
				Status:  http.StatusText(http.StatusForbidden),
				Message: "Access denied",
			},
			statusCode: http.StatusForbidden,
		})

		return

	}

	JsonResponseWriter(JsonResponse{
		ResponseWriter: w,
		Data: responseError{
			Error:   http.StatusText(http.StatusMethodNotAllowed),
			Status:  http.StatusMethodNotAllowed,
			Message: "Wrong HTTP VERB",
			Detail:  "Make use of the appropriate HTTP verb/method",
		},
		statusCode: http.StatusMethodNotAllowed,
	})
}

func (app *Handler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		checker, _ := app.RDsession.CheckSession(w, r)
		if !(*checker) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusUnauthorized),
					Status:  http.StatusUnauthorized,
					Message: "Unauthorize",
					Detail:  "you are not allowed to this resource, account requred",
				},
				statusCode: http.StatusUnauthorized,
			})

			return
		}

		usr, err := app.RDsession.GetUserFromSession(w, r)
		if err != nil {
			app.Logger.ErrorLog.Println(err)
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

		del, err := app.Db.DeleteProfile(*usr)

		if err == sql.ErrNoRows {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusNotFound,
					Status:  http.StatusText(http.StatusNotFound),
					Message: "Account does not exist",
				},
				statusCode: http.StatusNotFound,
			})

			return
		}

		if err != nil {
			app.Logger.ErrorLog.Println(err)
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

		if *del == "deleted" {
			err := app.RDsession.DeleteSession(w, r)
			if err != nil {
				app.Logger.ErrorLog.Println(err)
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
				Data: responseMessage{
					Status:  http.StatusOK,
					Message: "Account deleted",
				},
				statusCode: http.StatusOK,
			})

			return
		}

	}

	JsonResponseWriter(JsonResponse{
		ResponseWriter: w,
		Data: responseError{
			Error:   http.StatusText(http.StatusMethodNotAllowed),
			Status:  http.StatusMethodNotAllowed,
			Message: "Wrong HTTP VERB",
			Detail:  "Make use of the appropriate HTTP verb/method",
		},
		statusCode: http.StatusMethodNotAllowed,
	})
}
