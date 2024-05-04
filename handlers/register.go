package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"

	//"math/rand"
	"net/http"
	"strings"

	"github.com/ftsog/ecom/customerrors"
	"github.com/ftsog/ecom/mailer"
	"github.com/ftsog/ecom/utils"
)

func (app *Handler) Register(w http.ResponseWriter, r *http.Request) {
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
					Path:     "/api/auth/register",
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

		if len(jr.Data) <= 0 || len(jr.Data) < 6 || len(jr.Data) > 6 {
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

		us, ok1 := GetValue(jr, "username")
		em, ok2 := GetValue(jr, "email")
		fn, ok3 := GetValue(jr, "firstName")
		ln, ok4 := GetValue(jr, "lastName")
		pw, ok5 := GetValue(jr, "password")
		pn, ok6 := GetValue(jr, "phone")

		if ok1 != true || ok2 != true || ok3 != true || ok4 != true || ok5 != true || ok6 != true {
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

		ud.Username = us
		ud.Email = em
		ud.FirstName = fn
		ud.LastName = ln
		ud.Password = pw
		ud.PhoneNumber = pn

		err = utils.IsEmpty(ud.Username, ud.Email, ud.FirstName, ud.LastName, ud.Password, ud.PhoneNumber)
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Detail:  "make sure to provide all json field and their appropriate value",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		username, err := utils.ValidateUsername(ud.Username)
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Detail:  "Username must be greater than 5 and less than 20 in length",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		err = app.Db.UserExists(ud.Username)
		if err == nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "User Exists",
					Detail:  "Username must be unique, we already have an accounted linked to this username",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		if err == sql.ErrNoRows {
			// nop
		} else {
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

		email, err := utils.ValidateEmail(ud.Email)
		if err != nil {
			app.Logger.ErrorLog.Println(err.Error())
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Detail:  "make sure your email is in correct format",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		err = app.Db.EmailExists(ud.Email)
		if err == nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "Email Exists",
					Detail:  "Email must be unique, we already have an accounted linked to this email",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		if err == sql.ErrNoRows {
			// nop
		} else {
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

		firstName, err := utils.ValidateFirstName(ud.FirstName)
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Detail:  "First Name can't be greater than 100 in length",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		lastName, err := utils.ValidateLastName(ud.LastName)
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Detail:  "Last Name can't be greater than 100 in length",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		passWord, err := utils.ValidatePassword(ud.Password)
		if err != nil {
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

		salt, hashedPassword := app.RegisterHashedPassword(*passWord)

		phone, err := utils.ValidatePhone(ud.PhoneNumber)
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Detail:  "make sure you phone numer is a valid phone number",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		token := app.Token()

		err = app.Db.CreateUser(*username, *email, *firstName, *lastName, hashedPassword, *phone, salt, token, false)
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

		subject := "Verify your email"

		err = mailer.NewMail(*email, subject, mailer.VerifyMessage, *username, token)
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
				Message:  "Account Creation Successfull",
				Detail:   "check your email for verification",
				Path:     "/api/auth/register",
				Redirect: "/api/auth/login",
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

func (app *Handler) IV() (result string) {
	//seed := time.Now().UTC().UnixNano()
	salt := make([]byte, 32)
	//rand.Seed(seed)
	rand.Read(salt)
	output := hex.EncodeToString(salt)
	result = strings.TrimSpace(output)
	return result
}

func (app *Handler) RegisterHashedPassword(password string) (iv string, hashedPassword string) {

	iv = app.IV()

	newPassword := password + iv

	hashPassword := sha256.Sum256([]byte(newPassword))
	hPassToString := hex.EncodeToString(hashPassword[:])
	hashedPassword = strings.TrimSpace(hPassToString)

	return iv, hashedPassword
}

func (app *Handler) Token() (result string) {
	salt := make([]byte, 32)
	rand.Read(salt)
	output := hex.EncodeToString(salt)
	result = strings.TrimSpace(output)
	return result
}
