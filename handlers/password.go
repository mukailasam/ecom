package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/ftsog/ecom/customerrors"
	"github.com/ftsog/ecom/mailer"
	"github.com/ftsog/ecom/utils"
	"github.com/go-chi/chi/v5"
)

func (app *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

		currentPassword, ok1 := GetValue(jr, "CurrentPassword")
		newPassword, ok2 := GetValue(jr, "NewPassword")
		confirmNewPassword, ok3 := GetValue(jr, "ConfirmPassword")

		if ok1 != true && ok2 != true && ok3 != true {
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

		if newPassword != confirmNewPassword {
			app.Logger.ErrorLog.Println("Bad Request")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "Password does not match",
				},
				statusCode: http.StatusBadRequest,
			})
			return
		}

		ud, err := app.Db.GetByUsername(*usr)
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

		cPassword := app.LoginHashedPassword(currentPassword, strings.TrimSpace(ud.Salt))
		if ud.Password != cPassword {
			app.Logger.ErrorLog.Println("Bad Request")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "Incorrect current password",
					Detail:  "",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		npassword := app.LoginHashedPassword(newPassword, strings.TrimSpace(ud.Salt))
		err = app.Db.UpdatePassword(npassword, *usr)
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
				Message: "Password Updated Successfull",
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

func (app *Handler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

		email, ok := GetValue(jr, "email")
		if !ok {
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

		if len(jr.Data) <= 0 || len(jr.Data) < 1 || len(jr.Data) > 1 {
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

		if email == "" {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Detail:  "Make sure to provide required json fields and their appropriate value",
				},
				statusCode: http.StatusBadRequest,
			})

			return

		}

		eml, err := utils.ValidateEmail(email)
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "Invalid Email",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		ud, err := app.Db.GetByEmail(*eml)
		if err == sql.ErrNoRows {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "user does not exist",
					Detail:  "No account link with this email",
				},
				statusCode: http.StatusBadRequest,
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

		token := app.Token()
		err = app.Db.ResetPassword(token, false, *eml)
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

		subject := "Reset Password"

		err = mailer.NewMail(*eml, subject, mailer.ForgetMessage, ud.Username, token)
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
			Data: responseMessage{
				Status:  http.StatusOK,
				Message: "Check email to reset your password",
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

func (app *Handler) NewPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

		newPassword, ok1 := GetValue(jr, "NewPassword")
		confirmNewPassword, ok2 := GetValue(jr, "ConfirmPassword")
		if ok1 != true && ok2 != true {
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

		if len(jr.Data) <= 0 || len(jr.Data) < 2 || len(jr.Data) > 2 {
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

		newPassword = strings.TrimSpace(newPassword)
		confirmNewPassword = strings.TrimSpace(confirmNewPassword)

		if newPassword == "" || confirmNewPassword == "" {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: customerrors.FieldRequiredError.Error(),
					Detail:  "Make sure to provide required json fields and their appropriate value",
				},
				statusCode: http.StatusBadRequest,
			})
			return
		}

		if newPassword != confirmNewPassword {
			app.Logger.ErrorLog.Println("Bad Request")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "password not match",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		newPasswd, err := utils.ValidatePassword(newPassword)
		if err != nil {
			app.Logger.ErrorLog.Println("Bad Request")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Detail:  "make sure you your password has every infomation mention in the message",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		user := chi.URLParam(r, "username")
		token := chi.URLParam(r, "token")

		if strings.TrimSpace(user) == "" || strings.TrimSpace(token) == "" {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:  http.StatusText(http.StatusBadRequest),
					Status: http.StatusBadRequest,
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		salt, err := app.Db.GetSalt(user)
		if err != nil {
			if err == sql.ErrNoRows {
				app.Logger.ErrorLog.Println(err.Error())
				app.Logger.ErrorLog.Println(err.Error())
				JsonResponseWriter(JsonResponse{
					ResponseWriter: w,
					Data: responseError{
						Error:   http.StatusText(http.StatusNotFound),
						Message: "User does not exist",
					},
					statusCode: http.StatusNotFound,
				})
				return
			}

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

		ud, err := app.Db.Getprt(user)
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

		if ud.PrtExpired == true {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "Expired token",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		if token != ud.Prt {
			app.Logger.ErrorLog.Println("Bad Request")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "wrong token",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		passwd := app.LoginHashedPassword(*newPasswd, *salt)
		err = app.Db.Setprt(true, user, passwd)
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
			Data: responseError{
				Error:   http.StatusText(http.StatusOK),
				Status:  http.StatusOK,
				Message: "password reset sucessfully",
			},
			statusCode: http.StatusBadRequest,
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
