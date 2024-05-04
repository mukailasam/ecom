package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ftsog/ecom/customerrors"
	"github.com/go-chi/chi/v5"
)

func (app *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		checker, _ := app.RDsession.CheckSession(w, r)
		if !(*checker) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:    http.StatusText(http.StatusUnauthorized),
					Status:   http.StatusUnauthorized,
					Message:  "Unauthorize",
					Detail:   "you are not allowed to this resource, login required",
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

		username, err := app.RDsession.GetUserFromSession(w, r)
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

		productID := chi.URLParam(r, "id")
		pID, err := strconv.ParseInt(productID, 10, 64)
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

		pd, err := app.Db.ReadProduct(pID)
		if err == sql.ErrNoRows {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusNotFound),
					Status:  http.StatusNotFound,
					Message: "Invalid object Id",
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

		if pd.OwnerName == *username {
			err := app.Db.UpdateProduct(nm, ct, pr, ds, other, pID)
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
					Message: "object successfully updated",
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
