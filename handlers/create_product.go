package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ftsog/ecom/customerrors"
	"github.com/ftsog/ecom/utils"
)

func (app *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		checker, _ := app.RDsession.CheckSession(w, r)
		if !(*checker) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:    http.StatusText(http.StatusUnauthorized),
					Status:   http.StatusUnauthorized,
					Message:  "Unauthorize",
					Detail:   "you are not allowed to this resource, account requred",
					Path:     "/api/product/create",
					Redirect: "/api/auth/register",
				},
				statusCode: http.StatusUnauthorized,
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

		name := jr.Data["name"]
		category := jr.Data["category"]
		price := jr.Data["price"]
		description := jr.Data["description"]
		other := jr.Data["other"]

		nm := name.(string)
		ct := category.(string)
		pr := price.(float64)
		ds := description.(string)

		user, err := app.RDsession.GetUserFromSession(w, r)
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

		usr, err := app.Db.GetByUsername(*user)
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

		err = utils.ProductEmpty(nm, ct, pr, ds)
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

		if len(jr.Data) <= 0 || len(jr.Data) < 5 || len(jr.Data) > 5 {
			app.Logger.ErrorLog.Println("Bad Request")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusBadRequest),
					Status:  http.StatusBadRequest,
					Message: "check the json data you are sending",
					Detail:  "make sure the json data contains seven field and not less or more",
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		pId, err := app.Db.CreateProduct(nm, ct, pr, ds, other, usr.Phone, usr.UserId, usr.Username)
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

		postIDCookie := &http.Cookie{
			Name:    "postID",
			Value:   fmt.Sprintf("%d", *pId),
			Expires: time.Now().Add(time.Duration(60 * time.Second)),
			Path:    "/api/product",
		}

		http.SetCookie(w, postIDCookie)

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data: responseMessage{
				Status:  http.StatusOK,
				Message: "Successfully Created",
				Path:    "/api/product/create",
			},

			statusCode: http.StatusOK,
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
