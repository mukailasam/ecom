package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ftsog/ecom/customerrors"
	"github.com/go-chi/chi/v5"
)

func (app *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		checker, _ := app.RDsession.CheckSession(w, r)
		if !(*checker) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:    http.StatusText(http.StatusUnauthorized),
					Status:   http.StatusUnauthorized,
					Message:  "Unauthorize",
					Detail:   "you are not allowed to this resource",
					Path:     "/api/product/create",
					Redirect: "/api/auth/register",
				},
				statusCode: http.StatusUnauthorized,
			})
			return
		}

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
		if err != nil {
			if err == sql.ErrNoRows {
				app.Logger.ErrorLog.Println(err.Error())
				JsonResponseWriter(JsonResponse{
					ResponseWriter: w,
					Data: responseError{
						Error:   http.StatusText(http.StatusNotFound),
						Status:  http.StatusNotFound,
						Message: "Product Not Found",
						Detail:  "You can't delete a product that does not exist",
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

		if pd.OwnerName == *username {
			err = app.Db.DeleteProduct(pID)
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
					Message: "object successfully deleted",
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
