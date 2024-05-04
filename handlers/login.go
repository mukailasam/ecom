package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/ftsog/ecom/customerrors"
)

func (app *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		checker, _ := app.RDsession.CheckSession(w, r)
		if *checker {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:    http.StatusText(http.StatusBadRequest),
					Status:   http.StatusBadRequest,
					Message:  "already login",
					Detail:   "you are already login",
					Path:     "/api/auth/login",
					Redirect: "/api/index/hello",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		ud := userDetails{}
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

		em, ok1 := GetValue(jr, "email")
		paswd, ok2 := GetValue(jr, "password")

		if ok1 != true || ok2 != true {
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

		ud.Email = em
		ud.Password = paswd

		if ud.Email == "" || ud.Password == "" {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data:           responseError{Error: "All field Required"},
				statusCode:     http.StatusBadRequest,
			})

			return
		}

		usr, err := app.Db.GetByEmail(ud.Email)

		if err == sql.ErrNoRows {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data:           responseError{Error: "Incorrect email or password"},
				statusCode:     http.StatusBadRequest,
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

		check, err := app.Db.IsVerify(ud.Email)
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

		if !(*check) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:    http.StatusText(http.StatusBadRequest),
					Status:   http.StatusBadRequest,
					Message:  "Unverified Email",
					Detail:   "you can't login unverified, verify your email",
					Path:     "/login",
					Redirect: "/login",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		password := app.LoginHashedPassword(ud.Password, strings.TrimSpace(usr.Salt))
		if password == strings.TrimSpace(usr.Password) {
			err = app.RDsession.CreateSession(w, r, strings.TrimSpace(usr.Username))
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
					Status:   http.StatusOK,
					Message:  "Successfully Login",
					Path:     "/api/auth/login",
					Redirect: "/api/index/home",
				},

				statusCode: http.StatusOK,
			})

			return
		}

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data: responseError{
				Error:  "Incorrect email or password",
				Status: http.StatusText(http.StatusBadRequest),
			},
			statusCode: http.StatusBadRequest,
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

func (app *Handler) LoginHashedPassword(password string, iv string) (hashedPassword string) {

	iv = strings.TrimSpace(iv)

	output := password + iv
	output = strings.TrimSpace(output)

	hashPassword := sha256.Sum256([]byte(output))
	hPassToString := hex.EncodeToString(hashPassword[:])
	hashedPassword = strings.TrimSpace(hPassToString)

	return hashedPassword
}
