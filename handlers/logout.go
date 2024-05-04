package handlers

import (
	"net/http"

	"github.com/ftsog/ecom/customerrors"
)

func (app *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := app.RDsession.DeleteSession(w, r)
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
				Message:  "successfully logout",
				Detail:   "you successfully logout your account",
				Path:     "/api/auth/logout",
				Redirect: "/api/auth/login",
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
